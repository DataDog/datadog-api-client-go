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

	"gopkg.in/h2non/gock.v1"
)

func TestListOrgs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationListResponse
	json.Unmarshal(setupGock(ctx, t, "orgs/org_list.json", "get", "/org"), &orgsFixture)
	orgFixture := orgsFixture.GetOrgs()[0]

	// Get mocked request data
	orgs, _, err := Client(ctx).OrganizationsApi.ListOrgs(ctx).Execute()
	if err != nil {
		t.Errorf("Failed to Get the test org %s", err)
	}
	org := orgs.GetOrgs()[0]
	assert.Equal(orgFixture.GetName(), org.GetName())
	assert.Equal(orgFixture.GetPublicId(), org.GetPublicId())
	assert.Equal(orgFixture.GetSettings(), org.GetSettings())
	assert.Empty(org.GetCreated())
	assert.Empty((org.GetDescription()))

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := org.GetBilling()
	assert.Equal(orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubFixture := orgFixture.GetBilling()
	orgSubResp := org.GetBilling()
	assert.Equal(orgSubFixture.GetType(), orgSubResp.GetType())
}

func TestCreateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationCreateResponse
	json.Unmarshal(setupGock(ctx, t, "orgs/org_create.json", "post", "/org"), &orgsFixture)
	orgCreateBody := orgsFixture.GetOrg()

	// Get mocked request data
	createBody := datadog.OrganizationCreateBody{
		Name: orgCreateBody.GetName(),
	}
	createBody.SetSubscription(orgCreateBody.GetSubscription())
	createBody.SetBilling(orgCreateBody.GetBilling())
	getOrgResp, _, err := Client(ctx).OrganizationsApi.CreateChildOrg(ctx).Body(createBody).Execute()
	if err != nil {
		t.Errorf("Failed to create the test org %s", err)
	}

	// Assert User
	orgUserResp := getOrgResp.GetUser()
	orgUserFixture := orgsFixture.GetUser()
	assert.Equal(orgUserFixture.GetName(), orgUserResp.GetName())
	assert.Equal(orgUserFixture.GetHandle(), orgUserResp.GetHandle())

	// Assert API key fields
	apiKeyFixture := orgsFixture.GetApiKey()
	apiKeyResp := getOrgResp.GetApiKey()
	assert.Equal(apiKeyFixture.GetKey(), apiKeyResp.GetKey())
	assert.Equal(apiKeyFixture.GetCreated(), apiKeyResp.GetCreated())

	// Assert Application Key fields
	appKeyFixture := orgsFixture.GetApplicationKey()
	appKeyResp := getOrgResp.GetApplicationKey()
	assert.Equal(appKeyFixture.GetHash(), appKeyResp.GetHash())
	assert.Equal(appKeyFixture.GetName(), appKeyResp.GetName())
	assert.Equal(appKeyFixture.GetOwner(), appKeyResp.GetOwner())

	// Assert Org fields
	orgFixture := orgsFixture.GetOrg()
	orgResp := getOrgResp.GetOrg()
	assert.Equal(orgFixture.GetName(), orgResp.GetName())
	assert.Equal(orgFixture.GetPublicId(), orgResp.GetPublicId())
	assert.Equal(orgFixture.GetSettings(), orgResp.GetSettings())
	assert.Empty(orgResp.GetCreated())
	assert.Empty(orgResp.GetDescription())
	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := orgResp.GetBilling()
	assert.Equal(orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubFixture := orgFixture.GetBilling()
	orgSubResp := orgResp.GetBilling()
	assert.Equal(orgSubFixture.GetType(), orgSubResp.GetType())
}

func TestUpdateOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationResponse
	json.Unmarshal(setupGock(ctx, t, "orgs/org_update.json", "put", "/org"), &orgsFixture)

	// Get mocked request data
	updateOrgResp, _, err := Client(ctx).OrganizationsApi.UpdateOrg(ctx, *orgsFixture.GetOrg().PublicId).Body(datadog.Organization{Settings: orgsFixture.GetOrg().Settings}).Execute()
	if err != nil {
		t.Errorf("Failed to update the test org %s", err)
	}

	// Assert basic org fields
	orgFixture := orgsFixture.GetOrg()
	updateResp := updateOrgResp.GetOrg()
	assert.Equal(orgFixture.GetName(), updateResp.GetName())
	assert.Equal(orgFixture.GetPublicId(), updateResp.GetPublicId())
	assert.Equal(orgFixture.GetCreated(), updateResp.GetCreated())
	assert.Empty(updateResp.GetDescription())

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := updateResp.GetBilling()
	assert.Equal(orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubnFixture := orgFixture.GetSubscription()
	orgSubResp := updateResp.GetSubscription()
	assert.Equal(orgSubnFixture.GetType(), orgSubResp.GetType())

	// Assert Org Settings
	settingFixture := orgFixture.GetSettings()
	settingResp := updateResp.GetSettings()
	assert.Equal(settingFixture.GetSamlCanBeEnabled(), settingResp.GetSamlCanBeEnabled())

	assert.Equal(*settingFixture.GetSamlIdpInitiatedLogin().Enabled, *settingResp.GetSamlIdpInitiatedLogin().Enabled)

	assert.Equal(*settingFixture.GetSaml().Enabled, *settingResp.GetSaml().Enabled)
	assert.Equal(settingFixture.GetSamlIdpEndpoint(), settingResp.GetSamlIdpEndpoint())
	assert.Equal(settingFixture.GetSamlLoginUrl(), settingResp.GetSamlLoginUrl())
	assert.Equal(settingFixture.GetSamlIdpMetadataUploaded(), settingResp.GetSamlIdpMetadataUploaded())
	assert.Equal(*settingFixture.GetSamlStrictMode().Enabled, *settingResp.GetSamlStrictMode().Enabled)
	assert.Equal(*settingFixture.GetSamlAutocreateUsersDomains().Enabled, *settingResp.GetSamlAutocreateUsersDomains().Enabled)
	assert.Equal(len(*settingFixture.GetSamlAutocreateUsersDomains().Domains), len(*settingResp.GetSamlAutocreateUsersDomains().Domains))
}

func TestGetOrg(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	// Setup fixture data
	var orgsFixture datadog.OrganizationResponse
	json.Unmarshal(setupGock(ctx, t, "orgs/org_get.json", "get", "/org"), &orgsFixture)

	// Get mocked request data
	getOrgResp, _, err := Client(ctx).OrganizationsApi.GetOrg(ctx, *orgsFixture.GetOrg().PublicId).Execute()
	if err != nil {
		t.Errorf("Failed to get the test org %s", err)
	}

	// Assert basic org fields
	orgFixture := orgsFixture.GetOrg()
	getResp := getOrgResp.GetOrg()
	assert.Equal(orgFixture.GetName(), getResp.GetName())
	assert.Equal(orgFixture.GetPublicId(), getResp.GetPublicId())
	assert.Equal(orgFixture.GetCreated(), getResp.GetCreated())
	assert.Empty(getResp.GetDescription())

	orgBillingFixture := orgFixture.GetBilling()
	orgBillingResp := getResp.GetBilling()
	assert.Equal(orgBillingFixture.GetType(), orgBillingResp.GetType())

	orgSubnFixture := orgFixture.GetSubscription()
	orgSubResp := getResp.GetSubscription()
	assert.Equal(orgSubnFixture.GetType(), orgSubResp.GetType())

	// Assert Org Settings
	settingFixture := orgFixture.GetSettings()
	settingResp := getResp.GetSettings()
	assert.Equal(settingFixture.GetSamlCanBeEnabled(), settingResp.GetSamlCanBeEnabled())

	assert.Equal(*settingFixture.GetSamlIdpInitiatedLogin().Enabled, *settingResp.GetSamlIdpInitiatedLogin().Enabled)

	assert.Equal(*settingFixture.GetSaml().Enabled, *settingResp.GetSaml().Enabled)
	assert.Equal(settingFixture.GetSamlIdpEndpoint(), settingResp.GetSamlIdpEndpoint())
	assert.Equal(settingFixture.GetSamlLoginUrl(), settingResp.GetSamlLoginUrl())
	assert.Equal(settingFixture.GetSamlIdpMetadataUploaded(), settingResp.GetSamlIdpMetadataUploaded())
	assert.Equal(*settingFixture.GetSamlStrictMode().Enabled, *settingResp.GetSamlStrictMode().Enabled)
	assert.Equal(*settingFixture.GetSamlAutocreateUsersDomains().Enabled, *settingResp.GetSamlAutocreateUsersDomains().Enabled)
	assert.Equal(len(*settingFixture.GetSamlAutocreateUsersDomains().Domains), len(*settingResp.GetSamlAutocreateUsersDomains().Domains))
}

func TestUploadOrgIdpMeta(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	// Setup fixture data
	orgPubID := "12345"
	var idpResponseFixture datadog.IdpResponse
	json.Unmarshal(setupGock(ctx, t, "orgs/org_idp_upload.json", "post", fmt.Sprintf("/org/%s/idp_metadata", orgPubID)), &idpResponseFixture)

	// Get empty file object. This fixture doesn't exist since we don't need it to.
	file, _ := os.Open("test_go/idp_data.xml")

	idpResp, _, err := Client(ctx).OrganizationsApi.UploadIdPForOrg(ctx, orgPubID).IdpFile(file).Execute()
	if err != nil {
		t.Fatalf("Failed to update the test org's IDP meta %s", err)
	}

	assert.Equal(idpResponseFixture.GetMessage(), idpResp.GetMessage())
}

func TestOrgsCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.OrganizationCreateBody
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.OrganizationCreateBody{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.OrganizationCreateBody{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).OrganizationsApi.CreateChildOrg(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestOrgsListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).OrganizationsApi.ListOrgs(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestOrgsGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).OrganizationsApi.GetOrg(ctx, "lsqdkjf").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestOrgsUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).OrganizationsApi.UpdateOrg(ctx, "lsqdkjf").Body(datadog.Organization{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestOrgsUploadIdpErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		tcc := tc
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tcc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			// Get random file
			file, _ := os.Open("fixtures/orgs/error_415.json")

			_, httpresp, err := Client(ctx).OrganizationsApi.UploadIdPForOrg(ctx, *tests.UniqueEntityName(ctx, t)).IdpFile(file).Execute()
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tcc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestOrgsUploadIdp415Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/orgs/error_415.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/org/id/idp_metadata").Reply(415).JSON(res)
	defer gock.Off()

	// Get empty file object. This fixture doesn't exist since we don't need it to.
	file, _ := os.Open("test_go/idp_data.xml")

	_, httpresp, err := Client(ctx).OrganizationsApi.UploadIdPForOrg(ctx, "id").IdpFile(file).Execute()
	assert.Equal(415, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}
