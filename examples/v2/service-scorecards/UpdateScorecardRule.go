// Update an existing rule returns "Rule updated successfully" response

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
	// there is a valid "create_scorecard_rule" in the system
	CreateScorecardRuleDataAttributesName := os.Getenv("CREATE_SCORECARD_RULE_DATA_ATTRIBUTES_NAME")
	CreateScorecardRuleDataAttributesScorecardName := os.Getenv("CREATE_SCORECARD_RULE_DATA_ATTRIBUTES_SCORECARD_NAME")
	CreateScorecardRuleDataID := os.Getenv("CREATE_SCORECARD_RULE_DATA_ID")

	body := datadogV2.UpdateRuleRequest{
		Data: &datadogV2.UpdateRuleRequestData{
			Attributes: &datadogV2.RuleAttributes{
				Enabled:       datadog.PtrBool(true),
				Name:          datadog.PtrString(CreateScorecardRuleDataAttributesName),
				ScorecardName: datadog.PtrString(CreateScorecardRuleDataAttributesScorecardName),
				Description:   datadog.PtrString("Updated description via test"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateScorecardRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	resp, r, err := api.UpdateScorecardRule(ctx, CreateScorecardRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.UpdateScorecardRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceScorecardsApi.UpdateScorecardRule`:\n%s\n", responseContent)
}
