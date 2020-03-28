/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

func TestIncidentConfigFieldsLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testIncidentConfigFieldData := datadog.NewIncidentConfigFieldWithDefaults()
	testIncidentConfigFieldData.SetType("field")
	testIncidentConfigFieldData.SetAttributes(*datadog.NewIncidentConfigFieldAttributesWithDefaults())
	testIncidentConfigFieldData.Attributes.SetType(datadog.INCIDENTCONFIGFIELDTYPE_DROPDOWN)
	testIncidentConfigFieldData.Attributes.SetTableId(datadog.INCIDENTCONFIGFIELDTABLE_INCIDENT)
	testIncidentConfigFieldData.Attributes.SetName("Test-ConfigField")

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.CreateIncidentConfigField(TESTAUTH).
		Body(*datadog.NewIncidentConfigFieldPayload(*testIncidentConfigFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigField %v: Response %s: %v", testIncidentConfigFieldData, err.Error(), bStr)
	}
	incidentConfigField := incidentConfigRsp.GetData()
	incidentConfigFieldAttrs := incidentConfigField.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentConfigField if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentConfigField datadog.IncidentConfigField) {
		//Delete IncidentConfig
		httpresp, err := test_client.IncidentsConfigApi.DeleteIncidentConfigField(TESTAUTH, incidentConfigField.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentConfigField)

	assert.Equal(t, incidentConfigField.GetType(), testIncidentConfigFieldData.GetType())
	assert.Equal(t, incidentConfigFieldAttrs.GetType(), testIncidentConfigFieldData.Attributes.GetType())
	assert.Equal(t, incidentConfigFieldAttrs.GetTableId(), testIncidentConfigFieldData.Attributes.GetTableId())
	assert.Equal(t, incidentConfigFieldAttrs.GetName(), testIncidentConfigFieldData.Attributes.GetName())

	// Edit IncidentConfig
	incidentConfigField.Attributes.SetName("Test-ConfigField-Updated")
	incidentConfigUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		PatchIncidentConfigField(TESTAUTH, incidentConfigField.GetId()).
		Body(*datadog.NewIncidentConfigFieldPayload(incidentConfigField)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	incidentConfigFieldUpdated := incidentConfigUpdatedRsp.GetData()
	incidentConfigFieldAttrsUpdated := incidentConfigFieldUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentConfigFieldAttrsUpdated.GetName(), incidentConfigField.Attributes.GetName())

	// Get IncidentConfig
	incidentConfigGetResp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		GetIncidentConfigField(TESTAUTH, incidentConfigField.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigField %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentConfigFieldGet := incidentConfigGetResp.GetData()
	incidentConfigFieldAttrsGet := incidentConfigFieldGet.GetAttributes()
	assert.Equal(t, incidentConfigFieldGet.GetId(), incidentConfigFieldUpdated.GetId())
	assert.Equal(t, incidentConfigField.GetType(), testIncidentConfigFieldData.GetType())
	assert.Equal(t, incidentConfigFieldAttrsGet.GetName(), incidentConfigFieldAttrsUpdated.GetName())
	assert.Equal(t, incidentConfigFieldAttrsGet.GetType(), incidentConfigFieldAttrsUpdated.GetType())
	assert.Equal(t, incidentConfigFieldAttrsGet.GetTableId(), incidentConfigFieldAttrsUpdated.GetTableId())

	// Get IncidentConfigs
	incidentConfigsGetResp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.GetIncidentConfigFields(TESTAUTH).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFields %v: Response %s: %v", incidentConfigField, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentConfigsGetResp.GetData()) >= 1)

}

func TestIncidentConfigChoicesLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	FIELD := createIncidentConfigField(t)
	defer deleteIncidentConfigField(t, FIELD.GetId())
	testIncidentConfigFieldChoiceData := datadog.NewIncidentConfigFieldChoiceWithDefaults()
	testIncidentConfigFieldChoiceData.SetType("choice")
	testIncidentConfigFieldChoiceData.SetAttributes(*datadog.NewIncidentConfigFieldChoiceAttributesWithDefaults())
	testIncidentConfigFieldChoiceData.Attributes.SetName("Test-Choice-Name")
	testIncidentConfigFieldChoiceData.Attributes.SetValue("Test-Choice-Value")

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		CreateIncidentConfigFieldChoice(TESTAUTH, FIELD.GetId()).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(*testIncidentConfigFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigFieldChoice %v: Response %s: %v", testIncidentConfigFieldChoiceData, err.Error(), bStr)
	}
	incidentConfigFieldChoice := incidentConfigRsp.GetData()
	incidentConfigFieldChoiceAttrs := incidentConfigFieldChoice.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentConfigFieldChoice if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentConfigFieldChoice datadog.IncidentConfigFieldChoice) {
		//Delete IncidentConfig
		httpresp, err := test_client.IncidentsConfigApi.
			DeleteIncidentConfigFieldChoice(TESTAUTH, FIELD.GetId(), incidentConfigFieldChoice.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
		}
		assert.Equal(t, httpresp.StatusCode, 204)
	}(TESTAPICLIENT, incidentConfigFieldChoice)

	assert.Equal(t, incidentConfigFieldChoice.GetType(), testIncidentConfigFieldChoiceData.GetType())
	assert.Equal(t, incidentConfigFieldChoiceAttrs.GetName(), testIncidentConfigFieldChoiceData.Attributes.GetName())
	assert.Equal(t, incidentConfigFieldChoiceAttrs.GetValue(), testIncidentConfigFieldChoiceData.Attributes.GetValue())

	// Edit IncidentConfig
	incidentConfigFieldChoice.Attributes.SetName("Test-ConfigField-Updated")
	incidentConfigFieldChoice.Attributes.SetValue("Test-ConfigField-Updated")
	incidentConfigUpdatedRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		PatchIncidentConfigFieldChoice(TESTAUTH, FIELD.GetId(), incidentConfigFieldChoice.GetId()).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(incidentConfigFieldChoice)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
	}
	incidentConfigFieldChoiceUpdated := incidentConfigUpdatedRsp.GetData()
	incidentConfigFieldChoiceAttrsUpdated := incidentConfigFieldChoiceUpdated.GetAttributes()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, incidentConfigFieldChoiceAttrsUpdated.GetName(), incidentConfigFieldChoice.Attributes.GetName())
	assert.Equal(t, incidentConfigFieldChoiceAttrsUpdated.GetValue(), incidentConfigFieldChoice.Attributes.GetValue())

	// Get IncidentConfig
	incidentConfigGetResp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		GetIncidentConfigFieldChoice(TESTAUTH, FIELD.GetId(), incidentConfigFieldChoice.GetId()).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFieldChoice %v: Response %s: %v", incidentConfigFieldChoice, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	incidentConfigFieldChoiceGet := incidentConfigGetResp.GetData()
	incidentConfigFieldChoiceAttrsGet := incidentConfigFieldChoiceGet.GetAttributes()
	assert.Equal(t, incidentConfigFieldChoiceGet.GetId(), incidentConfigFieldChoiceUpdated.GetId())
	assert.Equal(t, incidentConfigFieldChoice.GetType(), testIncidentConfigFieldChoiceData.GetType())
	assert.Equal(t, incidentConfigFieldChoiceAttrsGet.GetName(), incidentConfigFieldChoiceAttrsUpdated.GetName())
	assert.Equal(t, incidentConfigFieldChoiceAttrsGet.GetValue(), incidentConfigFieldChoiceAttrsUpdated.GetValue())

	// Get IncidentConfigs
	incidentConfigFieldChoicesGetResp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		GetIncidentConfigFieldChoices(TESTAUTH, FIELD.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentConfigFieldChoices %v: Response %s: %v",
			incidentConfigFieldChoice, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.True(t, len(incidentConfigFieldChoicesGetResp.GetData()) >= 1)

}

const incidentsCharSet = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateUniqueString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = incidentsCharSet[seededRand.Intn(len(incidentsCharSet))]
	}
	return string(b)
}

func createIncidentConfigField(t *testing.T) datadog.IncidentConfigField {
	testIncidentConfigFieldData := datadog.NewIncidentConfigFieldWithDefaults()
	testIncidentConfigFieldData.SetType("field")
	testIncidentConfigFieldData.SetAttributes(*datadog.NewIncidentConfigFieldAttributesWithDefaults())
	testIncidentConfigFieldData.Attributes.SetType(datadog.INCIDENTCONFIGFIELDTYPE_DROPDOWN)
	testIncidentConfigFieldData.Attributes.SetTableId(datadog.INCIDENTCONFIGFIELDTABLE_INCIDENT)
	testIncidentConfigFieldData.Attributes.SetName(fmt.Sprintf("TestConfigField%s", generateUniqueString(32)))

	// Create IncidentConfig
	incidentConfigRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.CreateIncidentConfigField(TESTAUTH).
		Body(*datadog.NewIncidentConfigFieldPayload(*testIncidentConfigFieldData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigField %v: Response %s: %v", testIncidentConfigFieldData, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	return incidentConfigRsp.GetData()

}

func deleteIncidentConfigField(t *testing.T, id string) {
	//Delete IncidentConfig
	httpresp, err := TESTAPICLIENT.IncidentsConfigApi.DeleteIncidentConfigField(TESTAUTH, id).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting IncidentConfigField: %v failed with %v, "+
			"Another test may have already deleted this Config Field: %v",
			id, httpresp.StatusCode, err)
	}

}

func createIncidentConfigFieldChoice(t *testing.T, fieldID string) datadog.IncidentConfigFieldChoice {
	testIncidentConfigFieldChoiceData := datadog.NewIncidentConfigFieldChoiceWithDefaults()
	testIncidentConfigFieldChoiceData.SetType("choice")
	testIncidentConfigFieldChoiceData.SetAttributes(*datadog.NewIncidentConfigFieldChoiceAttributesWithDefaults())
	testIncidentConfigFieldChoiceData.Attributes.
		SetName(fmt.Sprintf("TestConfigFieldChoice%s", generateUniqueString(32)))
	testIncidentConfigFieldChoiceData.Attributes.SetValue("TestConfigChoiceValue")

	incidentConfigFieldChoiceRsp, httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		CreateIncidentConfigFieldChoice(TESTAUTH, fieldID).
		Body(*datadog.NewIncidentConfigFieldChoicePayload(*testIncidentConfigFieldChoiceData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentConfigFieldChoice %v: Response %s: %v", testIncidentConfigFieldChoiceData, err.Error(), bStr)
	}
	assert.Equal(t, httpresp.StatusCode, 201)
	return incidentConfigFieldChoiceRsp.GetData()

}

func deleteIncidentConfigFieldChoice(t *testing.T, fieldID, choiceID string) {
	httpresp, err := TESTAPICLIENT.IncidentsConfigApi.
		DeleteIncidentConfigFieldChoice(TESTAUTH, fieldID, choiceID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting IncidentConfigFieldChoice: %v failed with %v, "+
			"Another test may have already deleted this Config Field Choice: %v",
			choiceID, httpresp.StatusCode, err)
	}

}
