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

var randSuffix = fmt.Sprintf("%d", time.Now().Unix())[0:7]

var UNIQUEAWSACCOUNTID = fmt.Sprintf("dd-go%s", randSuffix)

var TESTAWSACC = datadog.AwsAccount{
	AccountId:                     datadog.PtrString(UNIQUEAWSACCOUNTID),
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
	FilterTags:                    &[]string{"testTag", "test:Tag2"},
	HostTags:                      &[]string{"filter:one", "filtertwo"},
}

var TESTUPDATEAWSACC = datadog.AwsAccount{
	AccountId:                     datadog.PtrString(UNIQUEAWSACCOUNTID),
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
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, TESTAWSACC)

	awsAcctOpts := datadog.GetAllAWSAccountsOpts{
		AccountId: optional.NewString(TESTAWSACC.GetAccountId()),
		RoleName:  optional.NewString(TESTAWSACC.GetRoleName()),
	}
	awsAccts, _, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH, &awsAcctOpts)
	if err != nil {
		t.Errorf("Error Getting AWS Account: %v", err)
	}
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
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, TESTAWSACC)

	awsAcctOpts := datadog.UpdateAWSAccountOpts{
		AccountId: optional.NewString(TESTAWSACC.GetAccountId()),
		RoleName:  optional.NewString(TESTAWSACC.GetRoleName()),
	}
	_, _, err := TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(TESTAUTH, TESTUPDATEAWSACC, &awsAcctOpts)
	if err != nil {
		t.Errorf("Error Updating AWS Account: %v", err)
	}

	// Assert AWS Account Get with proper fields
	awsGetAcctOpts := datadog.GetAllAWSAccountsOpts{
		AccountId: optional.NewString(TESTAWSACC.GetAccountId()),
		RoleName:  optional.NewString(TESTAWSACC.GetRoleName()),
	}
	awsAccts, _, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH, &awsGetAcctOpts)
	if err != nil {
		t.Errorf("Error Getting AWS Account: %v", err)
	}
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
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, TESTAWSACC)

	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH, TESTAWSACC)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error disabling AWS Account: %v: %v", httpresp, err)
	}
}

func uninstallAWSIntegration(account datadog.AwsAccount) {
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling AWS Account: %v, Another test may have already removed this account.", account)
	}
}
