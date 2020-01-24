package datadog_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueAwsAccount() datadog.AwsAccount {
	accountID := fmt.Sprintf("go_%09d", time.Now().UnixNano()%1000000000)
	return datadog.AwsAccount{
		AccountId:                     &accountID,
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
	retryCreateAccount(t, testAwsAccount)
	defer retryDeleteAccount(t, testAwsAccount)

	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.
		GetAllAWSAccounts(TESTAUTH).
		AccountId(testAwsAccount.GetAccountId()).
		RoleName(testAwsAccount.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
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
	retryCreateAccount(t, testAwsAccount)
	defer retryDeleteAccount(t, testAwsAccount)

	retryUpdateAccount(t, TESTUPDATEAWSACC, testAwsAccount.GetAccountId(), testAwsAccount.GetRoleName())
	UPDATEDAWSACCT := datadog.AwsAccount{
		AccountId: testAwsAccount.AccountId,
		RoleName:  TESTUPDATEAWSACC.RoleName,
	}
	defer retryDeleteAccount(t, UPDATEDAWSACCT)

	// Assert AWS Account Get with proper fields
	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH).
		AccountId(UPDATEDAWSACCT.GetAccountId()).
		RoleName(UPDATEDAWSACCT.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
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
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test
	testAwsAccount := generateUniqueAwsAccount()

	// Lets first create the account of us to delete
	retryCreateAccount(t, testAwsAccount)

	retryDeleteAccount(t, testAwsAccount)
}

func TestGenerateNewExternalId(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAwsAccount := generateUniqueAwsAccount()
	// Lets first create the account for us to generate a new id against
	retryCreateAccount(t, testAwsAccount)
	defer retryDeleteAccount(t, testAwsAccount)

	apiResp, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GenerateNewAWSExternalID(TESTAUTH).Body(testAwsAccount).Execute()
	if err != nil {
		t.Fatalf("Error generating new AWS External ID: Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, apiResp.GetExternalId() != "")
}

func TestListNamespaces(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	namespaces, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.ListAvailableAWSNamespaces(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing AWS Namespaces: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	namespacesCheck := make(map[string]bool)
	for _, namespace := range namespaces {
		namespacesCheck[namespace] = true
	}

	// Check that a few examples are in the response
	// Full list - https://docs.datadoghq.com/api/?lang=bash#list-namespace-rules
	assert.Assert(t, namespacesCheck["api_gateway"], true)
	assert.Assert(t, namespacesCheck["cloudsearch"], true)
	assert.Assert(t, namespacesCheck["directconnect"], true)
	assert.Assert(t, namespacesCheck["xray"], true)
}

func retryDeleteAccount(t *testing.T, awsAccount datadog.AwsAccount) {
	err := retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH).Body(awsAccount).Execute()
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error deleting AWS Account: Response %s", err)
	}
}

func retryCreateAccount(t *testing.T, awsAccount datadog.AwsAccount) {
	err := retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).Body(awsAccount).Execute()
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s", err)
	}
}

func retryUpdateAccount(t *testing.T, body datadog.AwsAccount, accountID string, roleName string) {
	err := retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(TESTAUTH).
			Body(body).
			AccountId(accountID).
			RoleName(roleName).
			Execute()
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error updating AWS Account: Response %s", err)
	}
}
