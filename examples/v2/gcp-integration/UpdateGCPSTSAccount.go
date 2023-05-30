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
	body := datadogV2.GCPSTSServiceAccountUpdateRequest{
		Data: &datadogV2.GCPSTSServiceAccountUpdateRequestData{
			Attributes: &datadogV2.GCPSTSServiceAccountAttributes{
				ClientEmail: datadog.PtrString("datadog-service-account@test-project.iam.gserviceaccount.com"),
				HostFilters: []string{},
			},
			Id:   datadog.PtrString("d291291f-12c2-22g4-j290-123456678897"),
			Type: datadogV2.GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.UpdateGCPSTSAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPSTSAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.UpdateGCPSTSAccount`:\n%s\n", responseContent)
}
