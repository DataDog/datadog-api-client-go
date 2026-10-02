// Get SPA recommendations v2 returns "OK" response

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
	body := datadogV2.RecommendationV2RequestBody{
		Data: datadogV2.RecommendationV2RequestData{
			Attributes: datadogV2.RecommendationV2RequestAttributes{
				Arguments: []string{
					"",
				},
			},
			Type: datadogV2.RECOMMENDATIONV2REQUESTTYPE_RECOMMENDATION_V2_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetSPARecommendationsV2", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpaApi(apiClient)
	resp, r, err := api.GetSPARecommendationsV2(ctx, "service", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaApi.GetSPARecommendationsV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpaApi.GetSPARecommendationsV2`:\n%s\n", responseContent)
}
