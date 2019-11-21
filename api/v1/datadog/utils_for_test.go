package datadog_test

import (
	"context"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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
