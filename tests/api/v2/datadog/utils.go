/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonboulle/clockwork"
)

// FakeAuth avoids issue of API returning `text/html` instead of `application/json`
var FakeAuth = context.WithValue(
	context.Background(),
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

func createWithDir(path string) (*os.File, error) {
	dirName := filepath.Dir(path)
	_, err := os.Stat(dirName)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

func setClock(t *testing.T) clockwork.FakeClock {
	os.MkdirAll("cassettes", 0755)

	f, err := createWithDir(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not set clock: %v", err)
	}
	defer f.Close()
	now := clockwork.NewRealClock().Now()
	f.WriteString(now.Format(time.RFC3339Nano))
	return clockwork.NewFakeClockAt(now)
}

func restoreClock(t *testing.T) clockwork.FakeClock {
	data, err := ioutil.ReadFile(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not load clock: %v", err)
	}
	now, err := time.Parse(time.RFC3339Nano, string(data))
	if err != nil {
		t.Fatalf("Could not parse clock date: %v", err)
	}
	return clockwork.NewFakeClockAt(now)
}

func removeURLSecrets(u *url.URL) *url.URL {
	query := u.Query()
	query.Del("api_key")
	query.Del("application_key")
	u.RawQuery = query.Encode()
	return u
}

// NewConfiguration return configuration with known options.
func NewConfiguration() *datadog.Configuration {
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	return config
}

// NewClientAuthContext returns authenticated context.
func NewClientAuthContext() context.Context {
	return context.WithValue(
		context.Background(),
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

// Client keeps track for APIClient and Auth Context
type Client struct {
	Client *datadog.APIClient
	Ctx    context.Context
	Clock  clockwork.FakeClock
	close  *func()
}

// NewClient returns client for unit tests.
func NewClient() *Client {
	c := Client{}
	c.Ctx = FakeAuth
	c.Client = datadog.NewAPIClient(NewConfiguration())
	return &c
}

// NewClientWithRecording returns configured client with recorder.
func NewClientWithRecording(t *testing.T) *Client {
	// Configure recorder
	var mode recorder.Mode
	if os.Getenv("RECORD") == "true" {
		mode = recorder.ModeRecording
	} else {
		mode = recorder.ModeReplaying
	}
	r, err := recorder.NewAsMode(fmt.Sprintf("cassettes/%s", t.Name()), mode, nil)
	if err != nil {
		log.Fatal(err)
	}
	close := func() { r.Stop() }

	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		return r.Method == i.Method && removeURLSecrets(r.URL).String() == i.URL
	})

	r.AddFilter(func(i *cassette.Interaction) error {
		u, err := url.Parse(i.URL)
		if err != nil {
			return err
		}
		i.URL = removeURLSecrets(u).String()
		i.Request.Headers.Del("Dd-Api-Key")
		i.Request.Headers.Del("Dd-Application-Key")
		return nil
	})

	// Create configuration
	config := NewConfiguration()
	config.HTTPClient = &http.Client{
		Transport: r, // Inject as transport!
	}

	// Configure client
	c := Client{}
	c.Ctx = NewClientAuthContext()
	c.Client = datadog.NewAPIClient(config)
	c.close = &close

	// Configure clock
	if os.Getenv("RECORD") == "true" {
		c.Clock = setClock(t)
	} else {
		c.Clock = restoreClock(t)
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
