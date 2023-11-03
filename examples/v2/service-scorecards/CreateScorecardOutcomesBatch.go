// Create outcomes batch returns "OK" response

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
	CreateScorecardRuleDataID := os.Getenv("CREATE_SCORECARD_RULE_DATA_ID")

	body := datadogV2.OutcomesBatchRequest{
		Data: &datadogV2.OutcomesBatchRequestData{
			Attributes: &datadogV2.OutcomesBatchAttributes{
				Results: []datadogV2.OutcomesBatchRequestItem{
					{
						Remarks:     datadog.PtrString(`See: <a href="https://app.datadoghq.com/services">Services</a>`),
						RuleId:      CreateScorecardRuleDataID,
						ServiceName: "my-service",
						State:       datadogV2.STATE_PASS,
					},
				},
			},
			Type: datadogV2.OUTCOMESBATCHTYPE_BATCHED_OUTCOME.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateScorecardOutcomesBatch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	resp, r, err := api.CreateScorecardOutcomesBatch(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.CreateScorecardOutcomesBatch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceScorecardsApi.CreateScorecardOutcomesBatch`:\n%s\n", responseContent)
}
