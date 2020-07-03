package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func TestScenarios(t *testing.T) {
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags([]string{"@todo"}),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			cctx, finish := WithClient(
				context.WithValue(
					context.Background(),
					datadog.ContextAPIKeys,
					map[string]datadog.APIKey{},
				),
				ct.(*testing.T),
			)
			tests.SetCtx(ctx, cctx)
			tests.SetData(ctx, make(map[string]interface{}))
			tests.SetCleanup(ctx, map[string]func(){"99-finish": finish})
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			tests.RunCleanup(ctx)
		}),
		gobdd.WithBeforeStep(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			parts := strings.Split(ct.(*testing.T).Name(), "/")
			if parent, ok := tracer.SpanFromContext(tests.GetCtx(ctx)); ok {
				ctx.Set("parentSpan", parent)
			}
			_, cctx := tracer.StartSpanFromContext(
				tests.GetCtx(ctx),
				"step",
				tracer.SpanType("step"),
				tracer.ResourceName(parts[len(parts)-1]),
			)
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

				if parent, err := ctx.Get("parentSpan"); err == nil {
					tests.SetCtx(ctx, tracer.ContextWithSpan(cctx, parent.(ddtrace.Span)))
				}
			}
		}),
	)
	tests.ConfigureSteps(s)

	s.AddStep(`a valid "apiKeyAuth" key in the system`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key in the system`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
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
	// use only first 3 segments from test name TestScenarios/<Feature Name>/<Scenario Name>
	path := tests.SecurePath(strings.Join(strings.Split(t.(*testing.T).Name(), "/")[0:3], "/"))

	cctx, err := tests.WithClock(tests.GetCtx(ctx), path)
	if err != nil {
		log.Fatal(err)
	}
	tests.SetCtx(ctx, cctx)

	r, err := tests.Recorder(tests.GetCtx(ctx), path)
	if err != nil {
		log.Fatal(err)
	}
	tests.GetCleanup(ctx)["90-recorder"] = func() { r.Stop() }
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	ct := reflect.ValueOf(client)
	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid API name %s", name))
	}
	tests.SetAPI(ctx, f)
}
