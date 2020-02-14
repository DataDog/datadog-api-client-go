// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestApiKeyFunctions(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create API Key
	// ----------------------------------
	testAPIKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	apiKeyData, httpresp, err := TESTAPICLIENT.KeysApi.CreateAPIKey(TESTAUTH).Body(datadog.ApiKey{Name: &testAPIKeyName}).Execute()
	if err != nil {
		t.Errorf("Error creating api key %v: Response %s: %v", testAPIKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAPIKey(apiKeyData.ApiKey.GetKey())
	assert.Equal(t, httpresp.StatusCode, 200)

	createAPIKeyReturned := apiKeyData.GetApiKey()
	createAPIKeyName := createAPIKeyReturned.GetName()
	createAPIKeyCreated := createAPIKeyReturned.GetCreated()
	createAPIKeyCreatedBy := createAPIKeyReturned.GetCreatedBy()
	createAPIKeyValue := createAPIKeyReturned.GetKey()

	// none of these values should be empty
	assert.Assert(t, createAPIKeyName != "")
	assert.Assert(t, createAPIKeyCreated != "")
	assert.Assert(t, createAPIKeyCreatedBy != "")
	assert.Assert(t, createAPIKeyValue != "")
	assert.Equal(t, createAPIKeyName, testAPIKeyName)

	// Get API Key
	// ----------------------------------
	apiKeyData, httpresp, err = TESTAPICLIENT.KeysApi.GetAPIKey(TESTAUTH, createAPIKeyValue).Execute()
	if err != nil {
		t.Errorf("Error getting api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	getAPIKeyReturned := apiKeyData.GetApiKey()
	getAPIKeyName := getAPIKeyReturned.GetName()
	getAPIKeyCreated := getAPIKeyReturned.GetCreated()
	getAPIKeyCreatedBy := getAPIKeyReturned.GetCreatedBy()
	getAPIKeyValue := getAPIKeyReturned.GetKey()

	// should be the same as the create
	assert.Equal(t, createAPIKeyName, getAPIKeyName)
	assert.Equal(t, createAPIKeyCreated, getAPIKeyCreated)
	assert.Equal(t, createAPIKeyCreatedBy, getAPIKeyCreatedBy)
	assert.Equal(t, createAPIKeyValue, getAPIKeyValue)

	// Get All API Keys
	// ----------------------------------
	respListData, httpresp, err := TESTAPICLIENT.KeysApi.GetAllAPIKeys(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error getting all api keys: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	getAllAPIKeyReturned := respListData.GetApiKeys()

	// should have more than 1 at least
	assert.Assert(t, len(getAllAPIKeyReturned) > 1)

	// Edit API Key
	// ----------------------------------
	newAPIKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	apiKeyData, httpresp, err = TESTAPICLIENT.KeysApi.EditAPIKey(TESTAUTH, createAPIKeyValue).Body(datadog.ApiKey{Name: &newAPIKeyName}).Execute()
	if err != nil {
		t.Errorf("Error editing api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	editAPIKeyReturned := apiKeyData.GetApiKey()
	editAPIKeyName := editAPIKeyReturned.GetName()
	editAPIKeyCreated := editAPIKeyReturned.GetCreated()
	editAPIKeyCreatedBy := editAPIKeyReturned.GetCreatedBy()
	editAPIKeyValue := editAPIKeyReturned.GetKey()

	// everything should be the same except the name
	assert.Assert(t, editAPIKeyName != getAPIKeyName)
	assert.Equal(t, editAPIKeyCreated, getAPIKeyCreated)
	assert.Equal(t, editAPIKeyCreatedBy, getAPIKeyCreatedBy)
	assert.Equal(t, editAPIKeyValue, getAPIKeyValue)

	// Delete API Key
	// ----------------------------------
	apiKeyData, httpresp, err = TESTAPICLIENT.KeysApi.DeleteAPIKey(TESTAUTH, createAPIKeyValue).Execute()
	if err != nil {
		t.Errorf("Error deleting api key %v: Response %s: %v", createAPIKeyValue, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	deleteAPIKeyReturned := apiKeyData.GetApiKey()
	deleteAPIKeyName := deleteAPIKeyReturned.GetName()
	deleteAPIKeyCreated := deleteAPIKeyReturned.GetCreated()
	deleteAPIKeyCreatedBy := deleteAPIKeyReturned.GetCreatedBy()
	deleteAPIKeyValue := deleteAPIKeyReturned.GetKey()

	// should return the key thats been deleted
	assert.Equal(t, deleteAPIKeyName, editAPIKeyName)
	assert.Equal(t, deleteAPIKeyCreated, editAPIKeyCreated)
	assert.Equal(t, deleteAPIKeyCreatedBy, editAPIKeyCreatedBy)
	assert.Equal(t, deleteAPIKeyValue, editAPIKeyValue)
}

func TestApplicationKeyFunctions(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create Application Key
	// ----------------------------------
	testAppKeyName := fmt.Sprintf("%s:%d", t.Name(), time.Now().UnixNano())
	appKeyData, httpresp, err := TESTAPICLIENT.KeysApi.CreateApplicationKey(TESTAUTH).Body(datadog.ApplicationKey{Name: &testAppKeyName}).Execute()
	if err != nil {
		t.Errorf("Error creating api key %v: Response %s: %v", testAppKeyName, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteAppKey(appKeyData.ApplicationKey.GetHash())
	assert.Equal(t, httpresp.StatusCode, 200)

	createAppKeyReturned := appKeyData.GetApplicationKey()
	createAppKeyOwner := createAppKeyReturned.GetOwner()
	createAppKeyHash := createAppKeyReturned.GetHash()
	createAppKeyName := createAppKeyReturned.GetName()

	// all values should not be nil
	assert.Assert(t, createAppKeyOwner != "")
	assert.Assert(t, createAppKeyHash != "")
	assert.Assert(t, createAppKeyName != "")
	assert.Equal(t, createAppKeyName, testAppKeyName)

	// Get Application Key
	// ----------------------------------
	appKeyData, httpresp, err = TESTAPICLIENT.KeysApi.GetApplicationKey(TESTAUTH, createAppKeyHash).Execute()
	if err != nil {
		t.Errorf("Error getting app key %v: Response %s: %v", createAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	getAppKeyReturned := appKeyData.GetApplicationKey()
	getAppKeyOwner := getAppKeyReturned.GetOwner()
	getAppKeyHash := getAppKeyReturned.GetHash()
	getAppKeyName := getAppKeyReturned.GetName()

	// should be the same as the create
	assert.Equal(t, createAppKeyOwner, getAppKeyOwner)
	assert.Equal(t, createAppKeyHash, getAppKeyHash)
	assert.Equal(t, createAppKeyName, getAppKeyName)

	// Get All Application Keys
	// ----------------------------------
	respListData, httpresp, err := TESTAPICLIENT.KeysApi.GetAllApplicationKeys(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error getting all app keys: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	getAllAppKeyReturned := respListData.GetApplicationKeys()

	// should have more than one at least
	assert.Assert(t, len(getAllAppKeyReturned) > 1)

	// Edit Application Key
	// ----------------------------------
	newAppKeyName := fmt.Sprintf("New %s:%d", t.Name(), time.Now().UnixNano())
	appKeyData, httpresp, err = TESTAPICLIENT.KeysApi.EditApplicationKey(TESTAUTH, getAppKeyHash).Body(datadog.ApplicationKey{Name: &newAppKeyName}).Execute()
	if err != nil {
		t.Errorf("Error editing app key %v: Response %s: %v", getAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	editAppKeyReturned := appKeyData.GetApplicationKey()
	editAppKeyOwner := editAppKeyReturned.GetOwner()
	editAppKeyHash := editAppKeyReturned.GetHash()
	editAppKeyName := editAppKeyReturned.GetName()

	// everything should be the same except the name
	assert.Assert(t, editAppKeyName != getAppKeyName)
	assert.Equal(t, editAppKeyOwner, getAppKeyOwner)
	assert.Equal(t, editAppKeyHash, getAppKeyHash)

	// Delete Application Key
	// ----------------------------------
	appKeyData, httpresp, err = TESTAPICLIENT.KeysApi.DeleteApplicationKey(TESTAUTH, createAppKeyHash).Execute()
	if err != nil {
		t.Errorf("Error deleting app key %v: Response %s: %v", createAppKeyHash, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	deleteAppKeyReturned := appKeyData.GetApplicationKey()
	deleteAppKeyOwner := deleteAppKeyReturned.GetOwner()
	deleteAppKeyHash := deleteAppKeyReturned.GetHash()
	deleteAppKeyName := deleteAppKeyReturned.GetName()

	// should return the deleted app key
	assert.Equal(t, deleteAppKeyOwner, editAppKeyOwner)
	assert.Equal(t, deleteAppKeyHash, editAppKeyHash)
	assert.Equal(t, deleteAppKeyName, editAppKeyName)
}

func deleteAPIKey(apiKeyValue string) {
	_, httpresp, err := TESTAPICLIENT.KeysApi.DeleteAPIKey(TESTAUTH, apiKeyValue).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting api key: %v failed with %v, Another test may have already deleted this api key.", apiKeyValue[len(apiKeyValue)-4:], httpresp.StatusCode)
	}
}

func deleteAppKey(appKeyHash string) {
	_, httpresp, err := TESTAPICLIENT.KeysApi.DeleteApplicationKey(TESTAUTH, appKeyHash).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting app key: %v failed with %v, Another test may have already deleted this app key.", appKeyHash[len(appKeyHash)-4:], httpresp.StatusCode)
	}
}
