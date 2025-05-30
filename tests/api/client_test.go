package api

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests/api/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const FAKE_TOKEN = "fake-token"

func TestPreAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	awsAuth.EXPECT().Authenticate(gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextDelegatedToken,
		&datadog.DelegatedTokenConfig{
			ProviderAuth: awsAuth,
			Provider:     "aws",
		},
	)
	config := datadog.NewConfiguration()
	testAPIClient := datadog.NewAPIClient(config)
	token, err := testAPIClient.GetDelegatedToken(testAuthCtx)
	assert.Nil(t, err)
	assert.Equal(t, token.DelegatedToken, FAKE_TOKEN)

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

func TestNoPreAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	awsAuth.EXPECT().Authenticate(gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextDelegatedToken,
		&datadog.DelegatedTokenConfig{
			ProviderAuth: awsAuth,
			Provider:     "aws",
		},
	)

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

func TestReAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	firstCall := awsAuth.EXPECT().Authenticate(gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(-16) * time.Minute),
		}, nil)
	awsAuth.EXPECT().Authenticate(gomock.Any()).After(firstCall).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextDelegatedToken,
		&datadog.DelegatedTokenConfig{
			ProviderAuth: awsAuth,
			Provider:     "aws",
		},
	)

	config := datadog.NewConfiguration()
	testAPIClient := datadog.NewAPIClient(config)
	testAPIClient.GetDelegatedToken(testAuthCtx)

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
