package datadog_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueGcpAccount() (datadog.GcpAccount, datadog.GcpAccount) {
	projectId := fmt.Sprintf("go_%s", time.Now())
	var testGcpAccount = datadog.GcpAccount{
		Type:                    datadog.PtrString("service_account"),
		ProjectId:               datadog.PtrString(projectId),
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

	var testUpdateGcpAccount = datadog.GcpAccount{
		ProjectId:   datadog.PtrString(projectId),
		ClientEmail: testGcpAccount.ClientEmail,
		HostFilters: datadog.PtrString("fake:update,example:update"),
		Automute:    datadog.PtrBool(true),
	}

	return testGcpAccount, testUpdateGcpAccount
}

func TestGcpCreate(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, _ := generateUniqueGcpAccount()
	defer uninstallGcpIntegration(testGCPAcct)

	_, httpResp, err := TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, testGCPAcct)
	if err != nil {
		t.Errorf("Error creating GCP intg: %v", err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)
}

func TestGcpListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, _ := generateUniqueGcpAccount()
	defer uninstallGcpIntegration(testGCPAcct)

	// Setup Gcp Account to List
	TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, testGCPAcct)

	gcpListOutput, _, err := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH)
	if err != nil {
		t.Errorf("Error listing GCP intg: %v", err)
	}
	var x datadog.GcpAccount
	for _, Account := range gcpListOutput {
		if Account.GetProjectId() == *testGCPAcct.ProjectId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientEmail(), *testGCPAcct.ClientEmail)
	assert.Equal(t, x.GetHostFilters(), *testGCPAcct.HostFilters)

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(gcpListOutput) >= 1)

	// Test account deletion as well
	_, httpResp, err := TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH, testGCPAcct)
	if httpResp.StatusCode != 200 || err != nil {
		t.Errorf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", testGCPAcct)
	}
	assert.Equal(t, httpResp.StatusCode, 200)
}

func TestUpdateGcpAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, testGCPUpdateAcct := generateUniqueGcpAccount()
	defer uninstallGcpIntegration(testGCPAcct)

	// Setup Gcp Account to Update
	TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, testGCPAcct)

	_, httpResp, err := TESTAPICLIENT.GCPIntegrationApi.UpdateGCPIntegration(TESTAUTH, testGCPUpdateAcct)
	if err != nil {
		t.Errorf("Error Updating GCP Account: %v", err)
	}

	assert.Equal(t, httpResp.StatusCode, 200)

	// List account to ensure update worked.
	gcpListOutput, _, err := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH)
	if err != nil {
		t.Errorf("Error listing GCP intg: %v", err)
	}
	var x datadog.GcpAccount
	for _, Account := range gcpListOutput {
		if Account.GetClientEmail() == *testGCPAcct.ClientEmail {
			x = Account
		}
	}
	assert.Equal(t, x.GetAutomute(), true)
	assert.Equal(t, x.GetHostFilters(), *testGCPUpdateAcct.HostFilters)
}

func uninstallGcpIntegration(account datadog.GcpAccount) {
	_, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", account)
	}
}
