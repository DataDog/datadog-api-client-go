/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	v1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	v2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	testsV1 "github.com/DataDog/datadog-api-client-go/tests/api/v1/datadog"
	testsV2 "github.com/DataDog/datadog-api-client-go/tests/api/v2/datadog"

	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
)

// UndoAction describes undo action.
type UndoAction struct {
	Tag  string `json:"tag"`
	Undo *struct {
		Type        string `json:"type"`
		OperationID string `json:"operationId"`
		Parameters  []struct {
			Name     string `json:"name"`
			Source   string `json:"source"`
			Template string `json:"template"`
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
func SetRequestsUndo(ctx gobdd.Context, value map[string]map[string]UndoAction) {
	ctx.Set(ctxRequestsUndoKey{}, value)
}

type clientKeys struct{}
type ctxRequestsUndoKey struct{}

func ConfigureClients(ctx gobdd.Context, cctx context.Context) (context.Context, func()) {
	t := GetT(ctx)
	c1 := v1.NewAPIClient(v1.NewConfiguration())
	c2 := v2.NewAPIClient(v2.NewConfiguration())

	cctx = context.WithValue(
		cctx,
		v1.ContextAPIKeys,
		map[string]v1.APIKey{},
	)
	cctx = context.WithValue(
		cctx,
		v2.ContextAPIKeys,
		map[string]v2.APIKey{},
	)

	cctx, err := tests.WithClock(cctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup clock: %v", err)
	}

	r, err := tests.Recorder(cctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup recorder: %v", err)
	}
	httpClient := http.Client{Transport: tests.WrapRoundTripper(r)}
	c1.GetConfig().HTTPClient = &httpClient
	c2.GetConfig().HTTPClient = &httpClient

	clients := map[string]reflect.Value{
		"v1": reflect.ValueOf(c1),
		"v2": reflect.ValueOf(c2),
	}
	ctx.Set(clientKeys{}, clients)
	return cctx, func() {
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

func GetClientVersion(ctx gobdd.Context, version string) (reflect.Value, error) {
	clientInterfaces, err := ctx.Get(clientKeys{})
	if err != nil {
		return reflect.Value{}, err
	}
	if clientInterface, ok := clientInterfaces.(map[string]reflect.Value); ok {
		if client, ok := clientInterface[version]; ok {
			return client, nil
		}
	}
	return reflect.Value{}, fmt.Errorf("could not get client")
}

// RegisterSuite adds step implementation.
func (s GivenStep) RegisterSuite(suite *gobdd.Suite, version string) {
	given := func(t gobdd.StepTest, ctx gobdd.Context) {
		// get an instance of Client
		clientInterface, err := GetClientVersion(ctx, version)
		if err != nil {
			t.Fatalf("could not find client: %v", err)
		}

		// Enable unstable operation
		configInterface := clientInterface.MethodByName("GetConfig").Call(nil)[0]
		if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{reflect.ValueOf(s.OperationID)})[0].Bool() {
			configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
				reflect.ValueOf(s.OperationID), reflect.ValueOf(true),
			})
		}

		// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
		tag := strings.Replace(s.Tag, " ", "", -1) + "Api"
		api := reflect.Indirect(clientInterface).FieldByName(tag)
		if !api.IsValid() {
			t.Fatalf("invalid API name %s", tag)
		}

		// get method from the service API: `client.{{ undo.Tag }}Api.{{ undo.Undo.OperationID }}`
		operation := api.MethodByName(s.OperationID)
		if !operation.IsValid() {
			t.Fatalf("invalid method name %s", s.OperationID)
		}

		// Assemble method arguments
		in := make([]reflect.Value, operation.Type().NumIn())
		// first argument is always context.Context
		in[0] = reflect.ValueOf(getAuthenticatedContext(ctx))
		for i := 1; i < operation.Type().NumIn(); i++ {
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
	undoClientInterface, err := GetClientVersion(ctx, version)
	if err != nil {
		return nil, err
	}

	// find API service based on undo tag value: `client.{{ undo.Tag }}Api`
	tag := strings.Replace(undo.Tag, " ", "", -1) + "Api"
	undoAPI := reflect.Indirect(undoClientInterface).FieldByName(tag)
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
			configInterface := undoClientInterface.MethodByName("GetConfig").Call(nil)[0]
			if configInterface.MethodByName("IsUnstableOperation").Call([]reflect.Value{reflect.ValueOf(undo.Undo.OperationID)})[0].Bool() {
				configInterface.MethodByName("SetUnstableOperationEnabled").Call([]reflect.Value{
					reflect.ValueOf(undo.Undo.OperationID), reflect.ValueOf(true),
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

			for i := 1; i < undoOperation.Type().NumIn() && i <= len(undo.Undo.Parameters); i++ {
				if undo.Undo.Parameters[i-1].Source != "" {
					object, err := lookup.LookupStringI(responseJSON, undo.Undo.Parameters[i-1].Source)
					if err != nil {
						t.Fatalf("%v", err)
					}
					in[i] = object.Convert(undoOperation.Type().In(i))
				} else if undo.Undo.Parameters[i-1].Template != "" {
					data := Templated(t, responseJSON.(map[string]interface{}), undo.Undo.Parameters[i-1].Template)
					object := reflect.New(undoOperation.Type().In(i))
					err := json.Unmarshal([]byte(data), object.Interface())
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
