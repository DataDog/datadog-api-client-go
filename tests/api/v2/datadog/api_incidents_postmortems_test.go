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

func TestIncidentPostmortemsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(ctx, t, client)
	defer deleteTestIncident(ctx, t, client, incident.GetId())

	testIncidentPostmortemData := datadog.NewIncidentPostmortemWithDefaults()
	testIncidentPostmortemData.SetAttributes(*datadog.NewIncidentPostmortemAttributesWithDefaults())
	testIncidentPostmortemData.Attributes.SetIncidentId(incident.GetId())

	// Create IncidentPostmortem
	incidentPostmortemRsp, httpresp, err := client.IncidentsApi.CreateIncidentPostmortem(ctx).
		Body(*datadog.NewIncidentPostmortemRequest(*testIncidentPostmortemData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentPostmortem %v: Response %s: %v", testIncidentPostmortemData, err.Error(), bStr)
	}
	incidentPostmortem := incidentPostmortemRsp.GetData()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentPostmortem if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentPostmortem datadog.IncidentPostmortem) {
		//Delete IncidentPostmortem
		httpresp, err := test_client.IncidentsApi.DeleteIncidentPostmortem(ctx, incidentPostmortem.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentPostmortem)

	assert.Equal(incidentPostmortem.GetType(), testIncidentPostmortemData.GetType())
	// Edit IncidentPostmortem
	incidentPostmortemUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchIncidentPostmortem(ctx, incidentPostmortem.GetId()).
		Body(*datadog.NewIncidentPostmortemRequest(incidentPostmortem)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	incidentPostmortemUpdated := incidentPostmortemUpdatedRsp.GetData()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentPostmortemUpdated.GetId(), incidentPostmortem.GetId())

	// Get IncidentPostmortem
	incidentPostmortemGetResp, httpresp, err := client.IncidentsApi.
		GetIncidentPostmortem(ctx, incidentPostmortem.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	incidentPostmortemGet := incidentPostmortemGetResp.GetData()
	assert.Equal(incidentPostmortemGet.GetId(), incidentPostmortemUpdated.GetId())

	// Get IncidentPostmortems
	incidentPostmortemsGetResp, httpresp, err := client.IncidentsApi.GetIncidentPostmortems(ctx).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentPostmortems %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentPostmortemsGetResp.GetData()) >= 1)

}
