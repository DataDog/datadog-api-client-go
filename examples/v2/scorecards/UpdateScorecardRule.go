// Update an existing scorecard rule returns "Rule updated successfully" response

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
	body := datadogV2.UpdateRuleRequest{
		Data: &datadogV2.UpdateRuleRequestData{
			Attributes: &datadogV2.RuleAttributesRequest{
				Enabled:       datadog.PtrBool(true),
				Level:         datadog.PtrInt32(2),
				Name:          datadog.PtrString("Team Defined"),
				ScopeQuery:    datadog.PtrString("kind:service"),
				ScorecardName: datadog.PtrString("Deployments automated via Deployment Trains"),
			},
			Type: datadogV2.RULETYPE_RULE.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewScorecardsApi(apiClient)
	resp, r, err := api.UpdateScorecardRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.UpdateScorecardRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ScorecardsApi.UpdateScorecardRule`:\n%s\n", responseContent)
}
