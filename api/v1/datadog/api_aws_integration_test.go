package datadog_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/antihax/optional"
	"gotest.tools/assert"
)

func generateUniqueAwsAccount() datadog.AwsAccount {
	accountId := fmt.Sprintf("go_%09d", time.Now().UnixNano()%1000000000)
	return datadog.AwsAccount{
		AccountId:                     &accountId,
		RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
		AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
		FilterTags:                    &[]string{"testTag", "test:Tag2"},
		HostTags:                      &[]string{"filter:one", "filtertwo"},
	}
}

var TESTUPDATEAWSACC = datadog.AwsAccount{
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRoleUpdated"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": false},
	FilterTags:                    &[]string{"testTagUpdate", "testUpdated:Tag2"},
	HostTags:                      &[]string{"filter:foo", "bar"},
}

func TestCreateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAwsAccount := generateUniqueAwsAccount()

	// Assert AWS Integration Created with proper fields
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, testAwsAccount)
	if err != nil {
		t.Errorf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	defer uninstallAWSIntegration(testAwsAccount)

	awsAcctOpts := datadog.GetAllAWSAccountsOpts{
		AccountId: optional.NewString(testAwsAccount.GetAccountId()),
		RoleName:  optional.NewString(testAwsAccount.GetRoleName()),
	}
	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH, &awsAcctOpts)
	if err != nil {
		t.Errorf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	awsAccounts := awsAccts.GetAccounts()
	assert.Assert(t, len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	assert.Equal(t, awsAcct.GetAccountId(), testAwsAccount.GetAccountId())
	assert.Equal(t, awsAcct.GetRoleName(), testAwsAccount.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), testAwsAccount.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), testAwsAccount.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), testAwsAccount.GetAccountSpecificNamespaceRules()), true)
}

func TestUpdateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAwsAccount := generateUniqueAwsAccount()

	// Assert AWS Integration Created with proper fields
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, testAwsAccount)
	if err != nil {
		t.Errorf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	defer uninstallAWSIntegration(testAwsAccount)

	awsAcctOpts := datadog.UpdateAWSAccountOpts{
		AccountId: optional.NewString(testAwsAccount.GetAccountId()),
		RoleName:  optional.NewString(testAwsAccount.GetRoleName()),
	}
	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(TESTAUTH, TESTUPDATEAWSACC, &awsAcctOpts)
	if err != nil {
		t.Errorf("Error Updating AWS Account: %v", err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	defer uninstallAWSIntegration(TESTUPDATEAWSACC)

	// Assert AWS Account Get with proper fields
	awsGetAcctOpts := datadog.GetAllAWSAccountsOpts{
		AccountId: optional.NewString(testAwsAccount.GetAccountId()),
		RoleName:  optional.NewString(TESTUPDATEAWSACC.GetRoleName()),
	}
	awsAccts, _, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH, &awsGetAcctOpts)
	if err != nil {
		t.Errorf("Error Getting AWS Account: %v", err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	awsAccounts := awsAccts.GetAccounts()
	assert.Assert(t, len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	// Test fields were updated
	assert.Equal(t, awsAcct.GetRoleName(), TESTUPDATEAWSACC.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), TESTUPDATEAWSACC.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), TESTUPDATEAWSACC.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), TESTUPDATEAWSACC.GetAccountSpecificNamespaceRules()), true)
}

func TestDisableAWSAcct(t *testing.T) {
	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test
	testAwsAccount := generateUniqueAwsAccount()

	// Lets first create the account of us to delete
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, testAwsAccount)
	if err != nil {
		t.Errorf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	defer uninstallAWSIntegration(testAwsAccount)

	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH, testAwsAccount)
	if err != nil {
		t.Errorf("Error disabling AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200, "Error disabling AWS Account: %v", httpresp)
}

func uninstallAWSIntegration(account datadog.AwsAccount) {
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling AWS Account: %v, Another test may have already removed this account.", account)
	}
}
