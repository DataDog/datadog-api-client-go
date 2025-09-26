// Create ruleset returns "OK" response

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
	body := datadogV2.CreateRulesetRequest{
		Data: &datadogV2.CreateRulesetRequestData{
			Attributes: &datadogV2.CreateRulesetRequestDataAttributes{
				Enabled: datadog.PtrBool(true),
				Rules: []datadogV2.CreateRulesetRequestDataAttributesRulesItems{
					{
						Enabled: true,
						Mapping: *datadogV2.NewNullableCreateRulesetRequestDataAttributesRulesItemsMapping(nil),
						Name:    "Add Cost Center Tag",
						Query: *datadogV2.NewNullableCreateRulesetRequestDataAttributesRulesItemsQuery(&datadogV2.CreateRulesetRequestDataAttributesRulesItemsQuery{
							Addition: *datadogV2.NewNullableCreateRulesetRequestDataAttributesRulesItemsQueryAddition(&datadogV2.CreateRulesetRequestDataAttributesRulesItemsQueryAddition{
								Key:   "cost_center",
								Value: "engineering",
							}),
							CaseInsensitivity: datadog.PtrBool(false),
							IfNotExists:       true,
							Query:             `account_id:"123456789" AND service:"web-api"`,
						}),
						ReferenceTable: *datadogV2.NewNullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable(nil),
					},
				},
			},
			Id:   datadog.PtrString("New Ruleset"),
			Type: datadogV2.CREATERULESETREQUESTDATATYPE_CREATE_RULESET,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateRuleset(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateRuleset`:\n%s\n", responseContent)
}
