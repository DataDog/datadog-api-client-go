// Create a new entry for your service account returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.GCPSTSServiceAccountCreateRequest{
		Data: &datadogV2.GCPSTSServiceAccountData{
			Attributes: &datadogV2.GCPSTSServiceAccountAttributes{
				ClientEmail: datadog.PtrString("Test-252bf553ef04b351@test-project.iam.gserviceaccount.com"),
				HostFilters: []string{},
			},
			Type: datadogV2.GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.CreateGCPSTSAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.CreateGCPSTSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.CreateGCPSTSAccount`:\n%s\n", responseContent)
}
