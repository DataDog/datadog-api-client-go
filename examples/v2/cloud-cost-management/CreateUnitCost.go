// Create a unit cost returns "Created" response

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
	body := datadogV2.UnitCostCreateRequest{
		Data: datadogV2.UnitCostCreateRequestData{
			Attributes: datadogV2.UnitCostRequestAttributes{
				DenominatorQuery: datadogV2.UnitCostQueryDefinition{
					Formulas: []map[string]interface{}{
						map[string]interface{}{
							"formula": "numerator",
						},
					},
					Queries: []map[string]interface{}{
						map[string]interface{}{
							"data_source": "cloud_cost",
							"name":        "numerator",
							"query":       "sum:aws.cost.net.amortized.shared.resources.allocated{*}.rollup(sum, daily)",
						},
					},
				},
				Description: *datadog.NewNullableString(datadog.PtrString("Amortized cloud spend divided by the number of active users.")),
				Name:        "Cloud cost per active user",
				NumeratorQuery: datadogV2.UnitCostQueryDefinition{
					Formulas: []map[string]interface{}{
						map[string]interface{}{
							"formula": "numerator",
						},
					},
					Queries: []map[string]interface{}{
						map[string]interface{}{
							"data_source": "cloud_cost",
							"name":        "numerator",
							"query":       "sum:aws.cost.net.amortized.shared.resources.allocated{*}.rollup(sum, daily)",
						},
					},
				},
				UnitLabel: "user",
			},
			Type: datadogV2.UNITCOSTTYPE_UNIT_COST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateUnitCost", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateUnitCost(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateUnitCost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateUnitCost`:\n%s\n", responseContent)
}
