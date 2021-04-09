/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func generateUniqueAWSAccount(ctx context.Context, t *testing.T) datadog.AWSAccount {
	return datadog.AWSAccount{
		AccountId:                     tests.UniqueEntityName(ctx, t),
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAccount := generateUniqueAWSAccount(ctx, t)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(ctx, t, testAWSAccount)
	defer retryDeleteAccount(ctx, t, testAWSAccount)

	awsAccts, httpresp, err := Client(ctx).AWSIntegrationApi.
		ListAWSAccounts(ctx, *datadog.NewListAWSAccountsParameters().
			WithAccountId(testAWSAccount.GetAccountId()).
			WithRoleName(testAWSAccount.GetRoleName()))
	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	awsAccounts := awsAccts.GetAccounts()
	assert.True(len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	assert.Equal(testAWSAccount.GetAccountId(), awsAcct.GetAccountId())
	assert.Equal(testAWSAccount.GetRoleName(), awsAcct.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(testAWSAccount.GetHostTags(), awsAcct.GetHostTags())
	assert.Equal(testAWSAccount.GetFilterTags(), awsAcct.GetFilterTags())
	assert.Equal(testAWSAccount.GetExcludedRegions(), awsAcct.GetExcludedRegions())
	assert.Equal(testAWSAccount.GetAccountSpecificNamespaceRules(), awsAcct.GetAccountSpecificNamespaceRules())
}

func TestUpdateAWSAccount(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAccount := generateUniqueAWSAccount(ctx, t)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(ctx, t, testAWSAccount)
	defer retryDeleteAccount(ctx, t, testAWSAccount)

	retryUpdateAccount(ctx, t, TESTUPDATEAWSACCWITHEXCLUDEDREGION, testAWSAccount.GetAccountId(), testAWSAccount.GetRoleName())
	UPDATEDAWSACCT := datadog.AWSAccount{
		AccountId: testAWSAccount.AccountId,
		RoleName:  TESTUPDATEAWSACCWITHEXCLUDEDREGION.RoleName,
	}
	defer retryDeleteAccount(ctx, t, UPDATEDAWSACCT)

	// Assert AWS Account Get with proper fields
	awsAccts, httpresp, err := Client(ctx).AWSIntegrationApi.ListAWSAccounts(ctx, *datadog.NewListAWSAccountsParameters().
		WithAccountId(UPDATEDAWSACCT.GetAccountId()).
		WithRoleName(UPDATEDAWSACCT.GetRoleName()))

	if err != nil {
		t.Fatalf("Error getting AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	awsAccounts := awsAccts.GetAccounts()
	assert.True(len(awsAccounts) > 0, "No AWS accounts found")

	awsAcct := awsAccounts[0]
	// Test fields were updated
	assert.Equal(TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetRoleName(), awsAcct.GetRoleName())
	// Golang doesn't have an equality operator defined for slices
	assert.Equal(TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetHostTags(), awsAcct.GetHostTags())
	assert.Equal(TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetFilterTags(), awsAcct.GetFilterTags())
	assert.Equal(TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetExcludedRegions(), awsAcct.GetExcludedRegions())
	assert.Equal(TESTUPDATEAWSACCWITHEXCLUDEDREGION.GetAccountSpecificNamespaceRules(), awsAcct.GetAccountSpecificNamespaceRules())
}

func TestDisableAWSAcct(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test
	testAWSAccount := generateUniqueAWSAccount(ctx, t)

	// Lets first create the account of us to delete
	retryCreateAccount(ctx, t, testAWSAccount)

	retryDeleteAccount(ctx, t, testAWSAccount)
}

func TestGenerateNewExternalId(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testAWSAccount := generateUniqueAWSAccount(ctx, t)
	// Lets first create the account for us to generate a new id against
	retryCreateAccount(ctx, t, testAWSAccount)
	defer retryDeleteAccount(ctx, t, testAWSAccount)

	apiResp, httpresp, err := Client(ctx).AWSIntegrationApi.CreateNewAWSExternalID(ctx, testAWSAccount)
	if err != nil {
		t.Fatalf("Error generating new AWS External ID: Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.NotEmpty(apiResp.GetExternalId())
}

func TestListNamespaces(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	namespaces, httpresp, err := Client(ctx).AWSIntegrationApi.ListAvailableAWSNamespaces(ctx)
	if err != nil {
		t.Fatalf("Error listing AWS Namespaces: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	namespacesCheck := make(map[string]bool)
	for _, namespace := range namespaces {
		namespacesCheck[namespace] = true
	}

	// Check that a few examples are in the response
	// Full list - https://docs.datadoghq.com/api/?lang=bash#list-namespace-rules
	assert.True(namespacesCheck["api_gateway"])
	assert.True(namespacesCheck["cloudsearch"])
	assert.True(namespacesCheck["directconnect"])
	assert.True(namespacesCheck["xray"])
}

func TestAWSIntegrationGenerateExternalIDErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSIntegrationApi.CreateNewAWSExternalID(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSIntegrationApi.CreateAWSAccount(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationCreateConflictErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	testAWSAccount := generateUniqueAWSAccount(ctx, t)
	retryCreateAccount(ctx, t, testAWSAccount)
	defer retryDeleteAccount(ctx, t, testAWSAccount)

	role := *testAWSAccount.RoleName
	for name, test := range map[string]*string{
		"Same":      datadog.PtrString(role),
		"Different": datadog.PtrString("Different-Role"),
	} {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(ctx, t)
			defer finish()
			assert := tests.Assert(ctx, t)

			testAWSAccount.RoleName = test
			_, httpresp, _ := Client(ctx).AWSIntegrationApi.CreateAWSAccount(ctx, testAWSAccount)
			assert.Equal(409, httpresp.StatusCode)
		})
	}
}

func TestAWSIntegrationDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSIntegrationApi.DeleteAWSAccount(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationGetAll403Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSIntegrationApi.ListAWSAccounts(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSIntegrationGetAll400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(ctx)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "AWSIntegrationApiService.ListAWSAccounts")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/integration/aws").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).AWSIntegrationApi.ListAWSAccounts(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSIntegrationListNamespacesErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// 403 Forbidden
	_, httpresp, err := Client(ctx).AWSIntegrationApi.ListAvailableAWSNamespaces(context.Background())
	assert.Equal(403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestAWSIntegrationUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.AWSAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.AWSAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.AWSAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).AWSIntegrationApi.UpdateAWSAccount(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func retryDeleteAccount(ctx context.Context, t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := Client(ctx).AWSIntegrationApi.DeleteAWSAccount(ctx, awsAccount)
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error deleting AWS Account: Response %s", err)
	}
}

func retryCreateAccount(ctx context.Context, t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := Client(ctx).AWSIntegrationApi.CreateAWSAccount(ctx, awsAccount)
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s", err)
	}
}

func retryUpdateAccount(ctx context.Context, t *testing.T, body datadog.AWSAccount, accountID string, roleName string) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := Client(ctx).AWSIntegrationApi.UpdateAWSAccount(ctx, body, *datadog.NewUpdateAWSAccountParameters().
			WithAccountId(accountID).
			WithRoleName(roleName))
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error updating AWS Account: Response %s", err)
	}
}
