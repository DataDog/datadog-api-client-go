package datadog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var UNIQUEACCOUNTID = fmt.Sprintf("12345%s", time.Now())

var TESTAWSACCLOGS = datadog.AwsAccount{
	AccountId: datadog.PtrString(UNIQUEACCOUNTID),
	RoleName:  datadog.PtrString("DatadogAWSIntegrationRole"),
}

var TESTADDLAMBDAARN = datadog.AwsAccountAndLambdaInput{
	AccountId: UNIQUEACCOUNTID,
	LambdaArn: "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest",
}

var TESTSAVESERVICES = datadog.AwsLogsServicesInput{
	AccountId: UNIQUEACCOUNTID,
	Services:  []string{"s3", "elb", "elbv2", "cloudfront", "redshift", "lambda"},
}

// Test AddAWSLambdaARN and EnableServices endpoints
func TestAddAndSaveAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAWSIntegration(TESTAWSACCLOGS)

	// Assert AWS Integration Created with proper fields
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, TESTAWSACCLOGS)
	_, httpResp, _ := TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH, TESTADDLAMBDAARN)
	assert.Equal(t, httpResp.StatusCode, 200)

	_, httpResp, _ = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH, TESTSAVESERVICES)
	assert.Equal(t, httpResp.StatusCode, 200)
}

func TestListAWSLogsServices(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	list_services_output, _, _ := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsServicesList(TESTAUTH)

	// Assert returned list has the expected length of 6
	assert.Equal(t, len(list_services_output), len(TESTSAVESERVICES.Services))
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallAWSIntegration(TESTAWSACCLOGS)
	// Create the AWS integration.
	TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, TESTAWSACCLOGS)

	// Add Lambda to Account
	add_output, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH, TESTADDLAMBDAARN)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error Adding Lambda %v: Status: %v: %v", add_output, httpresp.StatusCode, err)
	}

	// Enable services for Lambda
	TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH, TESTSAVESERVICES)

	// List AWS Logs integrations before deleting
	list_output_1, _, _ := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH)
	// Iterate over output and list Lambdas
	var list_of_arns []string
	for _, Account := range list_output_1 {
		for _, lambda := range Account.GetLambdas() {
			list_of_arns = append(list_of_arns, *lambda.Arn)
		}
	}
	// Set a variable to check if ARN exists
	var x = false
	// Check if ARN exists, if so, set true.
	for _, arn := range list_of_arns {
		if arn == "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest" {
			x = true
		}
	}
	// Test that variable is true as expected
	assert.Equal(t, x, true)

	// Delete newly added Lambda
	delete_output, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.DeleteAWSLambdaARN(TESTAUTH, TESTADDLAMBDAARN)
	if err != nil || httpResp.StatusCode != 200 {
		t.Errorf("Error Deleting Lambda %v: Status: %v: %v", delete_output, httpResp.StatusCode, err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// List AWS logs integrations after deleting
	list_output_2, _, _ := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH)

	var list_of_arns_2 []datadog.AwsLogsListOutputLambdas
	x = false
	for _, Account := range list_output_2 {
		if Account.GetAccountId() == UNIQUEACCOUNTID {
			list_of_arns_2 = Account.GetLambdas()
		}
	}
	for _, lambda := range list_of_arns_2 {
		if lambda.GetArn() == "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest" {
			x = true
		}
	}
	// Check that ARN no longer exists after delete
	assert.Assert(t, x != true)
}
