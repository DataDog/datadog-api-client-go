package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

type refreshTestProvider struct {
	calls   atomic.Int32
	failing atomic.Bool
	release chan struct{}
}

func (provider *refreshTestProvider) Authenticate(ctx context.Context, config *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
	call := provider.calls.Add(1)
	if provider.release != nil {
		select {
		case <-provider.release:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if provider.failing.Load() {
		return nil, errors.New("token exchange unavailable")
	}
	return &datadog.DelegatedTokenCredentials{
		OrgUUID:        config.OrgUUID,
		DelegatedToken: fmt.Sprintf("refreshed-token-%d", call),
		Expiration:     time.Now().Add(15 * time.Minute),
	}, nil
}

type refreshTestAPI struct {
	server   *httptest.Server
	requests atomic.Int32
	mu       sync.Mutex
	bearers  []string
}

func newRefreshTestAPI(t *testing.T) *refreshTestAPI {
	t.Helper()
	api := &refreshTestAPI{}
	api.server = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		api.requests.Add(1)
		api.mu.Lock()
		api.bearers = append(api.bearers, request.Header.Get("Authorization"))
		api.mu.Unlock()
		response.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(response, `{"data":{"type":"users","id":"00000000-0000-0000-0000-000000000001"}}`)
	}))
	t.Cleanup(api.server.Close)
	return api
}

func (api *refreshTestAPI) lastBearer() string {
	api.mu.Lock()
	defer api.mu.Unlock()
	if len(api.bearers) == 0 {
		return ""
	}
	return api.bearers[len(api.bearers)-1]
}

func refreshTestClient(api *refreshTestAPI, provider datadog.DelegatedTokenProvider, credentials *datadog.DelegatedTokenCredentials) (*datadogV2.UsersApi, context.Context) {
	configuration := datadog.NewConfiguration()
	configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{OrgUUID: testOrgUUID, ProviderAuth: provider}
	ctx := context.WithValue(delegatedTokenContext(api.server.URL), datadog.ContextDelegatedToken, credentials)
	return datadogV2.NewUsersApi(datadog.NewAPIClient(configuration)), ctx
}

func TestDelegatedTokenRefreshesBeforeExpiry(t *testing.T) {
	tests := []struct {
		name        string
		remaining   time.Duration
		wantRefresh bool
	}{
		{name: "well before expiry", remaining: 10 * time.Minute},
		{name: "inside the refresh window", remaining: time.Minute, wantRefresh: true},
		{name: "expired", remaining: -time.Second, wantRefresh: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			api := newRefreshTestAPI(t)
			provider := &refreshTestProvider{}
			credentials := &datadog.DelegatedTokenCredentials{DelegatedToken: "current-token", Expiration: time.Now().Add(test.remaining)}
			users, ctx := refreshTestClient(api, provider, credentials)

			if _, _, err := users.GetCurrentUser(ctx); err != nil {
				t.Fatalf("API call: %v", err)
			}
			want, wantCalls := "Bearer current-token", int32(0)
			if test.wantRefresh {
				want, wantCalls = "Bearer refreshed-token-1", 1
			}
			if got := api.lastBearer(); got != want || provider.calls.Load() != wantCalls {
				t.Fatalf("authorization = %q with %d refreshes, want %q with %d", got, provider.calls.Load(), want, wantCalls)
			}
		})
	}
}

func TestDelegatedTokenConcurrentRequestsShareOneRefresh(t *testing.T) {
	api := newRefreshTestAPI(t)
	provider := &refreshTestProvider{release: make(chan struct{})}
	users, ctx := refreshTestClient(api, provider, &datadog.DelegatedTokenCredentials{})

	const callers = 20
	var wg sync.WaitGroup
	errs := make(chan error, callers)
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _, err := users.GetCurrentUser(ctx)
			errs <- err
		}()
	}
	// Give every caller time to reach the refresh before the provider answers.
	time.Sleep(100 * time.Millisecond)
	close(provider.release)
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("API call: %v", err)
		}
	}
	if provider.calls.Load() != 1 {
		t.Fatalf("provider refreshed %d times for %d concurrent callers, want 1", provider.calls.Load(), callers)
	}
	api.mu.Lock()
	defer api.mu.Unlock()
	for _, bearer := range api.bearers {
		if bearer != "Bearer refreshed-token-1" {
			t.Fatalf("a request carried %q, want the shared refreshed token", bearer)
		}
	}
}

func TestDelegatedTokenKeepsValidTokenWhenRefreshFails(t *testing.T) {
	api := newRefreshTestAPI(t)
	provider := &refreshTestProvider{}
	provider.failing.Store(true)
	credentials := &datadog.DelegatedTokenCredentials{DelegatedToken: "current-token", Expiration: time.Now().Add(time.Minute)}
	users, ctx := refreshTestClient(api, provider, credentials)

	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call with a still-valid token: %v", err)
	}
	if got := api.lastBearer(); got != "Bearer current-token" || provider.calls.Load() != 1 {
		t.Fatalf("authorization = %q after %d refreshes, want the current token after one failed refresh", got, provider.calls.Load())
	}

	provider.failing.Store(false)
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call after the provider recovered: %v", err)
	}
	if got := api.lastBearer(); got != "Bearer refreshed-token-2" {
		t.Fatalf("authorization = %q, want the token from the next successful refresh", got)
	}
}

func TestDelegatedTokenFailsWhenExpiredAndRefreshFails(t *testing.T) {
	api := newRefreshTestAPI(t)
	provider := &refreshTestProvider{}
	provider.failing.Store(true)
	credentials := &datadog.DelegatedTokenCredentials{DelegatedToken: "expired-token", Expiration: time.Now().Add(-time.Second)}
	users, ctx := refreshTestClient(api, provider, credentials)

	if _, _, err := users.GetCurrentUser(ctx); err == nil {
		t.Fatal("API call succeeded with an expired token and a failed refresh")
	}
	if api.requests.Load() != 0 {
		t.Fatal("the request was sent without a valid delegated token")
	}
	if credentials.DelegatedToken != "" {
		t.Fatalf("expired token %q was kept after a failed refresh", credentials.DelegatedToken)
	}

	provider.failing.Store(false)
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call after the provider recovered: %v", err)
	}
}

func TestDelegatedTokenRejectsProviderWithoutToken(t *testing.T) {
	api := newRefreshTestAPI(t)
	users, ctx := refreshTestClient(api, emptyDelegatedTokenProvider{}, &datadog.DelegatedTokenCredentials{})
	if _, _, err := users.GetCurrentUser(ctx); err == nil || api.requests.Load() != 0 {
		t.Fatalf("error = %v with %d requests, want an error before any request", err, api.requests.Load())
	}
}

type emptyDelegatedTokenProvider struct{}

func (emptyDelegatedTokenProvider) Authenticate(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
	return nil, nil
}

func TestDelegatedTokenWaiterHonorsItsContext(t *testing.T) {
	api := newRefreshTestAPI(t)
	provider := &refreshTestProvider{release: make(chan struct{})}
	users, ctx := refreshTestClient(api, provider, &datadog.DelegatedTokenCredentials{})

	shortCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	if _, _, err := users.GetCurrentUser(shortCtx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want the caller's deadline while the refresh is pending", err)
	}

	// The abandoned refresh still completes and serves the next caller.
	close(provider.release)
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call after the refresh completed: %v", err)
	}
	if provider.calls.Load() != 1 || api.lastBearer() != "Bearer refreshed-token-1" {
		t.Fatalf("refreshes = %d, authorization = %q; want the abandoned refresh's token reused", provider.calls.Load(), api.lastBearer())
	}
}

func TestDelegatedTokenExplicitAuthenticateIsSafeDuringRequests(t *testing.T) {
	api := newRefreshTestAPI(t)
	provider := &refreshTestProvider{}
	credentials := &datadog.DelegatedTokenCredentials{}
	users, ctx := refreshTestClient(api, provider, credentials)
	configuration := datadog.NewConfiguration()
	configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{OrgUUID: testOrgUUID, ProviderAuth: provider}
	client := datadog.NewAPIClient(configuration)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			_, _, _ = users.GetCurrentUser(ctx)
		}()
		go func() {
			defer wg.Done()
			_, _ = client.GetDelegatedToken(ctx)
		}()
	}
	wg.Wait()
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call after concurrent explicit authentication: %v", err)
	}
}
