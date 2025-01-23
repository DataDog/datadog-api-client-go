// Patch a mute rule returns "Mute rule successfully patched" response

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
	// there is a valid "valid_mute_rule" in the system
	ValidMuteRuleDataID := uuid.MustParse(os.Getenv("VALID_MUTE_RULE_DATA_ID"))

	body := datadogV2.PatchMuteRuleParameters{
		Data: &datadogV2.PatchMuteRuleParametersData{
			Attributes: datadogV2.PatchMuteRuleParametersDataAttributes{
				Action: &datadogV2.ActionMute{
					ExpireAt:          datadog.PtrInt64(1893452400000),
					Reason:            datadogV2.MUTEREASON_DUPLICATE,
					ReasonDescription: datadog.PtrString("Muting for a while"),
				},
				Enabled: datadog.PtrBool(true),
				Name:    datadog.PtrString("Rule 1"),
				Rule: &datadogV2.Rule{
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
			Id:   ValidMuteRuleDataID,
			Type: datadogV2.MUTERULESTYPE_MUTE_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.PatchMuteRule(ctx, ValidMuteRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.PatchMuteRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.PatchMuteRule`:\n%s\n", responseContent)
}
