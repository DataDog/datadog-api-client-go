/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestIncidentUserDefinedFieldLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	// Set up field data
	testIncidentUserDefinedFieldData := datadog.NewUserDefinedFieldWithDefaults()
	testIncidentUserDefinedFieldData.SetAttributes(*datadog.NewUserDefinedFieldAttributesWithDefaults())
	testIncidentUserDefinedFieldData.Attributes.SetType(datadog.USERDEFINEDFIELDTYPEID_DROPDOWN)
	testIncidentUserDefinedFieldData.Attributes.SetTableId(datadog.USERDEFINEDFIELDTABLE_INCIDENT)
	testIncidentUserDefinedFieldData.Attributes.SetName("Test-ConfigField")

	// Create field
	userDefinedFieldRsp, httpresp, err := client.IncidentsApi.CreateUserDefinedField(ctx).
		Body(*datadog.NewUserDefinedFieldRequest(*testIncidentUserDefinedFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigField %v: Response %s: %v", testIncidentUserDefinedFieldData, err.Error(), bStr)
	}
	userDefinedField := userDefinedFieldRsp.GetData()
	userDefinedFieldAttrs := userDefinedField.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Ensure that we always delete the incidentConfigField if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, udf datadog.UserDefinedField) {
		//Delete field
		httpresp, err := test_client.IncidentsApi.DeleteUserDefinedField(ctx, udf.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentConfigField %v: Response %s: %v", udf, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, userDefinedField)

	assert.Equal(userDefinedField.GetType(), testIncidentUserDefinedFieldData.GetType())
	assert.Equal(userDefinedFieldAttrs.GetType(), testIncidentUserDefinedFieldData.Attributes.GetType())
	assert.Equal(userDefinedFieldAttrs.GetTableId(), testIncidentUserDefinedFieldData.Attributes.GetTableId())
	assert.Equal(userDefinedFieldAttrs.GetName(), testIncidentUserDefinedFieldData.Attributes.GetName())

	// Edit field
	userDefinedField.Attributes.SetName("Test-ConfigField-Updated")
	userDefinedFieldUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchUserDefinedField(ctx, userDefinedField.GetId()).
		Body(*datadog.NewUserDefinedFieldRequest(userDefinedField)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentConfigField %v: Response %s: %v", userDefinedField, err.Error(), bStr)
	}
	userDefinedFieldUpdated := userDefinedFieldUpdatedRsp.GetData()
	userDefinedFieldAttrsUpdated := userDefinedFieldUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(userDefinedFieldAttrsUpdated.GetName(), userDefinedField.Attributes.GetName())

	// Get field
	userDefinedFieldGetResp, httpresp, err := client.IncidentsApi.
		GetUserDefinedField(ctx, userDefinedField.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigField %v: Response %s: %v", userDefinedField, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	userDefinedFieldGet := userDefinedFieldGetResp.GetData()
	userDefinedAttrsGet := userDefinedFieldGet.GetAttributes()
	assert.Equal(userDefinedFieldGet.GetId(), userDefinedFieldUpdated.GetId())
	assert.Equal(userDefinedFieldGet.GetType(), testIncidentUserDefinedFieldData.GetType())
	assert.Equal(userDefinedAttrsGet.GetName(), userDefinedFieldAttrsUpdated.GetName())
	assert.Equal(userDefinedAttrsGet.GetType(), userDefinedFieldAttrsUpdated.GetType())
	assert.Equal(userDefinedAttrsGet.GetTableId(), userDefinedFieldAttrsUpdated.GetTableId())

	// Get field list
	userDefinedFieldListGetResp, httpresp, err := client.IncidentsApi.GetUserDefinedFields(ctx).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFields %v: Response %s: %v", userDefinedField, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(userDefinedFieldListGetResp.GetData()) >= 1)

}
