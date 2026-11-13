// Update an environment to preserve enabled full evaluation data when omitted returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "environment_with_full_evaluation_data" in the system
	EnvironmentWithFullEvaluationDataDataID := uuid.MustParse(os.Getenv("ENVIRONMENT_WITH_FULL_EVALUATION_DATA_DATA_ID"))

	body := datadogV2.UpdateEnvironmentRequest{
		Data: datadogV2.UpdateEnvironmentData{
			Type: datadogV2.UPDATEENVIRONMENTDATATYPE_ENVIRONMENTS,
			Attributes: datadogV2.UpdateEnvironmentAttributes{
				Name: datadog.PtrString("Updated Privacy Example-Feature-Flag"),
				Queries: []string{
					"updated-privacy-example-feature-flag",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateFeatureFlagsEnvironment(ctx, EnvironmentWithFullEvaluationDataDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateFeatureFlagsEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateFeatureFlagsEnvironment`:\n%s\n", responseContent)
}
