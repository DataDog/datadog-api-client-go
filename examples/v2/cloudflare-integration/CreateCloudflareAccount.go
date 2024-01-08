// Add Cloudflare account returns "CREATED" response

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
	body := datadogV2.CloudflareAccountCreateRequest{
		Data: datadogV2.CloudflareAccountCreateRequestData{
			Attributes: datadogV2.CloudflareAccountCreateRequestAttributes{
				ApiKey: "fakekey",
				Email:  datadog.PtrString("dev@datadoghq.com"),
				Name:   "examplecloudflareintegration",
			},
			Type: datadogV2.CLOUDFLAREACCOUNTTYPE_CLOUDFLARE_ACCOUNTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudflareIntegrationApi(apiClient)
	resp, r, err := api.CreateCloudflareAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareIntegrationApi.CreateCloudflareAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudflareIntegrationApi.CreateCloudflareAccount`:\n%s\n", responseContent)
}
