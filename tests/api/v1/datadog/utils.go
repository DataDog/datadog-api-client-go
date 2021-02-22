/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"
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
func WithClient(ctx context.Context) context.Context {
	return context.WithValue(ctx, clientKey, datadog.NewAPIClient(NewConfiguration()))
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
	ctx = WithClient(ctx)
	client := Client(ctx)

	ctx, err := tests.WithClock(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		log.Fatal(err)
	}

	r, err := tests.Recorder(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		log.Fatal(err)
	}
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	return ctx, func() {
		r.Stop()
	}
}

// SendRequest sends request to endpoints without specification.
func SendRequest(ctx context.Context, method, url string, payload []byte) (*http.Response, []byte, error) {
	baseURL := ""
	if !strings.HasPrefix(url, "https://") {
		var err error
		baseURL, err = Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
		if err != nil {
			return nil, []byte{}, fmt.Errorf("Failed to get base URL for Datadog API: %s", err.Error())
		}
	}

	request, err := http.NewRequest(method, baseURL+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Failed to create request for Datadog API: %s", err.Error())
	}
	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	request.Header.Add("DD-API-KEY", keys["apiKeyAuth"].Key)
	request.Header.Add("DD-APPLICATION-KEY", keys["appKeyAuth"].Key)
	request.Header.Set("Content-Type", "application/json")

	resp, respErr := Client(ctx).GetConfig().HTTPClient.Do(request)
	body, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		respErr = fmt.Errorf("Failed reading response body: %s", rerr.Error())
	}
	return resp, body, respErr
}

func setupGock(ctx context.Context, t *testing.T, fixtureFile string, method string, uriPath string) []byte {
	fixturePath, _ := filepath.Abs(fmt.Sprintf("fixtures/%s", fixtureFile))
	dat, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	if err != nil {
		t.Errorf("Failed to generate URL: %s", err)
	}
	x := gock.New(URL)
	switch strings.ToLower(method) {
	case "get":
		x.Get(uriPath)
	case "post":
		x.Post(uriPath)
	case "put":
		x.Put(uriPath)
	case "delete":
		x.Put(uriPath)
	}

	x.Reply(200).JSON(dat)
	return dat
}
