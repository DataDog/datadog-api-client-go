// Update an environment returns "OK" response

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
	// there is a valid "environment" in the system
	EnvironmentDataID := uuid.MustParse(os.Getenv("ENVIRONMENT_DATA_ID"))

	body := datadogV2.UpdateEnvironmentRequest{
		Data: datadogV2.UpdateEnvironmentData{
			Type: datadogV2.UPDATEENVIRONMENTDATATYPE_ENVIRONMENTS,
			Attributes: datadogV2.UpdateEnvironmentAttributes{
				Name: datadog.PtrString("Updated Test Environment Example-Feature-Flag"),
				Queries: []string{
					"updated-Example-Feature-Flag",
					"live-Example-Feature-Flag",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateFeatureFlagsEnvironment(ctx, EnvironmentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateFeatureFlagsEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateFeatureFlagsEnvironment`:\n%s\n", responseContent)
}
