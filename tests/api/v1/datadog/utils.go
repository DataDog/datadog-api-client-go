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
	"github.com/jonboulle/clockwork"
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

// SendRequest sends request to endpoints without specification.
func (c *Client) SendRequest(method, url string, payload []byte) (*http.Response, []byte, error) {
	baseURL := ""
	if !strings.HasPrefix(url, "https://") {
		var err error
		baseURL, err = c.Client.GetConfig().ServerURLWithContext(c.Ctx, "")
		if err != nil {
			return nil, []byte{}, fmt.Errorf("Failed to get base URL for Datadog API: %s", err.Error())
		}
	}

	request, err := http.NewRequest(method, baseURL+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Failed to create request for Datadog API: %s", err.Error())
	}
	keys := c.Ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	request.Header.Add("DD-API-KEY", keys["apiKeyAuth"].Key)
	request.Header.Add("DD-APPLICATION-KEY", keys["appKeyAuth"].Key)
	request.Header.Set("Content-Type", "application/json")

	resp, respErr := c.Client.GetConfig().HTTPClient.Do(request)
	body, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		respErr = fmt.Errorf("Failed reading response body: %s", rerr.Error())
	}
	return resp, body, respErr
}

func setupGock(t *testing.T, fixtureFile string, method string, uriPath string) []byte {
	fixturePath, _ := filepath.Abs(fmt.Sprintf("fixtures/%s", fixtureFile))
	dat, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}
	x := gock.New("https://api.datadoghq.com/api/v1")
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
