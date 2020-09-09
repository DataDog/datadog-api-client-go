/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestServicesLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	testServiceData := datadog.NewService()
	testServiceData.SetAttributes(datadog.ServiceAttributes{
		Name: datadog.PtrString("TestServicesLifecycle-generated-client-service"),
	})

	testServiceAttributes := testServiceData.GetAttributes()
	// Create Service
	serviceRsp, httpresp, err := client.ServicesApi.CreateService(ctx).Body(*datadog.NewServiceRequest(*testServiceData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Service %v: Response %s: %v", testServiceData, err.Error(), err)
	}
	service := serviceRsp.GetData()
	serviceAttrs := service.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the service if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, service datadog.Service) {
		//Delete Service
		httpresp, err := test_client.ServicesApi.DeleteService(ctx, service.GetId()).Execute()
		if err != nil {
			t.Fatalf("Error deleting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, service)

	assert.Equal(service.GetType(), testServiceData.GetType())
	assert.Equal(serviceAttrs.GetName(), testServiceAttributes.GetName())

	// Edit Service
	service.Attributes.SetName("Test-Service-Updated")
	serviceUpdatedRsp, httpresp, err := client.ServicesApi.PatchService(ctx, service.GetId()).Body(*datadog.NewServiceRequest(service)).Execute()
	if err != nil {
		t.Fatalf("Error updating Service %v: Response %s: %v", service, err.Error(), err)
	}
	serviceUpdated := serviceUpdatedRsp.GetData()
	serviceAttrsUpdated := serviceUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(serviceAttrsUpdated.GetName(), serviceAttrsUpdated.GetName())

	// Get Service
	serviceGetResp, httpresp, err := client.ServicesApi.GetService(ctx, service.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	serviceGet := serviceGetResp.GetData()
	serviceAttrsGet := serviceGet.GetAttributes()
	assert.Equal(serviceGet.GetId(), serviceUpdated.GetId())
	assert.Equal(serviceAttrsGet.GetName(), serviceAttrsUpdated.GetName())

	// Get Services

	servicesGetResp, httpresp, err := client.ServicesApi.GetServices(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(servicesGetResp.GetData()) >= 1)

}

func createTestService(ctx context.Context, t *testing.T, client *datadog.APIClient, name string) datadog.Service {
	testServiceData := datadog.NewService()
	testServiceData.SetType("services")
	testServiceData.SetAttributes(*datadog.NewServiceAttributesWithDefaults())
	testServiceData.Attributes.SetName(fmt.Sprintf("createTestService-%s", name))

	// Create Service
	serviceRsp, httpresp, err := client.ServicesApi.CreateService(ctx).Body(*datadog.NewServiceRequest(*testServiceData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Service %v: Response %s: %v", testServiceData, err.Error(), err)
	}
	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return serviceRsp.GetData()
}

func deleteTestService(ctx context.Context, t *testing.T, client *datadog.APIClient, id string) {
	httpresp, err := client.ServicesApi.DeleteService(ctx, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Service: %v failed with %v, Another test may have already deleted this service: %v",
			id, httpresp.StatusCode, err)
	}

}
