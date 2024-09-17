// Update Cloudflare account returns "OK" response

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
	// there is a valid "cloudflare_account" in the system
	CloudflareAccountDataID := os.Getenv("CLOUDFLARE_ACCOUNT_DATA_ID")

	body := datadogV2.CloudflareAccountUpdateRequest{
		Data: datadogV2.CloudflareAccountUpdateRequestData{
			Attributes: &datadogV2.CloudflareAccountUpdateRequestAttributes{
				ApiKey: "fakekey",
				Email:  datadog.PtrString("dev@datadoghq.com"),
				Zones: []string{
					"zone-id-3",
				},
			},
			Type: datadogV2.CLOUDFLAREACCOUNTTYPE_CLOUDFLARE_ACCOUNTS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudflareIntegrationApi(apiClient)
	resp, r, err := api.UpdateCloudflareAccount(ctx, CloudflareAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareIntegrationApi.UpdateCloudflareAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudflareIntegrationApi.UpdateCloudflareAccount`:\n%s\n", responseContent)
}
