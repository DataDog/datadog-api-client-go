package datadog_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	gock "gopkg.in/h2non/gock.v1"
)

// TESTAPICLIENT is the api client to use for tests
var TESTAPICLIENT *datadog.APIClient

// TESTAUTH is the authentication context to use with each test API call
var TESTAUTH context.Context

func setupTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"api_key": datadog.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"application_key": datadog.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadog.NewConfiguration()
	TESTAPICLIENT = datadog.NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}

func setupUnitTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"api_key": datadog.APIKey{
				Key: "FAKE_KEY",
			},
			"application_key": datadog.APIKey{
				Key: "FAKE_KEY",
			},
		},
	)
	config := datadog.NewConfiguration()
	TESTAPICLIENT = datadog.NewAPIClient(config)
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
