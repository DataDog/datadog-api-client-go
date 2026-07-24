/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-bdd/gobdd"
)

const testRunnerTemplateKey = "$openapi_transformer_template"

type testRunnerManifest struct {
	Scenarios []struct {
		Version  string `json:"version"`
		Feature  string `json:"feature"`
		Scenario string `json:"scenario"`
		File     string `json:"file"`
	} `json:"scenarios"`
}

type testRunnerPlan struct {
	API         string `json:"api"`
	OperationID string `json:"operation_id"`
	Request     struct {
		Body *struct {
			Value interface{} `json:"value"`
		} `json:"body"`
		Parameters []struct {
			Name     string `json:"name"`
			In       string `json:"in"`
			Required bool   `json:"required"`
			Source   struct {
				Type  string      `json:"type"`
				Path  string      `json:"path"`
				Value interface{} `json:"value"`
			} `json:"source"`
		} `json:"parameters"`
		Pagination bool `json:"pagination"`
	} `json:"request"`
}

type testServerSession struct {
	Session  string    `json:"session"`
	FrozenAt time.Time `json:"frozen_at"`
}

type testFeatureKey struct{}
type testScenarioKey struct{}
type testServerSessionKey struct{}

func testRunnerEnabled() bool {
	return os.Getenv("DD_TEST_RUNNER_DATA") != ""
}

func testServerEnabled() bool {
	return os.Getenv("DD_TEST_SERVER_URL") != ""
}

func testServerRequest(endpoint string, payload interface{}, result interface{}) error {
	var body io.Reader
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewReader(encoded)
	}
	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("DD_TEST_SERVER_URL")+"/__openapi_transformer__"+endpoint,
		body,
	)
	if err != nil {
		return err
	}
	if payload != nil {
		request.Header.Set("content-type", "application/json")
	}
	request.Close = true
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("test server POST %s failed (%d): %s", endpoint, response.StatusCode, responseBody)
	}
	if result == nil {
		return nil
	}
	return json.Unmarshal(responseBody, result)
}

func startTestServerSession(ctx gobdd.Context, feature string, scenario string) (testServerSession, error) {
	ctx.Set(testFeatureKey{}, feature)
	ctx.Set(testScenarioKey{}, scenario)
	var session testServerSession
	err := testServerRequest("/sessions", map[string]string{
		"version":  GetVersion(ctx),
		"feature":  feature,
		"scenario": scenario,
	}, &session)
	if err == nil {
		ctx.Set(testServerSessionKey{}, session.Session)
	}
	return session, err
}

func stopTestServerSession(ctx gobdd.Context) error {
	value, err := ctx.Get(testServerSessionKey{})
	if err != nil {
		return nil
	}
	var result struct {
		Complete          bool `json:"complete"`
		Interactions      int  `json:"interactions"`
		TotalInteractions int  `json:"total_interactions"`
	}
	if err := testServerRequest(fmt.Sprintf("/sessions/%s/stop", value.(string)), nil, &result); err != nil {
		return err
	}
	if !result.Complete {
		return fmt.Errorf(
			"test server session consumed %d of %d interactions",
			result.Interactions,
			result.TotalInteractions,
		)
	}
	return nil
}

func markMainRequestComplete(ctx gobdd.Context) error {
	value, err := ctx.Get(testServerSessionKey{})
	if err != nil {
		return nil
	}
	return testServerRequest(
		fmt.Sprintf("/sessions/%s/main-complete", value.(string)),
		nil,
		nil,
	)
}

func loadTestRunnerPlan(ctx gobdd.Context) (*testRunnerPlan, error) {
	root := os.Getenv("DD_TEST_RUNNER_DATA")
	manifestBody, err := os.ReadFile(filepath.Join(root, "manifest.json"))
	if err != nil {
		return nil, err
	}
	var manifest testRunnerManifest
	if err := json.Unmarshal(manifestBody, &manifest); err != nil {
		return nil, err
	}
	feature, _ := ctx.Get(testFeatureKey{})
	scenario, _ := ctx.Get(testScenarioKey{})
	for _, item := range manifest.Scenarios {
		if item.Version == GetVersion(ctx) && item.Feature == feature && item.Scenario == scenario {
			planBody, err := os.ReadFile(filepath.Join(root, item.File))
			if err != nil {
				return nil, err
			}
			var plan testRunnerPlan
			if err := json.Unmarshal(planBody, &plan); err != nil {
				return nil, err
			}
			return &plan, nil
		}
	}
	return nil, fmt.Errorf(
		"generated request plan not found for %s/%s/%s",
		GetVersion(ctx),
		feature,
		scenario,
	)
}

func materializeTestRunnerValue(t gobdd.StepTest, ctx gobdd.Context, value interface{}) (interface{}, error) {
	switch typed := value.(type) {
	case map[string]interface{}:
		if len(typed) == 1 {
			if template, ok := typed[testRunnerTemplateKey].(string); ok {
				rendered := Templated(t, GetData(ctx), template)
				var result interface{}
				if err := json.Unmarshal([]byte(rendered), &result); err != nil {
					return nil, err
				}
				return result, nil
			}
		}
		result := make(map[string]interface{}, len(typed))
		for key, item := range typed {
			materialized, err := materializeTestRunnerValue(t, ctx, item)
			if err != nil {
				return nil, err
			}
			result[key] = materialized
		}
		return result, nil
	case []interface{}:
		result := make([]interface{}, len(typed))
		for index, item := range typed {
			materialized, err := materializeTestRunnerValue(t, ctx, item)
			if err != nil {
				return nil, err
			}
			result[index] = materialized
		}
		return result, nil
	case string:
		return Templated(t, GetData(ctx), typed), nil
	default:
		return value, nil
	}
}

func applyTestRunnerPlan(t gobdd.StepTest, ctx gobdd.Context, pagination bool) {
	if !testRunnerEnabled() {
		return
	}
	plan, err := loadTestRunnerPlan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if plan.Request.Pagination != pagination {
		t.Fatalf("generated request plan pagination mismatch")
	}
	anInstanceOf(t, ctx, plan.API)
	newRequestNative(t, ctx, plan.OperationID)

	applyParameter := func(parameter struct {
		Name     string `json:"name"`
		In       string `json:"in"`
		Required bool   `json:"required"`
		Source   struct {
			Type  string      `json:"type"`
			Path  string      `json:"path"`
			Value interface{} `json:"value"`
		} `json:"source"`
	}) {
		if parameter.Source.Type == "fixture" {
			addParameterFromNative(t, ctx, parameter.Name, parameter.Source.Path)
			return
		}
		value, err := materializeTestRunnerValue(t, ctx, parameter.Source.Value)
		if err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		addParameterWithValueNative(t, ctx, parameter.Name, string(encoded))
	}
	for _, parameter := range plan.Request.Parameters {
		if parameter.In == "path" || parameter.Required {
			applyParameter(parameter)
		}
	}
	if plan.Request.Body != nil {
		value, err := materializeTestRunnerValue(t, ctx, plan.Request.Body.Value)
		if err != nil {
			t.Fatal(err)
		}
		body, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		GetRequestParameters(ctx)["body"] = string(body)
		pathCount, _ := ctx.Get(pathParamCountKey{})
		ctx.Set(pathParamCountKey{}, pathCount.(int)+1)
	}
	for _, parameter := range plan.Request.Parameters {
		if parameter.In != "path" && !parameter.Required {
			applyParameter(parameter)
		}
	}
}

type testServerTransport struct {
	session string
}

func (transport testServerTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	cloned := request.Clone(request.Context())
	cloned.Close = true
	cloned.Header.Set("x-openapi-test-session", transport.session)
	return http.DefaultTransport.RoundTrip(cloned)
}
