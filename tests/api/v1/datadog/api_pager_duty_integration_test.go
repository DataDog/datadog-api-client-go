/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func generatePagerDutyService(ctx context.Context, t *testing.T) datadog.PagerDutyService {
	serviceName := *tests.UniqueEntityName(ctx, t)
	return datadog.PagerDutyService{
		ServiceName: serviceName,
		ServiceKey:  "deadbeef",
	}
}

func TestPagerDutyLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Add single service object to the PagerDuty Integration
	serviceBody := generatePagerDutyService(ctx, t)
	serviceResponse, httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx, serviceBody)
	defer deletePagerDutyService(ctx, t, serviceBody.ServiceName)
	assert.NoError(err)
	assert.Equal(201, httpresp.StatusCode)
	assert.Equal(serviceBody.GetServiceName(), serviceResponse.GetServiceName())

	// Get created Service object
	serviceName, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, serviceBody.GetServiceName())
	assert.NoError(err)
	assert.Equal(serviceBody.GetServiceName(), serviceName.GetServiceName())
	assert.Equal(200, httpresp.StatusCode)

	// Update service object
	serviceKey := datadog.PagerDutyServiceKey{
		ServiceKey: "newkey",
	}
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, serviceBody.GetServiceName(), serviceKey)
	assert.NoError(err)
	assert.Equal(200, httpresp.StatusCode)

	// Delete service object
	httpresp, err = Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, serviceBody.GetServiceName())
	assert.NoError(err)
	assert.Equal(200, httpresp.StatusCode)

	// Check service object
	_, httpresp, err = Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, serviceBody.GetServiceName())
	assert.Error(err)
	assert.Equal(404, httpresp.StatusCode)
}

func TestPagerDutyServicesCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(ctx, t)
	defer finish()

	pgService := generatePagerDutyService(ctx, t)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.PagerDutyService
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.PagerDutyService{}, 400},
		"403 Forbidden":   {WithFakeAuth, pgService, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			service, httpresp, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx, tc.Body)
			if err == nil {
				defer deletePagerDutyService(ctx, t, service.GetServiceName())
			}
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			pgService := generatePagerDutyService(ctx, t)

			_, httpresp, err := Client(ctx).PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, pgService.GetServiceName())
			assert.NotNil(httpresp)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			errors, _ := apiError.GetErrorsOk()
			assert.NotNil(errors)
		})
	}
}

func TestPagerDutyServicesUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	service := generatePagerDutyService(ctx, t)
	_, _, err := Client(ctx).PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx, service)
	defer deletePagerDutyService(ctx, t, service.ServiceName)
	if err != nil {
		t.Errorf("could not create service %v: %v", service, err)
	}

	serviceKey := *datadog.NewPagerDutyServiceKeyWithDefaults()
	serviceKey.SetServiceKey("lalaa")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ServiceName        string
		Body               datadog.PagerDutyServiceKey
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, service.ServiceName, datadog.PagerDutyServiceKey{}, 400},
		"403 Forbidden":   {WithFakeAuth, service.ServiceName, serviceKey, 403},
		"404 Not Found":   {WithTestAuth, service.ServiceName + "-invalid", serviceKey, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, tc.ServiceName, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestPagerDutyServicesDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			pgService := generatePagerDutyService(ctx, t)

			httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, pgService.GetServiceName())
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			errors, _ := apiError.GetErrorsOk()
			assert.NotNil(errors)
		})
	}
}

func deletePagerDutyService(ctx context.Context, t *testing.T, serviceName string) {
	httpresp, err := Client(ctx).PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, serviceName)
	if httpresp.StatusCode != 204 || err != nil {
		t.Logf("Error deleting PagerDuty Service %s: Another test may have already removed this account: %v", serviceName, err)
	}
}
