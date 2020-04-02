/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
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
		ListAWSAccounts(TESTAUTH).
		AccountId(testAWSAccount.GetAccountId()).
		RoleName(testAWSAccount.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	awsAccounts := awsAccts.GetAccounts()
	assert.True(t, len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	assert.Equal(t, testAWSAccount.GetAccountId(), awsAcct.GetAccountId())
	assert.Equal(t, testAWSAccount.GetRoleName(), awsAcct.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, testAWSAccount.GetHostTags(), awsAcct.GetHostTags())
	assert.Equal(t, testAWSAccount.GetFilterTags(), awsAcct.GetFilterTags())
	assert.Equal(t, testAWSAccount.GetExcludedRegions(), awsAcct.GetExcludedRegions())
	assert.Equal(t, testAWSAccount.GetAccountSpecificNamespaceRules(), awsAcct.GetAccountSpecificNamespaceRules())
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
	awsAccts, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.ListAWSAccounts(TESTAUTH).
		AccountId(UPDATEDAWSACCT.GetAccountId()).
		RoleName(UPDATEDAWSACCT.GetRoleName()).
		Execute()
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	awsAccounts := awsAccts.GetAccounts()
	assert.True(t, len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	// Test fields were updated
	assert.Equal(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetRoleName(), awsAcct.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetHostTags(), awsAcct.GetHostTags())
	assert.Equal(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetFilterTags(), awsAcct.GetFilterTags())
	assert.Equal(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetExcludedRegions(), awsAcct.GetExcludedRegions())
	assert.Equal(t, TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetAccountSpecificNamespaceRules(), awsAcct.GetAccountSpecificNamespaceRules())
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

	apiResp, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateNewAWSExternalID(TESTAUTH).Body(testAWSAccount).Execute()
	if err != nil {
		t.Fatalf("Error generating new AWS External ID: Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotEmpty(t, apiResp.GetExternalId())
}

func TestListNamespaces(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	namespaces, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.ListAvailableAWSNamespaces(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing AWS Namespaces: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	namespacesCheck := make(map[string]bool)
	for _, namespace := range namespaces {
		namespacesCheck[namespace] = true
	}

	// Check that a few examples are in the response
	// Full list - https://docs.datadoghq.com/api/?lang=bash#list-namespace-rules
	assert.True(t, namespacesCheck["api_gateway"])
	assert.True(t, namespacesCheck["cloudsearch"])
	assert.True(t, namespacesCheck["directconnect"])
	assert.True(t, namespacesCheck["xray"])
}

func TestAWSIntegrationGenerateExternalIDErrors(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GenerateNewAWSExternalID(TESTAUTH).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())

	// 403 Forbidden
	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.GenerateNewAWSExternalID(context.Background()).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok = err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationCreateErrors(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())

	// 403 Forbidden
	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(context.Background()).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok = err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationDeleteErrors(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(TESTAUTH).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())

	// 403 Forbidden
	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.DeleteAWSAccount(context.Background()).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok = err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationGetAll403Error(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 403 Forbidden
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationGetAll400Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/aws/get_all_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/integration/aws").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.GetAllAWSAccounts(TESTAUTH).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationListNamespacesErrors(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 403 Forbidden
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.ListAvailableAWSNamespaces(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationUpdateErrors(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(TESTAUTH).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())

	// 403 Forbidden
	_, httpresp, err = TESTAPICLIENT.AWSIntegrationApi.UpdateAWSAccount(context.Background()).Body(datadog.AWSAccount{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok = err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
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
