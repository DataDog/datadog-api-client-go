// Get a campaign returns "OK" response

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
	resp, r, err := api.GetScorecardCampaign(ctx, "c10ODp0VCrrIpXmz", *datadogV2.NewGetScorecardCampaignOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.GetScorecardCampaign`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ScorecardsApi.GetScorecardCampaign`:\n%s\n", responseContent)
}
