// Delete a campaign returns "No Content" response

package main

import (
	"context"
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
	r, err := api.DeleteScorecardCampaign(ctx, "c10ODp0VCrrIpXmz")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.DeleteScorecardCampaign`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
