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
	"os"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	ddtesting "github.com/DataDog/dd-sdk-go-testing"
	"github.com/go-bdd/gobdd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	msgs "github.com/cucumber/messages-go/v12"
)

func TestScenarios(t *testing.T) {
	// Load undo defintions
	requestsUndoV1, err := LoadRequestsUndo("features/v1/undo.json")
	if err != nil {
		t.Fatalf("could not load undo actions: %v", err)
	}
	requestsUndoV2, err := LoadRequestsUndo("features/v2/undo.json")
	if err != nil {
		t.Fatalf("could not load undo actions: %v", err)
	}
	requestsUndo := map[string]map[string]UndoAction{
		"v1": requestsUndoV1,
		"v2": requestsUndoV2,
	}

	stepsV1, err := LoadGivenSteps("features/v1/given.json")
	if err != nil {
		t.Fatalf("could not load given steps: %v", err)
	}

	stepsV2, err := LoadGivenSteps("features/v2/given.json")
	if err != nil {
		t.Fatalf("could not load given steps: %v", err)
	}

	var bddTags []string
	if tags, ok := os.LookupEnv("BDD_TAGS"); ok {
		bddTags = strings.Split(tags, ",")
	}
	for _, version := range []string{"v1", "v2"} {
		t.Run(version, func(t *testing.T) {
			s := gobdd.NewSuite(
				t,
				gobdd.WithFeaturesPath(fmt.Sprintf("features/%s/*.feature", version)),
				gobdd.WithTags(bddTags...),
				gobdd.WithIgnoredTags(GetIgnoredTags()...),
				gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
					SetVersion(ctx, version)
					ct, _ := ctx.Get(gobdd.TestingTKey{})
					tt := ct.(*testing.T)
					cf, _ := ctx.Get(gobdd.FeatureKey{})
					feature := cf.(*msgs.GherkinDocument_Feature)
					cs, _ := ctx.Get(gobdd.ScenarioKey{})
					scenario := cs.(*msgs.GherkinDocument_Feature_Scenario)

					// Extract code owners from tags
					codeowners := make([]string, 0)
					for _, tag := range append(scenario.Tags, feature.Tags...) {
						prefix := "@team:"
						t := string(tag.GetName())
						if strings.HasPrefix(t, prefix) {
							codeowners = append(codeowners, fmt.Sprintf("@%s", t[len(prefix):]))
						}
					}
					testCodeowners, _ := json.Marshal(codeowners)
					
					cctx, closeSpan := ddtesting.StartTestWithContext(
						datadog.NewDefaultContext(context.Background()),
						tt,
						ddtesting.WithSpanOptions(
							// Override the default tags set by ddtesting package for bdd
							tracer.Tag("test.name", scenario.Name),
							tracer.Tag("test.suite", fmt.Sprintf("features/%s/%s.feature", version, strings.ToLower(strings.Replace(feature.Name, " ", "-", -1)))),
							tracer.Tag("test.framework", "github.com/go-bdd/gobdd"),
							// Set resource name to TestName
							tracer.ResourceName(tt.Name()),
							// Set the codeowners
							tracer.Tag("test.codeowners", string(testCodeowners)),
						),
					)
					cctx, closeRecorder := ConfigureClients(ctx, cctx)
					SetCtx(ctx, cctx)
					SetRequestsUndo(ctx, requestsUndo)
					SetData(ctx, make(map[string]interface{}))
					SetCleanup(ctx, map[string]func(){"99-finish": func() {
						closeRecorder()
						closeSpan()
					}})
				}),
				gobdd.WithAfterScenario(RunCleanup),
				gobdd.WithBeforeStep(StartSpanBeforeStep),
				gobdd.WithBeforeStep(SetFixtureData),
				gobdd.WithAfterStep(FinishSpanAfterStep),
			)
			ConfigureSteps(s)

			for _, givenStep := range stepsV1 {
				givenStep.RegisterSuite(s, "v1")
			}

			for _, givenStep := range stepsV2 {
				givenStep.RegisterSuite(s, "v2")
			}

			s.Run()
		})
	}
}
