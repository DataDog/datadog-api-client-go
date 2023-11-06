// Create a new rule returns "Created" response

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
	body := datadogV2.CreateRuleRequest{
		Data: &datadogV2.CreateRuleRequestData{
			Attributes: &datadogV2.RuleAttributes{
				Enabled:       datadog.PtrBool(true),
				Name:          datadog.PtrString("Example-Service-Scorecard"),
				ScorecardName: datadog.PtrString("Observability Best Practices"),
			},
			Type: datadogV2.RULETYPE_RULE.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateScorecardRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	resp, r, err := api.CreateScorecardRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.CreateScorecardRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceScorecardsApi.CreateScorecardRule`:\n%s\n", responseContent)
}
