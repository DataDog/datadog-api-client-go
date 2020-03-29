/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"gotest.tools/assert"
)

func generateUniqueAWSAccount() datadog.AWSAccount {
	accountID := fmt.Sprintf("go_%09d", TESTCLOCK.Now().UnixNano()%1000000000)
	return datadog.AWSAccount{
		AccountId:                     &accountID,
		RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
		AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
		FilterTags:                    &[]string{"testTag", "test:Tag2"},
		HostTags:                      &[]string{"filter:one", "filtertwo"},
		ExcludedRegions:               &[]string{"us-east-1", "us-west-1"},
	}
}

var TESTUPDATEAWSACC = datadog.AWSAccount{
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRoleUpdated"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": false},
	FilterTags:                    &[]string{"testTagUpdate", "testUpdated:Tag2"},
	HostTags:                      &[]string{"filter:foo", "bar"},
}

var TESTUPDATEAWSACCWITHEXCLUDEDREGION = datadog.AWSAccount{
	RoleName:                      datadog.PtrString("DatadogAWSIntegrationRoleUpdated"),
	AccountSpecificNamespaceRules: &map[string]bool{"opsworks": false},
	FilterTags:                    &[]string{"testTagUpdate", "testUpdated:Tag2"},
	HostTags:                      &[]string{"filter:foo", "bar"},
	ExcludedRegions:               &[]string{"us-east-1", "us-west-1"},

}

func TestCreateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAWSAccount := generateUniqueAWSAccount()

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAccount)
	defer retryDeleteAccount(t, testAWSAccount)

	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.
		GetAllAWSAccounts(TESTAUTH).
		AccountId(testAWSAccount.GetAccountId()).
		RoleName(testAWSAccount.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	awsAccounts := awsAccts.GetAccounts()
	assert.Assert(t, len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	assert.Equal(t, awsAcct.GetAccountId(), testAWSAccount.GetAccountId())
	assert.Equal(t, awsAcct.GetRoleName(), testAWSAccount.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), testAWSAccount.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), testAWSAccount.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetExcludedRegions(), testAWSAccount.GetExcludedRegions()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), testAWSAccount.GetAccountSpecificNamespaceRules()), true)
}

func TestUpdateAWSAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAWSAccount := generateUniqueAWSAccount()

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAccount)
	defer retryDeleteAccount(t, testAWSAccount)

	retryUpdateAccount(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION, testAWSAccount.GetAccountId(), testAWSAccount.GetRoleName())
	UPDATEDAWSACCT := datadog.AWSAccount{
		AccountId: testAWSAccount.AccountId,
		RoleName:  TESTUPDATEAWSACCWITHEXCLUDEDREGION.RoleName,
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
	assert.Equal(t, awsAcct.GetRoleName(), TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetHostTags(), TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetHostTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetFilterTags(), TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetFilterTags()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetExcludedRegions(), TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetExcludedRegions()), true)
	assert.Equal(t, reflect.DeepEqual(awsAcct.GetAccountSpecificNamespaceRules(), TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetAccountSpecificNamespaceRules()), true)
}

func TestDisableAWSAcct(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test
	testAWSAccount := generateUniqueAWSAccount()

	// Lets first create the account of us to delete
	retryCreateAccount(t, testAWSAccount)

	retryDeleteAccount(t, testAWSAccount)
}

func TestGenerateNewExternalId(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAWSAccount := generateUniqueAWSAccount()
	// Lets first create the account for us to generate a new id against
	retryCreateAccount(t, testAWSAccount)
	defer retryDeleteAccount(t, testAWSAccount)

	apiResp, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GenerateNewAWSExternalID(TESTAUTH).Body(testAWSAccount).Execute()
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

func retryDeleteAccount(t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
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

func retryCreateAccount(t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
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

func retryUpdateAccount(t *testing.T, body datadog.AWSAccount, accountID string, roleName string) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
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
