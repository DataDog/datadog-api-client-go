// Search cost recommendations returns "OK" response

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
	body := datadogV2.RecommendationsFilterRequest{
		Filter: datadog.PtrString("@resource_table:aws_ec2_instance"),
		Sort: []datadogV2.RecommendationsFilterRequestSortItems{
			{
				Expression: datadog.PtrString("potential_daily_savings.amount"),
				Order:      datadog.PtrString("DESC"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchCostRecommendations", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.SearchCostRecommendations(ctx, body, *datadogV2.NewSearchCostRecommendationsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.SearchCostRecommendations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.SearchCostRecommendations`:\n%s\n", responseContent)
}
