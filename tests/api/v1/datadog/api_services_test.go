/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

func TestServicesLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testServiceData := datadog.Service{
		Attributes: &datadog.ServiceAttributes{
			Name: datadog.PtrString("Test-Service"),
		},
		Type: datadog.PtrString("services"),
	}
	testServiceAttributes := testServiceData.GetAttributes()
	// Create Service
	serviceRsp, httpresp, err := TESTAPICLIENT.ServicesApi.CreateService(TESTAUTH).Body(*datadog.NewServicePayload(testServiceData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Service %v: Response %s: %v", testServiceData, err.Error(), err)
	}
	service := serviceRsp.GetData()
	serviceAttrs := service.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the service if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, service datadog.Service) {
		//Delete Service
		httpresp, err := test_client.ServicesApi.DeleteService(TESTAUTH, service.GetId()).Execute()
		if err != nil {
			t.Fatalf("Error deleting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, service)

	assert.Equal(t, service.GetType(), testServiceData.GetType())
	assert.Equal(t, serviceAttrs.GetName(), testServiceAttributes.GetName())

	// Edit Service
	service.Attributes.SetName("Test-Service-Updated")
	serviceUpdatedRsp, httpresp, err := TESTAPICLIENT.ServicesApi.PatchService(TESTAUTH, service.GetId()).Body(*datadog.NewServicePayload(service)).Execute()
	if err != nil {
		t.Fatalf("Error updating Service %v: Response %s: %v", service, err.Error(), err)
	}
	serviceUpdated := serviceUpdatedRsp.GetData()
	serviceAttrsUpdated := serviceUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, serviceAttrsUpdated.GetName(), serviceAttrsUpdated.GetName())

	// Get Service
	serviceGetResp, httpresp, err := TESTAPICLIENT.ServicesApi.GetService(TESTAUTH, service.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	serviceGet := serviceGetResp.GetData()
	serviceAttrsGet := serviceGet.GetAttributes()
	assert.Equal(t, serviceGet.GetId(), serviceUpdated.GetId())
	assert.Equal(t, serviceAttrsGet.GetName(), serviceAttrsUpdated.GetName())

	// Get Services

	servicesGetResp, httpresp, err := TESTAPICLIENT.ServicesApi.GetServices(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting Service %s: Response %s: %v", service.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(servicesGetResp.GetData()) >= 1)

}

var serviceNameCounter int64 = time.Now().UnixNano()

func createTestService(t *testing.T) datadog.Service {
	atomic.AddInt64(&serviceNameCounter, 1)
	testServiceData := datadog.NewService()
	testServiceData.SetType("services")
	testServiceData.SetAttributes(*datadog.NewServiceAttributesWithDefaults())
	testServiceData.Attributes.SetName(fmt.Sprintf("Test-Service-%d", serviceNameCounter))
	// Create Service
	serviceRsp, httpresp, err := TESTAPICLIENT.ServicesApi.CreateService(TESTAUTH).Body(*datadog.NewServicePayload(*testServiceData)).Execute()
	if err != nil {
		t.Fatalf("Error creating Service %v: Response %s: %v", testServiceData, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	return serviceRsp.GetData()
}

func deleteTestService(t *testing.T, id string) {
	httpresp, err := TESTAPICLIENT.ServicesApi.DeleteService(TESTAUTH, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Service: %v failed with %v, Another test may have already deleted this service: %v",
			id, httpresp.StatusCode, err)
	}

}
