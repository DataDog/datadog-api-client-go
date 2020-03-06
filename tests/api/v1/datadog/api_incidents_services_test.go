/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func TestIncidentServicesLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())
	SERVICE_1 := createTestService(t)
	defer deleteTestService(t, SERVICE_1.GetId())

	// Create IncidentService
	incidentServiceRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.AddServiceToIncident(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewServicePayload(SERVICE_1)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error adding service to incident %v: Response %s: %v", SERVICE_1, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentService if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentService datadog.Service) {
		//Delete IncidentService
		httpresp, err := test_client.IncidentsApi.
			RemoveServiceFromIncident(TESTAUTH, INCIDENT.GetId(), incidentService.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error removeing service for incident %v: Response %s: %v", incidentService, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, SERVICE_1)

	assert.True(t, len(incidentServiceRsp.GetData()) >= 1)
	assert.NotNil(t, incidentServiceRsp.GetData()[0].GetId())
	assert.Equal(t, incidentServiceRsp.GetData()[0].GetType(), SERVICE_1.GetType())
	assert.Equal(t, incidentServiceRsp.GetData()[0].GetAttributes().Name, SERVICE_1.GetAttributes().Name)

	// Get IncidentServices
	incidentServicesGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetServicesForIncident(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Services for Incident %v: Response %s: %v", SERVICE_1, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentServicesGetResp.GetData()) >= 1)
	assert.NotNil(t, incidentServiceRsp.GetData()[0].GetId())
	assert.Equal(t, incidentServiceRsp.GetData()[0].GetType(), SERVICE_1.GetType())
	assert.Equal(t, incidentServiceRsp.GetData()[0].GetAttributes().Name, SERVICE_1.GetAttributes().Name)

}
