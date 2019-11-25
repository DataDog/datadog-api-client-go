package datadog_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var UNIQUEPROJECTID = fmt.Sprintf("datadog-%s", time.Now())

var TESTGCPACCT = datadog.GcpAccount{
	Type:                    datadog.PtrString("service_account"),
	ProjectId:               datadog.PtrString(UNIQUEPROJECTID),
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

var TESTUPDATEGCPACCT = datadog.GcpAccount{
	ProjectId:   datadog.PtrString(UNIQUEPROJECTID),
	ClientEmail: datadog.PtrString("api-test@fake-sandbox.iam.gserviceaccount.com"),
	HostFilters: datadog.PtrString("fake:update,example:update"),
	Automute:    datadog.PtrBool(true),
}

func TestGcpCreate(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallGcpIntegration(TESTGCPACCT)

	createOutput, _, _ := TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, TESTGCPACCT)
	assert.Equal(t, len(createOutput), 0)
}

func TestGcpListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallGcpIntegration(TESTGCPACCT)

	// Setup Gcp Account to List
	TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, TESTGCPACCT)

	gcpListOutput, _, _ := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH)
	var x datadog.GcpAccount
	for _, Account := range gcpListOutput {
		if Account.GetProjectId() == UNIQUEPROJECTID {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientEmail(), "api-test@fake-sandbox.iam.gserviceaccount.com")
	assert.Equal(t, x.GetHostFilters(), "fake:tag,example:test")

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(gcpListOutput) >= 1)

	// Test account deletion as well
	deleteOutput, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH, TESTGCPACCT)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", TESTAZUREACCT)
	}
	assert.Equal(t, len(deleteOutput), 0)
}

func TestUpdateGcpAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	defer uninstallGcpIntegration(TESTGCPACCT)

	// Setup Gcp Account to Update
	TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH, TESTGCPACCT)

	gcpUpdateOutput, _, err := TESTAPICLIENT.GCPIntegrationApi.UpdateGCPIntegration(TESTAUTH, TESTUPDATEGCPACCT)
	if err != nil {
		t.Errorf("Error Updating GCP Account: %v", err)
	}

	assert.Equal(t, len(gcpUpdateOutput), 0)

	// List account to ensure update worked.
	gcpListOutput, _, _ := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH)
	var x datadog.GcpAccount
	for _, Account := range gcpListOutput {
		if Account.GetClientEmail() == "api-test@fake-sandbox.iam.gserviceaccount.com" {
			x = Account
		}
	}
	assert.Equal(t, x.GetAutomute(), true)
	assert.Equal(t, x.GetHostFilters(), "fake:update,example:update")
}

func uninstallGcpIntegration(account datadog.GcpAccount) {
	_, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH, account)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", account)
	}
}
