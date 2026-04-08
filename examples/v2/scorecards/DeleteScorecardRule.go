// Delete a rule returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewScorecardsApi(apiClient)
	r, err := api.DeleteScorecardRule(ctx, CreateScorecardRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.DeleteScorecardRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
