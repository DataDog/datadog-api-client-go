/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func generateUniqueAzureAccount(ctx context.Context) (datadog.AzureAccount, datadog.AzureAccount, datadog.AzureAccount) {
	tenantName := fmt.Sprintf("go_test-1234-5678-9101-%d", tests.ClockFromContext(ctx).Now().Unix())
	var testAzureAcct = datadog.AzureAccount{
		ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
		ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
		TenantName:   datadog.PtrString(tenantName),
	}

	var testUpdateAzureAcct = datadog.AzureAccount{
		ClientId:      testAzureAcct.ClientId,
		ClientSecret:  testAzureAcct.ClientSecret,
		TenantName:    testAzureAcct.TenantName,
		NewClientId:   datadog.PtrString("testc7f6-1234-5678-9101-3fcbf4update"),
		NewTenantName: datadog.PtrString("testc44-1234-5678-9101-cc0073update"),
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
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, _, _ := generateUniqueAzureAccount(ctx)
	defer uninstallAzureIntegration(ctx, testAzureAcct)

	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestAzureListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, _, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx)
	defer uninstallAzureIntegration(ctx, testAzureAcct)
	defer uninstallAzureIntegration(ctx, testUpdateAzureHostFilters)

	// Setup Azure Account to List
	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	azureListOutput, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx).Execute()
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
	_, httpresp, err = Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error deleting Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestUpdateAzureAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAzureAcct, testUpdateAzureAcct, testUpdateAzureHostFilters := generateUniqueAzureAccount(ctx)
	defer uninstallAzureIntegration(ctx, testAzureAcct)

	// Setup Azure Account to Update
	_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).AzureIntegrationApi.UpdateAzureIntegration(ctx).Body(testUpdateAzureAcct).Execute()
	defer uninstallAzureIntegration(ctx, testUpdateAzureAcct)
	if err != nil {
		t.Fatalf("Error updating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(200, httpresp.StatusCode)

	// List account to ensure update worked.
	azureListOutput, _, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx).Execute()
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
	_, httpresp, err = Client(ctx).AzureIntegrationApi.UpdateAzureHostFilters(ctx).Body(testUpdateAzureHostFilters).Execute()
	if err != nil {
		t.Fatalf("Error updating Azure Host Filters: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	HFListOutput, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx).Execute()
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
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/azure/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the azure integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/integration/azure").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(ctx).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAzure403Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AzureIntegrationApi.ListAzureIntegration(context.Background()).Execute()
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAzureCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).AzureIntegrationApi.CreateAzureIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).AzureIntegrationApi.UpdateAzureIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAzureUpdateHostFiltersErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).AzureIntegrationApi.UpdateAzureHostFilters(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func uninstallAzureIntegration(ctx context.Context, account datadog.AzureAccount) {
	_, httpresp, err := Client(ctx).AzureIntegrationApi.DeleteAzureIntegration(ctx).Body(account).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
