/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
	testsV1 "github.com/DataDog/datadog-api-client-go/v2/tests/api/datadogV1"
	testsV2 "github.com/DataDog/datadog-api-client-go/v2/tests/api/datadogV2"

	"github.com/go-bdd/gobdd"
)

// UndoAction describes undo action.
type UndoAction struct {
	Tag  string `json:"tag"`
	Undo *struct {
		Tag         *string `json:"tag"`
		Type        string  `json:"type"`
		OperationID string  `json:"operationId"`
		Parameters  []struct {
			Name     string  `json:"name"`
			Source   *string `json:"source"`
			Template *string `json:"template"`
			Origin   *string `json:"origin"`
		} `json:"parameters"`
	} `json:"undo"`
}

type operationParameter struct {
	Name   string  `json:"name"`
	Source *string `json:"source"`
	Value  *string `json:"value"`
	Origin *string `json:"origin"`
}

func (p operationParameter) Resolve(t gobdd.StepTest, ctx gobdd.Context, tp reflect.Type) reflect.Value {
	if p.Value != nil {
		if tp.Kind() == reflect.Slice {
			field, ok := tp.Elem().FieldByName(SnakeToCamelCase(p.Name))
			if ok && field.Type == reflect.TypeOf((*io.Reader)(nil)) {
				version := GetVersion(ctx)
				var basePath string
				tpl := Templated(t, GetData(ctx), *p.Value)
				if err := datadog.Unmarshal([]byte(tpl), &basePath); err != nil {
					t.Fatalf("can't unmarshal parameter value for %s: %v", p.Name, err)
				}
				filepath := fmt.Sprintf("./features/%s/%s", version, basePath)
				fp, err := os.Open(filepath)
				if err != nil {
					t.Error(err)
				}
				reader := io.Reader(fp)
				v := reflect.New(tp.Elem())
				v.Elem().FieldByName(SnakeToCamelCase(p.Name)).Set(reflect.ValueOf(&reader))
				return v.Elem()
			}
		}
		tpl := Templated(t, GetData(ctx), *p.Value)
		v := reflect.New(tp)
		err := datadog.Unmarshal([]byte(tpl), v.Interface())
		if err != nil {
			t.Fatalf("can't unmarshal parameter value for %s: %v", p.Name, err)
		}
		return v.Elem()
	}
	v, _ := tests.LookupStringI(GetData(ctx), *p.Source)
	return v.Convert(tp)
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

// LoadRequestsUndo load undo configuration.
func LoadRequestsUndo(file string) (map[string]UndoAction, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)

	var value map[string]UndoAction
	datadog.Unmarshal(byteValue, &value)
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
	byteValue, _ := io.ReadAll(f)
	err = datadog.Unmarshal(byteValue, &value)
	return value, err
}

// SetRequestsUndo sets map with undo function for each request.
func SetRequestsUndo(ctx gobdd.Context, value map[string]map[string]UndoAction) {
	ctx.Set(ctxRequestsUndoKey{}, value)
}

type clientKeys struct{}
type ctxRequestsUndoKey struct{}

// ConfigureClients sets up the Datadog API client for testing
func ConfigureClients(ctx context.Context, bddCTX gobdd.Context) (context.Context, func()) {
	t := GetT(bddCTX)
	debug := os.Getenv("DEBUG") == "true"
	config := datadog.NewConfiguration()
	config.Debug = debug
	config.RetryConfiguration.EnableRetry = true
	c := datadog.NewAPIClient(config)

	ctx = context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{},
	)

	ctx, err := tests.WithClock(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup clock: %v", err)
	}

	r, err := tests.Recorder(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup recorder: %v", err)
	}
	httpClient := http.Client{Transport: tests.WrapRoundTripper(r)}
	c.GetConfig().HTTPClient = &httpClient

	client := reflect.ValueOf(c)
	bddCTX.Set(clientKeys{}, client)
	return ctx, func() {
		r.Stop()
	}
}

func getAuthenticatedContext(ctx gobdd.Context) context.Context {
	return testsV1.WithTestAuth(
		testsV2.WithTestAuth(
			testsV1.NewDefaultContext(
				testsV2.NewDefaultContext(GetCtx(ctx)),
			),
		),
	)
}

// RegisterSuite adds step implementation.
func (s GivenStep) RegisterSuite(suite *gobdd.Suite, version string) {
	given := func(t gobdd.StepTest, ctx gobdd.Context) {
		// get an instance of Client
		clientInterface := GetClient(ctx)

		// Enable unstable operation
		configInterface := clientInterface.MethodByName("GetConfig").Call(nil)[0]
		unstableOperationId := reflect.ValueOf(fmt.Sprintf("%s.%s", version, s.OperationID))
		if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{unstableOperationId})[0].Bool() {
			configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
				unstableOperationId, reflect.ValueOf(true),
			})
		}

		// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
		tag := strings.Replace(s.Tag, " ", "", -1) + "Api"
		api := GetApiByVersionAndName(ctx, version, tag).Call([]reflect.Value{clientInterface})[0]
		if !api.IsValid() {
			t.Fatalf("invalid API name %s", tag)
		}

		// get method from the service API: `client.{{ undo.Tag }}Api.{{ undo.Undo.OperationID }}`
		operation := api.MethodByName(s.OperationID)
		if !operation.IsValid() {
			t.Fatalf("invalid method name %s", s.OperationID)
		}

		// Assemble method arguments
		numArgs := operation.Type().NumIn()
		in := make([]reflect.Value, numArgs)
		// first argument is always context.Context
		in[0] = reflect.ValueOf(getAuthenticatedContext(ctx))

		// Handle operations that accept optional variadic arguments
		if operation.Type().IsVariadic() {
			optionalParams := reflect.New(operation.Type().In(numArgs - 1).Elem())
			in[numArgs-1] = optionalParams.Elem()
		}

		for i := 1; i < numArgs && i <= len(s.Parameters); i++ {
			in[i] = s.Parameters[i-1].Resolve(t, ctx, operation.Type().In(i))
		}

		undo, err := GetRequestsUndo(ctx, version, s.OperationID)
		if err != nil {
			t.Errorf("missing undo for %s: %v", s.OperationID, err)
		}

		result := operation.Call(in)
		if result[len(result)-1].Interface() != nil {
			t.Fatal(result[len(result)-1])
		}

		responseJSON, _ := toJSON(result[0])

		if undo != nil {
			GetCleanup(ctx)[fmt.Sprintf("00-given-%s", s.Key)] = undo(result)
		}

		if s.Source != nil {
			responseJSON, err = tests.LookupStringI(responseJSON, *s.Source)
			if err != nil {
				t.Error(err)
			}
			responseJSON = responseJSON.(reflect.Value).Interface()
		}

		GetData(ctx)[s.Key] = responseJSON

		sleepAfterRequest()
	}
	suite.AddStep(s.Step, given)
}

// GetRequestsUndo gets map with undo function for each request.
func GetRequestsUndo(ctx gobdd.Context, version string, operationID string) (func([]reflect.Value) func(), error) {
	versionsUndo, err := ctx.Get(ctxRequestsUndoKey{})
	if err != nil {
		return nil, err
	}
	t := GetT(ctx)

	// find undo definitions for a given version
	requestsUndo, ok := versionsUndo.(map[string]map[string]UndoAction)[version]
	if !ok {
		return nil, fmt.Errorf("undefined undo for version: %s", version)
	}
	// find undo definition
	undo, ok := requestsUndo[operationID]
	if !ok {
		return nil, fmt.Errorf("undefined undo for operation: %s", operationID)
	}

	if undo.Undo.Type != "unsafe" {
		return nil, nil
	}

	// get an instance of Client
	undoClientInterface := GetClient(ctx)

	// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
	tag := strings.Replace(undo.Tag, " ", "", -1) + "Api"
	if undo.Undo.Tag != nil {
		tag = strings.Replace(*undo.Undo.Tag, " ", "", -1) + "Api"
	}
	undoAPI := GetApiByVersionAndName(ctx, version, tag).Call([]reflect.Value{undoClientInterface})[0]
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

			// No object is created when it's a Bad Request
			if responseStatusCode >= 400 {
				return
			}

			responseJSON, err := toJSON(response[0])
			if err != nil {
				t.Fatalf("%v", err)
			}

			// enable unstable operation
			configInterface := undoClientInterface.MethodByName("GetConfig").Call(nil)[0]
			undoOperationID := reflect.ValueOf(fmt.Sprintf("%s.%s", version, undo.Undo.OperationID))
			if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{undoOperationID})[0].Bool() {
				configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
					undoOperationID, reflect.ValueOf(true),
				})
			}

			// Assemble undo method as follows: Client(cctx).UsersApi.DisableUser(cctx, response.Data.GetId())
			numArgs := undoOperation.Type().NumIn()
			in := make([]reflect.Value, numArgs)
			// first argument is always context.Context
			in[0] = reflect.ValueOf(getAuthenticatedContext(ctx))

			// Handle undoOperations that accept optional variadic arguments
			if undoOperation.Type().IsVariadic() {
				optionalParams := reflect.New(undoOperation.Type().In(numArgs - 1).Elem())
				in[numArgs-1] = optionalParams.Elem()
			}

			for i := 1; i < numArgs && i <= len(undo.Undo.Parameters); i++ {
				var sourceJSONData interface{}
				if undo.Undo.Parameters[i-1].Origin == nil {
					sourceJSONData = responseJSON
				} else if *undo.Undo.Parameters[i-1].Origin == "request" {
					requestParams := GetRequestParameters(ctx)
					requestJSON := make(map[string]interface{})
					if requestParams["body"] != nil {
						requestData := requestParams["body"].(string)
						if err := datadog.Unmarshal([]byte(requestData), &requestJSON); err != nil {
							t.Fatalf("Error unmarshalling request: %v", err)
							return
						}
					}
					sourceJSONData = requestJSON
				} else if *undo.Undo.Parameters[i-1].Origin == "path" {
					// Extract from stored path parameters
					pathParams := GetPathParameters(ctx)
					if undo.Undo.Parameters[i-1].Source != nil {
						paramName := *undo.Undo.Parameters[i-1].Source
						if val, ok := pathParams[paramName]; ok {
							// Convert the stored value to the expected type
							object := reflect.New(undoOperation.Type().In(i))
							sourceJSON, err := datadog.Marshal(val)
							if err != nil {
								t.Fatalf("Error marshalling path parameter '%s': %v", paramName, err)
								return
							}
							err = datadog.Unmarshal(sourceJSON, object.Interface())
							if err != nil {
								t.Fatalf("Error unmarshalling path parameter '%s': %v", paramName, err)
								return
							}
							in[i] = object.Elem()
							continue // Skip template/source processing for path parameters
						} else {
							t.Fatalf("Path parameter '%s' not found", paramName)
							return
						}
					} else {
						t.Fatalf("Path origin requires 'source' field")
						return
					}
				} else if *undo.Undo.Parameters[i-1].Origin == "response" {
					sourceJSONData = responseJSON
				}
				if undo.Undo.Parameters[i-1].Template != nil {
					data := Templated(t, sourceJSONData.(map[string]interface{}), *undo.Undo.Parameters[i-1].Template)
					object := reflect.New(undoOperation.Type().In(i))
					err := datadog.Unmarshal([]byte(data), object.Interface())
					if err != nil {
						t.Fatalf("%v", err)
					}
					in[i] = object.Elem()
				} else if undo.Undo.Parameters[i-1].Source != nil {
					source, err := tests.LookupStringI(sourceJSONData, *undo.Undo.Parameters[i-1].Source)
					if err != nil {
						t.Fatalf("%v", err)
					}
					sourceJSON, err := datadog.Marshal(source.Interface())
					if err != nil {
						t.Fatalf("%v", err)
					}
					object := reflect.New(undoOperation.Type().In(i))
					err = datadog.Unmarshal(sourceJSON, object.Interface())
					if err != nil {
						t.Fatalf("%v", err)
					}
					in[i] = object.Elem()
				}
			}

			result := undoOperation.Call(in)

			if result[len(result)-1].Interface() != nil {
				t.Logf("error in undo %v", result[len(result)-1])
			}
		}
	}, nil
}

func sleepAfterRequest() {
	record, ok := os.LookupEnv("RECORD")
	if !ok {
		record = "false"
	}

	sleep := 0
	if sleepAfterRequestDuration, ok := os.LookupEnv("SLEEP_AFTER_REQUEST"); ok {
		sleep, _ = strconv.Atoi(sleepAfterRequestDuration)
	}

	if record != "false" && sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}
