package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	gock.New("https://api.datadoghq.com").
		Get("integration/aws").
		MatchHeader("DD-OPERATION-ID", "ListAWSAccounts").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.ListAWSAccounts(TESTAUTH).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 299, httpresp.StatusCode)
}
