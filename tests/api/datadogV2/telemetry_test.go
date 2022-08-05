package test

import (
	"context"
	"gopkg.in/h2non/gock.v1"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestTelemetryHeaders(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV2.NewDashboardListsApi(Client(ctx))

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "DashboardListsApiService.GetDashboardListItems")
	assert.NoError(err)
	gock.New(URL).
		Get("dashboard/lists/manual/1234/dashboards").
		MatchHeader("User-Agent", "^datadog-api-client-go\\/\\d+\\.\\d+\\.\\d+.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := api.GetDashboardListItems(ctx, 1234)
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
}
