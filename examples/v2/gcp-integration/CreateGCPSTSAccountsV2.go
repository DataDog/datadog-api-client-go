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
	body := datadogV2.ServiceAccountToBeCreatedData{
		Data: &datadogV2.ServiceAccountMetadata{
			Attributes: &datadogV2.AttributeMetadata{
				ClientEmail: datadog.PtrString("dd-integration@datadog-staging.iam.gserviceaccount.com"),
				HostFilters: []string{},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.CreateGCPSTSAccountsV2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.CreateGCPSTSAccountsV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.CreateGCPSTSAccountsV2`:\n%s\n", responseContent)
}
