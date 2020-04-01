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

func TestIncidentPostmortemsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	SetUnstableIncidentsAPIs(true)
	defer SetUnstableIncidentsAPIs(false)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())

	testIncidentPostmortemData := datadog.NewIncidentPostmortemWithDefaults()
	testIncidentPostmortemData.SetType("incident_postmortems")
	testIncidentPostmortemData.SetAttributes(*datadog.NewIncidentPostmortemAttributesWithDefaults())
	testIncidentPostmortemData.Attributes.SetState("not_started")

	// Create IncidentPostmortem
	incidentPostmortemRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncidentPostmortem(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewIncidentPostmortemPayload(*testIncidentPostmortemData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentPostmortem %v: Response %s: %v", testIncidentPostmortemData, err.Error(), bStr)
	}
	incidentPostmortem := incidentPostmortemRsp.GetData()
	incidentPostmortemAttrs := incidentPostmortem.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentPostmortem if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentPostmortem datadog.IncidentPostmortem) {
		//Delete IncidentPostmortem
		httpresp, err := test_client.IncidentsApi.DeleteIncidentPostmortem(TESTAUTH, INCIDENT.GetId(), incidentPostmortem.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentPostmortem)

	assert.Equal(t, incidentPostmortem.GetType(), testIncidentPostmortemData.GetType())
	assert.Equal(t, incidentPostmortemAttrs.GetState(), testIncidentPostmortemData.Attributes.GetState())
	// Edit IncidentPostmortem
	testIncidentPostmortemData.Attributes.SetState("draft")
	incidentPostmortemUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		PatchIncidentPostmortem(TESTAUTH, INCIDENT.GetId(), incidentPostmortem.GetId()).
		Body(*datadog.NewIncidentPostmortemPayload(incidentPostmortem)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	incidentPostmortemUpdated := incidentPostmortemUpdatedRsp.GetData()
	incidentPostmortemAttrsUpdated := incidentPostmortemUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentPostmortemUpdated.GetId(), incidentPostmortem.GetId())
	assert.Equal(t, incidentPostmortemAttrsUpdated.GetState(), incidentPostmortem.Attributes.GetState())

	// Get IncidentPostmortem
	incidentPostmortemGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		GetIncidentPostmortem(TESTAUTH, INCIDENT.GetId(), incidentPostmortem.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentPostmortem %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentPostmortemGet := incidentPostmortemGetResp.GetData()
	incidentPostmortemAttrsGet := incidentPostmortemGet.GetAttributes()
	assert.Equal(t, incidentPostmortemGet.GetId(), incidentPostmortemUpdated.GetId())
	assert.Equal(t, incidentPostmortemAttrsGet.GetState(), incidentPostmortemAttrsUpdated.GetState())

	// Get IncidentPostmortems
	incidentPostmortemsGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncidentPostmortems(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentPostmortems %v: Response %s: %v", incidentPostmortem, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentPostmortemsGetResp.GetData()) >= 1)

}
