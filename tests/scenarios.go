/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
	is "gotest.tools/assert/cmp"
)

// UndoParameter defines how to populate operation parameters.
type UndoParameter struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// Undo holds information about undo operation.
type Undo struct {
	Type string `json:"type"`
	OperationID string `json:"operationId"`
	Parameters []UndoParameter `json:"parameters"`
}

// UndoAction describes undo action.
type UndoAction struct {
	Tag string `json:"tag"`
	Undo *Undo `json:"undo"`
}

type ctxKey struct{}
type clientKey struct{}
type apiKey struct{}
type requestKey struct{}
type requestNameKey struct{}
type requestArgsKey struct{}
type requestParamsKey struct{}
type ctxRequestsUndoKey struct{}
type responseKey struct{}
type dataKey struct{}
type bodyKey struct{}
type cleanupKey struct{}

// Templated replaces {{ path }} in source with value from data[path].
func Templated(data interface{}, source string) string {
	re := regexp.MustCompile(`{{ ?([^}])+ ?}}`)
	replace := func(source string) string {
		path := strings.Trim(source, "{ }")
		v, err := lookup.LookupStringI(data, path)
		if err != nil {
			panic(fmt.Sprintf("problem with replacement of %s: %v", source, err))
		}
		return v.String()
	}
	return re.ReplaceAllStringFunc(source, replace)
}

// GetCtx returns Go context.
func GetCtx(ctx gobdd.Context) context.Context {
	c, err := ctx.Get(ctxKey{})
	if err != nil {
		panic(err)
	}
	return c.(context.Context)
}

// GetResponse returns request response.
func GetResponse(ctx gobdd.Context) []reflect.Value {
	r, err := ctx.Get(responseKey{})
	if err != nil {
		panic(err)
	}
	return r.([]reflect.Value)
}

// LoadRequestsUndo load undo configuration.
func LoadRequestsUndo(file string) map[string]UndoAction {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var value map[string]UndoAction
	json.Unmarshal(byteValue, &value)
	return value
}

// SetRequestsUndo sets map with undo function for each request.
func SetRequestsUndo(ctx gobdd.Context, value map[string]UndoAction) {
	ctx.Set(ctxRequestsUndoKey{}, value)
}

// GetRequestsUndo gets map with undo function for each request.
func GetRequestsUndo(ctx gobdd.Context, operationID string) (func(), error) {
	requestsUndo, err := ctx.Get(ctxRequestsUndoKey{})
	if err != nil {
		return nil, err
	}

	// find undo definition
	undo, ok := requestsUndo.(map[string]UndoAction)[operationID]
	if !ok {
		return nil, fmt.Errorf("undefined undo for operation: %s", operationID)
	}

	if undo.Undo.Type != "unsafe" {
		return nil, nil
	}

	// get an instance of Client
	undoClientInterface, err := ctx.Get(clientKey{})
	if err != nil {
		return nil, err
	}

	// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
	tag := strings.Replace(undo.Tag, " ", "", -1) + "Api"
	undoAPI := reflect.Indirect(undoClientInterface.(reflect.Value)).FieldByName(tag)
	if !undoAPI.IsValid() {
		return nil, fmt.Errorf("invalid API name %sApi", tag)
	}

	// get undo method from the service API: `client.{{ undo.Tag }}Api.{{ undo.Undo.OperationID }}`
	undoOperation := undoAPI.MethodByName(undo.Undo.OperationID)
	if !undoOperation.IsValid() {
		return nil, fmt.Errorf("invalid method name %s", undo.Undo.OperationID)
	}

	return func() {
		// Assemble undo method as follows: Client(cctx).UsersApi.DisableUser(cctx, response.Data.GetId()).Execute()
		response := GetResponse(ctx)[0].Interface()

		in := make([]reflect.Value, undoOperation.Type().NumIn())
		// first argument is always context.Context
		in[0] = reflect.ValueOf(GetCtx(ctx))
		for i := 1; i < undoOperation.Type().NumIn(); i++ {
			object, err := lookup.LookupStringI(response, undo.Undo.Parameters[i-1].Source)
			if err != nil {
				panic(err)
			}
			in[i] = object
		}
		request := undoOperation.Call(in)[0]
		request.MethodByName("Execute").Call(nil)
	}, nil
}

// SetCtx sets Go context in BDD context.
func SetCtx(ctx gobdd.Context, value context.Context) {
	ctx.Set(ctxKey{}, value)
}

// SetClient sets client reflection.
func SetClient(ctx gobdd.Context, value interface{}) {
	ctx.Set(clientKey{}, value)
}

// SetAPI sets client API.
func SetAPI(ctx gobdd.Context, value interface{}) {
	ctx.Set(apiKey{}, value)
}

// GetData returns feature context.
func GetData(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(dataKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

// SetData sets feature context.
func SetData(ctx gobdd.Context, value map[string]interface{}) {
	ctx.Set(dataKey{}, value)
}

// GetRequestParameters helps to build a request.
func GetRequestParameters(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(requestParamsKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

// GetRequestArguments helps to build a request.
func GetRequestArguments(ctx gobdd.Context) []interface{} {
	c, err := ctx.Get(requestArgsKey{})
	if err != nil {
		panic(err)
	}
	return c.([]interface{})
}

// SetCleanup functions.
func SetCleanup(ctx gobdd.Context, value map[string]func()) {
	ctx.Set(cleanupKey{}, value)
}

// GetCleanup functions.
func GetCleanup(ctx gobdd.Context) map[string]func() {
	c, err := ctx.Get(cleanupKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]func())
}

// RunCleanup executes cleanup functions in controlled ordered based on registered keys.
func RunCleanup(ctx gobdd.Context) {
	c, _ := ctx.Get(cleanupKey{})
	cc := c.(map[string]func())
	keys := make([]string, 0, len(cc))
	for k := range cc {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		cc[name]()
	}
}

// newRequest sets callable operation to requestKey{}
func newRequest(t gobdd.StepTest, ctx gobdd.Context, name string) {
	c, err := ctx.Get(apiKey{})
	if err != nil {
		panic(err)
	}
	r := c.(reflect.Value)
	f := r.MethodByName(name)
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid method name %s on API", name))
	}

	ctx.Set(requestKey{}, f)
	ctx.Set(requestNameKey{}, name)
	ctx.Set(requestParamsKey{}, make(map[string]interface{}))
	ctx.Set(requestArgsKey{}, make([]interface{}, 0))
}

// getRequestBuilder returns the reflect value of the current request in ctx
func getRequestBuilder(ctx gobdd.Context) (reflect.Value, error) {
	c, err := ctx.Get(requestKey{})
	if err != nil {
		return reflect.Value{}, fmt.Errorf("Missing requestKey{}")
	}
	f := c.(reflect.Value)

	in := make([]reflect.Value, f.Type().NumIn())

	// first argument is always context.Context
	in[0] = reflect.ValueOf(GetCtx(ctx))
	for i := 1; i < f.Type().NumIn(); i++ {
		object := GetRequestArguments(ctx)[i-1]
		in[i] = object.(reflect.Value)
	}
	return f.Call(in)[0], nil
}

func statusIs(t gobdd.StepTest, ctx gobdd.Context, expected int, text string) {
	// Execute() returns tripples -> 2nd value is *http.Response -> get StatusCode
	resp := GetResponse(ctx)
	code := resp[len(resp)-2].Interface().(*http.Response).StatusCode
	if expected != code {
		t.Fatalf("Excepted %d got %d", expected, code)
	}
}

func addParameterFrom(t gobdd.StepTest, ctx gobdd.Context, name string, path string) {
	value, err := lookup.LookupStringI(GetData(ctx), path)
	if err != nil {
		panic(err)
	}
	GetRequestParameters(ctx)[name] = value
	ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), value))
}

func addParameterWithValue(t gobdd.StepTest, ctx gobdd.Context, param string, value string) {
	// Get request builder for the current scenario
	request, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}
	name := ToVarName(param)
	// Get the method for setting the current parameter
	method := request.MethodByName(name)

	templatedValue := Templated(GetData(ctx), value)

	if method.IsValid() {
		at := reflect.New(method.Type().In(0))
		// Unmarshall the value specified in the step to the proper type and add it to the parameters
		json.Unmarshal([]byte(templatedValue), at.Interface())
		GetRequestParameters(ctx)[param] = at.Elem()
	}
}

func requestIsSent(t gobdd.StepTest, ctx gobdd.Context) {

	request, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}

	requestParameters := GetRequestParameters(ctx)

	for param, value := range requestParameters {
		name := ToVarName(param)
		method := request.MethodByName(name)
		if method.IsValid() {
			if param == "body" {
				at := reflect.New(method.Type().In(0))
				json.Unmarshal([]byte(value.(string)), at.Interface())
				request = method.Call([]reflect.Value{at.Elem()})[0]
			} else {
				request = method.Call([]reflect.Value{value.(reflect.Value)})[0]
			}
		}
	}

	operationID, err := ctx.GetString(requestNameKey{})
	if err != nil {
		t.Errorf("could not get a request name: %v", err)
	}
	if undo, err := GetRequestsUndo(ctx, operationID); err == nil {
		if undo != nil {
			GetCleanup(ctx)["01-undo"] = undo
		}
	} else {
		t.Errorf("missing undo for %s: %v", operationID, err)
	}

	result := request.MethodByName("Execute").Call(nil)
	ctx.Set(responseKey{}, result)
}

func body(t gobdd.StepTest, ctx gobdd.Context, body string) {
	GetRequestParameters(ctx)["body"] = Templated(GetData(ctx), body)
}

func stringToType(s string, t interface{}) (interface{}, error) {
	switch t.(type) {
	case int:
		return strconv.Atoi(s)
	case int64:
		return strconv.ParseInt(s, 10, 64)
	case string:
		return strconv.Unquote(s)
	case bool:
		return strconv.ParseBool(s)
	default:
		return nil, errors.New("Unknown type to convert")
	}
}

func expectEqual(t gobdd.StepTest, ctx gobdd.Context, responsePath string, value string) {
	responseValue, err := lookup.LookupStringI(GetResponse(ctx)[0].Interface(), SnakeToCamelCase(responsePath))
	if err != nil {
		t.Errorf("could not lookup response value %s in %v: %v", responsePath, GetResponse(ctx)[0].Interface(), err)
	}

	templatedValue := Templated(GetData(ctx), value)
	testValue, err := stringToType(templatedValue, responseValue.Interface())
	if err != nil {
		t.Errorf("%v", err)
	}

	cmp := is.DeepEqual(testValue, responseValue.Interface())()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectEqualValue(t gobdd.StepTest, ctx gobdd.Context, responsePath string, fixturePath string) {
	fixtureValue, err := lookup.LookupStringI(GetData(ctx), SnakeToCamelCase(fixturePath))
	if err != nil {
		t.Fatalf("could not lookup fixture value %s: %v", fixturePath, err)
	}
	responseValue, err := lookup.LookupStringI(GetResponse(ctx)[0].Interface(), SnakeToCamelCase(responsePath))
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", responsePath, err)
	}
	cmp := is.DeepEqual(responseValue.Interface(), fixtureValue.Interface())()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectLengthEqual(t gobdd.StepTest, ctx gobdd.Context, responsePath string, fixtureLength string) {
	lengthInt, err := strconv.Atoi(fixtureLength)
	if err != nil {
		t.Fatalf("assertion length value is not a number %s: %v", fixtureLength, err)
	}
	responseValue, err := lookup.LookupStringI(GetResponse(ctx)[0].Interface(), SnakeToCamelCase(responsePath))
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", responsePath, err)
	}
	cmp := is.Len(responseValue.Interface(), lengthInt)()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectFalse(t gobdd.StepTest, ctx gobdd.Context, responsePath string) {
	responseValue, err := lookup.LookupStringI(GetResponse(ctx)[0].Interface(), SnakeToCamelCase(responsePath))
	if err != nil {
		t.Errorf("could not lookup value: %v", err)
	}
	if responseValue.Bool() {
		t.Errorf("%v should be false", responseValue)
	}
}

// ConfigureSteps on given suite.
func ConfigureSteps(s *gobdd.Suite) {
	steps := map[string]interface{}{
		`new "([^"]+)" request`:                                  newRequest,
		`request contains "([^"]+)" parameter from "([^"]+)"`:    addParameterFrom,
		`request contains "([^"]+)" parameter with value (.+)`:   addParameterWithValue,
		`the request is sent`:                                    requestIsSent,
		`the response status is (\d+) (.*)`:                      statusIs,
		`body (.*)`:                                              body,
		`the response "([^"]+)" is equal to (.*)`:                expectEqual,
		`the response "([^"]+)" has the same value as "([^"]+)"`: expectEqualValue,
		`the response "([^"]+)" has length ([0-9]+)`:             expectLengthEqual,
		`the response "([^"]+)" is false`:                        expectFalse,
	}
	for expr, step := range steps {
		s.AddStep(expr, step)
	}
}
