// Script to paginate all containers and print the total count.
// Resumes from the last successful page cursor on network errors (up to maxFailures
// consecutive failures). Rate-limit responses are waited out and retried for free.
//
// Env vars:
//   DD_API_KEY  - Datadog API key
//   DD_APP_KEY  - Datadog app key
//   DD_HOST     - Base URL (default: https://app.datad0g.com)
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

const (
	pageSize    = int32(1000)
	maxFailures = 5
)

var retryDelays = []time.Duration{1 * time.Second, 5 * time.Second, 10 * time.Second, 20 * time.Second, 30 * time.Second}

func main() {
	host := os.Getenv("DD_HOST")
	if host == "" {
		host = "https://app.datad0g.com"
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, datadog.ContextAPIKeys, map[string]datadog.APIKey{
		"apiKeyAuth": {Key: os.Getenv("DD_API_KEY")},
		"appKeyAuth": {Key: os.Getenv("DD_APP_KEY")},
	})

	cfg := datadog.NewConfiguration()
	cfg.Servers = datadog.ServerConfigurations{{URL: host}}
	api := datadogV2.NewContainersApi(datadog.NewAPIClient(cfg))

	fmt.Printf("Fetching containers from %s (page_size=%d)...\n", host, pageSize)

	total := 0
	for range fetchWithRetry(ctx, api) {
		total++
		fmt.Printf("\rTotal so far: %d", total)
	}

	fmt.Printf("\n\nDone. Total containers: %d\n", total)
}

// fetchWithRetry returns a channel of ContainerItems spanning all pages.
// On a network error it retries the same page (cursor hasn't advanced) with
// incremental backoff. Rate-limit responses are waited out without counting
// as a failure. The channel is closed when all items have been delivered or
// maxFailures consecutive errors occur.
func fetchWithRetry(ctx context.Context, api *datadogV2.ContainersApi) <-chan datadogV2.ContainerItem {
	ch := make(chan datadogV2.ContainerItem, pageSize)
	go func() {
		defer close(ch)
		var cursor *string
		failures := 0
		for {
			params := datadogV2.NewListContainersOptionalParameters().WithPageSize(pageSize)
			if cursor != nil {
				params = params.WithPageCursor(*cursor)
			}

			resp, httpResp, err := api.ListContainers(ctx, *params)
			if err != nil {
				// Rate limited: wait for reset, then retry without counting as a failure.
				if httpResp != nil && httpResp.StatusCode == http.StatusTooManyRequests {
					wait := rateLimitWait(httpResp)
					fmt.Printf("\n[RATE LIMITED] waiting %v for reset...\n", wait)
					time.Sleep(wait)
					continue
				}

				// Network / other error: incremental backoff from the same cursor.
				failures++
				if failures > maxFailures {
					fmt.Printf("\n[FATAL] %d consecutive failures, giving up: %v\n", maxFailures, err)
					return
				}
				delay := retryDelays[failures-1]
				fmt.Printf("\n[ERROR] failure %d/%d, retrying in %v: %v\n", failures, maxFailures, delay, err)
				time.Sleep(delay)
				continue
			}
			failures = 0

			data, ok := resp.GetDataOk()
			if !ok {
				return
			}
			for _, item := range *data {
				select {
				case ch <- item:
				case <-ctx.Done():
					return
				}
			}
			if len(*data) == 0 {
				return
			}

			// Advance cursor only after a successful page.
			meta, ok := resp.GetMetaOk()
			if !ok {
				return
			}
			pg, ok := meta.GetPaginationOk()
			if !ok {
				return
			}
			next, ok := pg.GetNextCursorOk()
			if !ok {
				return
			}
			cursor = next
		}
	}()
	return ch
}

// rateLimitWait returns the duration to wait based on X-RateLimit-Reset, defaulting to 60s.
func rateLimitWait(resp *http.Response) time.Duration {
	if secs, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Reset")); err == nil && secs > 0 {
		return time.Duration(secs) * time.Second
	}
	return 60 * time.Second
}
