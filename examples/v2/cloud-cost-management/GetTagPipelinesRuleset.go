// Get a tag pipeline ruleset returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.GetTagPipelinesRuleset(ctx, "a1e9de9b-b88e-41c6-a0cd-cc0ebd7092de")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.GetTagPipelinesRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.GetTagPipelinesRuleset`:\n%s\n", responseContent)
}
