// Add Fastly service returns "CREATED" response

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
	body := datadogV2.FastlyServiceRequest{
		Data: datadogV2.FastlyServiceData{
			Attributes: &datadogV2.FastlyServiceAttributes{
				Tags: []string{
					"myTag",
					"myTag2:myValue",
				},
			},
			Id:   "abc123",
			Type: datadogV2.FASTLYSERVICETYPE_FASTLY_SERVICES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFastlyIntegrationApi(apiClient)
	resp, r, err := api.CreateFastlyService(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.CreateFastlyService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FastlyIntegrationApi.CreateFastlyService`:\n%s\n", responseContent)
}
