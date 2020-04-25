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

func generateUniqueAWSAccount(c *Client) datadog.AWSAccount {
	accountID := fmt.Sprintf("go_%09d", c.Clock.Now().UnixNano()%1000000000)
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
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testAWSAccount := generateUniqueAWSAccount(c)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(c, t, testAWSAccount)
	defer retryDeleteAccount(c, t, testAWSAccount)

	awsAccts, httpresp, err := c.Client.AWSIntegrationApi.
		ListAWSAccounts(c.Ctx).
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
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testAWSAccount := generateUniqueAWSAccount(c)

	// Assert AWS Integration Created with proper fields
	retryCreateAccount(c, t, testAWSAccount)
	defer retryDeleteAccount(c, t, testAWSAccount)

	retryUpdateAccount(c, t, TESTUPDATEAWSACCWITHEXCLUDEDREGION, testAWSAccount.GetAccountId(), testAWSAccount.GetRoleName())
	UPDATEDAWSACCT := datadog.AWSAccount{
		AccountId: testAWSAccount.AccountId,
		RoleName:  TESTUPDATEAWSACCWITHEXCLUDEDREGION.RoleName,
	}
	defer retryDeleteAccount(c, t, UPDATEDAWSACCT)

	// Assert AWS Account Get with proper fields
	awsAccts, httpresp, err := c.Client.AWSIntegrationApi.ListAWSAccounts(c.Ctx).
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
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// We already test this in the disableAWSAccount cleanup function, but good to have an explicit test
	testAWSAccount := generateUniqueAWSAccount(c)

	// Lets first create the account of us to delete
	retryCreateAccount(c, t, testAWSAccount)

	retryDeleteAccount(c, t, testAWSAccount)
}

func TestGenerateNewExternalId(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testAWSAccount := generateUniqueAWSAccount(c)
	// Lets first create the account for us to generate a new id against
	retryCreateAccount(c, t, testAWSAccount)
	defer retryDeleteAccount(c, t, testAWSAccount)

	apiResp, httpresp, err := c.Client.AWSIntegrationApi.CreateNewAWSExternalID(c.Ctx).Body(testAWSAccount).Execute()
	if err != nil {
		t.Fatalf("Error generating new AWS External ID: Response: %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotEmpty(t, apiResp.GetExternalId())
}

func TestListNamespaces(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	namespaces, httpresp, err := c.Client.AWSIntegrationApi.ListAvailableAWSNamespaces(c.Ctx).Execute()
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
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.AWSIntegrationApi.CreateNewAWSExternalID(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.AWSIntegrationApi.CreateAWSAccount(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.AWSIntegrationApi.DeleteAWSAccount(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestAWSIntegrationGetAll403Error(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSIntegrationApi.ListAWSAccounts(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationGetAll400Error(t *testing.T) {
	c := NewClient(context.Background(), t)
	defer c.Close()

	res, err := tests.ReadFixture("fixtures/aws/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/integration/aws").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := c.Client.AWSIntegrationApi.ListAWSAccounts(c.Ctx).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationListNamespacesErrors(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// 403 Forbidden
	_, httpresp, err := c.Client.AWSIntegrationApi.ListAvailableAWSNamespaces(context.Background()).Execute()
	assert.Equal(t, 403, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestAWSIntegrationUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.AWSIntegrationApi.UpdateAWSAccount(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func retryDeleteAccount(c *Client, t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := c.Client.AWSIntegrationApi.DeleteAWSAccount(c.Ctx).Body(awsAccount).Execute()
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error deleting AWS Account: Response %s", err)
	}
}

func retryCreateAccount(c *Client, t *testing.T, awsAccount datadog.AWSAccount) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := c.Client.AWSIntegrationApi.CreateAWSAccount(c.Ctx).Body(awsAccount).Execute()
		if httpresp.StatusCode == 502 {
			return false
		}
		return true
	})
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s", err)
	}
}

func retryUpdateAccount(c *Client, t *testing.T, body datadog.AWSAccount, accountID string, roleName string) {
	err := tests.Retry(time.Duration(rand.Intn(10))*time.Second, 10, func() bool {
		_, httpresp, _ := c.Client.AWSIntegrationApi.UpdateAWSAccount(c.Ctx).
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
