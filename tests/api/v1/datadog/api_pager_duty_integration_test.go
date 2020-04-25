/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
)

func ensureNoPagerDuty(t *testing.T) {
	// Make sure that there is not parallel execution
	err := tests.Retry(time.Duration(5)*time.Second, 10, func() bool {
		_, httpresp, _ := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegration(c.Ctx).Execute()
		return httpresp.StatusCode != 200
	})
	assert.Nil(t, err)
}

func TestPagerDutyDelete204(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	ensureNoPagerDuty(t)

	httpresp, err := c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegration(c.Ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestPagerDutyLifecycle(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	ensureNoPagerDuty(t)

	pagerDutyRequest := datadog.PagerDutyIntegration{
		Subdomain: datadog.PtrString("_deadbeef"),
		ApiToken:  datadog.PtrString("y_NbAkKc66ryYTWUXYEu"),
	}

	// Create a new PagerDuty integration
	httpresp, err := c.Client.PagerDutyIntegrationApi.CreatePagerDutyIntegration(c.Ctx).Body(pagerDutyRequest).Execute()
	defer deletePagerDuty()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)

	// Get PagerDuty integration
	pagerDuty, httpresp, err := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegration(c.Ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)
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
	httpresp, err = c.Client.PagerDutyIntegrationApi.UpdatePagerDutyIntegration(c.Ctx).Body(servicesAndSchedules).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)

	// Get PagerDuty Integration
	updatedPagerDuty, httpresp, err := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegration(c.Ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, pagerDutyRequest.GetSubdomain(), updatedPagerDuty.GetSubdomain())
	assert.Equal(t, "*****", updatedPagerDuty.GetApiToken())
	assert.Equal(t, 1, len(updatedPagerDuty.GetSchedules()))
	assert.Equal(t, servicesAndSchedules.GetSchedules()[0], updatedPagerDuty.GetSchedules()[0])
	assert.Equal(t, 1, len(updatedPagerDuty.GetServices()))
	assert.Equal(t, servicesAndSchedules.GetServices()[0].GetServiceName(), updatedPagerDuty.GetServices()[0].GetServiceName())
	assert.Equal(t, "*****", updatedPagerDuty.GetServices()[0].GetServiceKey())

	// Add single service object to the PagerDuty Integration
	serviceBody := datadog.PagerDutyService{
		ServiceKey:  "deadbeef",
		ServiceName: "test_go_2",
	}
	serviceNameResponse, httpresp, err := c.Client.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(c.Ctx).Body(serviceBody).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, serviceBody.GetServiceName(), serviceNameResponse.GetServiceName())

	// Get created Service object
	serviceName, httpresp, err := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(c.Ctx, "test_go_2").Execute()
	assert.Nil(t, err)
	assert.Equal(t, serviceBody.GetServiceName(), serviceName.GetServiceName())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Get previously added service item
	serviceName, httpresp, err = c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(c.Ctx, "test_go").Execute()
	assert.Nil(t, err)
	assert.Equal(t, "test_go", serviceName.GetServiceName())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Update service object
	serviceKey := datadog.PagerDutyServiceKey{
		ServiceKey: "newkey",
	}
	httpresp, err = c.Client.PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(c.Ctx, "test_go_2").Body(serviceKey).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)

	// Delete Service Object
	httpresp, err = c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(c.Ctx, "test_go").Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)

	// Get created Service object
	pagerDuty, httpresp, err = c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegration(c.Ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 1, len(pagerDuty.GetSchedules()))
	assert.Equal(t, servicesAndSchedules.GetSchedules()[0], pagerDuty.GetSchedules()[0])
	assert.Equal(t, 1, len(pagerDuty.GetServices()))
	assert.Equal(t, "test_go_2", pagerDuty.GetServices()[0].GetServiceName())
	assert.Equal(t, "*****", pagerDuty.GetServices()[0].GetServiceKey())

	// Delete Pager Duty Integration
	httpresp, err = c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegration(c.Ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestPagerDutyGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	ensureNoPagerDuty(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", c.Ctx, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegration(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	pg := *datadog.NewPagerDutyIntegrationWithDefaults()
	pg.SetApiToken("apitoken")
	pg.SetSchedules([]string{"schedule"})
	pg.SetSubdomain("subdomain")
	pg.SetServices([]datadog.PagerDutyService{})

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := c.Client.PagerDutyIntegrationApi.CreatePagerDutyIntegration(tc.Ctx).Body(pg).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	pg := *datadog.NewPagerDutyServicesAndSchedulesWithDefaults()
	pg.SetSchedules([]string{"schedule"})

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := c.Client.PagerDutyIntegrationApi.UpdatePagerDutyIntegration(tc.Ctx).Body(pg).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegration(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	ensureNoPagerDuty(t)

	pgService := *datadog.NewPagerDutyServiceWithDefaults()
	pgService.SetServiceKey("lalaa")
	pgService.SetServiceName("lalasa")

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.PagerDutyService
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.PagerDutyService{}, 400},
		{"403 Forbidden", fake_auth, pgService, 403},
		{"404 Not Found", c.Ctx, pgService, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	pgService := *datadog.NewPagerDutyServiceWithDefaults()
	pgService.SetServiceKey("lalaa")
	pgService.SetServiceName("lalasa")

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", c.Ctx, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(tc.Ctx, "service").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	pgService := *datadog.NewPagerDutyServiceKeyWithDefaults()
	pgService.SetServiceKey("lalaa")

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.PagerDutyServiceKey
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.PagerDutyServiceKey{}, 400},
		{"403 Forbidden", fake_auth, pgService, 403},
		{"404 Not Found", c.Ctx, pgService, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := c.Client.PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(tc.Ctx, "qoisdfhanniq").Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", c.Ctx, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(tc.Ctx, "lqnioiuyzbefnkje").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func deletePagerDuty() {
	httpresp, err := c.Client.PagerDutyIntegrationApi.DeletePagerDutyIntegration(c.Ctx).Execute()
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error uninstalling PagerDuty integration: Another test may have already removed this account: %v", err)
	}
}
