// Update tag pipeline ruleset with if_tag_exists returns "OK" response

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
				LastVersion: datadog.PtrInt64(3611102),
				Rules: []datadogV2.UpdateRulesetRequestDataAttributesRulesItems{
					{
						Enabled: true,
						Mapping: *datadogV2.NewNullableDataAttributesRulesItemsMapping(&datadogV2.DataAttributesRulesItemsMapping{
							DestinationKey: "team_owner",
							IfTagExists:    datadogV2.DATAATTRIBUTESRULESITEMSIFTAGEXISTS_REPLACE.Ptr(),
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
			Id:   datadog.PtrString("New Ruleset"),
			Type: datadogV2.UPDATERULESETREQUESTDATATYPE_UPDATE_RULESET,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateTagPipelinesRuleset(ctx, "ee10c3ff-312f-464c-b4f6-46adaa6d00a1", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateTagPipelinesRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateTagPipelinesRuleset`:\n%s\n", responseContent)
}
