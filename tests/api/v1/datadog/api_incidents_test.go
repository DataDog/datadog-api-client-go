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

func TestIncidentsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	time := TESTCLOCK.Now()
	formattedClock := *datadog.NewNullableTime(&time)
	testIncidentData := datadog.NewIncidentWithDefaults()
	testIncidentData.SetType("incidents")
	testIncidentData.SetAttributes(*datadog.NewIncidentAttributesWithDefaults())
	testIncidentData.Attributes.SetTitle("Test-Incident")
	testIncidentData.Attributes.SetState("unkown")
	testIncidentData.Attributes.SetCustomerImpacted(false)
	testIncidentData.Attributes.SetCustomerImpactScope("none")
	testIncidentData.Attributes.SetCustomerImpactStart(formattedClock)
	testIncidentData.Attributes.SetCustomerImpactEnd(formattedClock)
	testIncidentData.Attributes.SetDetected(formattedClock)

	// Create Incident
	incidentRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncident(TESTAUTH).Body(*datadog.NewIncidentPayload(*testIncidentData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating Incident %v: Response %s: %v", testIncidentData, err.Error(), bStr)
	}
	incident := incidentRsp.GetData()
	incidentAttrs := incident.GetAttributes()

	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incident if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incident datadog.Incident) {
		//Delete Incident
		httpresp, err := test_client.IncidentsApi.DeleteIncident(TESTAUTH, incident.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting Incident %v: Response %s: %v", incident, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incident)

	assert.Equal(t, incident.GetType(), testIncidentData.GetType())
	assert.Equal(t, incidentAttrs.GetTitle(), testIncidentData.Attributes.GetTitle())
	assert.Equal(t, incidentAttrs.GetState(), testIncidentData.Attributes.GetState())
	assert.Equal(t, incidentAttrs.GetCustomerImpacted(), testIncidentData.Attributes.GetCustomerImpacted())
	assert.Equal(t, incidentAttrs.GetCustomerImpactScope(), testIncidentData.Attributes.GetCustomerImpactScope())

	// Edit Incident
	incident.Attributes.SetTitle("Test-Incident-Updated")
	incidentUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.PatchIncident(TESTAUTH, incident.GetId()).Body(*datadog.NewIncidentPayload(incident)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating Incident %v: Response %s: %v", incident, err.Error(), bStr)
	}
	incidentUpdated := incidentUpdatedRsp.GetData()
	incidentAttrsUpdated := incidentUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentAttrsUpdated.GetTitle(), incidentAttrsUpdated.GetTitle())

	// Get Incident
	incidentGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncident(TESTAUTH, incident.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Incident %v: Response %s: %v", incident, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentGet := incidentGetResp.GetData()
	incidentAttrsGet := incidentGet.GetAttributes()
	assert.Equal(t, incidentGet.GetId(), incidentUpdated.GetId())
	assert.Equal(t, incidentAttrsGet.GetTitle(), incidentAttrsUpdated.GetTitle())

	// Get Incidents
	incidentsGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncidents(TESTAUTH).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting Incidents %v: Response %s: %v", incident, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentsGetResp.GetData()) >= 1)

}

func createTestIncident(t *testing.T) datadog.Incident {
	testIncidentData := datadog.NewIncidentWithDefaults()
	testIncidentData.SetType("incidents")
	testIncidentData.SetAttributes(*datadog.NewIncidentAttributesWithDefaults())
	testIncidentData.Attributes.SetTitle("Test-Incident")
	testIncidentData.Attributes.SetState("unkown")
	testIncidentData.Attributes.SetCustomerImpacted(false)
	testIncidentData.Attributes.SetCustomerImpactScope("none")

	incidentRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncident(TESTAUTH).Body(*datadog.NewIncidentPayload(*testIncidentData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating Incident %v: Response %s: %v", testIncidentData, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	return incidentRsp.GetData()

}

func deleteTestIncident(t *testing.T, id string) {
	httpresp, err := TESTAPICLIENT.IncidentsApi.DeleteIncident(TESTAUTH, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting Incident: %v failed with %v, Another test may have already deleted this Incident: %v",
			id, httpresp.StatusCode, err)
	}

}
