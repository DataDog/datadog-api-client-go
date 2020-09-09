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

func createUserDefinedField(ctx context.Context, t *testing.T, client *datadog.APIClient) datadog.UserDefinedField {
	testUserDefinedFieldData := datadog.NewUserDefinedFieldWithDefaults()
	testUserDefinedFieldData.SetType("user_defined_field")
	testUserDefinedFieldData.SetAttributes(*datadog.NewUserDefinedFieldAttributesWithDefaults())
	testUserDefinedFieldData.Attributes.SetType(datadog.USERDEFINEDFIELDTYPEID_DROPDOWN)
	testUserDefinedFieldData.Attributes.SetTableId(datadog.USERDEFINEDFIELDTABLE_INCIDENT)
	testUserDefinedFieldData.Attributes.SetName("createUserDefinedField-for_generated_tests")

	// Create UserDefinedField
	UserDefinedRsp, httpresp, err := client.IncidentsApi.CreateUserDefinedField(ctx).
		Body(*datadog.NewUserDefinedFieldRequest(*testUserDefinedFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating UserDefinedField %v: Response %s: %v", testUserDefinedFieldData, err.Error(), bStr)
	}
	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return UserDefinedRsp.GetData()

}

func deleteUserDefinedField(ctx context.Context, t *testing.T, client *datadog.APIClient, id string) {
	//Delete UserDefinedField
	httpresp, err := client.IncidentsApi.DeleteUserDefinedField(ctx, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting UserDefinedField: %v failed with %v, "+
			"Another test may have already deleted this Config Field: %v",
			id, httpresp.StatusCode, err)
	}

}

func createUserDefinedFieldChoice(ctx context.Context, t *testing.T, client *datadog.APIClient, fieldID string, name string) datadog.UserDefinedFieldChoice {
	testUserDefinedFieldChoiceData := datadog.NewUserDefinedFieldChoiceWithDefaults()
	testUserDefinedFieldChoiceData.SetType("user_defined_choice")
	testUserDefinedFieldChoiceData.SetAttributes(*datadog.NewUserDefinedFieldChoiceAttributesWithDefaults())
	testUserDefinedFieldChoiceData.Attributes.SetName(fmt.Sprintf("createUserDefinedFieldChoice-%s", name))
	testUserDefinedFieldChoiceData.Attributes.SetValue("Test-Choice-Value")

	UserDefinedFieldChoiceRsp, httpresp, err := client.IncidentsApi.
		CreateUserDefinedFieldChoice(ctx, fieldID).
		Body(*datadog.NewUserDefinedFieldChoiceRequest(*testUserDefinedFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating UserDefinedFieldChoice %v: Response %s: %v", testUserDefinedFieldChoiceData, err.Error(), bStr)
	}

	assert := tests.Assert(ctx, t)
	assert.Equal(httpresp.StatusCode, 201)
	return UserDefinedFieldChoiceRsp.GetData()

}

func deleteUserDefinedFieldChoice(ctx context.Context, t *testing.T, client *datadog.APIClient, fieldID, choiceID string) {
	httpresp, err := client.IncidentsApi.
		DeleteUserDefinedFieldChoice(ctx, fieldID, choiceID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting UserDefinedFieldChoice: %v failed with %v, "+
			"Another test may have already deleted this Config Field Choice: %v",
			choiceID, httpresp.StatusCode, err)
	}

}
func TestUserDefinedFieldChoicesLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	field := createUserDefinedField(ctx, t, client)
	defer deleteUserDefinedField(ctx, t, client, field.GetId())
	testUserDefinedFieldChoiceData := datadog.NewUserDefinedFieldChoiceWithDefaults()
	testUserDefinedFieldChoiceData.SetType("user_defined_choice")
	testUserDefinedFieldChoiceData.SetAttributes(*datadog.NewUserDefinedFieldChoiceAttributesWithDefaults())
	testUserDefinedFieldChoiceData.Attributes.SetName("Test-Choice-Name")
	testUserDefinedFieldChoiceData.Attributes.SetValue("Test-Choice-Value")

	// Create UserDefined
	UserDefinedRsp, httpresp, err := client.IncidentsApi.
		CreateUserDefinedFieldChoice(ctx, field.GetId()).
		Body(*datadog.NewUserDefinedFieldChoiceRequest(*testUserDefinedFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating UserDefinedFieldChoice %v: Response %s: %v", testUserDefinedFieldChoiceData, err.Error(), bStr)
	}
	UserDefinedFieldChoice := UserDefinedRsp.GetData()
	UserDefinedFieldChoiceAttrs := UserDefinedFieldChoice.GetAttributes()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the UserDefinedFieldChoice if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, UserDefinedFieldChoice datadog.UserDefinedFieldChoice) {
		//Delete UserDefined
		httpresp, err := test_client.IncidentsApi.
			DeleteUserDefinedFieldChoice(ctx, field.GetId(), UserDefinedFieldChoice.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting UserDefinedFieldChoice %v: Response %s: %v", UserDefinedFieldChoice, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, UserDefinedFieldChoice)

	assert.Equal(UserDefinedFieldChoice.GetType(), testUserDefinedFieldChoiceData.GetType())
	assert.Equal(UserDefinedFieldChoiceAttrs.GetName(), testUserDefinedFieldChoiceData.Attributes.GetName())
	assert.Equal(UserDefinedFieldChoiceAttrs.GetValue(), testUserDefinedFieldChoiceData.Attributes.GetValue())

	// Edit UserDefined
	UserDefinedFieldChoice.Attributes.SetName("Test-ConfigField-Updated")
	UserDefinedFieldChoice.Attributes.SetValue("Test-ConfigField-Updated")
	UserDefinedUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchUserDefinedFieldChoice(ctx, field.GetId(), UserDefinedFieldChoice.GetId()).
		Body(*datadog.NewUserDefinedFieldChoiceRequest(UserDefinedFieldChoice)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating UserDefinedFieldChoice %v: Response %s: %v", UserDefinedFieldChoice, err.Error(), bStr)
	}
	UserDefinedFieldChoiceUpdated := UserDefinedUpdatedRsp.GetData()
	UserDefinedFieldChoiceAttrsUpdated := UserDefinedFieldChoiceUpdated.GetAttributes()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(UserDefinedFieldChoiceAttrsUpdated.GetName(), UserDefinedFieldChoice.Attributes.GetName())
	assert.Equal(UserDefinedFieldChoiceAttrsUpdated.GetValue(), UserDefinedFieldChoice.Attributes.GetValue())

	// Get UserDefined
	UserDefinedGetResp, httpresp, err := client.IncidentsApi.
		GetUserDefinedFieldChoice(ctx, field.GetId(), UserDefinedFieldChoice.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting UserDefinedFieldChoice %v: Response %s: %v", UserDefinedFieldChoice, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)

	UserDefinedFieldChoiceGet := UserDefinedGetResp.GetData()
	UserDefinedFieldChoiceAttrsGet := UserDefinedFieldChoiceGet.GetAttributes()
	assert.Equal(UserDefinedFieldChoiceGet.GetId(), UserDefinedFieldChoiceUpdated.GetId())
	assert.Equal(UserDefinedFieldChoice.GetType(), testUserDefinedFieldChoiceData.GetType())
	assert.Equal(UserDefinedFieldChoiceAttrsGet.GetName(), UserDefinedFieldChoiceAttrsUpdated.GetName())
	assert.Equal(UserDefinedFieldChoiceAttrsGet.GetValue(), UserDefinedFieldChoiceAttrsUpdated.GetValue())

	// Get UserDefineds
	UserDefinedFieldChoicesGetResp, httpresp, err := client.IncidentsApi.
		GetUserDefinedFieldChoices(ctx, field.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting UserDefinedFieldChoices %v: Response %s: %v",
			UserDefinedFieldChoice, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(UserDefinedFieldChoicesGetResp.GetData()) >= 1)

}
