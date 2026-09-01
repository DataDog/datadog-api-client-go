package datadog

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetryBackoffIsMinimum(t *testing.T) {
	cfg := NewConfiguration()
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.BackOffBase = 5
	cfg.RetryConfiguration.BackOffMultiplier = 2
	cfg.RetryConfiguration.MaxRetries = 5
	client := NewAPIClient(cfg)

	response := retryResponse(http.StatusTooManyRequests, "1")
	for retryCount, want := range []time.Duration{5 * time.Second, 10 * time.Second, 20 * time.Second, 40 * time.Second, 80 * time.Second} {
		delay, retry := client.shouldRetryRequest(response, retryCount)
		if !retry || *delay != want {
			t.Fatalf("retry %d: delay, retry = %v, %t; want %v, true", retryCount, *delay, retry, want)
		}
	}

	t.Run("longer reset extends delay", func(t *testing.T) {
		delay, retry := client.shouldRetryRequest(retryResponse(http.StatusTooManyRequests, "30"), 0)
		if !retry || *delay != 30*time.Second {
			t.Fatalf("delay, retry = %v, %t; want 30s, true", *delay, retry)
		}
	})

	for name, reset := range map[string]string{
		"missing":   "",
		"malformed": "later",
		"negative":  "-1",
		"overflow":  "9223372037",
	} {
		t.Run(name+" reset uses backoff", func(t *testing.T) {
			delay, retry := client.shouldRetryRequest(retryResponse(http.StatusTooManyRequests, reset), 0)
			if !retry || *delay != 5*time.Second {
				t.Fatalf("delay, retry = %v, %t; want 5s, true", *delay, retry)
			}
		})
	}

	t.Run("server error ignores reset", func(t *testing.T) {
		delay, retry := client.shouldRetryRequest(retryResponse(http.StatusInternalServerError, "30"), 0)
		if !retry || *delay != 5*time.Second {
			t.Fatalf("delay, retry = %v, %t; want 5s, true", *delay, retry)
		}
	})
}

func TestRetryJitter(t *testing.T) {
	tests := []struct {
		name     string
		response *http.Response
		minimum  time.Duration
	}{
		{
			name:     "server error backoff",
			response: retryResponse(http.StatusInternalServerError, ""),
			minimum:  2 * time.Second,
		},
		{
			name:     "backoff floors rate limit reset",
			response: retryResponse(http.StatusTooManyRequests, "1"),
			minimum:  2 * time.Second,
		},
		{
			name:     "rate limit reset extends backoff",
			response: retryResponse(http.StatusTooManyRequests, "3"),
			minimum:  3 * time.Second,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := NewConfiguration()
			cfg.RetryConfiguration.EnableRetry = true
			cfg.RetryConfiguration.RetryJitter = time.Second
			client := NewAPIClient(cfg)

			var observedJitter bool
			for i := 0; i < 10; i++ {
				delay, retry := client.shouldRetryRequest(tc.response, 0)
				if !retry {
					t.Fatal("expected request to be retried")
				}
				if *delay < tc.minimum || *delay >= tc.minimum+time.Second {
					t.Fatalf("retry delay %s outside [%s, %s)", *delay, tc.minimum, tc.minimum+time.Second)
				}
				observedJitter = observedJitter || *delay > tc.minimum
			}
			if !observedJitter {
				t.Fatal("expected a jittered retry delay")
			}
		})
	}
}

func TestRetryDelayOutsideBudgetReturnsImmediately(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		w.Header().Set(rateLimitResetHeader, "10")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.HTTPRetryTimeout = 100 * time.Millisecond
	client := NewAPIClient(cfg)
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	started := time.Now()
	response, err := client.CallAPI(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if elapsed := time.Since(started); elapsed >= time.Second {
		t.Fatalf("request took %s; expected an immediate response", elapsed)
	}
	if calls.Load() != 1 {
		t.Fatalf("calls = %d; want 1", calls.Load())
	}
}

func TestRetryMaxAttempts(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.MaxRetries = 5
	client := NewAPIClient(cfg)
	client.Cfg.RetryConfiguration.BackOffBase = 0 // Keep this attempt-count test instant.
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.CallAPI(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if calls.Load() != 6 {
		t.Fatalf("calls = %d; want 6", calls.Load())
	}
}

func TestRetryDeadlineAllowsInFlightAttemptToFinish(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.HTTPRetryTimeout = 50 * time.Millisecond
	client := NewAPIClient(cfg)
	client.Cfg.RetryConfiguration.BackOffBase = 0 // Start the retry immediately.
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.CallAPI(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d; want %d", response.StatusCode, http.StatusOK)
	}
	if calls.Load() != 2 {
		t.Fatalf("calls = %d; want 2", calls.Load())
	}
}

func retryResponse(statusCode int, reset string) *http.Response {
	header := http.Header{}
	if reset != "" {
		header.Set(rateLimitResetHeader, reset)
	}
	return &http.Response{StatusCode: statusCode, Header: header}
}
