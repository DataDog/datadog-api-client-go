// Update ruleset returns "OK" response

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
	body := datadogV2.UpdateRulesetRequest{
		Data: &datadogV2.UpdateRulesetRequestData{
			Attributes: &datadogV2.UpdateRulesetRequestDataAttributes{
				Enabled:     true,
				LastVersion: datadog.PtrInt64(3601919),
				Rules: []datadogV2.UpdateRulesetRequestDataAttributesRulesItems{
					{
						Enabled: true,
						Mapping: *datadogV2.NewNullableUpdateRulesetRequestDataAttributesRulesItemsMapping(&datadogV2.UpdateRulesetRequestDataAttributesRulesItemsMapping{
							DestinationKey: "team_owner",
							IfNotExists:    true,
							SourceKeys: []string{
								"account_name",
								"account_id",
							},
						}),
						Name:           "Account Name Mapping",
						Query:          *datadogV2.NewNullableUpdateRulesetRequestDataAttributesRulesItemsQuery(nil),
						ReferenceTable: *datadogV2.NewNullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable(nil),
					},
				},
			},
			Type: datadogV2.UPDATERULESETREQUESTDATATYPE_UPDATE_RULESET,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateRuleset(ctx, "1c5dae14-237d-4b9a-a515-aa55b3939142", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateRuleset`:\n%s\n", responseContent)
}
