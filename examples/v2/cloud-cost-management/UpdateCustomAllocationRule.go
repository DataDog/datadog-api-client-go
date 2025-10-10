// Update custom allocation rule returns "OK" response

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
	body := datadogV2.ArbitraryCostUpsertRequest{
		Data: &datadogV2.ArbitraryCostUpsertRequestData{
			Attributes: &datadogV2.ArbitraryCostUpsertRequestDataAttributes{
				CostsToAllocate: []datadogV2.ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems{
					{
						Condition: "is",
						Tag:       "account_id",
						Value:     datadog.PtrString("123456789"),
						Values:    *datadog.NewNullableList(&[]string{}),
					},
					{
						Condition: "in",
						Tag:       "environment",
						Value:     datadog.PtrString(""),
						Values: *datadog.NewNullableList(&[]string{
							"production",
							"staging",
						}),
					},
				},
				Enabled: datadog.PtrBool(true),
				OrderId: datadog.PtrInt64(1),
				Provider: []string{
					"aws",
					"gcp",
				},
				RuleName: "example-arbitrary-cost-rule",
				Strategy: datadogV2.ArbitraryCostUpsertRequestDataAttributesStrategy{
					AllocatedByTagKeys: []string{
						"team",
						"environment",
					},
					BasedOnCosts: []datadogV2.ArbitraryCostUpsertRequestDataAttributesStrategyBasedOnCostsItems{
						{
							Condition: "is",
							Tag:       "service",
							Value:     datadog.PtrString("web-api"),
							Values:    *datadog.NewNullableList(&[]string{}),
						},
						{
							Condition: "not in",
							Tag:       "team",
							Value:     datadog.PtrString(""),
							Values: *datadog.NewNullableList(&[]string{
								"legacy",
								"deprecated",
							}),
						},
					},
					Granularity: datadog.PtrString("daily"),
					Method:      "proportional",
				},
				Type: "shared",
			},
			Type: datadogV2.ARBITRARYCOSTUPSERTREQUESTDATATYPE_UPSERT_ARBITRARY_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCustomAllocationRule(ctx, 683, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCustomAllocationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCustomAllocationRule`:\n%s\n", responseContent)
}
