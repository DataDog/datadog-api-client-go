package datadog_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestConfigurationServers(t *testing.T) {
	configuration := datadog.NewConfiguration()

	testCases := []struct {
		URL       string
		Variables map[string]string
	}{{
		URL:       "https://api.datadoghq.com",
		Variables: nil,
	}, {
		URL:       "https://api.datadoghq.eu",
		Variables: map[string]string{"site": "datadoghq.eu"},
	}, {
		URL: "https://http-intake.logs.datadoghq.eu",
		Variables: map[string]string{
			"subdomain": "http-intake.logs",
			"site":      "datadoghq.eu",
		},
	},
	}

	for _, tc := range testCases {
		t.Run(tc.URL, func(t *testing.T) {
			assert := tests.Assert(context.Background(), t)
			url, err := configuration.ServerURL(0, tc.Variables)
			if err != nil {
				t.Errorf("Could not format URL: %v", err)
			}
			assert.Equal(url, tc.URL)
		})
	}
}

func TestConfigurationServersAccess(t *testing.T) {
	configuration := datadog.NewConfiguration()
	testCases := []struct {
		Index int
		Err   string
	}{{
		Index: -1,
		Err:   "Index -1 out of range 1",
	}, {
		Index: len(configuration.Servers),
		Err:   "Index 2 out of range 1",
	}}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Index %v", tc.Index), func(t *testing.T) {
			assert := tests.Assert(context.Background(), t)
			_, err := configuration.ServerURL(tc.Index, nil)
			assert.Error(err, tc.Err)
		})
	}
}
