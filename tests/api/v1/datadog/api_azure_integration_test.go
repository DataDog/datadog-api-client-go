/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"strconv"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func generateUniqueAzureAccount(ctx context.Context, t *testing.T) (datadog.AzureAccount, datadog.AzureAccount, datadog.AzureAccount) {
	tenantName := tests.UniqueEntityName(ctx, t)
	updatedTenantName := *tenantName + "-updated"
	clock := strconv.FormatInt(tests.ClockFromContext(ctx).Now().Unix(), 10)
	var testAzureAcct = datadog.AzureAccount{
		ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-tt" + clock),
		ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
		TenantName:   tenantName,
	}

	var testUpdateAzureAcct = datadog.AzureAccount{
		ClientId:      testAzureAcct.ClientId,
		ClientSecret:  testAzureAcct.ClientSecret,
		TenantName:    testAzureAcct.TenantName,
		NewClientId:   datadog.PtrString("testc7f6-1234-5678-9101-uu" + clock),
		NewTenantName: &updatedTenantName,
		HostFilters:   datadog.PtrString("filter:foo,test:bar"),
	}

	var testUpdateAzureHostFilters = datadog.AzureAccount{
		ClientId:    testUpdateAzureAcct.NewClientId,
		TenantName:  testUpdateAzureAcct.NewTenantName,
		HostFilters: datadog.PtrString("test:foo,test:bar"),
	}

	return testAzureAcct, testUpdateAzureAcct, testUpdateAzureHostFilters
}

func TestAzureCreate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, _, _ := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)

	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestAzureListandDelete(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, _, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)
	defer uninstallAzureIntegration(ctx, t, testUpdateAzureHostFilters)

	// Setup Azure Account to List
	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	azureListOutput, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadog.AzureAccount
	for _, Account := range azureListOutput {
		if Account.GetClientId() == *testAzureAcct.ClientId {
			x = Account
		}
	}
	assert.Equal(*testAzureAcct.ClientId, x.GetClientId())
	assert.Equal(*testAzureAcct.TenantName, x.GetTenantName())

	// Assert returned list is greater than or equal to 1
	assert.True(len(azureListOutput) >= 1)

	// Test account deletion as well
	_, httpresp, err = Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error deleting Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestUpdateAzureAccount(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, testUpdateAzureAcct, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)

	// Setup Azure Account to Update
	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).AzureIntegrationApi.UpdateAzureIntegration(ctx, testUpdateAzureAcct)
	defer uninstallAzureIntegration(ctx, t, testUpdateAzureAcct)
	if err != nil {
		t.Fatalf("Error updating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(200, httpresp.StatusCode)

	// List account to ensure update worked.
	azureListOutput, _, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadog.AzureAccount
	for _, Account := range azureListOutput {
		if Account.GetClientId() == *testUpdateAzureAcct.NewClientId {
			x = Account
		}
	}
	assert.Equal(*testUpdateAzureAcct.NewClientId, x.GetClientId())
	assert.Equal(*testUpdateAzureAcct.NewTenantName, x.GetTenantName())
	assert.Equal(*testUpdateAzureAcct.HostFilters, x.GetHostFilters())

	// Test update host filters endpoint
	_, httpresp, err = Client(ctx).AzureIntegrationApi.UpdateAzureHostFilters(ctx, testUpdateAzureHostFilters)
	if err != nil {
		t.Fatalf("Error updating Azure Host Filters: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	HFListOutput, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var y datadog.AzureAccount
	for _, Account := range HFListOutput {
		if Account.GetClientId() == *testUpdateAzureHostFilters.ClientId {
			y = Account
		}
	}
	assert.Equal(*testUpdateAzureHostFilters.HostFilters, y.GetHostFilters())
}

func TestAzureList400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/azure/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the azure integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AzureIntegrationApiService.ListAzureIntegration")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/integration/azure").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAzure403Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAzureCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AzureIntegrationApi.UpdateAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureUpdateHostFiltersErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AzureIntegrationApi.UpdateAzureHostFilters(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func uninstallAzureIntegration(ctx context.Context, t *testing.T, account datadog.AzureAccount) {
	toDelete := datadog.AzureAccount{ClientId: account.ClientId, TenantName: account.TenantName}
	if account.NewClientId != nil {
		// when we call this on an update request, make sure we actually delete the updated entity, not the original one
		toDelete.ClientId = account.NewClientId
		toDelete.TenantName = account.NewTenantName
	}
	_, httpresp, err := Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx, toDelete)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
