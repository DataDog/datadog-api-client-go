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

func TestIncidentNotesLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())

	testIncidentNoteData := datadog.NewIncidentNoteWithDefaults()
	testIncidentNoteData.SetType("incident_notes")
	testIncidentNoteData.SetAttributes(*datadog.NewIncidentNoteAttributesWithDefaults())
	testIncidentNoteData.Attributes.SetContent("TestNoteContent")
	testIncidentNoteData.Attributes.SetState("highlight")

	// Create IncidentNote
	incidentNoteRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncidentNote(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewIncidentNotePayload(*testIncidentNoteData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentNote %v: Response %s: %v", testIncidentNoteData, err.Error(), bStr)
	}
	incidentNote := incidentNoteRsp.GetData()
	incidentNoteAttrs := incidentNote.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentNote if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentNote datadog.IncidentNote) {
		//Delete IncidentNote
		httpresp, err := test_client.IncidentsApi.DeleteIncidentNote(TESTAUTH, INCIDENT.GetId(), incidentNote.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentNote %v: Response %s: %v", incidentNote, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentNote)

	assert.Equal(t, incidentNote.GetType(), testIncidentNoteData.GetType())
	assert.Equal(t, incidentNoteAttrs.GetContent(), testIncidentNoteData.Attributes.GetContent())
	assert.Equal(t, incidentNoteAttrs.GetState(), testIncidentNoteData.Attributes.GetState())

	// Edit IncidentNote
	incidentNote.Attributes.SetContent("new-content")
	incidentNote.Attributes.SetState("normal")
	incidentNoteUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		PatchIncidentNote(TESTAUTH, INCIDENT.GetId(), incidentNote.GetId()).
		Body(*datadog.NewIncidentNotePayload(incidentNote)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentNote %v: Response %s: %v", incidentNote, err.Error(), bStr)
	}
	incidentNoteUpdated := incidentNoteUpdatedRsp.GetData()
	incidentNoteAttrsUpdated := incidentNoteUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentNoteAttrsUpdated.GetContent(), incidentNote.Attributes.GetContent())
	assert.Equal(t, incidentNoteAttrsUpdated.GetState(), incidentNote.Attributes.GetState())

	// Get IncidentNote
	incidentNoteGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		GetIncidentNote(TESTAUTH, INCIDENT.GetId(), incidentNote.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentNote %v: Response %s: %v", incidentNote, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentNoteGet := incidentNoteGetResp.GetData()
	incidentNoteAttrsGet := incidentNoteGet.GetAttributes()
	assert.Equal(t, incidentNoteGet.GetId(), incidentNoteUpdated.GetId())
	assert.Equal(t, incidentNoteAttrsGet.GetContent(), incidentNoteAttrsUpdated.GetContent())
	assert.Equal(t, incidentNoteAttrsGet.GetState(), incidentNoteAttrsUpdated.GetState())

	// Get IncidentNotes
	incidentNotesGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncidentNotes(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentNotes %v: Response %s: %v", incidentNote, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentNotesGetResp.GetData()) >= 1)

}
