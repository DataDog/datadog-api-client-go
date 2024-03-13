// Update STS Service Account returns "OK" response

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
	// there is a valid "gcp_sts_account" in the system
	GcpStsAccountDataID := os.Getenv("GCP_STS_ACCOUNT_DATA_ID")

	body := datadogV2.GCPSTSServiceAccountUpdateRequest{
		Data: &datadogV2.GCPSTSServiceAccountUpdateRequestData{
			Attributes: &datadogV2.GCPSTSServiceAccountAttributes{
				ClientEmail: datadog.PtrString("Test-252bf553ef04b351@example.com"),
				HostFilters: []string{
					"foo:bar",
				},
			},
			Id:   datadog.PtrString(GcpStsAccountDataID),
			Type: datadogV2.GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.UpdateGCPSTSAccount(ctx, GcpStsAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPSTSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.UpdateGCPSTSAccount`:\n%s\n", responseContent)
}
