// Get SPA Recommendations with a shard parameter returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetSPARecommendationsWithShard", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpaApi(apiClient)
	resp, r, err := api.GetSPARecommendationsWithShard(ctx, "shard", "service", *datadogV2.NewGetSPARecommendationsWithShardOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaApi.GetSPARecommendationsWithShard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpaApi.GetSPARecommendationsWithShard`:\n%s\n", responseContent)
}
