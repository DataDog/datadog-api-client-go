package datadog_test

import (
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestConfigurationServers(t *testing.T) {
	configuration := datadog.NewConfiguration()

	testCases := []struct {
		Url       string
		Variables map[string]string
	}{{
		Url:       "https://api.datadoghq.com",
		Variables: nil,
	}, {
		Url:       "https://api.datadoghq.eu",
		Variables: map[string]string{"site": "datadoghq.eu"},
	}, {
		Url: "https://http-intake.logs.datadoghq.eu",
		Variables: map[string]string{
			"subdomain": "http-intake.logs",
			"site":      "datadoghq.eu",
		},
	},
	}

	for _, tc := range testCases {
		t.Run(tc.Url, func(t *testing.T) {
			url, err := configuration.ServerURL(0, tc.Variables)
			if err != nil {
				t.Errorf("Could not format URL: %v", err)
			}
			assert.Equal(t, url, tc.Url)
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
		Err:   "Index -1 out of range 0",
	}, {
		Index: len(configuration.Servers),
		Err:   "Index 1 out of range 0",
	},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Index %v", tc.Index), func(t *testing.T) {
			_, err := configuration.ServerURL(tc.Index, nil)
			assert.Error(t, err, tc.Err)
		})
	}
}
