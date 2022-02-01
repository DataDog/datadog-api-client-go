package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AWSIntegrationApiService.ListAWSAccounts")
	assert.NoError(err)
	gock.New(URL).
		Get("integration/aws").
		MatchHeader("User-Agent", "^datadog-api-client-go/\\d\\.\\d\\.\\d.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := Client(ctx).AWSIntegrationApi.ListAWSAccounts(ctx)
	assert.Nil(err)
	assert.Equal(299, httpresp.StatusCode)
}
