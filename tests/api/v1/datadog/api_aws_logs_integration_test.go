/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func generateUniqueAWSLambdaAccounts() (datadog.AWSAccount, datadog.AWSAccountAndLambdaRequest, datadog.AWSLogsServicesRequest) {
	accountID := fmt.Sprintf("go_%09d", c.Clock.Now().UnixNano()%1000000000)
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
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	testawsacc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testawsacc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testawsacc)

	_, httpresp, err := c.Client.AWSLogsIntegrationApi.CreateAWSLambdaARN(c.Ctx).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error adding lamda ARN: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	_, httpresp, err = c.Client.AWSLogsIntegrationApi.EnableAWSLogServices(c.Ctx).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestListAWSLogsServices(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	listServicesOutput, httpresp, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsServices(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// There are currently 6 supported AWS Logs services as noted in the docs
	// https://docs.datadoghq.com/api/?lang=bash#get-list-of-aws-log-ready-services
	assert.True(t, len(listServicesOutput) >= 6)
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	testAWSAcc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Create the AWS integration.
	retryCreateAccount(t, testAWSAcc)

	// Add Lambda to Account
	addOutput, httpresp, err := c.Client.AWSLogsIntegrationApi.CreateAWSLambdaARN(c.Ctx).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error Adding Lambda %v: Response %s: %v", addOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// Enable services for Lambda
	_, httpresp, err = c.Client.AWSLogsIntegrationApi.EnableAWSLogServices(c.Ctx).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// List AWS Logs integrations before deleting
	listOutput1, _, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsIntegrations(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
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
	assert.True(t, accountExists)

	// Delete newly added Lambda
	deleteOutput, httpresp, err := c.Client.AWSLogsIntegrationApi.DeleteAWSLambdaARN(c.Ctx).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error deleting Lambda %v: Response %s: %v", deleteOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// List AWS logs integrations after deleting
	listOutput2, httpresp, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsIntegrations(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

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
	assert.False(t, accountExistsAfterDelete)
}

func TestCheckLambdaAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testAWSAcc, testLambdaAcc, _ := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAcc)

	status, httpresp, err := c.Client.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(c.Ctx).Body(testLambdaAcc).Execute()
	if err != nil {
		t.Fatalf("Error checking the AWS Lambda Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	assert.Equal(t, 0, len(status.GetErrors()))
	assert.Equal(t, "created", status.GetStatus())

	// Give the async call time to finish
	tests.Retry(time.Duration(5*time.Second), 10, func() bool {
		status, httpresp, err = c.Client.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(c.Ctx).Body(testLambdaAcc).Execute()
		if err != nil {
			t.Logf("Error checking the AWS Lambda Response: %s %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		return httpresp.StatusCode == 200 && len(status.GetErrors()) > 0
	})

	assert.NotEmpty(t, status.GetErrors()[0].GetCode())
	assert.NotEmpty(t, status.GetErrors()[0].GetMessage())
	assert.Equal(t, "error", status.GetStatus())
}

func TestCheckServicesAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	testAWSAcc, _, testServices := generateUniqueAWSLambdaAccounts()
	defer retryDeleteAccount(t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(t, testAWSAcc)

	status, httpresp, err := c.Client.AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(c.Ctx).Body(testServices).Execute()
	if err != nil {
		t.Fatalf("Error checking the AWS Logs Services Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.NotEmpty(t, status.GetErrors()[0].GetCode())
	assert.NotEmpty(t, status.GetErrors()[0].GetMessage())
}

func TestAWSLogsList400Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer c.Close()

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsIntegrations(c.Ctx).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsList403Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsIntegrations(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsAdd400Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer c.Close()

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Post("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.CreateAWSLambdaARN(c.Ctx).Body(datadog.AWSAccountAndLambdaRequest{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsAdd403Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.CreateAWSLambdaARN(context.Background()).Body(datadog.AWSAccountAndLambdaRequest{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsDelete400Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer c.Close()

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Delete("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.DeleteAWSLambdaARN(c.Ctx).Body(datadog.AWSAccountAndLambdaRequest{}).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsDelete403Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.DeleteAWSLambdaARN(context.Background()).Body(datadog.AWSAccountAndLambdaRequest{}).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsServicesListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSLogsIntegrationApi.ListAWSLogsServices(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSLogsServicesEnableErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.AWSLogsServicesRequest
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.AWSLogsServicesRequest{}, 400},
		{"403 Forbidden", context.Background(), datadog.AWSLogsServicesRequest{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.AWSLogsIntegrationApi.EnableAWSLogServices(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAWSLogsServicesCheckErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.AWSLogsServicesRequest
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.AWSLogsServicesRequest{}, 400},
		{"403 Forbidden", context.Background(), datadog.AWSLogsServicesRequest{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

// FIXME: Right now we get 502s for these request instead of 400 or 403.
func TestAWSLogsLambdaCheckErrors(t *testing.T) {
	t.Skip("Receiving 502 instead of 400 or 403, so skipping")
	//// Setup the Client we'll use to interact with the Test account
	//c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	//defer c.Close()

	//testCases := []struct {
	//	Name               string
	//	Ctx                context.Context
	//	Body               datadog.AWSAccountAndLambdaRequest
	//	ExpectedStatusCode int
	//}{
	//	{"400 Bad Request", c.Ctx, datadog.AWSAccountAndLambdaRequest{}, 400},
	//	{"403 Forbidden", context.Background(), datadog.AWSAccountAndLambdaRequest{}, 403},
	//}

	//for _, tc := range testCases {
	//	t.Run(tc.Name, func(t *testing.T) {
	//		_, httpresp, err := c.Client.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(tc.Ctx).Body(tc.Body).Execute()
	//		assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
	//		apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	//		assert.True(t, ok)
	//		assert.NotEmpty(t, apiError.GetErrors())
	//	})
	//}
}
