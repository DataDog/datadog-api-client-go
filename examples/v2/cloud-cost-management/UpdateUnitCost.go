// Update a unit cost returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.UnitCostUpdateRequest{
		Data: datadogV2.UnitCostUpdateRequestData{
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
			Id:   uuid.MustParse("64aecd58-e355-4f07-9c3a-56ff6bda6cd8"),
			Type: datadogV2.UNITCOSTTYPE_UNIT_COST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateUnitCost", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateUnitCost(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateUnitCost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateUnitCost`:\n%s\n", responseContent)
}
