/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"
	"time"

	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func generateUniqueAWSLambdaAccounts(ctx context.Context, t *testing.T) (datadog.AWSAccount, datadog.AWSAccountAndLambdaRequest, datadog.AWSLogsServicesRequest) {
	var uniqueAWSAccount = datadog.AWSAccount{
		AccountId:                     tests.UniqueEntityName(ctx, t),
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testawsacc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts(ctx, t)
	defer retryDeleteAccount(ctx, t, testawsacc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(ctx, t, testawsacc)

	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CreateAWSLambdaARN(ctx, testLambdaAcc)
	if err != nil {
		t.Fatalf("Error adding lamda ARN: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).AWSLogsIntegrationApi.EnableAWSLogServices(ctx, testServices)
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestListAWSLogsServices(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	listServicesOutput, httpresp, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsServices(ctx)
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// There are currently 6 supported AWS Logs services as noted in the docs
	// https://docs.datadoghq.com/api/?lang=bash#get-list-of-aws-log-ready-services
	assert.True(len(listServicesOutput) >= 6)
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAcc, testLambdaAcc, testServices := generateUniqueAWSLambdaAccounts(ctx, t)
	defer retryDeleteAccount(ctx, t, testAWSAcc)

	// Create the AWS integration.
	retryCreateAccount(ctx, t, testAWSAcc)

	// Add Lambda to Account
	addOutput, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CreateAWSLambdaARN(ctx, testLambdaAcc)
	if err != nil {
		t.Fatalf("Error Adding Lambda %v: Response %s: %v", addOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// Enable services for Lambda
	_, httpresp, err = Client(ctx).AWSLogsIntegrationApi.EnableAWSLogServices(ctx, testServices)
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// List AWS Logs integrations before deleting
	listOutput1, _, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsIntegrations(ctx)
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
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
	assert.True(accountExists)

	// Delete newly added Lambda
	deleteOutput, httpresp, err := Client(ctx).AWSLogsIntegrationApi.DeleteAWSLambdaARN(ctx, testLambdaAcc)
	if err != nil {
		t.Fatalf("Error deleting Lambda %v: Response %s: %v", deleteOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// List AWS logs integrations after deleting
	listOutput2, httpresp, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsIntegrations(ctx)
	if err != nil {
		t.Fatalf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	var listOfARNs2 []datadog.AWSLogsLambda
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
	assert.False(accountExistsAfterDelete)
}

func TestCheckLambdaAsync(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAcc, testLambdaAcc, _ := generateUniqueAWSLambdaAccounts(ctx, t)
	defer retryDeleteAccount(ctx, t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(ctx, t, testAWSAcc)

	status, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(ctx, testLambdaAcc)
	if err != nil {
		t.Fatalf("Error checking the AWS Lambda Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(0, len(status.GetErrors()))
	assert.Equal("created", status.GetStatus())

	// Give the async call time to finish
	tests.Retry(time.Duration(5*time.Second), 10, func() bool {
		status, httpresp, err = Client(ctx).AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(ctx, testLambdaAcc)
		if err != nil {
			t.Logf("Error checking the AWS Lambda Response: %s %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		return httpresp.StatusCode == 200 && len(status.GetErrors()) > 0
	})

	assert.NotEmpty(status.GetErrors()[0].GetCode())
	assert.NotEmpty(status.GetErrors()[0].GetMessage())
	assert.Equal("error", status.GetStatus())
}

func TestCheckServicesAsync(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAcc, _, testServices := generateUniqueAWSLambdaAccounts(ctx, t)
	defer retryDeleteAccount(ctx, t, testAWSAcc)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(ctx, t, testAWSAcc)

	status, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(ctx, testServices)
	if err != nil {
		t.Fatalf("Error checking the AWS Logs Services Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	assert.NotEmpty(status.GetErrors()[0].GetCode())
	assert.NotEmpty(status.GetErrors()[0].GetMessage())
}

func TestAWSLogsList400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AWSLogsIntegrationApiService.ListAWSLogsIntegrations")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsIntegrations(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsList403Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsIntegrations(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsAdd400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AWSLogsIntegrationApiService.CreateAWSLambdaARN")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CreateAWSLambdaARN(ctx, datadog.AWSAccountAndLambdaRequest{})
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsAdd403Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CreateAWSLambdaARN(context.Background(), datadog.AWSAccountAndLambdaRequest{})
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsDelete400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AWSLogsIntegrationApiService.DeleteAWSLambdaARN")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/integration/aws/logs").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.DeleteAWSLambdaARN(ctx, datadog.AWSAccountAndLambdaRequest{})
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsDelete403Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.DeleteAWSLambdaARN(context.Background(), datadog.AWSAccountAndLambdaRequest{})
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsServicesListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.ListAWSLogsServices(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSLogsServicesEnableErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSLogsServicesRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSLogsServicesRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSLogsServicesRequest{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup the Client we'll use to interact with the Test account
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.EnableAWSLogServices(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAWSLogsServicesCheckErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSLogsServicesRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSLogsServicesRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSLogsServicesRequest{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

// FIXME: Right now we get 502s for these request instead of 400 or 403.
func TestAWSLogsLambdaCheckErrors(t *testing.T) {
	t.Skip("Receiving 502 instead of 400 or 403, so skipping")
	// ctx, close := tests.WithTestSpan(context.Background(), t)
	// defer close()

	//testCases := map[string]struct {
	//	Ctx func(context.Context) context.Context
	//	Body               datadog.AWSAccountAndLambdaRequest
	//	ExpectedStatusCode int
	//}{
	//	"400 Bad Request": {WithTestAuth, datadog.AWSAccountAndLambdaRequest{}, 400},
	//	"403 Forbidden":   {WithFakeAuth, datadog.AWSAccountAndLambdaRequest{}, 403},
	//}

	//for name, tc := range testCases {
	//	t.Run(name, func(t *testing.T) {
	//		_, httpresp, err := Client(ctx).AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(ctx, tc.Body)
	//		assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
	//		apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	//		assert.True(ok)
	//		assert.NotEmpty(apiError.GetErrors())
	//	})
	//}
}
