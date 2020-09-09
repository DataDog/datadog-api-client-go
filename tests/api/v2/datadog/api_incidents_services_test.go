// /*
//  * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
//  * This product includes software developed at Datadog (https://www.datadoghq.com/).
//  * Copyright 2019-Present Datadog, Inc.
//  */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestIncidentServicesLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(ctx, t, client)
	defer deleteTestIncident(ctx, t, client, incident.GetId())
	service1 := createTestService(ctx, t, client, "generated-test-client-1")
	defer deleteTestService(ctx, t, client, service1.GetId())

	// Create IncidentService
	incidentServiceRsp, httpresp, err := client.IncidentsApi.AddServiceToIncident(ctx, incident.GetId()).
		Body(*datadog.NewServiceRequest(service1)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error adding service to incident %v: Response %s: %v", service1, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentService if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentService datadog.Service) {
		//Delete IncidentService
		httpresp, err := test_client.IncidentsApi.
			RemoveServiceFromIncident(ctx, incident.GetId(), incidentService.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error removeing service for incident %v: Response %s: %v", incidentService, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, service1)

	assert.True(len(incidentServiceRsp.GetData()) >= 1)
	assert.NotNil(incidentServiceRsp.GetData()[0].GetId())
	assert.Equal(incidentServiceRsp.GetData()[0].GetType(), service1.GetType())
	assert.Equal(incidentServiceRsp.GetData()[0].GetAttributes().Name, service1.GetAttributes().Name)

	// Get IncidentServices
	incidentServicesGetResp, httpresp, err := client.IncidentsApi.GetServicesForIncident(ctx, incident.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Services for Incident %v: Response %s: %v", service1, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentServicesGetResp.GetData()) >= 1)
	assert.NotNil(incidentServiceRsp.GetData()[0].GetId())
	assert.Equal(incidentServiceRsp.GetData()[0].GetType(), service1.GetType())
	assert.Equal(incidentServiceRsp.GetData()[0].GetAttributes().Name, service1.GetAttributes().Name)

}
