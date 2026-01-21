// List Entity Risk Scores returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListEntityRiskScores", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEntityRiskScoresApi(apiClient)
	resp, r, err := api.ListEntityRiskScores(ctx, *datadogV2.NewListEntityRiskScoresOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityRiskScoresApi.ListEntityRiskScores`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EntityRiskScoresApi.ListEntityRiskScores`:\n%s\n", responseContent)
}
