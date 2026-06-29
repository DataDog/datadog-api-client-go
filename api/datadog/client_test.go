package datadog

import (
	"net/http"
	"testing"
)

// TestShouldRetryRequest_ZeroTimeoutProducesPositiveDuration verifies that when
// HTTPClient.Timeout == 0 (Go's "no timeout" sentinel, e.g. http.DefaultClient),
// the retry backoff is NOT clamped to zero. Before the fix, math.Min(0, retryVal)
// always returned 0, causing all retries to fire instantly.
func TestShouldRetryRequest_ZeroTimeoutProducesPositiveDuration(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{} // Timeout is zero by default
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.MaxRetries = 3
	cfg.RetryConfiguration.BackOffBase = 2
	cfg.RetryConfiguration.BackOffMultiplier = 2

	client := NewAPIClient(cfg)

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     http.Header{},
	}

	duration, shouldRetry := client.shouldRetryRequest(resp, 0)
	if !shouldRetry {
		t.Fatal("expected shouldRetry=true for a 500 response")
	}
	if duration == nil || *duration <= 0 {
		t.Fatalf("expected positive retry duration, got %v", duration)
	}
}

// TestShouldRetryRequest_TimeoutCapIsApplied verifies that when a real timeout is
// set, the retry duration is still capped to that timeout value.
func TestShouldRetryRequest_TimeoutCapIsApplied(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: 3 * 1e9} // 3 seconds
	cfg.RetryConfiguration.EnableRetry = true
	cfg.RetryConfiguration.MaxRetries = 3
	cfg.RetryConfiguration.BackOffBase = 2
	cfg.RetryConfiguration.BackOffMultiplier = 10 // would produce 200s without cap

	client := NewAPIClient(cfg)

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     http.Header{},
	}

	duration, shouldRetry := client.shouldRetryRequest(resp, 2)
	if !shouldRetry {
		t.Fatal("expected shouldRetry=true for a 500 response")
	}
	if duration == nil || *duration > 3*1e9 {
		t.Fatalf("expected retry duration <= 3s (the configured timeout), got %v", duration)
	}
}
