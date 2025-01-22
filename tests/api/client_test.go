package api

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/stretchr/testify/assert"
	"testing"
)

const FAKE_TOKEN = "fake-token"

type FakeAWSAuth struct {
	orgUUID string
}

func (f *FakeAWSAuth) Authenticate() (*datadog.CloudProviderCredentials, error) {
	return &datadog.CloudProviderCredentials{
		OrgUUID:            f.orgUUID,
		DatadogToken:       FAKE_TOKEN,
		CloudProviderProof: "proof",
	}, nil
}

func TestAuthenticate(t *testing.T) {
	awsAuth := FakeAWSAuth{
		orgUUID: "1234",
	}

	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextCloudProvider,
		&datadog.CloudProviderConfig{
			ProviderAuth: &awsAuth,
			Provider:     "aws",
		},
	)
	config := datadog.NewConfiguration()
	testAPIClient := datadog.NewAPIClient(config)

	token, err := testAPIClient.GetTokenUsingCloudProviderAuth(testAuthCtx)
	assert.Nil(t, err)
	assert.Equal(t, token.DatadogToken, FAKE_TOKEN)

	headers := map[string]string{}
	datadog.SetAuthKeys(
		testAuthCtx,
		&headers,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	assert.NotEmpty(t, headers)
	assert.Equal(t, headers["dd-auth-token"], FAKE_TOKEN)

}
