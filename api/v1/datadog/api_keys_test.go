package datadog_test

import (
	"testing"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/antihax/optional"
	"gotest.tools/assert"
)

func TestApiKeyFunctions(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create API Key
	// ----------------------------------
	var createOpts datadog.CreateAPIKeyOpts
	testApiKeyName := "api key name"
	createOpts.ApiKey = optional.NewInterface(datadog.ApiKey{Name: &testApiKeyName})
	respData, respCode, err := TESTAPICLIENT.KeysApi.CreateAPIKey(TESTAUTH, &createOpts)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error creating api key %v: Status: %v: %v", testApiKeyName, respCode.StatusCode, err)
	}

	createApiKeyReturned := respData.GetApiKey()
	createApiKeyName := createApiKeyReturned.GetName()
	createApiKeyCreated := createApiKeyReturned.GetCreated()
	createApiKeyCreatedBy := createApiKeyReturned.GetCreatedBy()
	createApiKeyValue := createApiKeyReturned.GetKey()

	// none of these values should be empty
	assert.Assert(t, createApiKeyName != "")
	assert.Assert(t, createApiKeyCreated != "")
	assert.Assert(t, createApiKeyCreatedBy != "")
	assert.Assert(t, createApiKeyValue != "")
	assert.Equal(t, createApiKeyName, testApiKeyName)

	// Get API Key
	// ----------------------------------
	respData, respCode, err = TESTAPICLIENT.KeysApi.GetAPIKey(TESTAUTH, createApiKeyValue)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error getting api key %v: Status: %v: %v", createApiKeyValue, respCode.StatusCode, err)
	}

	getApiKeyReturned := respData.GetApiKey()
	getApiKeyName := getApiKeyReturned.GetName()
	getApiKeyCreated := getApiKeyReturned.GetCreated()
	getApiKeyCreatedBy := getApiKeyReturned.GetCreatedBy()
	getApiKeyValue := getApiKeyReturned.GetKey()

	// should be the same as the create
	assert.Equal(t, createApiKeyName, getApiKeyName)
	assert.Equal(t, createApiKeyCreated, getApiKeyCreated)
	assert.Equal(t, createApiKeyCreatedBy, getApiKeyCreatedBy)
	assert.Equal(t, createApiKeyValue, getApiKeyValue)

	// Get All API Keys
	// ----------------------------------
	respListData, respCode, err := TESTAPICLIENT.KeysApi.GetAllAPIKeys(TESTAUTH)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error getting all api keys: Status: %v: %v", respCode.StatusCode, err)
	}
	getAllApiKeyReturned := respListData.GetApiKeys()

	// should have more than 1 at least
	assert.Assert(t, len(getAllApiKeyReturned) > 1)

	// Edit API Key
	// ----------------------------------
	var editOpts datadog.EditAPIKeyOpts
	newApiKeyName := "new api key name"
	editOpts.ApiKey = optional.NewInterface(datadog.ApiKey{Name: &newApiKeyName})
	respData, respCode, err = TESTAPICLIENT.KeysApi.EditAPIKey(TESTAUTH, createApiKeyValue, &editOpts)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error editing api key %v: Status: %v: %v", createApiKeyValue, respCode.StatusCode, err)
	}

	editApiKeyReturned := respData.GetApiKey()
	editApiKeyName := editApiKeyReturned.GetName()
	editApiKeyCreated := editApiKeyReturned.GetCreated()
	editApiKeyCreatedBy := editApiKeyReturned.GetCreatedBy()
	editApiKeyValue := editApiKeyReturned.GetKey()

	// everything should be the same except the name
	assert.Assert(t, editApiKeyName != getApiKeyName)
	assert.Equal(t, editApiKeyCreated, getApiKeyCreated)
	assert.Equal(t, editApiKeyCreatedBy, getApiKeyCreatedBy)
	assert.Equal(t, editApiKeyValue, getApiKeyValue)

	// Delete API Key
	// ----------------------------------
	respData, respCode, err = TESTAPICLIENT.KeysApi.DeleteAPIKey(TESTAUTH, createApiKeyValue)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error deleting api key %v: Status: %v: %v", createApiKeyValue, respCode.StatusCode, err)
	}

	deleteApiKeyReturned := respData.GetApiKey()
	deleteApiKeyName := deleteApiKeyReturned.GetName()
	deleteApiKeyCreated := deleteApiKeyReturned.GetCreated()
	deleteApiKeyCreatedBy := deleteApiKeyReturned.GetCreatedBy()
	deleteApiKeyValue := deleteApiKeyReturned.GetKey()

	// should return the key thats been deleted
	assert.Equal(t, deleteApiKeyName, editApiKeyName)
	assert.Equal(t, deleteApiKeyCreated, editApiKeyCreated)
	assert.Equal(t, deleteApiKeyCreatedBy, editApiKeyCreatedBy)
	assert.Equal(t, deleteApiKeyValue, editApiKeyValue)
}

func TestApplicationKeyFunctions(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create Application Key
	// ----------------------------------
	var createOpts datadog.CreateApplicationKeyOpts
	testAppKeyName := "app key name"
	createOpts.ApplicationKey = optional.NewInterface(datadog.ApplicationKey{Name: &testAppKeyName})
	respData, respCode, err := TESTAPICLIENT.KeysApi.CreateApplicationKey(TESTAUTH, &createOpts)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error creating api key %v: Status: %v: %v", testAppKeyName, respCode.StatusCode, err)
	}

	createAppKeyReturned := respData.GetApplicationKey()
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
	respData, respCode, err = TESTAPICLIENT.KeysApi.GetApplicationKey(TESTAUTH, createAppKeyHash)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error getting app key %v: Status: %v: %v", createAppKeyHash, respCode.StatusCode, err)
	}

	getAppKeyReturned := respData.GetApplicationKey()
	getAppKeyOwner := getAppKeyReturned.GetOwner()
	getAppKeyHash := getAppKeyReturned.GetHash()
	getAppKeyName := getAppKeyReturned.GetName()

	// should be the same as the create
	assert.Equal(t, createAppKeyOwner, getAppKeyOwner)
	assert.Equal(t, createAppKeyHash, getAppKeyHash)
	assert.Equal(t, createAppKeyName, getAppKeyName)

	// Get All Application Keys
	// ----------------------------------
	respListData, respCode, err := TESTAPICLIENT.KeysApi.GetAllApplicationKeys(TESTAUTH)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error getting all app keys: Status: %v: %v", respCode.StatusCode, err)
	}
	getAllAppKeyReturned := respListData.GetApplicationKeys()

	// should have more than one at least
	assert.Assert(t, len(getAllAppKeyReturned) > 1)

	// Edit Application Key
	// ----------------------------------
	var editOpts datadog.EditApplicationKeyOpts
	newAppKeyName := "new app key name"
	editOpts.ApplicationKey = optional.NewInterface(datadog.ApplicationKey{Name: &newAppKeyName})
	respData, respCode, err = TESTAPICLIENT.KeysApi.EditApplicationKey(TESTAUTH, getAppKeyHash, &editOpts)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error editing app key %v: Status: %v: %v", getAppKeyHash, respCode.StatusCode, err)
	}

	editAppKeyReturned := respData.GetApplicationKey()
	editAppKeyOwner := editAppKeyReturned.GetOwner()
	editAppKeyHash := editAppKeyReturned.GetHash()
	editAppKeyName := editAppKeyReturned.GetName()

	// everything should be the same except the name
	assert.Assert(t, editAppKeyName != getAppKeyName)
	assert.Equal(t, editAppKeyOwner, getAppKeyOwner)
	assert.Equal(t, editAppKeyHash, getAppKeyHash)

	// Delete Application Key
	// ----------------------------------
	respData, respCode, err = TESTAPICLIENT.KeysApi.DeleteApplicationKey(TESTAUTH, createAppKeyHash)
	if err != nil || respCode.StatusCode != 200 {
		t.Errorf("Error deleting app key %v: Status: %v: %v", createAppKeyHash, respCode.StatusCode, err)
	}

	deleteAppKeyReturned := respData.GetApplicationKey()
	deleteAppKeyOwner := deleteAppKeyReturned.GetOwner()
	deleteAppKeyHash := deleteAppKeyReturned.GetHash()
	deleteAppKeyName := deleteAppKeyReturned.GetName()

	// should return the deleted app key
	assert.Equal(t, deleteAppKeyOwner, editAppKeyOwner)
	assert.Equal(t, deleteAppKeyHash, editAppKeyHash)
	assert.Equal(t, deleteAppKeyName, editAppKeyName)
}
