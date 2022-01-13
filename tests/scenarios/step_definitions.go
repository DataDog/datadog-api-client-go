/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"

	v1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	v2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
	is "gotest.tools/assert/cmp"
)

func setAPIKey(ctx gobdd.Context, name string, value string) {
	version := GetVersion(ctx)
	if version == "v1" {
		if keys, ok := GetCtx(ctx).Value(v1.ContextAPIKeys).(map[string]v1.APIKey); ok {
			keys[name] = v1.APIKey{
				Key: value,
			}
		} else {
			GetT(ctx).Fatalf("could not set %s key", name)
		}
	} else if version == "v2" {
		if keys, ok := GetCtx(ctx).Value(v2.ContextAPIKeys).(map[string]v2.APIKey); ok {
			keys[name] = v2.APIKey{
				Key: value,
			}
		} else {
			GetT(ctx).Fatalf("could not set %s key", name)
		}
	} else {
		GetT(ctx).Fatalf("unknown version %s", version)
	}
}

func aValidAPIKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_API_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_API_KEY must be set")
	}
	setAPIKey(ctx, "apiKeyAuth", key)
}

func aValidAppKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_APP_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_APP_KEY must be set")
	}
	setAPIKey(ctx, "appKeyAuth", key)
}

// anInstanceOf sets API callable to apiKey{}
func anInstanceOf(t gobdd.StepTest, ctx gobdd.Context, name string) {
	ct := GetClient(ctx)
	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		t.Fatalf("invalid API name %s", name)
	}
	SetAPI(ctx, f)
}

// enableOperations sets unstable operations specific in this clause to enabled
func enableOperations(t gobdd.StepTest, ctx gobdd.Context, name string) {
	ct := GetClient(ctx)
	// client.GetConfig().SetUnstableOperationEnabled(name, true)
	config := ct.MethodByName("GetConfig").Call([]reflect.Value{})
	enable := config[0].MethodByName("SetUnstableOperationEnabled")
	enable.Call([]reflect.Value{reflect.ValueOf(name), reflect.ValueOf(true)})
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
	value, err := lookup.LookupStringI(GetData(ctx), CamelToSnakeCase(path))
	if err != nil {
		t.Errorf("key %s in %+v: %v", path, GetData(ctx), err)
	}
	request, _, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}
	numArgs := request.Type().NumIn()
	in := make([]reflect.Value, numArgs)
	in[0] = reflect.ValueOf(GetCtx(ctx))
	// The order of the path arguments in the scenario definition
	// must match the order of the arguments in the function signature
	// Here we keep track of which numbered path argument we're setting
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
	var varType reflect.Value

	if request.Type().IsVariadic() && pathCount.(int) > numArgs-2 {
		optionalParams := reflect.New(request.Type().In(numArgs - 1).Elem())
		field := reflect.Indirect(optionalParams.Elem()).FieldByName(ToVarName(name))
		if field.IsValid() {
			varType = reflect.New(field.Type().Elem())
		}
	} else {
		varType = reflect.New(request.Type().In(pathCount.(int)))
	}

	if varType.IsValid() {
		data, err := json.Marshal(value.Interface())
		if err != nil {
			t.Error(err)
		}
		json.Unmarshal(data, varType.Interface())
		GetRequestParameters(ctx)[name] = varType.Elem()
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), varType.Elem()))
	} else {
		GetRequestParameters(ctx)[name] = value
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), value))
	}
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
		field := reflect.Indirect(optionalParams.Elem()).FieldByName(ToVarName(param))
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

func requestIsSent(t gobdd.StepTest, ctx gobdd.Context) {
	request, in, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}

	operationID, err := ctx.GetString(requestNameKey{})
	if err != nil {
		t.Errorf("could not get a request name: %v", err)
	}

	undo, err := GetRequestsUndo(ctx, GetVersion(ctx), operationID)
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
		var responseJSON interface{}
		if err != nil {
			version := GetVersion(ctx)
			if version == "v1" {
				err := err.(v1.GenericOpenAPIError)
				if newErr := json.Unmarshal(err.Body(), &responseJSON); newErr != nil {
					t.Errorf("Unable to decode response object to JSON: %v", newErr)
				}
			} else {
				err := err.(v2.GenericOpenAPIError)
				if newErr := json.Unmarshal(err.Body(), &responseJSON); newErr != nil {
					t.Errorf("Unable to decode response object to JSON: %v", newErr)
				}
			}
		} else {
			// Store the unmarshalled JSON in context
			responseJSON, err = toJSON(result[0])
			if err != nil {
				t.Errorf("Unable to decode response object to JSON: %v", err)
			}
		}
		ctx.Set(jsonResponseKey{}, responseJSON)
	}

	if undo != nil {
		GetCleanup(ctx)["01-undo"] = undo(result)
	}
}

func body(t gobdd.StepTest, ctx gobdd.Context, body string) {
	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), body)
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
}

func bodyFromFile(t gobdd.StepTest, ctx gobdd.Context, bodyFile string) {
	version := GetVersion(ctx)
	body, err := ioutil.ReadFile(fmt.Sprintf("./features/%s/%s", version, bodyFile))
	if err != nil {
		t.Fatalf("Error reading file ./features/%s: %v", bodyFile, err)
	}

	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), string(body))
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
}

func expectEqual(t gobdd.StepTest, ctx gobdd.Context, responsePath string, value string) {
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), CamelToSnakeCase(responsePath))
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
	fixtureValue, err := lookup.LookupStringI(GetData(ctx), CamelToSnakeCase(fixturePath))
	if err != nil {
		t.Fatalf("could not lookup fixture value %s in %+v: %v", fixturePath, GetData(ctx), err)
	}
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), CamelToSnakeCase(responsePath))
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
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), CamelToSnakeCase(responsePath))
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", responsePath, err)
	}
	cmp := is.Len(responseValue.Interface(), lengthInt)()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectFalse(t gobdd.StepTest, ctx gobdd.Context, responsePath string) {
	responseValue, err := lookup.LookupStringI(GetJSONResponse(ctx), CamelToSnakeCase(responsePath))
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
		`a valid "apiKeyAuth" key in the system`:                 aValidAPIKeyAuth,
		`a valid "appKeyAuth" key in the system`:                 aValidAppKeyAuth,
		`an instance of "([^"]+)" API`:                           anInstanceOf,
		`operation "([^"]+)" enabled`:                            enableOperations,
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
