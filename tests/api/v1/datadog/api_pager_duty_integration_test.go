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

	httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)
}

func TestPagerDutyLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	ensureNoPagerDuty(t)

	pagerDutyRequest := datadog.PagerDutyIntegration{
		Subdomain: datadog.PtrString("_deadbeef"),
		ApiToken:  datadog.PtrString("y_NbAkKc66ryYTWUXYEu"),
	}

	// Create a new PagerDuty integration
	httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.CreatePagerDutyIntegration(TESTAUTH).Body(pagerDutyRequest).Execute()
	defer deletePagerDuty()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)

	// Get PagerDuty integration
	pagerDuty, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pagerDutyRequest.GetSubdomain(), pagerDuty.GetSubdomain())
	assert.Equal(t, "*****", pagerDuty.GetApiToken())
	assert.Equal(t, 0, len(pagerDuty.GetServices()))
	assert.Equal(t, 0, len(pagerDuty.GetSchedules()))

	servicesAndSchedules := datadog.PagerDutyServicesAndSchedules{
		Services: &[]datadog.PagerDutyService{{
			ServiceName: "test_go",
			ServiceKey:  "deadbeef",
		}},
		Schedules: &[]string{"https://_deadbeef.pagerduty.com/schedules#DEAD3F"},
	}
	// Add a Service and Schedules items by updating the PagerDuty Integration
	httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.UpdatePagerDutyIntegration(TESTAUTH).Body(servicesAndSchedules).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)

	// Get PagerDuty Integration
	updatedPagerDuty, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pagerDutyRequest.GetSubdomain(), updatedPagerDuty.GetSubdomain())
	assert.Equal(t, "*****", updatedPagerDuty.GetApiToken())
	assert.Equal(t, 1, len(updatedPagerDuty.GetSchedules()))
	assert.Equal(t, servicesAndSchedules.GetSchedules()[0], updatedPagerDuty.GetSchedules()[0])
	assert.Equal(t, 1, len(updatedPagerDuty.GetServices()))
	assert.Equal(t, servicesAndSchedules.GetServices()[0].GetServiceName(), updatedPagerDuty.GetServices()[0].GetServiceName())
	assert.Equal(t, "*****", updatedPagerDuty.GetServices()[0].GetServiceKey())

	// Add single service object to the PagerDuty Integration
	serviceBody := datadog.PagerDutyService{
		ServiceKey: "deadbeef",
		ServiceName: "test_java_2",
	}
	serviceNameResponse, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(TESTAUTH).Body(serviceBody).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 201)
	assert.Equal(t, serviceBody.GetServiceName(), serviceNameResponse.GetServiceName())

	// Get created Service object
	serviceName, httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(TESTAUTH, "test_java_2").Execute()
	assert.NilError(t, err)
	assert.Equal(t, serviceBody.GetServiceName(), serviceName.GetServiceName())
	assert.Equal(t, httpresp.StatusCode, 200)

	// Get previously added service item
	serviceName, httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(TESTAUTH, "test_java").Execute()
	assert.NilError(t, err)
	assert.Equal(t, "test_java", serviceName.GetServiceName())
	assert.Equal(t, httpresp.StatusCode, 200)

	// Update service object
	serviceKey := datadog.PagerDutyServiceKey{
		ServiceKey: "newkey",
	}
	httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(TESTAUTH, "test_java_2").Body(serviceKey).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)

	// Delete Service Object
	httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(TESTAUTH, "test_java").Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)

	// Get created Service object
	pagerDuty, httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.GetPagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, 1, len(pagerDuty.GetSchedules()))
	assert.Equal(t, servicesAndSchedules.GetSchedules()[0], pagerDuty.GetSchedules()[0])
	assert.Equal(t, 1, len(pagerDuty.GetServices()))
	assert.Equal(t, "test_java_2", pagerDuty.GetServices()[0].GetServiceName())
	assert.Equal(t, "*****", pagerDuty.GetServices()[0].GetServiceKey())

	// Delete Pager Duty Integration
	httpresp, err = TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegration(TESTAUTH).Execute()
	assert.NilError(t, err)
	assert.Equal(t, httpresp.StatusCode, 204)
}

func deletePagerDuty() {
	httpresp, err := TESTAPICLIENT.PagerDutyIntegrationApi.DeletePagerDutyIntegration(TESTAUTH).Execute()
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error uninstalling PagerDuty integration: Another test may have already removed this account: %v", err)
	}
}
