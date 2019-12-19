package datadog_test

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestLogHTTPIntake(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	httpLog := datadog.HttpLog{
		Ddsource: datadog.PtrString("source"),
		Ddtags:   datadog.PtrString("tags"),
		Hostname: datadog.PtrString("datadog-api-client-go-test"),
		Message:  datadog.PtrString("testing message"),
	}

	// Create log entry
	_, httpresp, err := TESTAPICLIENT.LogsHTTPIntakeApi.SendLog(TESTAUTH).HttpLog(httpLog).Execute()
	if err != nil {
		t.Errorf("Error creating log: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}
