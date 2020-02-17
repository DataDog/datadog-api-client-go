/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"gotest.tools/assert"
)

func ensureNoPagerDuty(t *testing.T) {
	// Make sure that there is not parallel execution
	err := tests.Retry(time.Duration(5)*time.Second, 10, func() bool {
		_, httpresp, _ := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
		return httpresp.StatusCode != 200
	})
	assert.NilError(t, err)
}

func TestPagerDutyGet404(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	ensureNoPagerDuty(t)

	_, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.ErrorContains(t, err, "404 Not Found")
	assert.Equal(t, httpresp.StatusCode, 404)
}

func TestPagerDutyDelete204(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	ensureNoPagerDuty(t)

	_, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)
}

func TestPagerDutyLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	ensureNoPagerDuty(t)

	pagerDutyRequest := datadog.PagerDuty{
		Subdomain: datadog.PtrString("_deadbeef"),
		ApiToken:  datadog.PtrString("y_NbAkKc66ryYTWUXYEu"),
	}

	httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.CreatePagerDutyIntegration(TESTAUTH).Body(pagerDutyRequest).Execute()
	defer deletePagerDuty()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)

	pagerDuty, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)

	servicesAndSchedules := datadog.PagerDutyServicesAndSchedules{
		Services: &[]datadog.PagerDutyService{{
			ServiceName: "test_go",
			ServiceKey:  "deadbeef",
		}},
		Schedules: &[]string{"https://_deadbeef.pagerduty.com/schedules#DEAD3F"},
	}
	httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.UpdatePagerDutyIntegration(TESTAUTH).Body(servicesAndSchedules).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)

	updatedPagerDuty, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, pagerDuty.GetSubdomain(), updatedPagerDuty.GetSubdomain())
	for index, service := range updatedPagerDuty.GetServices() {
		assert.Equal(t, service.GetServiceName(), servicesAndSchedules.GetServices()[index].GetServiceName())
	}
}

func deletePagerDuty() {
	_, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegration(TESTAUTH).Execute()
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error uninstalling PagerDuty integration: Another test may have already removed this account: %v", err)
	}
}
