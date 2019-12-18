package datadog_test

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

	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, testAzureAcct)
	if err != nil {
		t.Errorf("Error Creating Azure intg: %v", err)
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
	TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, testAzureAcct)

	azure_list_output, _, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	if err != nil {
		t.Errorf("Error listing azure intgs: %v", err)
	}
	var x datadog.AzureAccount
	for _, Account := range azure_list_output {
		if Account.GetClientId() == *testAzureAcct.ClientId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), *testAzureAcct.ClientId)
	assert.Equal(t, x.GetTenantName(), *testAzureAcct.TenantName)

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(azure_list_output) >= 1)

	// Test account deletion as well
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH, testAzureAcct)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", testAzureAcct)
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
	TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, testAzureAcct)

	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.UpdateAzureIntegration(TESTAUTH, testUpdateAzureAcct)
	defer uninstallAzureIntegration(testUpdateAzureAcct)
	if err != nil {
		t.Errorf("Error Updating Azure Account: %v", err)
	}

	assert.Equal(t, httpresp.StatusCode, 200)

	// List account to ensure update worked.
	azure_list_output, _, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	if err != nil {
		t.Errorf("Error listing Azure intgs: %v", err)
	}
	var x datadog.AzureAccount
	for _, Account := range azure_list_output {
		if Account.GetClientId() == *testUpdateAzureAcct.NewClientId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), *testUpdateAzureAcct.NewClientId)
	assert.Equal(t, x.GetTenantName(), *testUpdateAzureAcct.NewTenantName)
	assert.Equal(t, x.GetHostFilters(), *testUpdateAzureAcct.HostFilters)

	// Test update host filters endpoint
	_, httpresp, err = TESTAPICLIENT.AzureIntegrationApi.AzureUpdateHostFilters(TESTAUTH, testAzureUpdateHostFilters)
	if err != nil {
		t.Errorf("Error Updating Azure Host Filters: %v", err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	hf_list_output, _, err := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	if err != nil {
		t.Errorf("Error listing Azure intgs: %v", err)
	}
	var y datadog.AzureAccount
	for _, Account := range hf_list_output {
		if Account.GetClientId() == *testAzureUpdateHostFilters.ClientId {
			y = Account
		}
	}
	assert.Equal(t, y.GetHostFilters(), *testAzureUpdateHostFilters.HostFilters)
}

func uninstallAzureIntegration(account datadog.AzureAccount) {
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
