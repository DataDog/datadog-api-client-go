package datadog_v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

// TESTAPICLIENT is the api client to use for tests
var TESTAPICLIENT *APIClient

// TESTAUTH is the authentication context to use with each test API call
var TESTAUTH context.Context

func setupTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		ContextAPIKeys,
		map[string]APIKey{
			"api_key": APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"application_key": APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := NewConfiguration()
	TESTAPICLIENT = NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}

func setupUnitTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		ContextAPIKeys,
		map[string]APIKey{
			"api_key": APIKey{
				Key: "FAKE_KEY",
			},
			"application_key": APIKey{
				Key: "FAKE_KEY",
			},
		},
	)
	config := NewConfiguration()
	TESTAPICLIENT = NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}

func setupGock(t *testing.T, fixtureFile string, method string, uriPath string) []byte {
	fixturePath, _ := filepath.Abs(fmt.Sprintf("fixtures/%s", fixtureFile))
	dat, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}
	switch strings.ToLower(method) {
	case "get":
		gock.New("https://api.datadoghq.com/api/v1").
			Get(uriPath).
			Reply(200).
			JSON(dat)
	case "post":
		gock.New("https://api.datadoghq.com/api/v1").
			Post(uriPath).
			Reply(200).
			JSON(dat)
	case "put":
		gock.New("https://api.datadoghq.com/api/v1").
			Put(uriPath).
			Reply(200).
			JSON(dat)
	case "delete":
		gock.New("https://api.datadoghq.com/api/v1").
			Put(uriPath).
			Reply(200).
			JSON(dat)
	}

	return dat
}
