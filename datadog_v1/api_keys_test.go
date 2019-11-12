package datadog_v1_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	datadog "github.com/DataDog/datadog-api-client-go/datadog_v1"
	"github.com/antihax/optional"
	gock "gopkg.in/h2non/gock.v1"
	"gotest.tools/assert"
)

func TestGetApiKeys(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var apiKeyListFixture datadog.ApiKeyListResponse
	json.Unmarshal(setupGock(t, "keys/api_keys_get.json", "get", "/api_key"), &apiKeyListFixture)
	apiKeyListFixture := apiKeyListFixture.GetAllAPIKeys()[0]

	// Get mocked request data
	apiKeys, _, err := TESTAPICLIENT.KeysApi.GetAllAPIKeys(TESTAUTH)
	if err != nil {
		t.Errorf("Failed to Get the test org %s", err)
	}
	// org := orgs.GetOrgs()[0]
	// assert.Equal(t, orgFixture.GetName(), org.GetName())
	// assert.Equal(t, orgFixture.GetPublicId(), org.GetPublicId())
	// assert.Equal(t, reflect.DeepEqual(orgFixture.GetSubscription(), org.GetSubscription()), true)
	// assert.Equal(t, reflect.DeepEqual(orgFixture.GetBilling(), org.GetBilling()), true)
}
