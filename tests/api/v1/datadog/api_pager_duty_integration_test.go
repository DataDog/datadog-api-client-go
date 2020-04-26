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

func ensureNoPagerDuty(ctx context.Context, t *testing.T) {
	// Make sure that there is not parallel execution
	err := tests.Retry(time.Duration(5)*time.Second, 10, func() bool {
		_, httpresp, _ := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegration(ctx).Execute()
		return httpresp.StatusCode != 200
	})
	assert.Nil(t, err)
}

func TestPagerDutyDelete204(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	ensureNoPagerDuty(ctx, t)

	httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegration(ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestPagerDutyLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	ensureNoPagerDuty(ctx, t)

	pagerDutyRequest := datadog.PagerDutyIntegration{
		Subdomain: datadog.PtrString("_deadbeef"),
		ApiToken:  datadog.PtrString("y_NbAkKc66ryYTWUXYEu"),
	}

	// Create a new PagerDuty integration
	httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegration(ctx).Body(pagerDutyRequest).Execute()
	defer deletePagerDuty(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)

	// Get PagerDuty integration
	pagerDuty, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegration(ctx).Execute()
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
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegration(ctx).Body(servicesAndSchedules).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)

	// Get PagerDuty Integration
	updatedPagerDuty, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegration(ctx).Execute()
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
	serviceNameResponse, httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx).Body(serviceBody).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, serviceBody.GetServiceName(), serviceNameResponse.GetServiceName())

	// Get created Service object
	serviceName, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, "test_go_2").Execute()
	assert.Nil(t, err)
	assert.Equal(t, serviceBody.GetServiceName(), serviceName.GetServiceName())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Get previously added service item
	serviceName, httpresp, err = Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, "test_go").Execute()
	assert.Nil(t, err)
	assert.Equal(t, "test_go", serviceName.GetServiceName())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Update service object
	serviceKey := datadog.PagerDutyServiceKey{
		ServiceKey: "newkey",
	}
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, "test_go_2").Body(serviceKey).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)

	// Delete Service Object
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, "test_go").Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)

	// Get created Service object
	pagerDuty, httpresp, err = Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegration(ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 1, len(pagerDuty.GetSchedules()))
	assert.Equal(t, servicesAndSchedules.GetSchedules()[0], pagerDuty.GetSchedules()[0])
	assert.Equal(t, 1, len(pagerDuty.GetServices()))
	assert.Equal(t, "test_go_2", pagerDuty.GetServices()[0].GetServiceName())
	assert.Equal(t, "*****", pagerDuty.GetServices()[0].GetServiceKey())

	// Delete Pager Duty Integration
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegration(ctx).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestPagerDutyGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			ensureNoPagerDuty(ctx, t)

			_, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegration(ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	pg := *datadog.NewPagerDutyIntegrationWithDefaults()
	pg.SetApiToken("apitoken")
	pg.SetSchedules([]string{"schedule"})
	pg.SetSubdomain("subdomain")
	pg.SetServices([]datadog.PagerDutyService{})

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegration(ctx).Body(pg).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	pg := *datadog.NewPagerDutyServicesAndSchedulesWithDefaults()
	pg.SetSchedules([]string{"schedule"})

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegration(ctx).Body(pg).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegration(ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	pgService := *datadog.NewPagerDutyServiceWithDefaults()
	pgService.SetServiceKey("lalaa")
	pgService.SetServiceName("lalasa")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.PagerDutyService
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.PagerDutyService{}, 400},
		"403 Forbidden":   {WithFakeAuth, pgService, 403},
		"404 Not Found":   {WithTestAuth, pgService, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			ensureNoPagerDuty(ctx, t)

			_, httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	pgService := *datadog.NewPagerDutyServiceWithDefaults()
	pgService.SetServiceKey("lalaa")
	pgService.SetServiceName("lalasa")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, "service").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	pgService := *datadog.NewPagerDutyServiceKeyWithDefaults()
	pgService.SetServiceKey("lalaa")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.PagerDutyServiceKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.PagerDutyServiceKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, pgService, 403},
		"404 Not Found":   {WithTestAuth, pgService, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, "qoisdfhanniq").Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, "lqnioiuyzbefnkje").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func deletePagerDuty(ctx context.Context) {
	httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegration(ctx).Execute()
	if httpresp.StatusCode != 204 || err != nil {
		log.Printf("Error uninstalling PagerDuty integration: Another test may have already removed this account: %v", err)
	}
}
