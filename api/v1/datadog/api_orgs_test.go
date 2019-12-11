package datadog_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	gock "gopkg.in/h2non/gock.v1"
	"gotest.tools/assert"
)

func TestGetOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrgListResponse
	json.Unmarshal(setupGock(t, "orgs/org_get.json", "get", "/org"), &orgsFixture)
	orgFixture := orgsFixture.GetOrgs()[0]

	// Get mocked request data
	orgs, _, err := TESTAPICLIENT.OrgsApi.GetOrg(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Failed to Get the test org %s", err)
	}
	org := orgs.GetOrgs()[0]
	assert.Equal(t, orgFixture.GetName(), org.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), org.GetPublicId())
	assert.Equal(t, orgFixture.GetSettings(), org.GetSettings())
	assert.Equal(t, org.GetCreated(), "")
	assert.Equal(t, (org.GetDescription()), "")

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := org.GetBilling()
	assert.Equal(t, orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubFixture := orgFixture.GetBilling()
	orgSubResp := org.GetBilling()
	assert.Equal(t, orgSubFixture.GetType(), orgSubResp.GetType())
}

func TestCreateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrgCreateResponse
	json.Unmarshal(setupGock(t, "orgs/org_create.json", "post", "/org"), &orgsFixture)
	orgCreateBody := orgsFixture.GetOrg()

	// Get mocked request data
	createBody := datadog.OrgCreateBody{
		Name: orgCreateBody.GetName(),
	}
	createBody.SetSubscription(orgCreateBody.GetSubscription())
	createBody.SetBilling(orgCreateBody.GetBilling())
	getOrgResp, _, err := TESTAPICLIENT.OrgsApi.CreateChildOrg(TESTAUTH).OrgCreateBody(&createBody).Execute()
	if err != nil {
		t.Errorf("Failed to create the test org %s", err)
	}

	// Assert User
	orgUserResp := getOrgResp.GetUser()
	orgUserFixture := orgsFixture.GetUser()
	assert.Equal(t, orgUserFixture.GetName(), orgUserResp.GetName())
	assert.Equal(t, orgUserFixture.GetHandle(), orgUserResp.GetHandle())

	// Assert API key fields
	apiKeyFixture := orgsFixture.GetApiKey()
	apiKeyResp := getOrgResp.GetApiKey()
	assert.Equal(t, apiKeyFixture.GetKey(), apiKeyResp.GetKey())
	assert.Equal(t, apiKeyFixture.GetCreated(), apiKeyResp.GetCreated())

	// Assert Application Key fields
	appKeyFixture := orgsFixture.GetApplicationKey()
	appKeyResp := getOrgResp.GetApplicationKey()
	assert.Equal(t, appKeyFixture.GetHash(), appKeyResp.GetHash())
	assert.Equal(t, appKeyFixture.GetName(), appKeyResp.GetName())
	assert.Equal(t, appKeyFixture.GetOwner(), appKeyResp.GetOwner())

	// Assert Org fields
	orgFixture := orgsFixture.GetOrg()
	orgResp := getOrgResp.GetOrg()
	assert.Equal(t, orgFixture.GetName(), orgResp.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), orgResp.GetPublicId())
	assert.Equal(t, orgFixture.GetSettings(), orgResp.GetSettings())
	assert.Equal(t, orgResp.GetCreated(), "")
	assert.Equal(t, (orgResp.GetDescription()), "")
	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := orgResp.GetBilling()
	assert.Equal(t, orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubFixture := orgFixture.GetBilling()
	orgSubResp := orgResp.GetBilling()
	assert.Equal(t, orgSubFixture.GetType(), orgSubResp.GetType())
}

func TestUpdateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrgResponse
	json.Unmarshal(setupGock(t, "orgs/org_update.json", "put", "/org"), &orgsFixture)

	// Get mocked request data
	updateOrgResp, _, err := TESTAPICLIENT.OrgsApi.UpdateOrg(TESTAUTH, *orgsFixture.GetOrg().PublicId).Org(&datadog.Org{Settings: orgsFixture.GetOrg().Settings}).Execute()
	if err != nil {
		t.Errorf("Failed to update the test org %s", err)
	}

	// Assert basic org fields
	orgFixture := orgsFixture.GetOrg()
	updateResp := updateOrgResp.GetOrg()
	assert.Equal(t, orgFixture.GetName(), updateResp.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), updateResp.GetPublicId())
	assert.Equal(t, updateResp.GetCreated(), updateResp.GetCreated())
	assert.Equal(t, (updateResp.GetDescription()), "")

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := updateResp.GetBilling()
	assert.Equal(t, orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubnFixture := orgFixture.GetSubscription()
	orgSubResp := updateResp.GetSubscription()
	assert.Equal(t, orgSubnFixture.GetType(), orgSubResp.GetType())

	// Assert Org Settings
	settingFixture := orgFixture.GetSettings()
	settingResp := updateResp.GetSettings()
	assert.Equal(t, settingFixture.GetSamlCanBeEnabled(), settingResp.GetSamlCanBeEnabled())

	assert.Equal(t, *settingFixture.GetSamlIdpInitiatedLogin().Enabled, *settingResp.GetSamlIdpInitiatedLogin().Enabled)

	assert.Equal(t, *settingFixture.GetSaml().Enabled, *settingResp.GetSaml().Enabled)
	assert.Equal(t, settingFixture.GetSamlIdpEndpoint(), settingResp.GetSamlIdpEndpoint())
	assert.Equal(t, settingFixture.GetSamlLoginUrl(), settingResp.GetSamlLoginUrl())
	assert.Equal(t, settingFixture.GetSamlIdpMetadataUploaded(), settingResp.GetSamlIdpMetadataUploaded())
	assert.Equal(t, *settingFixture.GetSamlStrictMode().Enabled, *settingResp.GetSamlStrictMode().Enabled)
	assert.Equal(t, *settingFixture.GetSamlAutocreateUsersDomains().Enabled, *settingResp.GetSamlAutocreateUsersDomains().Enabled)
	assert.Equal(t, len(*settingFixture.GetSamlAutocreateUsersDomains().Domains), len(*settingResp.GetSamlAutocreateUsersDomains().Domains))
}

func TestUploadOrgIdpMeta(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	orgPubID := "12345"
	var idpResponseFixture datadog.IdpResponse
	json.Unmarshal(setupGock(t, "orgs/org_idp_upload.json", "post", fmt.Sprintf("/org/%s/idp_metadata", orgPubID)), &idpResponseFixture)

	// Get empty file object. This fixture doesn't exist since we don't need it to.
	file, _ := os.Open("test_go/idp_data.xml")

	idpResp, _, err := TESTAPICLIENT.OrgsApi.UploadIdPForOrg(TESTAUTH, orgPubID).IdpFile(&file).Execute()
	if err != nil {
		t.Errorf("Failed to update the test org's IDP meta %s", err)
	}

	assert.Equal(t, idpResponseFixture.GetMessage(), idpResp.GetMessage())
}
