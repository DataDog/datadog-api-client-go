package api

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests/api/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

// TestDebugDumpRedactsAccessToken verifies that when Cfg.Debug is true and a
// bearer token is set via ContextAccessToken, the SDK's debug dump replaces
// the token with REDACTED before writing to the global logger. Without this,
// callers that enable debug logging would write the token to stderr, CI
// artifacts, and log shippers. See CRED-2625.
func TestDebugDumpRedactsAccessToken(t *testing.T) {
	var buf bytes.Buffer
	prev := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(prev)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	const token = "ddpat_supersecret_should_not_leak_a1b2c3d4e5"

	cfg := datadog.NewConfiguration()
	cfg.Debug = true
	cfg.RetryConfiguration.EnableRetry = false
	client := datadog.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), datadog.ContextAccessToken, token)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	assert.NoError(t, err)
	// In production code, PrepareRequest sets this header from ContextAccessToken
	// before CallAPI sees the request. Simulate that here so the dump path has
	// something to redact.
	req.Header.Set("Authorization", "Bearer "+token)
	_, err = client.CallAPI(req)
	assert.NoError(t, err)

	dumped := buf.String()
	assert.NotContains(t, dumped, token, "bearer token leaked into debug dump")
	assert.Contains(t, dumped, "REDACTED", "expected REDACTED marker in debug dump")
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
