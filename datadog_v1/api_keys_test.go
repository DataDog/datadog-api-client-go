package datadog_v1_test

import (
	"encoding/json"
//	"fmt"
	// "os"
	// "reflect"
	"testing"

	datadog "github.com/DataDog/datadog-api-client-go/datadog_v1"
	"github.com/antihax/optional"
	gock "gopkg.in/h2non/gock.v1"
	"gotest.tools/assert"
)

func TestCreateApiKey(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var apiKeyResponse datadog.ApiKeyResponse
	json.Unmarshal(setupGock(t, "keys/api_key_get.json", "delete", "/api_key/"), &apiKeyResponse)
	apiKey := apiKeyResponse.GetApiKey()

	// Get mocked request data
	var createOpts datadog.CreateAPIKeyOpts
	apiKeySettings := *apiKeyResponse.GetApiKey().Name
	createOpts.ApiKey = optional.NewInterface(datadog.ApiKey{Name: &apiKeySettings})
	apiKeyResp, _, _ := TESTAPICLIENT.KeysApi.CreateAPIKey(TESTAUTH, &createOpts)

	apiKeyReturned := apiKeyResp.GetApiKey()
	assert.Equal(t, apiKey.GetName(), apiKeyReturned.GetName())
	assert.Equal(t, apiKey.GetKey(), apiKeyReturned.GetKey())
}

func TestGetApiKey(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var apiKeyResponse datadog.ApiKeyResponse
	json.Unmarshal(setupGock(t, "keys/api_key_get.json", "get", "/api_key"), &apiKeyResponse)
	apiKey := apiKeyResponse.GetApiKey()

	// Get mocked request data
	apiKeyResponseReturned, _, _ := TESTAPICLIENT.KeysApi.GetAPIKey(TESTAUTH, "b65bee9d3578485794a5b6069d73dfb7")

	apiKeyReturned := apiKeyResponseReturned.GetApiKey()
	assert.Equal(t, apiKey.GetCreated(), apiKeyReturned.GetCreated())
	assert.Equal(t, apiKey.GetCreatedBy(), apiKeyReturned.GetCreatedBy())
	assert.Equal(t, apiKey.GetKey(), apiKeyReturned.GetKey())
	assert.Equal(t, apiKey.GetName(), apiKeyReturned.GetName())
}

func TestGetListApiKey(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var listApiKeyResponse datadog.ApiKeyListResponse
	json.Unmarshal(setupGock(t, "keys/api_key_list_get.json", "get", "/api_key"), &listApiKeyResponse)
	listApiKey := listApiKeyResponse.GetApiKeys()
	singleApiKey := listApiKey[0]

	// Get mocked request data
	listApiKeyResponseReturned, _, _ := TESTAPICLIENT.KeysApi.GetAllAPIKeys(TESTAUTH)

	listApiKeyReturned := listApiKeyResponseReturned.GetApiKeys()
	assert.Equal(t, len(listApiKeyReturned), 2)
	
	singleApiKeyReturned := listApiKeyReturned[0]

	assert.Equal(t, singleApiKeyReturned.GetCreated(), singleApiKey.GetCreated())
	assert.Equal(t, singleApiKeyReturned.GetCreatedBy(), singleApiKey.GetCreatedBy())
	assert.Equal(t, singleApiKeyReturned.GetKey(), singleApiKey.GetKey())
	assert.Equal(t, singleApiKeyReturned.GetName(), singleApiKey.GetName())
}

func TestDeleteApiKey(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var apiKeyResponse datadog.ApiKeyResponse
	json.Unmarshal(setupGock(t, "keys/api_key_get.json", "get", "/api_key"), &apiKeyResponse)
	apiKey := apiKeyResponse.GetApiKey()

	// Get mocked request data
	apiKeyResponseReturned, _, _ := TESTAPICLIENT.KeysApi.DeleteAPIKey(TESTAUTH, "b65bee9d3578485794a5b6069d73dfb7")

	apiKeyReturned := apiKeyResponseReturned.GetApiKey()
	assert.Equal(t, apiKey.GetCreated(), apiKeyReturned.GetCreated())
	assert.Equal(t, apiKey.GetCreatedBy(), apiKeyReturned.GetCreatedBy())
	assert.Equal(t, apiKey.GetKey(), apiKeyReturned.GetKey())
	assert.Equal(t, apiKey.GetName(), apiKeyReturned.GetName())
}

// func TestEditApiKey(t *testing.T) {
// 	// Setup the Client we'll use to interact with the Test account
// 	teardownTest := setupUnitTest(t)
// 	defer teardownTest(t)
// 	defer gock.Off()

// 	// Setup fixture data
// 	apiKeyId := "b65bee9d3578485794a5b6069d73dfb7"
// 	var apiKeyResponse datadog.ApiKeyResponse
// 	json.Unmarshal(setupGock(t, "keys/api_key_get.json", "delete", fmt.Sprintf("/api_key/%s", apiKeyId)), &apiKeyResponse)
// 	apiKey := apiKeyResponse.GetApiKey()

// 	// Get mocked request data
// 	var updateOpts datadog.EditAPIKeyOpts
// 	updateOpts = optional.NewInterface(datadog.EditAPIKeyOpts{ApiKey: apiKey})
// 	apiKeyResponseReturned, _, err := TESTAPICLIENT.KeysApi.EditAPIKey(TESTAUTH, "b65bee9d3578485794a5b6069d73dfb7", &updateOpts)

// 	apiKeyReturned := apiKeyResponseReturned.GetApiKey()
// 	assert.Equal(t, apiKey.GetCreated(), apiKeyReturned.GetCreated())
// 	assert.Equal(t, apiKey.GetCreatedBy(), apiKeyReturned.GetCreatedBy())
// 	assert.Equal(t, apiKey.GetKey(), apiKeyReturned.GetKey())
// 	assert.Equal(t, apiKey.GetName(), apiKeyReturned.GetName())
// }
