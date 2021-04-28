/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"bytes"
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
	"testing"
	"time"

	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
	is "gotest.tools/assert/cmp"
)

// UndoAction describes undo action.
type UndoAction struct {
	Tag  string `json:"tag"`
	Undo *struct {
		Type        string `json:"type"`
		OperationID string `json:"operationId"`
		Parameters  []struct {
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"parameters"`
	} `json:"undo"`
}

type operationParameter struct {
	Name   string  `json:"name"`
	Source *string `json:"source"`
	Value  *string `json:"value"`
}

func (p operationParameter) Resolve(t gobdd.StepTest, ctx gobdd.Context, tp reflect.Type) reflect.Value {
	if p.Value != nil {
		tpl := Templated(t, GetData(ctx), *p.Value)
		v := reflect.New(tp)
		err := json.Unmarshal([]byte(tpl), v.Interface())
		if err != nil {
			t.Fatalf("can't unmarshal parameter value for %s: %v", p.Name, err)
		}
		return v.Elem()
	}
	v, _ := lookup.LookupStringI(GetData(ctx), *p.Source)
	return v
}

// GivenStep defines a step.
type GivenStep struct {
	Tag         string               `json:"tag"`
	Key         string               `json:"key"`
	Step        string               `json:"step"`
	OperationID string               `json:"operationId"`
	Parameters  []operationParameter `json:"parameters"`
	Source      *string              `json:"source"`
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
type jsonResponseKey struct{}
type dataKey struct{}
type bodyKey struct{}
type cleanupKey struct{}
type pathParamCountKey struct{}

// GetIgnoredTags returns list of ignored tags.
func GetIgnoredTags() []string {
	tags := make([]string, 1)
	tags = append(tags, "@skip")
	if GetRecording() != ModeIgnore {
		tags = append(tags, "@integration-only")
	}
	return tags
}

// Templated replaces {{ path }} in source with value from data[path].
func Templated(t gobdd.StepTest, data interface{}, source string) string {
	re := regexp.MustCompile(`{{ ?([^}])+ ?}}`)
	replace := func(source string) string {
		path := strings.Trim(source, "{ }")
		v, err := lookup.LookupStringI(data, path)
		if err != nil {
			t.Fatalf("problem with replacement of %s: %v", source, err)
		}
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(v.Int(), 10)
		case reflect.Float64, reflect.Float32:
			return strconv.FormatFloat(v.Float(), 'f', -1, 10)
		case reflect.Bool:
			return strconv.FormatBool(v.Bool())
		}
		return v.String()
	}
	return re.ReplaceAllStringFunc(source, replace)
}

// GetT returns stored reference to the testing object.
func GetT(ctx gobdd.Context) *testing.T {
	t, err := ctx.Get(gobdd.TestingTKey{})
	if err != nil {
		return nil
	}
	tt, ok := t.(*testing.T)
	if !ok {
		return nil
	}
	return tt
}

// GetCtx returns Go context.
func GetCtx(ctx gobdd.Context) context.Context {
	c, err := ctx.Get(ctxKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a context: %v", err)
	}
	return c.(context.Context)
}

// GetResponse returns request response.
func GetResponse(ctx gobdd.Context) []reflect.Value {
	r, err := ctx.Get(responseKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a response: %v", err)
	}
	return r.([]reflect.Value)
}

// GetResponse returns request response.
func GetJSONResponse(ctx gobdd.Context) interface{} {
	r, err := ctx.Get(jsonResponseKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a json response: %v", err)
	}
	return r
}

// GetResponseStatusCode returns request response status code.
func GetResponseStatusCode(resp []reflect.Value) (int, error) {
	if len(resp) < 2 {
		return -1, fmt.Errorf("needs at least 2 values, got %d", len(resp))
	}
	// Execute() returns triples -> 2nd value is *http.Response -> get StatusCode]
	response := resp[len(resp)-2].Interface().(*http.Response)
	if response == nil {
		return -1, errors.New("response is nil")
	}
	return response.StatusCode, nil
}

// LoadRequestsUndo load undo configuration.
func LoadRequestsUndo(file string) (map[string]UndoAction, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var value map[string]UndoAction
	json.Unmarshal(byteValue, &value)
	return value, nil
}

// LoadGivenSteps load undo configuration.
func LoadGivenSteps(file string) ([]GivenStep, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var value []GivenStep
	byteValue, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &value)
	return value, err
}

// SetRequestsUndo sets map with undo function for each request.
func SetRequestsUndo(ctx gobdd.Context, value map[string]UndoAction) {
	ctx.Set(ctxRequestsUndoKey{}, value)
}

// RegisterSuite adds step implementation.
func (s GivenStep) RegisterSuite(suite *gobdd.Suite) {
	given := func(t gobdd.StepTest, ctx gobdd.Context) {
		// get an instance of Client
		clientInterface, err := ctx.Get(clientKey{})
		if err != nil {
			t.Fatalf("could not find client: %v", err)
		}

		// Enable unstable operation
		configInterface := reflect.Indirect(clientInterface.(reflect.Value)).Addr().MethodByName("GetConfig").Call(nil)[0]
		if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{reflect.ValueOf(s.OperationID)})[0].Bool() {
			configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
				reflect.ValueOf(s.OperationID), reflect.ValueOf(true),
			})
		}

		// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
		tag := strings.Replace(s.Tag, " ", "", -1) + "Api"
		api := reflect.Indirect(clientInterface.(reflect.Value)).FieldByName(tag)
		if !api.IsValid() {
			t.Fatalf("invalid API name %sApi", tag)
		}

		// get method from the service API: `client.{{ undo.Tag }}Api.{{ undo.Undo.OperationID }}`
		operation := api.MethodByName(s.OperationID)
		if !operation.IsValid() {
			t.Fatalf("invalid method name %s", s.OperationID)
		}

		// Assemble method arguments
		in := make([]reflect.Value, operation.Type().NumIn())
		// first argument is always context.Context
		in[0] = reflect.ValueOf(GetCtx(ctx))
		for i := 1; i < operation.Type().NumIn(); i++ {
			in[i] = s.Parameters[i-1].Resolve(t, ctx, operation.Type().In(i))
		}

		undo, err := GetRequestsUndo(ctx, s.OperationID)
		if err != nil {
			t.Errorf("missing undo for %s: %v", s.OperationID, err)
		}

		result := operation.Call(in)
		if result[len(result)-1].Interface() != nil {
			t.Fatal(result[len(result)-1])
		}

		responseJSON, err := toJSON(result[0])

		if undo != nil {
			GetCleanup(ctx)[fmt.Sprintf("00-given-%s", s.Key)] = undo(result)
		}

		if s.Source != nil {
			responseJSON, err = lookup.LookupStringI(responseJSON, *s.Source)
			if err != nil {
				t.Error(err)
			}
			responseJSON = responseJSON.(reflect.Value).Interface()
		}

		GetData(ctx)[s.Key] = responseJSON
	}
	suite.AddStep(s.Step, given)
}

// GetRequestsUndo gets map with undo function for each request.
func GetRequestsUndo(ctx gobdd.Context, operationID string) (func([]reflect.Value) func(), error) {
	requestsUndo, err := ctx.Get(ctxRequestsUndoKey{})
	if err != nil {
		return nil, err
	}
	t := GetT(ctx)

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

	return func(response []reflect.Value) func() {
		return func() {
			responseStatusCode, err := GetResponseStatusCode(response)
			if err != nil {
				t.Fatalf("%v", err)
			}
			responseJSON, err := toJSON(response[0])
			if err != nil {
				t.Fatalf("%v", err)
			}

			// No object is created when it's a Bad Request
			if responseStatusCode >= 400 {
				return
			}

			// enable unstable operation
			configInterface := reflect.Indirect(undoClientInterface.(reflect.Value)).Addr().MethodByName("GetConfig").Call(nil)[0]
			if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{reflect.ValueOf(undo.Undo.OperationID)})[0].Bool() {
				configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
					reflect.ValueOf(undo.Undo.OperationID), reflect.ValueOf(true),
				})
			}

			// Assemble undo method as follows: Client(cctx).UsersApi.DisableUser(cctx, response.Data.GetId())
			numArgs := undoOperation.Type().NumIn()
			in := make([]reflect.Value, numArgs)
			// first argument is always context.Context
			in[0] = reflect.ValueOf(GetCtx(ctx))

			// Handle undoOperations that accept optional variadic arguments
			if undoOperation.Type().IsVariadic() {
				optionalParams := reflect.New(undoOperation.Type().In(numArgs - 1).Elem())
				in[numArgs-1] = optionalParams.Elem()
			}

			for i := 1; i < undoOperation.Type().NumIn() && i <= len(undo.Undo.Parameters); i++ {
				object, err := lookup.LookupStringI(responseJSON, undo.Undo.Parameters[i-1].Source)
				if err != nil {
					t.Fatalf("%v", err)
				}
				in[i] = object
			}

			result := undoOperation.Call(in)

			if result[len(result)-1].Interface() != nil {
				t.Logf("error in undo %v", result[len(result)-1])
			}
		}
	}, nil
}

// SetCtx sets Go context in BDD context.
func SetCtx(ctx gobdd.Context, value context.Context) {
	ctx.Set(ctxKey{}, value)
}

// SetFixtureData sets the fixture data in BDD context
func SetFixtureData(ctx gobdd.Context) {
	ct, _ := ctx.Get(gobdd.TestingTKey{})
	cctx := GetCtx(ctx)
	testName := strings.SplitN(strings.Split(ct.(*testing.T).Name(), "/")[2], "_", 2)[1]
	unique := WithUniqueSurrounding(cctx, testName)
	alnum := regexp.MustCompile(`[^A-Za-z0-9]+`)
	data := GetData(ctx)
	data["unique"] = unique
	data["unique_lower"] = strings.ToLower(unique)
	data["unique_alnum"] = string(alnum.ReplaceAll([]byte(unique), []byte("")))
	data["unique_lower_alnum"] = strings.ToLower(data["unique_alnum"].(string))
	data["now_ts"] = ClockFromContext(cctx).Now().Unix()
	data["now_iso"] = ClockFromContext(cctx).Now().Format(time.RFC3339)
	data["hour_later_ts"] = ClockFromContext(cctx).Now().Add(time.Hour).Unix()
	data["hour_later_iso"] = ClockFromContext(cctx).Now().Add(time.Hour).Format(time.RFC3339)
	data["hour_ago_ts"] = ClockFromContext(cctx).Now().Add(-time.Hour).Unix()
	data["hour_ago_iso"] = ClockFromContext(cctx).Now().Add(-time.Hour).Format(time.RFC3339)
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
		GetT(ctx).Fatalf("could not get data: %v", err)
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
		GetT(ctx).Fatalf("could not get request parameters: %v", err)
	}
	return c.(map[string]interface{})
}

// GetRequestArguments helps to build a request.
func GetRequestArguments(ctx gobdd.Context) []interface{} {
	c, err := ctx.Get(requestArgsKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get request arguments: %v", err)
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
		GetT(ctx).Fatalf("could not get cleanup: %v", err)
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
		t.Error(err)
	}
	r := c.(reflect.Value)
	f := r.MethodByName(name)
	if !f.IsValid() {
		t.Errorf("invalid method name %s on API", name)
	}

	ctx.Set(requestKey{}, f)
	ctx.Set(requestNameKey{}, name)
	ctx.Set(requestParamsKey{}, make(map[string]interface{}))
	ctx.Set(requestArgsKey{}, make([]interface{}, 0))
	ctx.Set(pathParamCountKey{}, 1)

}

// getRequestBuilder returns the reflect value of the current request in ctx
func getRequestBuilder(ctx gobdd.Context) (reflect.Value, []reflect.Value, error) {
	c, err := ctx.Get(requestKey{})
	if err != nil {
		return reflect.Value{}, []reflect.Value{}, fmt.Errorf("Missing requestKey{}")
	}
	f := c.(reflect.Value)

	in := make([]reflect.Value, f.Type().NumIn())

	// first argument is always context.Context
	in[0] = reflect.ValueOf(GetCtx(ctx))
	requestArgs := GetRequestArguments(ctx)
	requestParams := GetRequestParameters(ctx)
	numArgs := f.Type().NumIn()

	if f.Type().IsVariadic() {
		for i := 1; i < numArgs-1 && i <= len(requestArgs); i++ {
			object := requestArgs[i-1]
			in[i] = object.(reflect.Value)
		}

		optionalParams := reflect.New(f.Type().In(numArgs - 1).Elem())
		for k, v := range requestParams {
			paramName := toVarName(k)
			if method := optionalParams.MethodByName("With" + paramName); method.IsValid() {
				at := reflect.New(method.Type().In(0))
				switch v.(type) {
				case reflect.Value:
					at.Elem().Set(v.(reflect.Value))
				case interface{}:
					json.Unmarshal([]byte(v.(string)), at.Interface())
				}
				optionalParams = method.Call([]reflect.Value{at.Elem()})[0]
			} else if k == "body" {
				a := reflect.New(f.Type().In(numArgs - 2))
				json.Unmarshal([]byte(v.(string)), a.Interface())
				in[numArgs-2] = a.Elem()
			}
		}

		in[len(in)-1] = optionalParams.Elem()
	} else {
		// Set request args
		for i := 1; i <= len(requestArgs); i++ {
			object := requestArgs[i-1]
			in[i] = object.(reflect.Value)
		}
		// Append body to args if it exists
		if val, ok := requestParams["body"]; ok {
			a := reflect.New(f.Type().In(numArgs - 1))
			json.Unmarshal([]byte(val.(string)), a.Interface())
			in[numArgs-1] = a.Elem()
		}
	}

	return f, in, nil
}

func statusIs(t gobdd.StepTest, ctx gobdd.Context, expected int, text string) {
	// Execute() returns triples -> 2nd value is *http.Response -> get StatusCode
	resp := GetResponse(ctx)
	code, err := GetResponseStatusCode(resp)
	if err != nil {
		t.Fatal(err)
	}
	if expected != code {
		t.Fatalf("Expected %d got %d", expected, code)
	}
}

func addParameterFrom(t gobdd.StepTest, ctx gobdd.Context, name string, path string) {
	value, err := lookup.LookupStringI(GetData(ctx), path)
	if err != nil {
		t.Errorf("key %s in %+v: %v", path, GetData(ctx), err)
	}

	GetRequestParameters(ctx)[name] = value
	ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), value))
}

// This function adds path arguments to the requestArgs key. This shouldn't be used as a step method, but just a helper util
func addPathArgumentWithValue(t gobdd.StepTest, ctx gobdd.Context, param string, value string, request reflect.Value) {
	numArgs := request.Type().NumIn()
	in := make([]reflect.Value, numArgs)
	in[0] = reflect.ValueOf(GetCtx(ctx))
	// The order of the path arguments in the scenario definition
	// must match the order of the arguments in the function signature
	// Here we keep track of which numbered path argument we're setting
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)

	templatedValue := Templated(t, GetData(ctx), value)
	var varType reflect.Value

	if request.Type().IsVariadic() && pathCount.(int) > numArgs-2 {
		optionalParams := reflect.New(request.Type().In(numArgs - 1).Elem())
		field := optionalParams.Elem().FieldByName(toVarName(param))
		if field.IsValid() {
			varType = reflect.New(field.Type().Elem())
		}
	} else {
		varType = reflect.New(request.Type().In(pathCount.(int)))
	}

	if varType.IsValid() {
		json.Unmarshal([]byte(templatedValue), varType.Interface())
		GetRequestParameters(ctx)[param] = varType.Elem()
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), varType.Elem()))
	} else {
		GetRequestParameters(ctx)[param] = reflect.ValueOf(templatedValue)
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), reflect.ValueOf(templatedValue)))
	}
}

func addParameterWithValue(t gobdd.StepTest, ctx gobdd.Context, param string, value string) {
	// Get request builder for the current scenario
	request, _, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}

	addPathArgumentWithValue(t, ctx, param, value, request)
}

func toJSON(o reflect.Value) (interface{}, error) {
	var jsonResult interface{}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(o.Interface()); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf.Bytes(), &jsonResult); err != nil {
		return nil, err
	}
	return jsonResult, nil
}

func requestIsSent(t gobdd.StepTest, ctx gobdd.Context) {
	request, in, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}

	operationID, err := ctx.GetString(requestNameKey{})
	if err != nil {
		t.Errorf("could not get a request name: %v", err)
	}
	undo, err := GetRequestsUndo(ctx, operationID)
	if err != nil {
		t.Errorf("missing undo for %s: %v", operationID, err)
	}

	result := request.Call(in)
	ctx.Set(responseKey{}, result)

	// Report probable serialization errors
	if len(result) > 2 {
		resp := result[len(result)-2].Interface().(*http.Response)
		if resp == nil {
			return
		}
		code := resp.StatusCode
		err := result[len(result)-1].Interface()
		if code < 300 && err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Store the unmarshalled JSON in context
		responseJSON, err := toJSON(result[0])
		if err != nil {
			t.Errorf("Unable to decode response object to JSON: %v", err)
		}
		ctx.Set(jsonResponseKey{}, responseJSON)
	}

	if undo != nil {
		GetCleanup(ctx)["01-undo"] = undo(result)
	}
}

func body(t gobdd.StepTest, ctx gobdd.Context, body string) {
	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), body)
}

func bodyFromFile(t gobdd.StepTest, ctx gobdd.Context, bodyFile string) {
	body, err := os.ReadFile(fmt.Sprintf("./features/%s", bodyFile))
	if err != nil {
		t.Fatalf("Error reading file ./features/%s: %v", bodyFile, err)
	}

	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), string(body))
}

func stringToType(s string, t interface{}) (interface{}, error) {
	switch t.(type) {
	case int:
		return strconv.Atoi(s)
	case int64:
		return strconv.ParseInt(s, 10, 64)
	case float64:
		return strconv.ParseFloat(s, 64)
	case string:
		return strconv.Unquote(s)
	case bool:
		return strconv.ParseBool(s)
	case map[string]interface{}, []interface{}:
		var res map[string]interface{}
		if err := json.Unmarshal([]byte(s), &res); err != nil {
			return nil, fmt.Errorf("error converting %s to %T: %v", s, t, err)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unknown type '%T' to convert", t)
	}
}

func expectEqual(t gobdd.StepTest, ctx gobdd.Context, responsePath string, value string) {
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup response value %s in %+v: %v", responsePath, GetJSONResponse(ctx), err)
	}

	templatedValue := Templated(t, GetData(ctx), value)
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
	fixtureValue, err := lookup.LookupStringI(GetData(ctx), fixturePath)
	if err != nil {
		t.Fatalf("could not lookup fixture value %s in %+v: %v", fixturePath, GetData(ctx), err)
	}
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", SnakeToCamelCase(responsePath), err)
	}
	if !responseValue.IsValid() && !fixtureValue.IsValid() {
		// Lookup was successful but both values are nil
		return
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
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", responsePath, err)
	}
	cmp := is.Len(responseValue.Interface(), lengthInt)()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectFalse(t gobdd.StepTest, ctx gobdd.Context, responsePath string) {
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), responsePath)
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
		`body from file "(.*)"`:                                  bodyFromFile,
		`body with value (.*)`:                                   body,
		`the response "([^"]+)" is equal to (.*)`:                expectEqual,
		`the response "([^"]+)" has the same value as "([^"]+)"`: expectEqualValue,
		`the response "([^"]+)" has length ([0-9]+)`:             expectLengthEqual,
		`the response "([^"]+)" is false`:                        expectFalse,
	}
	for expr, step := range steps {
		s.AddStep(expr, step)
	}
}
