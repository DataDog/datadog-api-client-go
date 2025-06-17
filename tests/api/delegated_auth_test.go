package api

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func TestGetDelegatedTokenUrl(t *testing.T) {
	testCases := []struct {
		name           string
		serverVars     map[string]string
		serverIndex    int
		expectedUrl    string
		expectedErrMsg string
	}{
		{
			name:           "Valid URL",
			serverVars:     map[string]string{"site": "datadoghq.com", "subdomain": "api"},
			expectedUrl:    "https://api.datadoghq.com/api/v2/delegated-token",
			expectedErrMsg: "",
		},
		{
			name:           "Missing Site",
			serverVars:     map[string]string{"subdomain": "api"},
			expectedUrl:    "https://api.datadoghq.com/api/v2/delegated-token",
			expectedErrMsg: "site not found in server variables",
		},
		{
			name:        "Missing Subdomain",
			serverVars:  map[string]string{"site": "datadoghq.com"},
			expectedUrl: "https://api.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:        "Valid URL from name",
			serverVars:  map[string]string{"name": "api.datadoghq.eu"},
			serverIndex: 1,
			expectedUrl: "https://api.datadoghq.eu/api/v2/delegated-token",
		},
		{
			name:        "Valid URL from name no subdomain",
			serverVars:  map[string]string{"name": "datadoghq.com"},
			serverIndex: 1,
			expectedUrl: "https://api.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:           "Invalid URL with missing site and subdomain",
			serverVars:     map[string]string{},
			expectedUrl:    "https://api.datadoghq.com/api/v2/delegated-token",
			expectedErrMsg: "site not found in server variables",
		},
		{
			name:        "Valid URL with custom subdomain",
			serverVars:  map[string]string{"site": "datadoghq.com", "subdomain": "custom"},
			expectedUrl: "https://custom.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:        "Valid URL with custom subdomain and name",
			serverVars:  map[string]string{"name": "datadoghq.com", "subdomain": "custom"},
			expectedUrl: "https://custom.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:        "Valid URL with custom subdomain and name with api prefix",
			serverVars:  map[string]string{"name": "api.datadoghq.com", "subdomain": "custom"},
			expectedUrl: "https://custom.datadoghq.com/api/v2/delegated-token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), datadog.ContextServerVariables, tc.serverVars)
			ctx = context.WithValue(ctx, datadog.ContextServerIndex, tc.serverIndex)
			url, err := datadog.GetDelegatedTokenUrl(ctx)
			if err != nil && err.Error() != tc.expectedErrMsg {
				t.Errorf("expected error '%s', got '%v'", tc.expectedErrMsg, err)
			}
			if url != tc.expectedUrl {
				t.Errorf("expected URL '%s', got '%s'", tc.expectedUrl, url)
			}
		})
	}
}

func TestParseTokenResponse(t *testing.T) {
	testOrgUuid := "12345678-1234-1234-1234-123456789012"
	testProof := "test_proof"
	expiration := time.Unix(1700000000, 0)
	type testCase struct {
		name             string
		tokenResponse    []byte
		expectedCreds    *datadog.DelegatedTokenCredentials
		expectedErrorMsg string
	}

	testCases := []testCase{
		{
			name: "Valid Token Response",
			tokenResponse: []byte(`{
				"data": {
					"type": "delegated_token",
					"id": "12345678-1234-1234-1234-123456789012",	
					"attributes": {	
						"access_token": "valid_token",
						"expires": "1700000000"
					}
				}}`),
			expectedCreds: &datadog.DelegatedTokenCredentials{
				OrgUUID:        testOrgUuid,
				DelegatedToken: "valid_token",
				DelegatedProof: testProof,
				Expiration:     expiration,
			},
		},
		{
			name: "Invalid Token Response - Missing Access token",
			tokenResponse: []byte(`{
				"data": {
					"type": "delegated_token",
					"id": "12345678-1234-1234-1234-123456789012",
					"attributes": {
						"expires": "1700000000"
					}
				}}`),
			expectedCreds:    nil,
			expectedErrorMsg: "failed to get token from response",
		},
		{
			name: "Invalid Token Response - Missing Expires; Default to 15 minutes",
			tokenResponse: []byte(`{
				"data": {
					"type": "delegated_token",
					"id": "12345678-1234-1234-1234-123456789012",
					"attributes": {
						"access_token": "valid_token"
					}
				}}`),
			expectedCreds: &datadog.DelegatedTokenCredentials{
				OrgUUID:        testOrgUuid,
				DelegatedToken: "valid_token",
				DelegatedProof: testProof,
				Expiration:     time.Now().Add(15 * time.Minute), // Default expiration
			},
		},
		{
			name: "Invalid Token Response - Empty Data",
			tokenResponse: []byte(`{
				"data": {}	
				}`),
			expectedCreds:    nil,
			expectedErrorMsg: "failed to get attributes from response",
		},
		{
			name: "Invalid Token Response - Empty Attributes",
			tokenResponse: []byte(`{
				"data": {
					"type": "delegated_token",
					"id": "12345678-1234-1234-1234-123456789012",
					"attributes": {}
				}}`),
			expectedCreds:    nil,
			expectedErrorMsg: "failed to get token from response",
		},
		{
			name: "Invalid Token Response - Invalid Expires Format; Default to 15 minutes",
			tokenResponse: []byte(`{
				"data": {
					"type": "delegated_token",
					"id": "12345678-1234-1234-1234-123456789012",
					"attributes": {
						"access_token": "valid_token",
						"expires": "invalid_expires"
					}
				}}`),
			expectedCreds: &datadog.DelegatedTokenCredentials{
				OrgUUID:        testOrgUuid,
				DelegatedToken: "valid_token",
				DelegatedProof: testProof,
				Expiration:     time.Now().Add(15 * time.Minute), // Default expiration
			},
		},
		{
			name:             "Invalid Token Response - Invalid JSON",
			tokenResponse:    []byte(`{invalid json`),
			expectedCreds:    nil,
			expectedErrorMsg: "invalid character 'i' looking for beginning of object key string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			creds, err := datadog.ParseDelegatedTokenResponse(tc.tokenResponse, testOrgUuid, testProof)
			if err != nil && tc.expectedErrorMsg == "" {
				t.Errorf("unexpected error: %v", err)
			}
			if tc.expectedErrorMsg != "" && (err == nil || !strings.Contains(err.Error(), tc.expectedErrorMsg)) {
				t.Errorf("expected error '%s', got '%v'", tc.expectedErrorMsg, err)
			}
			if creds != nil && tc.expectedCreds != nil {
				if creds.OrgUUID != tc.expectedCreds.OrgUUID {
					t.Errorf("expected OrgUUID '%s', got '%s'", tc.expectedCreds.OrgUUID, creds.OrgUUID)
				}
				if creds.DelegatedToken != tc.expectedCreds.DelegatedToken {
					t.Errorf("expected DelegatedToken '%s', got '%s'", tc.expectedCreds.DelegatedToken, creds.DelegatedToken)
				}
				if creds.DelegatedProof != tc.expectedCreds.DelegatedProof {
					t.Errorf("expected DelegatedProof '%s', got '%s'", tc.expectedCreds.DelegatedProof, creds.DelegatedProof)
				}
				if creds.Expiration.Sub(tc.expectedCreds.Expiration).Abs() > (5 * time.Second) {
					t.Errorf("expected Expiration within 5 second '%v', got '%v'", tc.expectedCreds.Expiration, creds.Expiration)
				}
			}
		})
	}
}
