package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests/api/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const FAKE_TOKEN = "fake-token"

func TestDelegatedPreAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	awsAuth.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := datadog.NewDefaultContext(context.Background())
	testAuthCtx = context.WithValue(
		testAuthCtx,
		datadog.ContextAWSVariables,
		map[string]string{
			"aws_access_key_id":     "fake-access-key-id",
			"aws_secret_access_key": "fake-secret-access-key",
			"aws_session_token":     "fake-session-token",
		},
	)

	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
		OrgUUID:      "1234",
		ProviderAuth: awsAuth,
		Provider:     "aws",
	}
	testAPIClient := datadog.NewAPIClient(config)
	token, err := testAPIClient.GetDelegatedToken(testAuthCtx)
	assert.Nil(t, err)
	assert.Equal(t, token.DelegatedToken, FAKE_TOKEN)

	headers := map[string]string{}
	datadog.UseDelegatedTokenAuth(testAuthCtx, &headers, config.DelegatedTokenConfig)
	assert.NotEmpty(t, headers)
	assert.Equal(t, headers["Authorization"], fmt.Sprintf("Bearer %s", FAKE_TOKEN))
}

func TestDelegatedNoPreAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	awsAuth.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := datadog.NewDefaultContext(context.Background())
	testAuthCtx = context.WithValue(
		testAuthCtx,
		datadog.ContextAWSVariables,
		map[string]string{
			"aws_access_key_id":     "fake-access-key-id",
			"aws_secret_access_key": "fake-secret-access-key",
			"aws_session_token":     "fake-session-token",
		},
	)

	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
		OrgUUID:      "1234",
		ProviderAuth: awsAuth,
		Provider:     "aws",
	}

	headers := map[string]string{}
	datadog.UseDelegatedTokenAuth(testAuthCtx, &headers, config.DelegatedTokenConfig)
	assert.NotEmpty(t, headers)
	assert.Equal(t, headers["Authorization"], fmt.Sprintf("Bearer %s", FAKE_TOKEN))
}

func TestDelegatedReAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	awsAuth := mocks.NewMockDelegatedTokenProvider(ctrl)
	firstCall := awsAuth.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(-16) * time.Minute),
		}, nil)
	awsAuth.EXPECT().Authenticate(gomock.Any(), gomock.Any()).After(firstCall).Return(
		&datadog.DelegatedTokenCredentials{
			OrgUUID:        "1234",
			DelegatedToken: FAKE_TOKEN,
			DelegatedProof: "proof",
			Expiration:     time.Now().Add(time.Duration(10) * time.Minute),
		}, nil)

	testAuthCtx := datadog.NewDefaultContext(context.Background())
	testAuthCtx = context.WithValue(
		testAuthCtx,
		datadog.ContextAWSVariables,
		map[string]string{
			"aws_access_key_id":     "fake-access-key-id",
			"aws_secret_access_key": "fake-secret-access-key",
			"aws_session_token":     "fake-session-token",
		},
	)

	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
		OrgUUID:      "1234",
		ProviderAuth: awsAuth,
		Provider:     "aws",
	}
	testAPIClient := datadog.NewAPIClient(config)
	testAPIClient.GetDelegatedToken(testAuthCtx)

	headers := map[string]string{}
	datadog.UseDelegatedTokenAuth(testAuthCtx, &headers, config.DelegatedTokenConfig)
	assert.NotEmpty(t, headers)
	assert.Equal(t, headers["Authorization"], fmt.Sprintf("Bearer %s", FAKE_TOKEN))
}

func TestApiAppKeyAuthenticate(t *testing.T) {
	apiKey := "api-key"
	appKey := "app-key"
	keys := map[string]datadog.APIKey{
		"apiKeyAuth": {Key: apiKey},
		"appKeyAuth": {Key: appKey},
	}
	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		keys,
	)

	headers := map[string]string{}
	datadog.SetAuthKeys(
		testAuthCtx,
		&headers,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	assert.NotEmpty(t, headers)
	assert.Equal(t, headers["DD-API-KEY"], apiKey)
	assert.Equal(t, headers["DD-APPLICATION-KEY"], appKey)
}

func TestBearerTokenAuthentication(t *testing.T) {
	pat := "ddapp_test-pat-token"
	keys := map[string]datadog.APIKey{
		"apiKeyAuth": {Key: "api-key"},
	}
	testAuthCtx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		keys,
	)
	testAuthCtx = context.WithValue(testAuthCtx, datadog.ContextBearerToken, pat)

	headers := map[string]string{}
	datadog.SetAuthKeys(
		testAuthCtx,
		&headers,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	assert.Equal(t, "Bearer ddapp_test-pat-token", headers["Authorization"])
	assert.Equal(t, "api-key", headers["DD-API-KEY"], "API key should still be sent alongside bearer token")
}

func TestNewDefaultContextWithBearerToken(t *testing.T) {
	t.Setenv("DD_API_KEY", "test-api-key")
	t.Setenv("DD_BEARER_TOKEN", "ddapp_test-pat-token")

	ctx := datadog.NewDefaultContext(context.Background())
	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	assert.Equal(t, "test-api-key", keys["apiKeyAuth"].Key)
	bearerToken := ctx.Value(datadog.ContextBearerToken).(string)
	assert.Equal(t, "ddapp_test-pat-token", bearerToken)
}

func TestNewDefaultContextAllAuthHeadersSentTogether(t *testing.T) {
	t.Setenv("DD_API_KEY", "test-api-key")
	t.Setenv("DD_APP_KEY", "test-app-key")
	t.Setenv("DD_BEARER_TOKEN", "ddapp_test-pat-token")

	ctx := datadog.NewDefaultContext(context.Background())
	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	assert.Equal(t, "test-app-key", keys["appKeyAuth"].Key)
	bearerToken := ctx.Value(datadog.ContextBearerToken).(string)
	assert.Equal(t, "ddapp_test-pat-token", bearerToken)

	// All auth headers should be sent together
	headers := map[string]string{}
	datadog.SetAuthKeys(
		ctx,
		&headers,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	assert.Equal(t, "Bearer ddapp_test-pat-token", headers["Authorization"])
	assert.Equal(t, "test-api-key", headers["DD-API-KEY"], "API key should be sent alongside bearer token")
	assert.Equal(t, "test-app-key", headers["DD-APPLICATION-KEY"], "App key should be sent alongside bearer token")
}

func TestNewDefaultContextAppKeyWithoutBearerToken(t *testing.T) {
	t.Setenv("DD_API_KEY", "test-api-key")
	t.Setenv("DD_APP_KEY", "test-app-key")

	ctx := datadog.NewDefaultContext(context.Background())
	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	assert.Equal(t, "test-app-key", keys["appKeyAuth"].Key)
	assert.Nil(t, ctx.Value(datadog.ContextBearerToken))
}

func TestPATBearerHeaderSentViaHTTPClient(t *testing.T) {
	pat := "ddapp_test-token"

	var capturedHeaders http.Header
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeaders = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, `{"data":[],"meta":{"page":{"total_count":0,"total_filtered_count":0}}}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	ctx = context.WithValue(ctx, datadog.ContextBearerToken, pat)

	cfg := datadog.NewConfiguration()
	cfg.Servers = datadog.ServerConfigurations{{URL: ts.URL}}
	client := datadog.NewAPIClient(cfg)
	usersAPI := datadogV2.NewUsersApi(client)

	_, httpResp, err := usersAPI.ListUsers(ctx, *datadogV2.NewListUsersOptionalParameters())
	require.NoError(t, err)
	defer httpResp.Body.Close()

	assert.Equal(t, 200, httpResp.StatusCode)
	assert.Equal(t, "Bearer "+pat, capturedHeaders.Get("Authorization"))
}
