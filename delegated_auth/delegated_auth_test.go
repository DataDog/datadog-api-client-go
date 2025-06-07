package delegated_auth

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"strings"
	"testing"
	"time"
)

func TestGetServerVariables(t *testing.T) {
	testOrgUuid := "12345678-1234-1234-1234-123456789012"
	serverVars := map[string]string{
		"api_subdomain": "api",
		"org_uuid":      testOrgUuid,
	}

	ctx := context.WithValue(context.Background(), datadog.ContextServerVariables, serverVars)

	retrievedVars := getServerVariables(ctx)
	if retrievedVars["api_subdomain"] != "api" {
		t.Errorf("expected api_subdomain to be 'api', got '%s'", retrievedVars["api_subdomain"])
	}
	if retrievedVars["org_uuid"] != testOrgUuid {
		t.Errorf("expected org_uuid to be '%s', got '%s'", testOrgUuid, retrievedVars["org_uuid"])
	}
}

func TestGetSite(t *testing.T) {
	testSite := "datadoghq.com"

	type testCase struct {
		name         string
		serverVars   map[string]string
		expectedSite string
	}

	testCases := []testCase{
		{
			name:         "Default Site",
			serverVars:   map[string]string{"site": testSite},
			expectedSite: testSite,
		},
		{
			name:         "Site from name",
			serverVars:   map[string]string{"name": testSite},
			expectedSite: testSite,
		},
		{
			name:         "Site from name with subdomain",
			serverVars:   map[string]string{"name": "dd.datadoghq.com"},
			expectedSite: "dd.datadoghq.com",
		},
		{
			name:         "Site from name using api subdomain",
			serverVars:   map[string]string{"name": "api.datadoghq.com"},
			expectedSite: testSite,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			site, err := getSite(tc.serverVars)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if site != tc.expectedSite {
				t.Errorf("expected site '%s', got '%s'", tc.expectedSite, site)
			}
		})
	}
}

func TestGetApiSubdomain(t *testing.T) {
	testCases := []struct {
		name           string
		serverVars     map[string]string
		expectedSubdom string
	}{
		{
			name:           "Default API Subdomain",
			serverVars:     map[string]string{"subdomain": "api"},
			expectedSubdom: "api",
		},
		{
			name:           "Custom API Subdomain",
			serverVars:     map[string]string{"subdomain": "custom"},
			expectedSubdom: "custom",
		},
		{
			name:           "Empty API Subdomain",
			serverVars:     map[string]string{},
			expectedSubdom: "api",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subdom := getApiSubdomain(tc.serverVars)
			if subdom != tc.expectedSubdom {
				t.Errorf("expected subdomain '%s', got '%s'", tc.expectedSubdom, subdom)
			}
		})
	}
}

func TestGetUrl(t *testing.T) {
	testCases := []struct {
		name           string
		serverVars     map[string]string
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
			expectedUrl:    "",
			expectedErrMsg: "site not found in server variables",
		},
		{
			name:        "Missing Subdomain",
			serverVars:  map[string]string{"site": "datadoghq.com"},
			expectedUrl: "https://api.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:        "Valid URL from name",
			serverVars:  map[string]string{"name": "api.datadoghq.com"},
			expectedUrl: "https://api.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:        "Valid URL from name no subdomain",
			serverVars:  map[string]string{"name": "datadoghq.com"},
			expectedUrl: "https://api.datadoghq.com/api/v2/delegated-token",
		},
		{
			name:           "Invalid URL with missing site and subdomain",
			serverVars:     map[string]string{},
			expectedUrl:    "",
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
			url, err := getUrl(ctx)
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
			creds, err := parseTokenResponse(tc.tokenResponse, testOrgUuid, testProof)
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
