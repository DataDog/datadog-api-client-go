// List all scores returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewScorecardsApi(apiClient)
	resp, r, err := api.ListScorecardScores(ctx, datadogV2.SCORECARDSCORESAGGREGATION_BY_ENTITY, *datadogV2.NewListScorecardScoresOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.ListScorecardScores`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ScorecardsApi.ListScorecardScores`:\n%s\n", responseContent)
}
