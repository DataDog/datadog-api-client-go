// Update Scorecard outcomes asynchronously returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "create_scorecard_rule" in the system
	CreateScorecardRuleDataID := os.Getenv("CREATE_SCORECARD_RULE_DATA_ID")

	body := datadogV2.UpdateOutcomesAsyncRequest{
		Data: &datadogV2.UpdateOutcomesAsyncRequestData{
			Attributes: &datadogV2.UpdateOutcomesAsyncAttributes{
				Results: []datadogV2.UpdateOutcomesAsyncRequestItem{
					{
						RuleId:          CreateScorecardRuleDataID,
						EntityReference: "service:my-service",
						Remarks:         datadog.PtrString(`See: <a href="https://app.datadoghq.com/services">Services</a>`),
						State:           datadogV2.STATE_PASS,
					},
				},
			},
			Type: datadogV2.UPDATEOUTCOMESASYNCTYPE_BATCHED_OUTCOME.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewScorecardsApi(apiClient)
	r, err := api.UpdateScorecardOutcomes(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.UpdateScorecardOutcomes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
