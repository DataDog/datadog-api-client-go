/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
	"github.com/go-bdd/gobdd"
	is "gotest.tools/assert/cmp"
)

// Copy from the client code
var jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)

func setAPIKey(ctx gobdd.Context, name string, value string) {
	if keys, ok := GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys[name] = datadog.APIKey{
			Key: value,
		}
	} else {
		GetT(ctx).Fatalf("could not set %s key", name)
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
	version := GetVersion(ctx)

	api := GetApiByVersionAndName(ctx, version, name+"Api")
	f := api.Call([]reflect.Value{ct})[0]

	SetAPI(ctx, f)
}

// enableOperations sets unstable operations specific in this clause to enabled
func enableOperations(t gobdd.StepTest, ctx gobdd.Context, name string) {
	ct := GetClient(ctx)
	version := GetVersion(ctx)
	// client.GetConfig().SetUnstableOperationEnabled(name, true)
	config := ct.MethodByName("GetConfig").Call([]reflect.Value{})
	enable := config[0].MethodByName("SetUnstableOperationEnabled")
	enable.Call([]reflect.Value{reflect.ValueOf(fmt.Sprintf("%s.%s", version, name)), reflect.ValueOf(true)})
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
	ClearPathParameters(ctx)

}
func statusIs(t gobdd.StepTest, ctx gobdd.Context, expected int, text string) {
	// Execute() returns triples -> 2nd value is *http.Response -> get StatusCode
	resp := GetResponse(ctx)
	code, err := GetResponseStatusCode(resp)
	if err != nil {
		t.Logf("could not extract status code: %v", err)
		return
	}
	if expected != code {
		t.Fatalf("Expected %d got %d", expected, code)
	}
}

func addParameterFrom(t gobdd.StepTest, ctx gobdd.Context, name string, path string) {
	value, err := tests.LookupStringI(GetData(ctx), CamelToSnakeCase(path))
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
		data, err := datadog.Marshal(value.Interface())
		if err != nil {
			t.Error(err)
		}
		datadog.Unmarshal(data, varType.Interface())
		GetRequestParameters(ctx)[name] = varType.Elem()
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), varType.Elem()))
		// Store path parameter for undo operations
		SetPathParameter(ctx, name, varType.Elem().Interface())
	} else {
		GetRequestParameters(ctx)[name] = value
		ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), value))
		// Store path parameter for undo operations
		SetPathParameter(ctx, name, value.Interface())
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

	// Store path parameter by name for undo operations
	SetPathParameter(ctx, param, templatedValue)

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
		switch varType.Interface().(type) {
		case *io.Reader:
			version := GetVersion(ctx)
			var basePath string
			datadog.Unmarshal([]byte(templatedValue), &basePath)
			filepath := fmt.Sprintf("./features/%s/%s", version, basePath)
			fp, err := os.Open(filepath)
			if err != nil {
				t.Error(err)
			}
			GetRequestParameters(ctx)[param] = reflect.ValueOf(fp)
			ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), reflect.ValueOf(fp)))
		default:
			datadog.Unmarshal([]byte(templatedValue), varType.Interface())
			GetRequestParameters(ctx)[param] = varType.Elem()
			ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), varType.Elem()))
		}
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
	if len(result) > 1 {
		resp := result[len(result)-2].Interface().(*http.Response)
		if resp == nil {
			t.Errorf("unexpected error: %v", result[len(result)-1])
			return
		}
		code := resp.StatusCode

		contentType := resp.Header.Get("Content-Type")
		if code < 300 && !jsonCheck.MatchString(contentType) {
			// We don't care about non-JSON responses
			t.Logf("Response is not JSON: %s", contentType)
			ctx.Set(jsonResponseKey{}, result[0])
			return
		}

		err := result[len(result)-1].Interface()
		if code < 300 && err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		var responseJSON interface{}
		if err != nil {
			err := err.(datadog.GenericOpenAPIError)
			if newErr := datadog.Unmarshal(err.Body(), &responseJSON); newErr != nil {
				responseJSON = string(err.Body())
			}
		} else if len(result) == 3 {
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

func requestWithPaginationIsSent(t gobdd.StepTest, ctx gobdd.Context) {
	// use WithPagination method
	newWithPagination, _ := ctx.GetString(requestNameKey{})
	newWithPagination = newWithPagination + "WithPagination"

	c, err := ctx.Get(apiKey{})
	if err != nil {
		t.Error(err)
	}
	r := c.(reflect.Value)
	f := r.MethodByName(newWithPagination)
	if !f.IsValid() {
		t.Errorf("invalid method name %s on API", newWithPagination)
	}

	ctx.Set(requestKey{}, f)
	// call the WithPagination method

	request, in, err := getRequestBuilder(ctx)
	if err != nil {
		t.Error(err)
	}

	result := request.Call(in)
	ctx.Set(responseKey{}, result)

	resp := result[0].Interface()
	if resp == nil {
		return
	}

	var responseJSON interface{}
	if err != nil {
		err := err.(datadog.GenericOpenAPIError)
		if newErr := datadog.Unmarshal(err.Body(), &responseJSON); newErr != nil {
			responseJSON = string(err.Body())
		}
	} else {
		// Store the unmarshalled JSON in context
		items, err := readChannel(resp)
		if err != nil {
			t.Errorf("Unable to read response channel: %v", err)
		}
		buf := &bytes.Buffer{}
		if err := datadog.NewEncoder(buf).Encode(items); err != nil {
			t.Errorf("Unable to decode response object to JSON: %v", err)
		}
		if err := datadog.Unmarshal(buf.Bytes(), &responseJSON); err != nil {
			t.Errorf("Unable to decode response object to JSON: %v", err)
		}

		if err != nil {
			t.Errorf("Unable to decode response object to JSON: %v", err)
		}
	}
	ctx.Set(jsonResponseKey{}, responseJSON)
}

func body(t gobdd.StepTest, ctx gobdd.Context, body string) {
	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), body)
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
}

func bodyFromFile(t gobdd.StepTest, ctx gobdd.Context, bodyFile string) {
	version := GetVersion(ctx)
	body, err := os.ReadFile(fmt.Sprintf("./features/%s/%s", version, bodyFile))
	if err != nil {
		t.Fatalf("Error reading file ./features/%s: %v", bodyFile, err)
	}

	GetRequestParameters(ctx)["body"] = Templated(t, GetData(ctx), string(body))
	pathCount, _ := ctx.Get(pathParamCountKey{})
	ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
}

func expectEqual(t gobdd.StepTest, ctx gobdd.Context, responsePath string, value string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup response value %s in %+v: %v", CamelToSnakeCase(responsePath), GetJSONResponse(ctx), err)
	}

	templatedValue := Templated(t, GetData(ctx), value)

	if (responseValue.Kind() == reflect.Interface && responseValue.IsNil()) || !responseValue.IsValid() {
		cmp := is.DeepEqual(templatedValue, "null")()
		if !cmp.Success() {
			t.Errorf("%v", cmp)
		}
	} else {
		testValue, err := stringToType(templatedValue, responseValue.Interface())
		if err != nil {
			t.Errorf("%v", err)
		}

		cmp := is.DeepEqual(testValue, responseValue.Interface())()
		if !cmp.Success() {
			t.Errorf("%v", cmp)
		}
	}
}

func expectEqualValue(t gobdd.StepTest, ctx gobdd.Context, responsePath string, fixturePath string) {
	fixtureValue, err := tests.LookupStringI(GetData(ctx), CamelToSnakeCase(fixturePath))
	if err != nil {
		t.Fatalf("could not lookup fixture value %s in %+v: %v", fixturePath, GetData(ctx), err)
	}
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
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
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Fatalf("could not lookup response value %s: %v", responsePath, err)
	}
	cmp := is.Len(responseValue.Interface(), lengthInt)()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectNumberOfItems(t gobdd.StepTest, ctx gobdd.Context, responseLength string) {
	lengthInt, err := strconv.Atoi(responseLength)
	if err != nil {
		t.Fatalf("assertion length value is not a number %s: %v", responseLength, err)
	}
	responseValue := GetJSONResponse(ctx)
	cmp := is.Len(responseValue, lengthInt)()
	if !cmp.Success() {
		t.Errorf("%v", cmp)
	}
}

func expectFalse(t gobdd.StepTest, ctx gobdd.Context, responsePath string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup value: %v", err)
	}
	if responseValue.Bool() {
		t.Errorf("%v should be false", responseValue)
	}
}

func expectResponseHasField(t gobdd.StepTest, ctx gobdd.Context, responsePath string, field string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup response value %s in %+v: %v", CamelToSnakeCase(responsePath), GetJSONResponse(ctx), err)
	}

	responseMap, ok := responseValue.Interface().(map[string]interface{})
	if !ok {
		t.Fatalf("%v is not a map. Cannot lookup field %s", responseValue, field)
	}

	_, ok = responseMap[field]
	if !ok {
		t.Errorf("%v should contain field %s", responseValue, field)
	}
}

func expectResponseDoesNotHaveField(t gobdd.StepTest, ctx gobdd.Context, responsePath string, field string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup response value %s in %+v: %v", CamelToSnakeCase(responsePath), GetJSONResponse(ctx), err)
	}

	responseMap, ok := responseValue.Interface().(map[string]interface{})
	if !ok {
		t.Fatalf("%v is not a map. Cannot lookup field %s", responseValue, field)
	}

	_, ok = responseMap[field]
	if ok {
		t.Errorf("%v should not contain field %s", responseValue, field)
	}
}

func expectArrayContainsObject(t gobdd.StepTest, ctx gobdd.Context, responsePath string, keyPath string, value string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	templatedValue := Templated(t, GetData(ctx), value)
	if err != nil {
		t.Errorf("could not lookup value: %v", err)
	}
	if responseMap, ok := responseValue.Interface().([]interface{}); ok {
		for _, responseItem := range responseMap {
			responseItemValue, err := tests.LookupStringI(responseItem, keyPath)
			if err == tests.ErrNotFound {
				continue
			} else if err != nil {
				t.Errorf("could not lookup key value: %v", err)
			}
			testValue, err := stringToType(templatedValue, responseItemValue.Interface())
			if err != nil {
				t.Errorf("%v", err)
			}
			if is.DeepEqual(responseItemValue.Interface(), testValue)().Success() {
				return
			}
		}
	}
	t.Errorf("could not find value: %v", value)
}

func expectArrayContainsValue(t gobdd.StepTest, ctx gobdd.Context, responsePath string, value string) {
	responseValue, err := tests.LookupStringI(GetJSONResponse(ctx), responsePath)
	if err != nil {
		t.Errorf("could not lookup value: %v", err)
	}
	templatedValue := Templated(t, GetData(ctx), value)
	if responseList, ok := responseValue.Interface().([]interface{}); ok {
		for _, responseItem := range responseList {
			testValue, err := stringToType(templatedValue, responseItem)
			if err != nil {
				t.Errorf("%v", err)
			}
			if is.DeepEqual(responseItem, testValue)().Success() {
				return
			}
		}
	}
	t.Errorf("could not find value: %v", templatedValue)
}

// ConfigureSteps on given suite.
func ConfigureSteps(s *gobdd.Suite) {
	steps := map[string]interface{}{
		`a valid "apiKeyAuth" key in the system`:                               aValidAPIKeyAuth,
		`a valid "appKeyAuth" key in the system`:                               aValidAppKeyAuth,
		`an instance of "([^"]+)" API`:                                         anInstanceOf,
		`operation "([^"]+)" enabled`:                                          enableOperations,
		`new "([^"]+)" request`:                                                newRequest,
		`request contains "([^"]+)" parameter from "([^"]+)"`:                  addParameterFrom,
		`request contains "([^"]+)" parameter with value (.+)`:                 addParameterWithValue,
		`the request is sent`:                                                  requestIsSent,
		`the request with pagination is sent`:                                  requestWithPaginationIsSent,
		`the response status is (\d+) (.*)`:                                    statusIs,
		`body from file "(.*)"`:                                                bodyFromFile,
		`body with value (.*)`:                                                 body,
		`the response "([^"]+)" is equal to (.*)`:                              expectEqual,
		`the response "([^"]+)" has the same value as "([^"]+)"`:               expectEqualValue,
		`the response "([^"]+)" has length ([0-9]+)`:                           expectLengthEqual,
		`the response has ([0-9]+) items`:                                      expectNumberOfItems,
		`the response "([^"]+)" is false`:                                      expectFalse,
		`the response "([^"]+)" has field "([^"]+)"`:                           expectResponseHasField,
		`the response "([^"]+)" does not have field "([^"]+)"`:                 expectResponseDoesNotHaveField,
		`the response "([^"]+)" has item with field "([^"]+)" with value (.*)`: expectArrayContainsObject,
		`the response "([^"]+)" array contains value (.*)`:                     expectArrayContainsValue,
	}
	for expr, step := range steps {
		s.AddStep(expr, step)
	}
}
