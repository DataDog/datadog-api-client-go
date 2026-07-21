// Create or replace a budget's custom forecast returns "OK" response

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
	body := datadogV2.CustomForecastUpsertRequest{
		Data: datadogV2.CustomForecastUpsertRequestData{
			Attributes: datadogV2.CustomForecastUpsertRequestDataAttributes{
				BudgetUid: "00000000-0000-0000-0000-000000000001",
				Entries: []datadogV2.CustomForecastEntry{
					{
						Amount: 400,
						Month:  202501,
						TagFilters: []datadogV2.CustomForecastEntryTagFilter{
							{
								TagKey:   "service",
								TagValue: "ec2",
							},
						},
					},
				},
			},
			Id:   datadog.PtrString(""),
			Type: datadogV2.CUSTOMFORECASTTYPE_CUSTOM_FORECAST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpsertCustomForecast(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpsertCustomForecast`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpsertCustomForecast`:\n%s\n", responseContent)
}
