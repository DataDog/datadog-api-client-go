/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func generateUniqueGCPAccount(ctx context.Context) (datadog.GCPAccount, datadog.GCPAccount) {
	projectID := fmt.Sprintf("go_%d", tests.ClockFromContext(ctx).Now().Unix())
	var testGCPAccount = datadog.GCPAccount{
		Type:                    datadog.PtrString("service_account"),
		ProjectId:               datadog.PtrString(projectID),
		PrivateKeyId:            datadog.PtrString("fake_private_key_id"),
		PrivateKey:              datadog.PtrString("fake_key"),
		ClientEmail:             datadog.PtrString("api-test@fake-sandbox.iam.gserviceaccount.com"),
		ClientId:                datadog.PtrString("123456712345671234567"),
		AuthUri:                 datadog.PtrString("fake_uri"),
		TokenUri:                datadog.PtrString("fake_uri"),
		AuthProviderX509CertUrl: datadog.PtrString("fake_url"),
		ClientX509CertUrl:       datadog.PtrString("fake_url"),
		HostFilters:             datadog.PtrString("fake:tag,example:test"),
		Automute:                datadog.PtrBool(false),
	}

	var testUpdateGCPAccount = datadog.GCPAccount{
		ProjectId:   datadog.PtrString(projectID),
		ClientEmail: testGCPAccount.ClientEmail,
		HostFilters: datadog.PtrString("fake:update,example:update"),
		Automute:    datadog.PtrBool(true),
	}

	return testGCPAccount, testUpdateGCPAccount
}

func TestGCPCreate(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	testGCPAcct, _ := generateUniqueGCPAccount(ctx)
	defer uninstallGCPIntegration(ctx, testGCPAcct)

	_, httpresp, err := Client(ctx).GCPIntegrationApi.CreateGCPIntegration(ctx).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestGCPListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	testGCPAcct, _ := generateUniqueGCPAccount(ctx)
	defer uninstallGCPIntegration(ctx, testGCPAcct)

	// Setup GCP Account to List
	Client(ctx).GCPIntegrationApi.CreateGCPIntegration(ctx).Body(testGCPAcct).Execute()

	gcpListOutput, httpresp, err := Client(ctx).GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing GCP Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetProjectId() == *testGCPAcct.ProjectId {
			x = Account
		}
	}
	assert.Equal(t, *testGCPAcct.ClientEmail, x.GetClientEmail())
	assert.Equal(t, *testGCPAcct.HostFilters, x.GetHostFilters())

	// Assert returned list is greater than or equal to 1
	assert.True(t, len(gcpListOutput) >= 1)

	// Test account deletion as well
	_, httpresp, err = Client(ctx).GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error uninstalling GCP Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestUpdateGCPAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	testGCPAcct, testGCPUpdateAcct := generateUniqueGCPAccount(ctx)
	defer uninstallGCPIntegration(ctx, testGCPAcct)

	// Setup GCP Account to Update
	_, httpresp, err := Client(ctx).GCPIntegrationApi.CreateGCPIntegration(ctx).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	_, httpresp, err = Client(ctx).GCPIntegrationApi.UpdateGCPIntegration(ctx).Body(testGCPUpdateAcct).Execute()
	if err != nil {
		t.Fatalf("Error updating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// List account to ensure update worked.
	gcpListOutput, _, err := Client(ctx).GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing GCP accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetClientEmail() == *testGCPAcct.ClientEmail {
			x = Account
		}
	}
	assert.Equal(t, true, x.GetAutomute())
	assert.Equal(t, *testGCPUpdateAcct.HostFilters, x.GetHostFilters())
}

func TestGCPList400Error(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer stop()

	res, err := tests.ReadFixture("fixtures/gcp/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the gcp integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/integration/gcp").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestGCPListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestGCPCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).GCPIntegrationApi.CreateGCPIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestGCPDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestGCPUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.GCPAccount
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.GCPAccount{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.GCPAccount{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).GCPIntegrationApi.UpdateGCPIntegration(ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func uninstallGCPIntegration(ctx context.Context, account datadog.GCPAccount) {
	_, httpresp, err := Client(ctx).GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(account).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", account)
	}
}
