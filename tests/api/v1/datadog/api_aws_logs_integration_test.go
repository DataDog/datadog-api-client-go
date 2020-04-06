/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueAWSLambdaAccounts() (datadog.AWSAccount, datadog.AWSAccountAndLambdaRequest, datadog.AWSLogsServicesRequest) {
	accountID := fmt.Sprintf("go_%09d", TESTCLOCK.Now().UnixNano()%1000000000)
	var uniqueAWSAccount = datadog.AWSAccount{
		AccountId:                     &accountID,
		RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
		AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
		FilterTags:                    &[]string{"testTag", "test:Tag2"},
		HostTags:                      &[]string{"filter:one", "filtertwo"},
	}

	var testLambdaArn = datadog.AWSAccountAndLambdaRequest{
		AccountId: *uniqueAWSAccount.AccountId,
		LambdaArn: "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest",
	}

	var savedServices = datadog.AWSLogsServicesRequest{
		AccountId: *uniqueAWSAccount.AccountId,
		Services:  []string{"s3", "elb", "elbv2", "cloudfront", "redshift", "lambda"},
	}

	return uniqueAWSAccount, testLambdaArn, savedServices
}

// Test CreateAWSLambdaARN and EnableServices endpoints
func TestAddAndSaveAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testawsacc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testawsacc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testawsacc)

	_, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.CreateAWSLambdaARN(TESTAUTH).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error adding lamda ARN: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	_, httpresp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestListAWSLogsServices(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	listServicesOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.ListAWSLogsServices(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// There are currently 6 supported AWS Logs services as noted in the docs
	// https://docs.datadoghq.com/api/?lang=bash#get-list-of-aws-log-ready-services
	assert.Check(t, len(listServicesOutput) >= 6)
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAWSAcc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Create the AWS integration.
	retryCreateAccount(t, testAWSAcc)

	// Add Lambda to Account
	addOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.CreateAWSLambdaARN(TESTAUTH).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error Adding Lambda %v: Response %s: %v", addOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// Enable services for Lambda
	_, httpresp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// List AWS Logs integrations before deleting
	listOutput1, _, err := TESTAPICLIENT.AWSLogsIntegrationApi.ListAWSLogsIntegrations(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	// Iterate over output and list Lambdas
	var accountExists = false
	for _, Account := range listOutput1 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			if Account.GetLambdas()[0].GetArn() == testLambdaAcc.LambdaArn {
				accountExists = true
			}
		}
	}
	// Test that variable is true as expected
	assert.Equal(t, accountExists, true)

	// Delete newly added Lambda
	deleteOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.DeleteAWSLambdaARN(TESTAUTH).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error deleting Lambda %v: Response %s: %v", deleteOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// List AWS logs integrations after deleting
	listOutput2, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.ListAWSLogsIntegrations(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	var listOfARNs2 []datadog.AWSLogsListResponseLambdas
	var accountExistsAfterDelete = false
	for _, Account := range listOutput2 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			listOfARNs2 = Account.GetLambdas()
		}
	}
	for _, lambda := range listOfARNs2 {
		if lambda.GetArn() == testLambdaAcc.LambdaArn {
			accountExistsAfterDelete = true
		}
	}
	// Check that ARN no longer exists after delete
	assert.Assert(t, accountExistsAfterDelete != true)
}

func TestCheckLambdaAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAWSAcc, testLambdaAcc, _ := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAcc)

	status, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(TESTAUTH).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error checking the AWS Lambda Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, len(status.GetErrors()), 0)
	assert.Equal(t, status.GetStatus(), "created")

	// Give the async call time to finish
	tests.Retry(time.Duration(5*time.Second), 10, func() bool {
		status, httpresp, err = TESTAPICLIENT.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(TESTAUTH).Body(testLambdaAcc).Execute()
		if err != nil {
			t.Logf("Error checking the AWS Lambda Response: %s %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		return httpresp.StatusCode == 200 && len(status.GetErrors()) > 0
	})

	assert.Assert(t, status.GetErrors()[0].GetCode() != "")
	assert.Assert(t, status.GetErrors()[0].GetMessage() != "")
	assert.Equal(t, status.GetStatus(), "error")
}

func TestCheckServicesAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAWSAcc, _, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAcc)

	status, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(TESTAUTH).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error checking the AWS Logs Services Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Assert(t, status.GetErrors()[0].GetCode() != "")
	assert.Assert(t, status.GetErrors()[0].GetMessage() != "")
}
