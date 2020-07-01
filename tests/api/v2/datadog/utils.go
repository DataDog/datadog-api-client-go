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

	// Set test site as default
	testSite, ok := os.LookupEnv("DD_TEST_SITE")
	if ok {
		server := config.Servers[0]
		site := server.Variables["site"]
		site.DefaultValue = testSite
		site.EnumValues = append(site.EnumValues, testSite)
		server.Variables["site"] = site
		config.Servers[0] = server

		for operationID, servers := range config.OperationServers {
			server := servers[0]
			site := server.Variables["site"]
			site.DefaultValue = testSite
			site.EnumValues = append(site.EnumValues, testSite)
			server.Variables["site"] = site
			servers[0] = server
			config.OperationServers[operationID] = servers
		}
	}
	return config
}

type contextKey string

var (
	clientKey = contextKey("client")
)

// WithClient sets client for unit tests in context.
func WithClient(ctx context.Context, t *testing.T) (context.Context, func()) {
	ctx, finish := tests.WithTestSpan(ctx, t)
	ctx = context.WithValue(ctx, clientKey, datadog.NewAPIClient(NewConfiguration()))
	return ctx, finish
}

// ClientFromContext returns client and indication if it was successful.
func ClientFromContext(ctx context.Context) (*datadog.APIClient, bool) {
	if ctx == nil {
		return nil, false
	}
	v := ctx.Value(clientKey)
	if c, ok := v.(*datadog.APIClient); ok {
		return c, true
	}
	return nil, false
}

// Client returns client from context.
func Client(ctx context.Context) *datadog.APIClient {
	c, ok := ClientFromContext(ctx)
	if !ok {
		log.Fatal("client is not configured")
	}
	return c
}

// WithRecorder configures client with recorder.
func WithRecorder(ctx context.Context, t *testing.T) (context.Context, func()) {
	ctx, finish := WithClient(ctx, t)
	client := Client(ctx)

	ctx, err := tests.WithClock(ctx, t.Name())
	if err != nil {
		log.Fatal(err)
	}

	r, err := tests.Recorder(ctx, t.Name())
	if err != nil {
		log.Fatal(err)
	}
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	return ctx, func() {
		r.Stop()
		finish()
	}
}
