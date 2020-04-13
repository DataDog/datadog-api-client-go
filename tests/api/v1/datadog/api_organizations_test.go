/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListOrgs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationListResponse
	json.Unmarshal(setupGock(t, "orgs/org_list.json", "get", "/org"), &orgsFixture)
	orgFixture := orgsFixture.GetOrgs()[0]

	// Get mocked request data
	orgs, _, err := TESTAPICLIENT.OrganizationsApi.ListOrgs(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Failed to Get the test org %s", err)
	}
	org := orgs.GetOrgs()[0]
	assert.Equal(t, orgFixture.GetName(), org.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), org.GetPublicId())
	assert.Equal(t, orgFixture.GetSettings(), org.GetSettings())
	assert.Empty(t, org.GetCreated())
	assert.Empty(t, (org.GetDescription()))

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
	var orgsFixture datadog.OrganizationCreateResponse
	json.Unmarshal(setupGock(t, "orgs/org_create.json", "post", "/org"), &orgsFixture)
	orgCreateBody := orgsFixture.GetOrg()

	// Get mocked request data
	createBody := datadog.OrganizationCreateBody{
		Name: orgCreateBody.GetName(),
	}
	createBody.SetSubscription(orgCreateBody.GetSubscription())
	createBody.SetBilling(orgCreateBody.GetBilling())
	getOrgResp, _, err := TESTAPICLIENT.OrganizationsApi.CreateChildOrg(TESTAUTH).Body(createBody).Execute()
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
	assert.Empty(t, orgResp.GetCreated())
	assert.Empty(t, orgResp.GetDescription())
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
	var orgsFixture datadog.OrganizationResponse
	json.Unmarshal(setupGock(t, "orgs/org_update.json", "put", "/org"), &orgsFixture)

	// Get mocked request data
	updateOrgResp, _, err := TESTAPICLIENT.OrganizationsApi.UpdateOrg(TESTAUTH, *orgsFixture.GetOrg().PublicId).Body(datadog.Organization{Settings: orgsFixture.GetOrg().Settings}).Execute()
	if err != nil {
		t.Errorf("Failed to update the test org %s", err)
	}

	// Assert basic org fields
	orgFixture := orgsFixture.GetOrg()
	updateResp := updateOrgResp.GetOrg()
	assert.Equal(t, orgFixture.GetName(), updateResp.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), updateResp.GetPublicId())
	assert.Equal(t, orgFixture.GetCreated(), updateResp.GetCreated())
	assert.Empty(t, updateResp.GetDescription())

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

func TestGetOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationResponse
	json.Unmarshal(setupGock(t, "orgs/org_get.json", "get", "/org"), &orgsFixture)

	// Get mocked request data
	getOrgResp, _, err := TESTAPICLIENT.OrganizationsApi.GetOrg(TESTAUTH, *orgsFixture.GetOrg().PublicId).Execute()
	if err != nil {
		t.Errorf("Failed to get the test org %s", err)
	}

	// Assert basic org fields
	orgFixture := orgsFixture.GetOrg()
	getResp := getOrgResp.GetOrg()
	assert.Equal(t, orgFixture.GetName(), getResp.GetName())
	assert.Equal(t, orgFixture.GetPublicId(), getResp.GetPublicId())
	assert.Equal(t, orgFixture.GetCreated(), getResp.GetCreated())
	assert.Empty(t, getResp.GetDescription())

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := getResp.GetBilling()
	assert.Equal(t, orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubnFixture := orgFixture.GetSubscription()
	orgSubResp := getResp.GetSubscription()
	assert.Equal(t, orgSubnFixture.GetType(), orgSubResp.GetType())

	// Assert Org Settings
	settingFixture := orgFixture.GetSettings()
	settingResp := getResp.GetSettings()
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

	idpResp, _, err := TESTAPICLIENT.OrganizationsApi.UploadIdPForOrg(TESTAUTH, orgPubID).IdpFile(file).Execute()
	if err != nil {
		t.Fatalf("Failed to update the test org's IDP meta %s", err)
	}

	assert.Equal(t, idpResponseFixture.GetMessage(), idpResp.GetMessage())
}

func TestOrgsCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.OrganizationCreateBody
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.OrganizationCreateBody{}, 400},
		{"403 Forbidden", fake_auth, datadog.OrganizationCreateBody{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.OrganizationsApi.CreateChildOrg(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestOrgsListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.OrganizationsApi.ListOrgs(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestOrgsGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.OrganizationsApi.GetOrg(tc.Ctx, "lsqdkjf").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestOrgsUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.OrganizationsApi.UpdateOrg(tc.Ctx, "lsqdkjf").Body(datadog.Organization{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestOrgsUploadIdpErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Get random file
	file, _ := os.Open("fixtures/orgs/error_415.json")

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.OrganizationsApi.UploadIdPForOrg(tc.Ctx, "lsqdkjf").IdpFile(file).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestOrgsUploadIdp415Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/orgs/error_415.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	gock.New("https://api.datadoghq.com").Post("/api/v1/org/id/idp_metadata").Reply(415).JSON(res)
	defer gock.Off()

	// Get empty file object. This fixture doesn't exist since we don't need it to.
	file, _ := os.Open("test_go/idp_data.xml")

	_, httpresp, err := TESTAPICLIENT.OrganizationsApi.UploadIdPForOrg(TESTAUTH, "id").IdpFile(file).Execute()
	assert.Equal(t, 415, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}
