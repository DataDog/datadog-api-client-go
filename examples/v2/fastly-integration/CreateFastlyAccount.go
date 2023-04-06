// Add Fastly account returns "CREATED" response

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
	body := datadogV2.FastlyAccountCreateRequest{
		Data: datadogV2.FastlyAccountCreateRequestData{
			Attributes: datadogV2.FastlyAccountCreateRequestAttributes{
				ApiKey:   "ExampleFastlyIntegration",
				Name:     "Example-Fastly-Integration",
				Services: []datadogV2.FastlyService{},
			},
			Type: datadogV2.FASTLYACCOUNTTYPE_FASTLY_ACCOUNTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFastlyIntegrationApi(apiClient)
	resp, r, err := api.CreateFastlyAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.CreateFastlyAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FastlyIntegrationApi.CreateFastlyAccount`:\n%s\n", responseContent)
}
