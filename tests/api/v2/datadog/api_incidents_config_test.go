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

func TestIncidentConfigFieldsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	testIncidentConfigFieldData := datadog.NewIncidentConfigFieldWithDefaults()
	testIncidentConfigFieldData.SetType("user_defined_field")
	testIncidentConfigFieldData.SetAttributes(*datadog.NewIncidentConfigFieldAttributesWithDefaults())
	testIncidentConfigFieldData.Attributes.SetType(datadog.INCIDENTCONFIGFIELDTYPE_DROPDOWN)
	testIncidentConfigFieldData.Attributes.SetTableId(datadog.INCIDENTCONFIGFIELDTABLE_INCIDENT)
	testIncidentConfigFieldData.Attributes.SetName("Test-ConfigField")

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := client.IncidentsApi.CreateIncidentConfigField(ctx).
		Body(*datadog.NewIncidentConfigFieldPayload(*testIncidentConfigFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigField %v: Response %s: %v", testIncidentConfigFieldData, err.Error(), bStr)
	}
	incidentConfigField := incidentConfigRsp.GetData()
	incidentConfigFieldAttrs := incidentConfigField.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentConfigField if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentConfigField datadog.IncidentConfigField) {
		//Delete IncidentConfig
		httpresp, err := test_client.IncidentsApi.DeleteIncidentConfigField(ctx, incidentConfigField.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentConfigField)

	assert.Equal(incidentConfigField.GetType(), testIncidentConfigFieldData.GetType())
	assert.Equal(incidentConfigFieldAttrs.GetType(), testIncidentConfigFieldData.Attributes.GetType())
	assert.Equal(incidentConfigFieldAttrs.GetTableId(), testIncidentConfigFieldData.Attributes.GetTableId())
	assert.Equal(incidentConfigFieldAttrs.GetName(), testIncidentConfigFieldData.Attributes.GetName())

	// Edit IncidentConfig
	incidentConfigField.Attributes.SetName("Test-ConfigField-Updated")
	incidentConfigUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchIncidentConfigField(ctx, incidentConfigField.GetId()).
		Body(*datadog.NewIncidentConfigFieldPayload(incidentConfigField)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	incidentConfigFieldUpdated := incidentConfigUpdatedRsp.GetData()
	incidentConfigFieldAttrsUpdated := incidentConfigFieldUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentConfigFieldAttrsUpdated.GetName(), incidentConfigField.Attributes.GetName())

	// Get IncidentConfig
	incidentConfigGetResp, httpresp, err := client.IncidentsApi.
		GetIncidentConfigField(ctx, incidentConfigField.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	incidentConfigFieldGet := incidentConfigGetResp.GetData()
	incidentConfigFieldAttrsGet := incidentConfigFieldGet.GetAttributes()
	assert.Equal(incidentConfigFieldGet.GetId(), incidentConfigFieldUpdated.GetId())
	assert.Equal(incidentConfigField.GetType(), testIncidentConfigFieldData.GetType())
	assert.Equal(incidentConfigFieldAttrsGet.GetName(), incidentConfigFieldAttrsUpdated.GetName())
	assert.Equal(incidentConfigFieldAttrsGet.GetType(), incidentConfigFieldAttrsUpdated.GetType())
	assert.Equal(incidentConfigFieldAttrsGet.GetTableId(), incidentConfigFieldAttrsUpdated.GetTableId())

	// Get IncidentConfigs
	incidentConfigsGetResp, httpresp, err := client.IncidentsApi.GetIncidentConfigFields(ctx).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFields %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentConfigsGetResp.GetData()) >= 1)

}

func TestIncidentConfigChoicesLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	field := createIncidentConfigField(t, client, ctx)
	defer deleteIncidentConfigField(t, client, ctx, field.GetId())
	testIncidentConfigFieldChoiceData := datadog.NewIncidentConfigFieldChoiceWithDefaults()
	testIncidentConfigFieldChoiceData.SetType("user_defined_choice")
	testIncidentConfigFieldChoiceData.SetAttributes(*datadog.NewIncidentConfigFieldChoiceAttributesWithDefaults())
	testIncidentConfigFieldChoiceData.Attributes.SetName("Test-Choice-Name")
	testIncidentConfigFieldChoiceData.Attributes.SetValue("Test-Choice-Value")

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := client.IncidentsApi.
		CreateIncidentConfigFieldChoice(ctx, field.GetId()).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(*testIncidentConfigFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigFieldChoice %v: Response %s: %v", testIncidentConfigFieldChoiceData, err.Error(), bStr)
	}
	incidentConfigFieldChoice := incidentConfigRsp.GetData()
	incidentConfigFieldChoiceAttrs := incidentConfigFieldChoice.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentConfigFieldChoice if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentConfigFieldChoice datadog.IncidentConfigFieldChoice) {
		//Delete IncidentConfig
		httpresp, err := test_client.IncidentsApi.
			DeleteIncidentConfigFieldChoice(ctx, field.GetId(), incidentConfigFieldChoice.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentConfigFieldChoice)

	assert.Equal(incidentConfigFieldChoice.GetType(), testIncidentConfigFieldChoiceData.GetType())
	assert.Equal(incidentConfigFieldChoiceAttrs.GetName(), testIncidentConfigFieldChoiceData.Attributes.GetName())
	assert.Equal(incidentConfigFieldChoiceAttrs.GetValue(), testIncidentConfigFieldChoiceData.Attributes.GetValue())

	// Edit IncidentConfig
	incidentConfigFieldChoice.Attributes.SetName("Test-ConfigField-Updated")
	incidentConfigFieldChoice.Attributes.SetValue("Test-ConfigField-Updated")
	incidentConfigUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchIncidentConfigFieldChoice(ctx, field.GetId(), incidentConfigFieldChoice.GetId()).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(incidentConfigFieldChoice)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
	}
	incidentConfigFieldChoiceUpdated := incidentConfigUpdatedRsp.GetData()
	incidentConfigFieldChoiceAttrsUpdated := incidentConfigFieldChoiceUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentConfigFieldChoiceAttrsUpdated.GetName(), incidentConfigFieldChoice.Attributes.GetName())
	assert.Equal(incidentConfigFieldChoiceAttrsUpdated.GetValue(), incidentConfigFieldChoice.Attributes.GetValue())

	// Get IncidentConfig
	incidentConfigGetResp, httpresp, err := client.IncidentsApi.
		GetIncidentConfigFieldChoice(ctx, field.GetId(), incidentConfigFieldChoice.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	incidentConfigFieldChoiceGet := incidentConfigGetResp.GetData()
	incidentConfigFieldChoiceAttrsGet := incidentConfigFieldChoiceGet.GetAttributes()
	assert.Equal(incidentConfigFieldChoiceGet.GetId(), incidentConfigFieldChoiceUpdated.GetId())
	assert.Equal(incidentConfigFieldChoice.GetType(), testIncidentConfigFieldChoiceData.GetType())
	assert.Equal(incidentConfigFieldChoiceAttrsGet.GetName(), incidentConfigFieldChoiceAttrsUpdated.GetName())
	assert.Equal(incidentConfigFieldChoiceAttrsGet.GetValue(), incidentConfigFieldChoiceAttrsUpdated.GetValue())

	// Get IncidentConfigs
	incidentConfigFieldChoicesGetResp, httpresp, err := client.IncidentsApi.
		GetIncidentConfigFieldChoices(ctx, field.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFieldChoices %v: Response %s: %v",
			incidentConfigFieldChoice, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentConfigFieldChoicesGetResp.GetData()) >= 1)

}

func createIncidentConfigField(t *testing.T, client *datadog.APIClient, ctx context.Context) datadog.IncidentConfigField {
	testIncidentConfigFieldData := datadog.NewIncidentConfigFieldWithDefaults()
	testIncidentConfigFieldData.SetType("user_defined_field")
	testIncidentConfigFieldData.SetAttributes(*datadog.NewIncidentConfigFieldAttributesWithDefaults())
	testIncidentConfigFieldData.Attributes.SetType(datadog.INCIDENTCONFIGFIELDTYPE_DROPDOWN)
	testIncidentConfigFieldData.Attributes.SetTableId(datadog.INCIDENTCONFIGFIELDTABLE_INCIDENT)
	testIncidentConfigFieldData.Attributes.SetName(fmt.Sprintf("createIncidentConfigField-%s", RandomString(16)))

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := client.IncidentsApi.CreateIncidentConfigField(ctx).
		Body(*datadog.NewIncidentConfigFieldPayload(*testIncidentConfigFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigField %v: Response %s: %v", testIncidentConfigFieldData, err.Error(), bStr)
	}
	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return incidentConfigRsp.GetData()

}

func deleteIncidentConfigField(t *testing.T, client *datadog.APIClient, ctx context.Context, id string) {
	//Delete IncidentConfig
	httpresp, err := client.IncidentsApi.DeleteIncidentConfigField(ctx, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting IncidentConfigField: %v failed with %v, "+
			"Another test may have already deleted this Config Field: %v",
			id, httpresp.StatusCode, err)
	}

}

func createIncidentConfigFieldChoice(t *testing.T, client *datadog.APIClient, ctx context.Context, fieldID string) datadog.IncidentConfigFieldChoice {
	testIncidentConfigFieldChoiceData := datadog.NewIncidentConfigFieldChoiceWithDefaults()
	testIncidentConfigFieldChoiceData.SetType("user_defined_choice")
	testIncidentConfigFieldChoiceData.SetAttributes(*datadog.NewIncidentConfigFieldChoiceAttributesWithDefaults())
	testIncidentConfigFieldChoiceData.Attributes.SetName(fmt.Sprintf("createIncidentConfigFieldChoice-%s", RandomString(16)))
	testIncidentConfigFieldChoiceData.Attributes.SetValue("Test-Choice-Value")

	incidentConfigFieldChoiceRsp, httpresp, err := client.IncidentsApi.
		CreateIncidentConfigFieldChoice(ctx, fieldID).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(*testIncidentConfigFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigFieldChoice %v: Response %s: %v", testIncidentConfigFieldChoiceData, err.Error(), bStr)
	}

	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return incidentConfigFieldChoiceRsp.GetData()

}

func deleteIncidentConfigFieldChoice(t *testing.T, client *datadog.APIClient, ctx context.Context, fieldID, choiceID string) {
	httpresp, err := client.IncidentsApi.
		DeleteIncidentConfigFieldChoice(ctx, fieldID, choiceID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting IncidentConfigFieldChoice: %v failed with %v, "+
			"Another test may have already deleted this Config Field Choice: %v",
			choiceID, httpresp.StatusCode, err)
	}

}
