// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueAzureAccount() (datadog.AzureAccount, datadog.AzureAccount, datadog.AzureAccount) {
	tenantName := fmt.Sprintf("go_test-1234-5678-9101-%s", time.Now())
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

	var testAzureUpdateHostFilters = datadog.AzureAccount{
		ClientId:    testUpdateAzureAcct.NewClientId,
		TenantName:  testUpdateAzureAcct.NewTenantName,
		HostFilters: datadog.PtrString("test:foo,test:bar"),
	}

	return testAzureAcct, testUpdateAzureAcct, testAzureUpdateHostFilters
}

func TestAzureCreate(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAzureAcct, _, _ := generateUniqueAzureAccount()
	defer uninstallAzureIntegration(testAzureAcct)

	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestAzureListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAzureAcct, _, testAzureUpdateHostFilters := generateUniqueAzureAccount()
	defer uninstallAzureIntegration(testAzureAcct)
	defer uninstallAzureIntegration(testAzureUpdateHostFilters)

	// Setup Azure Account to List
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	azureListOutput, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	var x datadog.AzureAccount
	for _, Account := range azureListOutput {
		if Account.GetClientId() == *testAzureAcct.ClientId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), *testAzureAcct.ClientId)
	assert.Equal(t, x.GetTenantName(), *testAzureAcct.TenantName)

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(azureListOutput) >= 1)

	// Test account deletion as well
	_, httpresp, err = TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error deleting Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestUpdateAzureAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAzureAcct, testUpdateAzureAcct, testAzureUpdateHostFilters := generateUniqueAzureAccount()
	defer uninstallAzureIntegration(testAzureAcct)

	// Setup Azure Account to Update
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH).Body(testAzureAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	_, httpresp, err = TESTAPICLIENT.AzureIntegrationApi.UpdateAzureIntegration(TESTAUTH).Body(testUpdateAzureAcct).Execute()
	defer uninstallAzureIntegration(testUpdateAzureAcct)
	if err != nil {
		t.Fatalf("Error updating Azure Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(t, httpresp.StatusCode, 200)

	// List account to ensure update worked.
	azureListOutput, _, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	var x datadog.AzureAccount
	for _, Account := range azureListOutput {
		if Account.GetClientId() == *testUpdateAzureAcct.NewClientId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), *testUpdateAzureAcct.NewClientId)
	assert.Equal(t, x.GetTenantName(), *testUpdateAzureAcct.NewTenantName)
	assert.Equal(t, x.GetHostFilters(), *testUpdateAzureAcct.HostFilters)

	// Test update host filters endpoint
	_, httpresp, err = TESTAPICLIENT.AzureIntegrationApi.AzureUpdateHostFilters(TESTAUTH).Body(testAzureUpdateHostFilters).Execute()
	if err != nil {
		t.Fatalf("Error updating Azure Host Filters: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	HFListOutput, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing Azure Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	var y datadog.AzureAccount
	for _, Account := range HFListOutput {
		if Account.GetClientId() == *testAzureUpdateHostFilters.ClientId {
			y = Account
		}
	}
	assert.Equal(t, y.GetHostFilters(), *testAzureUpdateHostFilters.HostFilters)
}

func uninstallAzureIntegration(account datadog.AzureAccount) {
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH).Body(account).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
