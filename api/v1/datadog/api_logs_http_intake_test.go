package datadog_test

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestLogHTTPIntake(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	ctx := TESTAUTH

	var key string
	if auth, ok := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		if auth, ok := auth["api_key"]; ok {
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
		}
	}

	httpLog := datadog.HttpLog{
		Ddsource: datadog.PtrString("source"),
		Ddtags:   datadog.PtrString("tags"),
		Hostname: datadog.PtrString("datadog-api-client-go-test"),
		Message:  datadog.PtrString("testing message"),
	}

	// Create log entry
	_, httpresp, err := TESTAPICLIENT.LogsHTTPIntakeApi.SendLog(TESTAUTH, key, httpLog)
	if err != nil {
		t.Errorf("Error creating log: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}
