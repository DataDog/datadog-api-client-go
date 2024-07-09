/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func generateUniqueAzureAccount(ctx context.Context, t *testing.T) (datadogV1.AzureAccount, datadogV1.AzureAccount, datadogV1.AzureAccount) {
	// Need something that looks like a UUID
	tenantName := fmt.Sprintf("aaaaaaaa-bbbb-cccc-dddd-%dee", tests.ClockFromContext(ctx).Now().Unix())
	clock := strconv.FormatInt(tests.ClockFromContext(ctx).Now().Unix(), 10)
	var testAzureAcct = datadogV1.AzureAccount{
		ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-tt" + clock),
		ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
		TenantName:   &tenantName,
	}

	var testUpdateAzureAcct = datadogV1.AzureAccount{
		ClientId:     testAzureAcct.ClientId,
		ClientSecret: testAzureAcct.ClientSecret,
		TenantName:   testAzureAcct.TenantName,
		HostFilters:  datadog.PtrString("filter:foo,test:bar"),
	}

	var testUpdateAzureHostFilters = datadogV1.AzureAccount{
		ClientId:    testUpdateAzureAcct.ClientId,
		TenantName:  testUpdateAzureAcct.TenantName,
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
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

	testAzureAcct, _, _ := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)

	_, httpresp, err := api.CreateAzureIntegration(ctx, testAzureAcct)
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
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

	testAzureAcct, _, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)
	defer uninstallAzureIntegration(ctx, t, testUpdateAzureHostFilters)

	// Setup Azure Account to List
	_, httpresp, err := api.CreateAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	azureListOutput, httpresp, err := api.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadogV1.AzureAccount
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
	_, httpresp, err = api.DeleteAzureIntegration(ctx, testAzureAcct)
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
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

	testAzureAcct, testUpdateAzureAcct, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx, t)
	defer uninstallAzureIntegration(ctx, t, testAzureAcct)

	// Setup Azure Account to Update
	_, httpresp, err := api.CreateAzureIntegration(ctx, testAzureAcct)
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = api.UpdateAzureIntegration(ctx, testUpdateAzureAcct)
	defer uninstallAzureIntegration(ctx, t, testUpdateAzureAcct)
	if err != nil {
		t.Fatalf("Error updating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(200, httpresp.StatusCode)

	// List account to ensure update worked.
	azureListOutput, _, err := api.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var x datadogV1.AzureAccount
	for i, Account := range azureListOutput {
		fmt.Println(i, Account)
		if Account.GetClientId() == testUpdateAzureAcct.GetClientId() {
			x = Account
		}
	}

	assert.Equal(*testUpdateAzureAcct.HostFilters, x.GetHostFilters())

	// Test update host filters endpoint
	_, httpresp, err = api.UpdateAzureHostFilters(ctx, testUpdateAzureHostFilters)
	if err != nil {
		t.Fatalf("Error updating Azure Host Filters: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	HFListOutput, httpresp, err := api.ListAzureIntegration(ctx)
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	var y datadogV1.AzureAccount
	for _, Account := range HFListOutput {
		if Account.GetClientId() == testUpdateAzureHostFilters.GetClientId() {
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
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

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
	_, httpresp, err := api.ListAzureIntegration(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
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
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

	// 403 Forbidden
	_, httpresp, err := api.ListAzureIntegration(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAzureCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewAzureIntegrationApi(Client(ctx))

			_, httpresp, err := api.CreateAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
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
		Body               datadogV1.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewAzureIntegrationApi(Client(ctx))

			_, httpresp, err := api.DeleteAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
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
		Body               datadogV1.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewAzureIntegrationApi(Client(ctx))

			_, httpresp, err := api.UpdateAzureIntegration(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
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
		Body               datadogV1.AzureAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.AzureAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.AzureAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewAzureIntegrationApi(Client(ctx))

			_, httpresp, err := api.UpdateAzureHostFilters(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func uninstallAzureIntegration(ctx context.Context, t *testing.T, account datadogV1.AzureAccount) {
	api := datadogV1.NewAzureIntegrationApi(Client(ctx))

	toDelete := datadogV1.AzureAccount{ClientId: account.ClientId, TenantName: account.TenantName}
	if account.NewClientId != nil {
		// when we call this on an update request, make sure we actually delete the updated entity, not the original one
		toDelete.ClientId = account.NewClientId
		toDelete.TenantName = account.NewTenantName
	}
	_, httpresp, err := api.DeleteAzureIntegration(ctx, toDelete)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
