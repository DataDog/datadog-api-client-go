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

func TestIncidentSelectionsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	INCIDENT := createTestIncident(t)
	defer deleteTestIncident(t, INCIDENT.GetId())
	FIELD := createIncidentConfigField(t)
	defer deleteIncidentConfigField(t, FIELD.GetId())
	CHOICE_1 := createIncidentConfigFieldChoice(t, FIELD.GetId())
	defer deleteIncidentConfigFieldChoice(t, FIELD.GetId(), CHOICE_1.GetId())
	CHOICE_2 := createIncidentConfigFieldChoice(t, FIELD.GetId())
	defer deleteIncidentConfigFieldChoice(t, FIELD.GetId(), CHOICE_2.GetId())

	testIncidentSelectionData := datadog.NewIncidentSelectionWithDefaults()
	testIncidentSelectionData.SetType("selections")
	testIncidentSelectionData.SetAttributes(*datadog.NewIncidentSelectionAttributesWithDefaults())
	testIncidentSelectionData.SetRelationships(*datadog.NewIncidentSelectionRelationshipsWithDefaults())
	testIncidentSelectionData.Attributes.SetObjectId(INCIDENT.GetId())
	testIncidentSelectionData.Relationships.SetField(*datadog.NewIncidentSelectionRelationshipsFieldWithDefaults())
	testIncidentSelectionData.Relationships.Field.SetData(FIELD)
	testIncidentSelectionData.Relationships.SetChoice(*datadog.NewIncidentSelectionRelationshipsChoiceWithDefaults())
	testIncidentSelectionData.Relationships.Choice.SetData(CHOICE_1)

	// Create IncidentSelection
	incidentSelectionRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.CreateIncidentSelection(TESTAUTH, INCIDENT.GetId()).
		Body(*datadog.NewIncidentSelectionPayload(*testIncidentSelectionData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentSelection %v: Response %s: %v", testIncidentSelectionData, err.Error(), bStr)
	}
	incidentSelection := incidentSelectionRsp.GetData()
	incidentSelectionAttrs := incidentSelection.GetAttributes()
	incidentSelectionRelationships := incidentSelection.GetRelationships()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentSelection if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentSelection datadog.IncidentSelection) {
		//Delete IncidentSelection
		httpresp, err := test_client.IncidentsApi.DeleteIncidentSelection(TESTAUTH, INCIDENT.GetId(), incidentSelection.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentSelection)

	assert.Equal(t, incidentSelection.GetType(), testIncidentSelectionData.GetType())
	assert.Equal(t, incidentSelectionAttrs.GetObjectId(), INCIDENT.GetId())
	assert.Equal(t, incidentSelectionRelationships.Choice.Data.GetId(), CHOICE_1.GetId())
	assert.Equal(t, incidentSelectionRelationships.Field.Data.GetId(), FIELD.GetId())

	// Edit IncidentSelection
	incidentSelection.Relationships.Choice.SetData(CHOICE_2)
	incidentSelectionUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsApi.
		PatchIncidentSelection(TESTAUTH, INCIDENT.GetId(), incidentSelection.GetId()).
		Body(*datadog.NewIncidentSelectionPayload(incidentSelection)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	incidentSelectionUpdated := incidentSelectionUpdatedRsp.GetData()
	incidentSelectionAttrsUpdated := incidentSelectionUpdated.GetAttributes()
	incidentSelectionRelationshipsUpdated := incidentSelectionUpdated.GetRelationships()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentSelectionAttrsUpdated.GetObjectId(), INCIDENT.GetId())
	assert.Equal(t, incidentSelectionRelationshipsUpdated.Choice.Data.GetId(), CHOICE_2.GetId())
	assert.Equal(t, incidentSelectionRelationshipsUpdated.Field.Data.GetId(), FIELD.GetId())

	// Get IncidentSelections
	incidentSelectionsGetResp, httpresp, err := TESTAPICLIENT.IncidentsApi.GetIncidentSelections(TESTAUTH, INCIDENT.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentSelections %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentSelectionsGetResp.GetData()) >= 1)

}
