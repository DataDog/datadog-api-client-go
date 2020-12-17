package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func TestScenarios(t *testing.T) {
	requestsUndo := tests.LoadRequestsUndo("./features/undo.json")
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags(tests.GetIgnoredTags()),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			cctx, finish := WithRecorder(
				context.WithValue(
					context.Background(),
					datadog.ContextAPIKeys,
					map[string]datadog.APIKey{},
				),
				ct.(*testing.T),
			)
			tests.SetCtx(ctx, cctx)
			tests.SetRequestsUndo(ctx, requestsUndo)
			tests.SetData(ctx, make(map[string]interface{}))
			tests.SetCleanup(ctx, map[string]func(){"99-finish": finish})
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			tests.RunCleanup(ctx)
		}),
		gobdd.WithBeforeStep(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			cctx := tests.GetCtx(ctx)
			parts := strings.Split(ct.(*testing.T).Name(), "/")
			if parent, ok := tracer.SpanFromContext(cctx); ok {
				ctx.Set("parentSpan", parent)
			} else {
				ctx.Set("parentSpan", nil)
			}
			_, cctx = tracer.StartSpanFromContext(
				cctx,
				"step",
				tracer.SpanType("step"),
				tracer.ResourceName(parts[len(parts)-1]),
			)

			testName := strings.Join(strings.Split(ct.(*testing.T).Name(), "/")[1:3], "/")
			unique := tests.WithUniqueSurrounding(cctx, testName)
			data := tests.GetData(ctx)
			data["unique"] = unique
			data["unique_lower"] = strings.ToLower(unique)

			tests.SetCtx(ctx, cctx)
		}),
		gobdd.WithAfterStep(func(ctx gobdd.Context) {
			cctx := tests.GetCtx(ctx)
			if span, ok := tracer.SpanFromContext(cctx); ok {
				ct, _ := ctx.Get(gobdd.TestingTKey{})
				failed := ct.(*testing.T).Failed()
				if failed {
					span.SetTag(ext.Error, failed)
				}
				span.Finish()

				if parent, err := ctx.Get("parentSpan"); err == nil && parent != nil {
					tests.SetCtx(ctx, tracer.ContextWithSpan(cctx, parent.(ddtrace.Span)))
				}
			}
		}),
	)
	tests.ConfigureSteps(s)

	s.AddStep(`a valid "apiKeyAuth" key in the system`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key in the system`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
	s.AddStep(`operation "([^"]+)" enabled`, enableOperations)

	for _, givenStep := range tests.LoadGivenSteps("./features/given.json") {
		givenStep.RegisterSuite(s)
	}

	s.Run()
}

func aValidAPIKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_API_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_API_KEY must be set")
	}
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["apiKeyAuth"] = datadog.APIKey{
			Key: key,
		}
	} else {
		t.Fatal("could not set API key")
	}
}

func aValidAppKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_APP_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_APP_KEY must be set")
	}
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["appKeyAuth"] = datadog.APIKey{
			Key: key,
		}
	} else {
		t.Fatal("could not set App key")
	}
}

// anInstanceOf sets API callable to apiKey{}
func anInstanceOf(t gobdd.StepTest, ctx gobdd.Context, name string) {
	client := Client(tests.GetCtx(ctx))
	ct := reflect.ValueOf(client)
	tests.SetClient(ctx, ct)

	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid API name %s", name))
	}
	tests.SetAPI(ctx, f)
}

// enableOperations sets unstable operations specific in this clause to enabled
func enableOperations(t gobdd.StepTest, ctx gobdd.Context, name string) {
	client := Client(tests.GetCtx(ctx))
	client.GetConfig().SetUnstableOperationEnabled(name, true)
}
