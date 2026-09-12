package api

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// delegatedTokenTestContext returns a context and DelegatedTokenConfig wired
// so that GetDelegatedTokenUrl routes its POST /api/v2/delegated-token call at
// the given test server, using the "{protocol}://{name}" server configuration
// (index 1).
func delegatedTokenTestContext(t *testing.T, server *httptest.Server, provider string) (context.Context, *datadog.DelegatedTokenConfig) {
	t.Helper()

	ctx := context.WithValue(context.Background(), datadog.ContextServerIndex, 1)
	ctx = context.WithValue(ctx, datadog.ContextServerVariables, map[string]string{
		"name":     server.URL[len("http://"):],
		"protocol": "http",
	})
	ctx = context.WithValue(ctx, datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})

	cfg := &datadog.DelegatedTokenConfig{
		OrgUUID:  "12345678-1234-1234-1234-123456789012",
		Provider: provider,
	}
	return ctx, cfg
}
