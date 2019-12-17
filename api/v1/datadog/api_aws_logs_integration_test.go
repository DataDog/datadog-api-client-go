package datadog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueAwsLambdaAccounts() (datadog.AwsAccount, datadog.AwsAccountAndLambdaInput, datadog.AwsLogsServicesInput) {
	accountID := fmt.Sprintf("go_%09d", time.Now().UnixNano()%1000000000)
	var uniqueAwsAccount = datadog.AwsAccount{
		AccountId:                     &accountID,
		RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
		AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
		FilterTags:                    &[]string{"testTag", "test:Tag2"},
		HostTags:                      &[]string{"filter:one", "filtertwo"},
	}

	var testLambdaArn = datadog.AwsAccountAndLambdaInput{
		AccountId: *uniqueAwsAccount.AccountId,
		LambdaArn: "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest",
	}

	var savedServices = datadog.AwsLogsServicesInput{
		AccountId: *uniqueAwsAccount.AccountId,
		Services:  []string{"s3", "elb", "elbv2", "cloudfront", "redshift", "lambda"},
	}

	return uniqueAwsAccount, testLambdaArn, savedServices
}

// Test AddAWSLambdaARN and EnableServices endpoints
func TestAddAndSaveAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testawsacc, testLambdaAcc, testServices := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testawsacc)

	// Assert AWS Integration Created with proper fields
	_, httpResp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(testawsacc).Execute()
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	_, httpResp, err = TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH).AwsAccountAndLambdaInput(testLambdaAcc).Execute()
	if err != nil {
		t.Errorf("Error adding lamda ARN: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	_, httpResp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).AwsLogsServicesInput(testServices).Execute()
	if err != nil {
		t.Errorf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)
}

func TestListAWSLogsServices(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	listServicesOutput, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsServicesList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// There are currently 6 supported AWS Logs services as noted in the docs
	// https://docs.datadoghq.com/api/?lang=bash#get-list-of-aws-log-ready-services
	assert.Check(t, len(listServicesOutput) >= 6)
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAWSAcc, testLambdaAcc, testServices := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testAWSAcc)

	// Create the AWS integration.
	_, httpResp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(testAWSAcc).Execute()
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// Add Lambda to Account
	addOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH).AwsAccountAndLambdaInput(testLambdaAcc).Execute()
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error Adding Lambda %v: Response %s: %v", addOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// Enable services for Lambda
	_, httpResp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).AwsLogsServicesInput(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// List AWS Logs integrations before deleting
	listOutput1, _, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)
	// Iterate over output and list Lambdas
	var x = false
	for _, Account := range listOutput1 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			if Account.GetLambdas()[0].GetArn() == testLambdaAcc.LambdaArn {
				x = true
			}
		}
	}
	// Test that variable is true as expected
	assert.Equal(t, x, true)

	// Delete newly added Lambda
	deleteOutput, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.DeleteAWSLambdaARN(TESTAUTH).AwsAccountAndLambdaInput(testLambdaAcc).Execute()
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error deleting Lambda %v: Response %s: %v", deleteOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// List AWS logs integrations after deleting
	listOutput2, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	var listOfARNs2 []datadog.AwsLogsListOutputLambdas
	x = false
	for _, Account := range listOutput2 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			listOfARNs2 = Account.GetLambdas()
		}
	}
	for _, lambda := range listOfARNs2 {
		if lambda.GetArn() == testLambdaAcc.LambdaArn {
			x = true
		}
	}
	// Check that ARN no longer exists after delete
	assert.Assert(t, x != true)
}
