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

func TestUserDefinedFieldSelectionsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(ctx, t, client)
	defer deleteTestIncident(ctx, t, client, incident.GetId())
	field := createUserDefinedField(ctx, t, client)
	defer deleteUserDefinedField(ctx, t, client, field.GetId())
	choice1 := createUserDefinedFieldChoice(ctx, t, client, field.GetId(), "generated-client-choice-1")
	defer deleteUserDefinedFieldChoice(ctx, t, client, field.GetId(), choice1.GetId())
	choice2 := createUserDefinedFieldChoice(ctx, t, client, field.GetId(), "generated-client-choice-2")
	defer deleteUserDefinedFieldChoice(ctx, t, client, field.GetId(), choice2.GetId())

	testIncidentSelectionData := datadog.NewUserDefinedFieldSelectionWithDefaults()
	testIncidentSelectionData.SetAttributes(*datadog.NewUserDefinedFieldSelectionAttributesWithDefaults())
	testIncidentSelectionData.SetRelationships(*datadog.NewUserDefinedFieldSelectionRelationshipsWithDefaults())
	testIncidentSelectionData.Attributes.SetObjectId(incident.GetId())
	testIncidentSelectionData.Relationships.SetField(*datadog.NewUserDefinedFieldSelectionRelationshipsFieldWithDefaults())
	fieldRelationship := datadog.NewUserDefinedFieldRelationship()
	fieldRelationship.SetId(field.Id)
	testIncidentSelectionData.Relationships.Field.SetData(*fieldRelationship)
	testIncidentSelectionData.Relationships.SetChoice(*datadog.NewUserDefinedFieldSelectionRelationshipsChoiceWithDefaults())
	choiceRelationship := datadog.NewUserDefinedFieldChoiceRelationship()
	choiceRelationship.SetId(choice1.Id)
	testIncidentSelectionData.Relationships.Choice.SetData(*choiceRelationship)

	// Create IncidentSelection
	incidentSelectionRsp, httpresp, err := client.IncidentsApi.CreateUserDefinedFieldSelection(ctx, incident.GetId()).
		Body(*datadog.NewUserDefinedFieldSelectionRequest(*testIncidentSelectionData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentSelection %v: Response %s: %v", testIncidentSelectionData, err.Error(), bStr)
	}
	incidentSelection := incidentSelectionRsp.GetData()
	incidentSelectionAttrs := incidentSelection.GetAttributes()
	incidentSelectionRelationships := incidentSelection.GetRelationships()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentSelection if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentSelection datadog.UserDefinedFieldSelection) {
		//Delete IncidentSelection
		httpresp, err := test_client.IncidentsApi.DeleteUserDefinedFieldSelection(ctx, incident.GetId(), incidentSelection.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentSelection)

	assert.Equal(incidentSelection.GetType(), testIncidentSelectionData.GetType())
	assert.Equal(incidentSelectionAttrs.GetObjectId(), incident.GetId())
	assert.Equal(incidentSelectionRelationships.Choice.Data.GetId(), choice1.GetId())
	assert.Equal(incidentSelectionRelationships.Field.Data.GetId(), field.GetId())

	// Edit IncidentSelection

	choice2Relationship := datadog.NewUserDefinedFieldChoiceRelationship()
	choice2Relationship.SetId(choice2.Id)
	incidentSelection.Relationships.Choice.SetData(*choice2Relationship)
	incidentSelectionUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchUserDefinedFieldSelection(ctx, incident.GetId(), incidentSelection.GetId()).
		Body(*datadog.NewUserDefinedFieldSelectionRequest(incidentSelection)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	incidentSelectionUpdated := incidentSelectionUpdatedRsp.GetData()
	incidentSelectionAttrsUpdated := incidentSelectionUpdated.GetAttributes()
	incidentSelectionRelationshipsUpdated := incidentSelectionUpdated.GetRelationships()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentSelectionAttrsUpdated.GetObjectId(), incident.GetId())
	assert.Equal(incidentSelectionRelationshipsUpdated.Choice.Data.GetId(), choice2.GetId())
	assert.Equal(incidentSelectionRelationshipsUpdated.Field.Data.GetId(), field.GetId())

	// Get IncidentSelections
	incidentSelectionsGetResp, httpresp, err := client.IncidentsApi.GetUserDefinedFieldSelections(ctx, incident.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentSelections %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentSelectionsGetResp.GetData()) >= 1)

}
