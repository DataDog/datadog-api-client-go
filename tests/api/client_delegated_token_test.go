package api

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

type delegatedTokenTestProvider func(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error)

func (f delegatedTokenTestProvider) Authenticate(ctx context.Context, config *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
	return f(ctx, config)
}

// Observe when a caller reaches the wait select, so tests need no scheduling sleeps.
type delegatedTokenWaitContext struct {
	context.Context
	waiting chan<- struct{}
}

func (c delegatedTokenWaitContext) Done() <-chan struct{} {
	c.waiting <- struct{}{}
	return c.Context.Done()
}

type delegatedTokenTestResult struct {
	creds *datadog.DelegatedTokenCredentials
	err   error
}

func delegatedTokenTestCall(ctx context.Context, client *datadog.APIClient, method int) delegatedTokenTestResult {
	switch method % 3 {
	case 0:
		creds, err := client.GetDelegatedToken(ctx)
		return delegatedTokenTestResult{creds, err}
	case 1:
		creds, err := datadog.CallDelegatedTokenAuthenticate(ctx, client.Cfg.DelegatedTokenConfig)
		return delegatedTokenTestResult{creds, err}
	default:
		headers := map[string]string{}
		err := datadog.UseDelegatedTokenAuth(ctx, &headers, client.Cfg.DelegatedTokenConfig)
		if err != nil && headers["Authorization"] != "" {
			return delegatedTokenTestResult{nil, fmt.Errorf("failed exchange set Authorization: %w", err)}
		}
		if err == nil && headers["Authorization"] != "Bearer fresh" {
			return delegatedTokenTestResult{nil, fmt.Errorf("unexpected Authorization: %q", headers["Authorization"])}
		}
		return delegatedTokenTestResult{nil, err}
	}
}

func waitForDelegatedTokenTest[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for delegated token authentication")
		var zero T
		return zero
	}
}

func TestDelegatedTokenConcurrentExchange(t *testing.T) {
	for _, state := range []string{"first use", "expired"} {
		for _, failure := range []string{"", "error", "nil credentials", "empty token"} {
			t.Run(state+"/"+failure, func(t *testing.T) {
				shared := &datadog.DelegatedTokenCredentials{}
				if state == "expired" {
					shared.DelegatedToken = "expired"
					shared.Expiration = time.Now().Add(-time.Hour)
				}
				ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, shared)
				entered, release := make(chan struct{}, 1), make(chan struct{})
				unblock := sync.OnceFunc(func() { close(release) })
				defer unblock()
				var calls atomic.Int32
				wantErr := errors.New("exchange rejected")
				expiration := time.Now().Add(time.Hour)
				config := datadog.NewConfiguration()
				config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{ProviderAuth: delegatedTokenTestProvider(func(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
					if calls.Add(1) == 1 {
						entered <- struct{}{}
						<-release
						switch failure {
						case "error":
							return nil, wantErr
						case "nil credentials":
							return nil, nil
						case "empty token":
							return &datadog.DelegatedTokenCredentials{}, nil
						}
					}
					return &datadog.DelegatedTokenCredentials{OrgUUID: "org", DelegatedToken: "fresh", DelegatedProof: "proof", Expiration: expiration}, nil
				})}
				client := datadog.NewAPIClient(config)
				const parallel = 16
				results := make(chan delegatedTokenTestResult, parallel)
				go func() { results <- delegatedTokenTestCall(ctx, client, 0) }()
				waitForDelegatedTokenTest(t, entered)
				waiting := make(chan struct{}, parallel)
				for i := 1; i < parallel; i++ {
					go func(method int) {
						results <- delegatedTokenTestCall(delegatedTokenWaitContext{ctx, waiting}, client, method)
					}(i)
				}
				for i := 1; i < parallel; i++ {
					waitForDelegatedTokenTest(t, waiting)
				}
				unblock()
				var snapshots []*datadog.DelegatedTokenCredentials
				for i := 0; i < parallel; i++ {
					result := waitForDelegatedTokenTest(t, results)
					if failure == "" {
						if result.err != nil {
							t.Fatal(result.err)
						}
						if result.creds != nil {
							if result.creds == shared || result.creds.DelegatedToken != "fresh" || result.creds.OrgUUID != "org" || result.creds.DelegatedProof != "proof" || !result.creds.Expiration.Equal(expiration) {
								t.Fatal("expected a consistent, independent credentials snapshot")
							}
							snapshots = append(snapshots, result.creds)
						}
					} else if result.err == nil {
						t.Fatal("failed exchange returned success")
					} else if failure == "error" && !errors.Is(result.err, wantErr) {
						t.Fatalf("error = %v; want %v", result.err, wantErr)
					}
				}
				if calls.Load() != 1 {
					t.Fatalf("got %d exchanges; want 1", calls.Load())
				}
				if failure != "" {
					if state == "expired" && shared.DelegatedToken != "expired" {
						t.Fatal("failure changed cached credentials")
					}
					// A later request can retry; a failed exchange is not cached forever.
					if result := delegatedTokenTestCall(ctx, client, 2); result.err != nil {
						t.Fatal(result.err)
					}
					if calls.Load() != 2 {
						t.Fatalf("got %d exchanges after recovery; want 2", calls.Load())
					}
				}
				before := calls.Load()
				for i := 0; i < parallel; i++ {
					go func() { results <- delegatedTokenTestCall(ctx, client, 2) }()
				}
				for i := 0; i < parallel; i++ {
					if result := waitForDelegatedTokenTest(t, results); result.err != nil {
						t.Fatal(result.err)
					}
				}
				if calls.Load() != before {
					t.Fatal("cached requests exchanged tokens")
				}
				// Returned credentials must neither alias the cache nor each other.
				for _, snapshot := range snapshots {
					if snapshot.DelegatedToken != "fresh" {
						t.Fatal("callers shared a mutable result")
					}
					snapshot.DelegatedToken = "caller mutation"
				}
				if result := delegatedTokenTestCall(ctx, client, 2); result.err != nil {
					t.Fatal(result.err)
				}
			})
		}
	}
}

func TestDelegatedTokenFailedForcedRefreshPreservesValidCache(t *testing.T) {
	shared := &datadog.DelegatedTokenCredentials{DelegatedToken: "fresh", Expiration: time.Now().Add(time.Hour)}
	ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, shared)
	entered, release := make(chan struct{}), make(chan struct{})
	defer close(release)
	config := datadog.NewConfiguration()
	wantErr := errors.New("refresh rejected")
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{ProviderAuth: delegatedTokenTestProvider(func(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
		close(entered)
		<-release
		return nil, wantErr
	})}
	client := datadog.NewAPIClient(config)
	results := make(chan delegatedTokenTestResult, 1)
	go func() { results <- delegatedTokenTestCall(ctx, client, 0) }()
	waitForDelegatedTokenTest(t, entered)
	// A valid cache remains usable even while an explicit refresh is pending.
	if result := delegatedTokenTestCall(ctx, client, 2); result.err != nil {
		t.Fatal(result.err)
	}
	release <- struct{}{}
	if result := waitForDelegatedTokenTest(t, results); !errors.Is(result.err, wantErr) {
		t.Fatalf("error = %v; want %v", result.err, wantErr)
	}
	if result := delegatedTokenTestCall(ctx, client, 2); result.err != nil {
		t.Fatal(result.err)
	}
}

func TestDelegatedTokenWaitingRequestCanCancel(t *testing.T) {
	shared := &datadog.DelegatedTokenCredentials{}
	ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, shared)
	entered, release := make(chan struct{}), make(chan struct{})
	defer close(release)
	config := datadog.NewConfiguration()
	var calls atomic.Int32
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{ProviderAuth: delegatedTokenTestProvider(func(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
		calls.Add(1)
		close(entered)
		<-release
		return &datadog.DelegatedTokenCredentials{DelegatedToken: "fresh", Expiration: time.Now().Add(time.Hour)}, nil
	})}
	client := datadog.NewAPIClient(config)
	leader := make(chan delegatedTokenTestResult, 1)
	go func() { leader <- delegatedTokenTestCall(ctx, client, 0) }()
	waitForDelegatedTokenTest(t, entered)
	waiting := make(chan struct{}, 1)
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	follower := make(chan delegatedTokenTestResult, 1)
	go func() { follower <- delegatedTokenTestCall(delegatedTokenWaitContext{cancelCtx, waiting}, client, 2) }()
	waitForDelegatedTokenTest(t, waiting)
	cancel()
	if result := waitForDelegatedTokenTest(t, follower); !errors.Is(result.err, context.Canceled) {
		t.Fatalf("error = %v; want context.Canceled", result.err)
	}
	release <- struct{}{}
	if result := waitForDelegatedTokenTest(t, leader); result.err != nil {
		t.Fatal(result.err)
	}
	if result := delegatedTokenTestCall(ctx, client, 2); result.err != nil {
		t.Fatal(result.err)
	}
	if calls.Load() != 1 {
		t.Fatal("cancelled waiter triggered another exchange")
	}
}

func TestDelegatedTokenCredentialSetsAreIndependent(t *testing.T) {
	entered, release := make(chan struct{}, 2), make(chan struct{})
	defer close(release)
	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{ProviderAuth: delegatedTokenTestProvider(func(context.Context, *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
		entered <- struct{}{}
		<-release
		return &datadog.DelegatedTokenCredentials{DelegatedToken: "fresh", Expiration: time.Now().Add(time.Hour)}, nil
	})}
	client := datadog.NewAPIClient(config)
	results := make(chan delegatedTokenTestResult, 2)
	for i := 0; i < 2; i++ {
		ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})
		go func() { results <- delegatedTokenTestCall(ctx, client, 2) }()
	}
	for i := 0; i < 2; i++ {
		waitForDelegatedTokenTest(t, entered)
	}
	for i := 0; i < 2; i++ {
		release <- struct{}{}
	}
	for i := 0; i < 2; i++ {
		if result := waitForDelegatedTokenTest(t, results); result.err != nil {
			t.Fatal(result.err)
		}
	}
}

func TestDelegatedTokenInvalidConfigurationFailsClosed(t *testing.T) {
	for _, config := range []*datadog.DelegatedTokenConfig{nil, {}} {
		ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})
		headers := map[string]string{}
		if err := datadog.UseDelegatedTokenAuth(ctx, &headers, config); err == nil {
			t.Fatal("missing provider returned success")
		}
		if len(headers) != 0 {
			t.Fatal("missing provider set an authorization header")
		}
	}
	for _, ctx := range []context.Context{nil, context.Background(), context.WithValue(context.Background(), datadog.ContextDelegatedToken, (*datadog.DelegatedTokenCredentials)(nil))} {
		headers := map[string]string{}
		if err := datadog.UseDelegatedTokenAuth(ctx, &headers, nil); err == nil {
			t.Fatal("missing credentials returned success")
		}
	}
}