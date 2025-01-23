// Create a new mute rule returns "Successfully created the mute rule" response

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
	body := datadogV2.CreateMuteRuleParameters{
		Data: &datadogV2.CreateMuteRuleParametersData{
			Attributes: datadogV2.CreateMuteRuleParametersDataAttributes{
				Action: datadogV2.ActionMute{
					ExpireAt:          datadog.PtrInt64(1893452400000),
					Reason:            datadogV2.MUTEREASON_DUPLICATE,
					ReasonDescription: datadog.PtrString("Muting for a while"),
				},
				Enabled: datadog.PtrBool(true),
				Name:    "Rule 1",
				Rule: datadogV2.Rule{
					IssueType: datadogV2.ISSUETYPE_VULNERABILITY,
					Query:     datadog.PtrString("key:val"),
					RuleIds: []string{
						"rule-id-1",
					},
					RuleTypes: []datadogV2.RuleTypesItems{
						datadogV2.RULETYPESITEMS_APPLICATION_CODE_VULNERABILITY,
					},
					Severities: []datadogV2.RuleSeverity{
						datadogV2.RULESEVERITY_CRITICAL,
					},
				},
			},
			Type: datadogV2.MUTERULESTYPE_MUTE_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateMuteRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateMuteRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateMuteRule`:\n%s\n", responseContent)
}
