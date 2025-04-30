// Create or update a budget returns "OK" response

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
	body := datadogV2.BudgetWithEntries{
		Data: &datadogV2.BudgetWithEntriesData{
			Attributes: &datadogV2.BudgetAttributes{
				CreatedAt: datadog.PtrInt64(1738258683590),
				CreatedBy: datadog.PtrString("00000000-0a0a-0a0a-aaa0-00000000000a"),
				EndMonth:  datadog.PtrInt64(202502),
				Entries: []datadogV2.BudgetEntry{
					{
						Amount: datadog.PtrFloat64(500),
						Month:  datadog.PtrInt64(202501),
						TagFilters: []datadogV2.TagFilter{
							{
								TagKey:   datadog.PtrString("service"),
								TagValue: datadog.PtrString("ec2"),
							},
						},
					},
				},
				MetricsQuery: datadog.PtrString("aws.cost.amortized{service:ec2} by {service}"),
				Name:         datadog.PtrString("my budget"),
				OrgId:        datadog.PtrInt64(123),
				StartMonth:   datadog.PtrInt64(202501),
				TotalAmount:  datadog.PtrFloat64(1000),
				UpdatedAt:    datadog.PtrInt64(1738258683590),
				UpdatedBy:    datadog.PtrString("00000000-0a0a-0a0a-aaa0-00000000000a"),
			},
			Id: datadog.PtrString("00000000-0a0a-0a0a-aaa0-00000000000a"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpsertBudget(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpsertBudget`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpsertBudget`:\n%s\n", responseContent)
}
