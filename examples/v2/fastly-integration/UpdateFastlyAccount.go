// Update Fastly account returns "OK" response

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
	// there is a valid "fastly_account" in the system
	FastlyAccountDataID := os.Getenv("FASTLY_ACCOUNT_DATA_ID")

	body := datadogV2.FastlyAccountUpdateRequest{
		Data: datadogV2.FastlyAccountUpdateRequestData{
			Attributes: &datadogV2.FastlyAccountUpdateRequestAttributes{
				ApiKey: datadog.PtrString("update-secret"),
			},
			Type: datadogV2.FASTLYACCOUNTTYPE_FASTLY_ACCOUNTS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFastlyIntegrationApi(apiClient)
	resp, r, err := api.UpdateFastlyAccount(ctx, FastlyAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.UpdateFastlyAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FastlyIntegrationApi.UpdateFastlyAccount`:\n%s\n", responseContent)
}
