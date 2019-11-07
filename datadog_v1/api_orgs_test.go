package datadog_v1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/antihax/optional"
	gock "gopkg.in/h2non/gock.v1"
	"gotest.tools/assert"
)

func TestGetOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture OrgListResponse
	json.Unmarshal(setupGock(t, "orgs/org_get.json", "get", "/org"), &orgsFixture)
	orgFixture := orgsFixture.GetOrgs()[0]

	// Get mocked request data
	orgs, _, err := TESTAPICLIENT.OrgsApi.GetOrg(TESTAUTH)
	if err != nil {
		t.Errorf("Failed to Get the test org %s", err)
	}
	org := orgs.GetOrgs()[0]
	assert.Equal(t, orgFixture.GetName(), org.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), org.GetPublicId())
	assert.Equal(t, reflect.DeepEqual(orgFixture.GetSubscription(), org.GetSubscription()), true)
	assert.Equal(t, reflect.DeepEqual(orgFixture.GetBilling(), org.GetBilling()), true)
}

func TestCreateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture OrgCreateResponse
	json.Unmarshal(setupGock(t, "orgs/org_create.json", "post", "/org"), &orgsFixture)
	orgCreateBody := orgsFixture.GetOrg()
	orgUserFixture := orgsFixture.GetUser()

	// Get mocked request data
	createBody := OrgCreateBody{
		Name: PtrString((&orgCreateBody).GetName()),
	}
	createBody.SetSubscription(orgCreateBody.GetSubscription())
	createBody.SetBilling(map[string]string{})
	orgResp, _, err := TESTAPICLIENT.OrgsApi.CreateChildOrg(TESTAUTH, createBody)
	if err != nil {
		t.Errorf("Failed to create the test org %s", err)
	}
	orgUserResp := orgResp.GetUser()

	assert.Equal(t, orgUserFixture.GetName(), orgUserResp.GetName())
	assert.Equal(t, orgUserFixture.GetHandle(), orgUserResp.GetHandle())
	assert.Equal(t, reflect.DeepEqual(orgsFixture.GetApiKey(), orgResp.GetApiKey()), true)
	assert.Equal(t, reflect.DeepEqual(orgsFixture.GetApplicationKey(), orgResp.GetApplicationKey()), true)
	assert.Equal(t, reflect.DeepEqual(orgsFixture.GetOrg(), orgResp.GetOrg()), true)
}

func TestUpdateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture OrgResponse
	json.Unmarshal(setupGock(t, "orgs/org_update.json", "put", "/org"), &orgsFixture)

	// Get mocked request data
	var updateOpts UpdateOrgOpts
	orgSettings := *orgsFixture.GetOrg().Settings
	updateOpts.Org = optional.NewInterface(Org{Settings: &orgSettings})
	orgResp, _, err := TESTAPICLIENT.OrgsApi.UpdateOrg(TESTAUTH, *orgsFixture.GetOrg().PublicId, &updateOpts)
	if err != nil {
		t.Errorf("Failed to update the test org %s", err)
	}

	assert.Equal(t, reflect.DeepEqual(orgsFixture, orgResp), true)
}

// [TODO] Implement this test
func TestUploadOrgIdpMeta(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	orgPubID := "12345"
	var iDPResponseFixture IdpResponse
	json.Unmarshal(setupGock(t, "orgs/org_idp_upload.json", "post", fmt.Sprintf("/org/%s/idp_metadata", orgPubID)), &iDPResponseFixture)

	// Get mocked request data
	var updateOpts UploadIdPForOrgOpts
	updateOpts.Body = optional.NewString("metadata_string")
	idpResp, _, err := TESTAPICLIENT.OrgsApi.UploadIdPForOrg(TESTAUTH, orgPubID, &updateOpts)
	if err != nil {
		t.Errorf("Failed to update the test org's IDP meta %s", err)
	}

	assert.Equal(t, reflect.DeepEqual(iDPResponseFixture, idpResp), true)
}
