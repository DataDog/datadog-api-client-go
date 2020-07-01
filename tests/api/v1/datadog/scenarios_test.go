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
)

func TestScenarios(t *testing.T) {
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags([]string{"@todo"}),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			client := datadog.NewAPIClient(NewConfiguration())
			tests.SetCtx(ctx, context.WithValue(context.WithValue(
				context.Background(),
				datadog.ContextAPIKeys,
				map[string]datadog.APIKey{},
			), clientKey, client))
			tests.SetData(ctx, make(map[string]interface{}))
			tests.SetCleanup(ctx, make(map[string]func()))
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			tests.RunCleanup(ctx)
		}),
	)
	tests.ConfigureSteps(s)

	s.AddStep(`a valid "apiKeyAuth" key`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
	s.Run()
}

func aValidAPIKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["apiKeyAuth"] = datadog.APIKey{
			Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
		}
	} else {
		panic("could not set API key")
	}
}

func aValidAppKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["appKeyAuth"] = datadog.APIKey{
			Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
		}
	} else {
		panic("could not set App key")
	}
}

// anInstanceOf sets API callable to apiKey{}
func anInstanceOf(t gobdd.StepTest, ctx gobdd.Context, name string) {
	client := Client(tests.GetCtx(ctx))
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
