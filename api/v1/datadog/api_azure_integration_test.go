package datadog_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var UNIQUETENANTNAME = fmt.Sprintf("testc44-1234-5678-9101-%s", time.Now())

var TESTAZUREACCT = datadog.AzureAccount{
	ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
	ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
	TenantName:   datadog.PtrString(UNIQUETENANTNAME),
}

var TESTUPDATEAZUREACC = datadog.AzureAccount{
	ClientId:      datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
	ClientSecret:  datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
	TenantName:    datadog.PtrString(UNIQUETENANTNAME),
	NewClientId:   datadog.PtrString("testc7f6-1234-5678-9101-3fcbf4update"),
	NewTenantName: datadog.PtrString("testc44-1234-5678-9101-cc0073update"),
	HostFilters:   datadog.PtrString("filter:foo,test:bar"),
}

var TESTUPDATEAZUREHOSTFILTERS = datadog.AzureAccount{
	ClientId:    datadog.PtrString("testc7f6-1234-5678-9101-3fcbf4update"),
	TenantName:  datadog.PtrString("testc44-1234-5678-9101-cc0073update"),
	HostFilters: datadog.PtrString("test:foo,test:bar"),
}

func TestAzureCreate(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAzureIntegration(TESTAZUREACCT)

	createOutput, _, _ := TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, TESTAZUREACCT)
	assert.Equal(t, len(createOutput), 0)
}

func TestAzureListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAzureIntegration(TESTAZUREACCT)

	// Setup Azure Account to List
	TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, TESTAZUREACCT)

	azure_list_output, _, _ := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	var x datadog.AzureAccount
	for _, Account := range azure_list_output {
		if Account.GetClientId() == "testc7f6-1234-5678-9101-3fcbf4update" {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), "testc7f6-1234-5678-9101-3fcbf4update")
	assert.Equal(t, x.GetTenantName(), "testc44-1234-5678-9101-cc0073update")

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(azure_list_output) >= 1)

	// Test account deletion as well
	delete_output, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH, TESTAZUREACCT)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", TESTAZUREACCT)
	}
	assert.Equal(t, len(delete_output), 0)
}

func TestUpdateAzureAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAzureIntegration(TESTAZUREACCT)

	// Setup Azure Account to Update
	TESTAPICLIENT.AzureIntegrationApi.CreateAzureIntegration(TESTAUTH, TESTAZUREACCT)

	azureUpdateOutput, _, err := TESTAPICLIENT.AzureIntegrationApi.UpdateAzureIntegration(TESTAUTH, TESTUPDATEAZUREACC)
	if err != nil {
		t.Errorf("Error Updating Azure Account: %v", err)
	}

	assert.Equal(t, len(azureUpdateOutput), 0)

	// List account to ensure update worked.
	azure_list_output, _, _ := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	var x datadog.AzureAccount
	for _, Account := range azure_list_output {
		if Account.GetClientId() == "testc7f6-1234-5678-9101-3fcbf4update" {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientId(), "testc7f6-1234-5678-9101-3fcbf4update")
	assert.Equal(t, x.GetTenantName(), "testc44-1234-5678-9101-cc0073update")
	assert.Equal(t, x.GetHostFilters(), "filter:foo,test:bar")

	// Test update host filters endpoint
	azureUpdateHFOutput, _, err := TESTAPICLIENT.AzureIntegrationApi.AzureUpdateHostFilters(TESTAUTH, TESTUPDATEAZUREHOSTFILTERS)
	if err != nil {
		t.Errorf("Error Updating Azure Host Filters: %v", err)
	}
	assert.Equal(t, len(azureUpdateHFOutput), 0)
	hf_list_output, _, _ := TESTAPICLIENT.AzureIntegrationApi.ListAzureIntegration(TESTAUTH)
	var y datadog.AzureAccount
	for _, Account := range hf_list_output {
		if Account.GetClientId() == "testc7f6-1234-5678-9101-3fcbf4update" {
			y = Account
		}
	}
	assert.Equal(t, y.GetHostFilters(), "test:foo,test:bar")
}

func uninstallAzureIntegration(account datadog.AzureAccount) {
	_, httpresp, err := TESTAPICLIENT.AzureIntegrationApi.DeleteAzureIntegration(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling Azure Account: %v, Another test may have already removed this account.", account)
	}
}
