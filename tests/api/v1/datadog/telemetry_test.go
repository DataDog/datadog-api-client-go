package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer stop()

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	gock.New("https://api.datadoghq.com").
		Get("integration/aws").
		MatchHeader("DD-OPERATION-ID", "ListAWSAccounts").
		MatchHeader("User-Agent", "^datadog-api-client-go/\\d\\.\\d\\.\\d.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := Client(ctx).AWSIntegrationApi.ListAWSAccounts(ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 299, httpresp.StatusCode)
}
