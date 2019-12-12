package datadog_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var TESTAWSACC = datadog.AwsAccount{
	AccountId:                     datadog.PtrString("123456789012"),
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
	FilterTags:                    &[]string{"testTag", "test:Tag2"},
	HostTags:                      &[]string{"filter:one", "filtertwo"},
}

var TESTUPDATEAWSACC = datadog.AwsAccount{
	AccountId:                     datadog.PtrString("123456789012"),
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": false},
	FilterTags:                    &[]string{"testTagUpdate", "testUpdated:Tag2"},
	HostTags:                      &[]string{"filter:foo", "bar"},
}

func TestCreateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAWSIntegration(TESTAWSACC)

	// Assert AWS Integration Created with proper fields
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(TESTAWSACC).Execute()

	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.
		GetAllAWSAccounts(TESTAUTH).
		AccountId(TESTAWSACC.GetAccountId()).
		RoleName(TESTAWSACC.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200, "Error disabling AWS Account: %v", httpresp)
	awsAcct := awsAccts.GetAccounts()[0]
	assert.Equal(t, awsAcct.GetAccountId(), TESTAWSACC.GetAccountId())
	assert.Equal(t, awsAcct.GetRoleName(), TESTAWSACC.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), TESTAWSACC.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), TESTAWSACC.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), TESTAWSACC.GetAccountSpecificNamespaceRules()), true)
}

func TestUpdateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAWSIntegration(TESTAWSACC)

	// Assert AWS Integration Created with proper fields
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(TESTAWSACC).Execute()

	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(TESTAUTH).
		AwsAccount(TESTUPDATEAWSACC).
		AccountId(TESTAWSACC.GetAccountId()).
		RoleName(TESTAWSACC.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error updating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200, "Error disabling AWS Account: %v", httpresp)

	// Assert AWS Account Get with proper fields
	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH).
		AccountId(TESTAWSACC.GetAccountId()).
		RoleName(TESTAWSACC.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200, "Error disabling AWS Account: %v", httpresp)

	awsAcct := awsAccts.GetAccounts()[0]
	// Test fields were updated
	assert.Equal(t, awsAcct.GetAccountId(), TESTUPDATEAWSACC.GetAccountId())
	assert.Equal(t, awsAcct.GetRoleName(), TESTUPDATEAWSACC.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), TESTUPDATEAWSACC.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), TESTUPDATEAWSACC.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), TESTUPDATEAWSACC.GetAccountSpecificNamespaceRules()), true)
}

func TestDisableAWSAcct(t *testing.T) {
	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test

	// Lets first create the account of us to delete
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(TESTAWSACC).Execute()

	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH).AwsAccount(TESTAWSACC).Execute()
	if err != nil {
		t.Fatalf("Error disabling AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200, "Error disabling AWS Account: %v", httpresp)
}

func uninstallAWSIntegration(account datadog.AwsAccount) {
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH).AwsAccount(account).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling AWS Account: %v, Another test may have already removed this account.", account)
	}
}
