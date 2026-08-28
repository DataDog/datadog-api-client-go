package datadog

import (
	"net/http"
	"testing"
	"time"
)

func TestRetryJitter(t *testing.T) {
	tests := []struct {
		name     string
		response *http.Response
		minimum  time.Duration
	}{
		{
			name:     "server error backoff",
			response: &http.Response{StatusCode: http.StatusInternalServerError, Header: http.Header{}},
			minimum:  2 * time.Second,
		},
		{
			name: "rate limit reset",
			response: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{rateLimitResetHeader: []string{"1"}},
			},
			minimum: time.Second,
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
