package datadog_v1

import (
	"context"
	"os"
	"testing"
)

// TESTAPICLIENT is the api client to use for tests
var TESTAPICLIENT *APIClient

// TESTAUTH is the authentication context to use with each test API call
var TESTAUTH context.Context

func setupTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		ContextAPIKey,
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
