package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "DashboardListsApiService.GetDashboardListItems")
	assert.NoError(err)
	gock.New(URL).
		Get("dashboard/lists/manual/1234/dashboards").
		MatchHeader("DD-OPERATION-ID", "GetDashboardListItems").
		MatchHeader("User-Agent", "^datadog-api-client-go/\\d\\.\\d\\.\\d.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := Client(ctx).DashboardListsApi.GetDashboardListItems(ctx, 1234).Execute()
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
}
