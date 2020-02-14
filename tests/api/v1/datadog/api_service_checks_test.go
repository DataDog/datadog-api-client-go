/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func createServiceCheck(t *testing.T) datadog.ServiceCheck {
	now := time.Now().Unix()
	return datadog.ServiceCheck{
		Check:    t.Name(),
		HostName: datadog.PtrString(fmt.Sprintf("go-client-test-host-%d", now)),
		Message:  datadog.PtrString("Go client integration test is running"),
		Status:   datadog.SERVICECHECKSTATUS_OK,
		Tags: &[]string{
			"test",
			"client:go",
			t.Name(),
		},
		Timestamp: datadog.PtrInt64(now),
	}
}

func TestServiceCheck(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	serviceCheck := createServiceCheck(t)
	// Submit a service check
	status, httpresp, err := TESTAPICLIENT.ServiceChecksApi.SubmitServiceCheck(TESTAUTH).Body([]datadog.ServiceCheck{serviceCheck}).Execute()
	if err != nil {
		t.Fatalf("Error creating Service Check %v: Response %s: %v", serviceCheck, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 202, httpresp.StatusCode)
	assert.Equal(t, "ok", status.GetStatus())
}
