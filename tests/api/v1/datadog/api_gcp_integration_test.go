/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/api/common"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func generateUniqueGCPAccount(ctx context.Context, t *testing.T) (datadog.GCPAccount, datadog.GCPAccount) {
	name := tests.UniqueEntityName(ctx, t)
	var testGCPAccount = datadog.GCPAccount{
		Type:                    common.PtrString("service_account"),
		ProjectId:               name,
		PrivateKeyId:            common.PtrString("fake_private_key_id"),
		PrivateKey:              common.PtrString("fake_key"),
		ClientEmail:             common.PtrString(fmt.Sprintf("%s@fake-sandbox.iam.gserviceaccount.com", *name)),
		ClientId:                common.PtrString("123456712345671234567"),
		AuthUri:                 common.PtrString("fake_uri"),
		TokenUri:                common.PtrString("fake_uri"),
		AuthProviderX509CertUrl: common.PtrString("fake_url"),
		ClientX509CertUrl:       common.PtrString("fake_url"),
		HostFilters:             common.PtrString("fake:tag,example:test"),
		Automute:                common.PtrBool(false),
	}

	var testUpdateGCPAccount = datadog.GCPAccount{
		ProjectId:   common.PtrString(testGCPAccount.GetProjectId()),
		ClientEmail: testGCPAccount.ClientEmail,
		HostFilters: common.PtrString("fake:update,example:update"),
		Automute:    common.PtrBool(true),
	}

	return testGCPAccount, testUpdateGCPAccount
}

func TestGCPCreate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadog.GCPIntegrationApi(Client(ctx))

	testGCPAcct, _ := generateUniqueGCPAccount(ctx, t)
	defer uninstallGCPIntegration(ctx, t, testGCPAcct)

	_, httpresp, err := api.CreateGCPIntegration(ctx, testGCPAcct)
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestGCPListandDelete(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadog.GCPIntegrationApi(Client(ctx))

	testGCPAcct, _ := generateUniqueGCPAccount(ctx, t)
	defer uninstallGCPIntegration(ctx, t, testGCPAcct)

	// Setup GCP Account to List
	api.CreateGCPIntegration(ctx, testGCPAcct)

	gcpListOutput, httpresp, err := api.ListGCPIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing GCP Accounts: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetProjectId() == *testGCPAcct.ProjectId {
			x = Account
		}
	}
	assert.Equal(*testGCPAcct.ClientEmail, x.GetClientEmail())
	assert.Equal(*testGCPAcct.HostFilters, x.GetHostFilters())

	// Assert returned list is greater than or equal to 1
	assert.True(len(gcpListOutput) >= 1)

	// Test account deletion as well
	_, httpresp, err = api.DeleteGCPIntegration(ctx, testGCPAcct)
	if err != nil {
		t.Fatalf("Error uninstalling GCP Account: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestUpdateGCPAccount(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadog.GCPIntegrationApi(Client(ctx))

	testGCPAcct, testGCPUpdateAcct := generateUniqueGCPAccount(ctx, t)
	defer uninstallGCPIntegration(ctx, t, testGCPAcct)

	// Setup GCP Account to Update
	_, httpresp, err := api.CreateGCPIntegration(ctx, testGCPAcct)
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = api.UpdateGCPIntegration(ctx, testGCPUpdateAcct)
	if err != nil {
		t.Fatalf("Error updating GCP integration: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// List account to ensure update worked.
	gcpListOutput, _, err := api.ListGCPIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing GCP accounts: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetClientEmail() == *testGCPAcct.ClientEmail {
			x = Account
		}
	}
	assert.Equal(true, x.GetAutomute())
	assert.Equal(*testGCPUpdateAcct.HostFilters, x.GetHostFilters())
}

func TestGCPList400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadog.GCPIntegrationApi(Client(ctx))

	res, err := tests.ReadFixture("fixtures/gcp/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the gcp integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "TODO")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/integration/gcp").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := api.ListGCPIntegration(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestGCPListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadog.GCPIntegrationApi(Client(ctx))

			_, httpresp, err := api.ListGCPIntegration(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGCPCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.GCPIntegrationApi(Client(ctx))

			_, httpresp, err := api.CreateGCPIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGCPDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.GCPIntegrationApi(Client(ctx))

			_, httpresp, err := api.DeleteGCPIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGCPUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.GCPIntegrationApi(Client(ctx))

			_, httpresp, err := api.UpdateGCPIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func uninstallGCPIntegration(ctx context.Context, t *testing.T, account datadog.GCPAccount) {
	api := datadog.GCPIntegrationApi(Client(ctx))

	_, httpresp, err := api.DeleteGCPIntegration(ctx, account)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", account)
	}
}
