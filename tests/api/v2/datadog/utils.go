/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

// TEST_API_CLIENT is the api client to use for tests
var TestApiClient *datadog.APIClient

// TESTAUTH is the authentication context to use with each test API call
var TestAuth context.Context

func setupTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TestAuth = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": datadog.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": datadog.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	TestApiClient = datadog.NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}

func setupUnitTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TestAuth = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": datadog.APIKey{
				Key: "FAKE_KEY",
			},
			"appKeyAuth": datadog.APIKey{
				Key: "FAKE_KEY",
			},
		},
	)
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	TestApiClient = datadog.NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}
