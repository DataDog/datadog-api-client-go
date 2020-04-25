/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/jonboulle/clockwork"
)

// WithFakeAuth avoids issue of API returning `text/html` instead of `application/json`
func WithFakeAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: "FAKE_KEY",
			},
			"appKeyAuth": {
				Key: "FAKE_KEY",
			},
		},
	)
}

// WithTestAuth returns authenticated context.
func WithTestAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
}

// NewConfiguration return configuration with known options.
func NewConfiguration() *datadog.Configuration {
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	return config
}

// Client keeps track for APIClient and Auth Context
type Client struct {
	Client *datadog.APIClient
	Ctx    context.Context
	Clock  clockwork.FakeClock
	close  *func()
}

// NewClient returns client for unit tests.
func NewClient(ctx context.Context, t *testing.T) *Client {
	ctx, close := tests.WithTestSpan(ctx, t)
	return &Client{Ctx: ctx, Client: datadog.NewAPIClient(NewConfiguration()), close: &close}
}

// NewClientWithRecording returns configured client with recorder.
func NewClientWithRecording(ctx context.Context, t *testing.T) *Client {
	ctx, close := tests.WithTestSpan(ctx, t)
	r, err := tests.Recorder(ctx, t)
	if err != nil {
		log.Fatal(err)
	}
	closeRecording := func() {
		r.Stop()
		close()
	}

	// Create configuration
	config := NewConfiguration()
	config.HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	// Configure client
	c := Client{Ctx: ctx, Client: datadog.NewAPIClient(config), close: &closeRecording}

	// Configure clock
	if tests.IsRecording() {
		c.Clock = tests.SetClock(t)
	} else {
		c.Clock = tests.RestoreClock(t)
	}

	return &c
}

// Close open resources.
func (c *Client) Close() {
	if c.close == nil {
		return
	}
	(*c.close)()
}
