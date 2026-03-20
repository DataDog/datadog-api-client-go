// Create an environment returns "Created" response

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
	body := datadogV2.CreateEnvironmentRequest{
		Data: datadogV2.CreateEnvironmentData{
			Type: datadogV2.CREATEENVIRONMENTDATATYPE_ENVIRONMENTS,
			Attributes: datadogV2.CreateEnvironmentAttributes{
				Name: "Test Environment Example-Feature-Flag",
				Queries: []string{
					"test-Example-Feature-Flag",
					"env-Example-Feature-Flag",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.CreateFeatureFlagsEnvironment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CreateFeatureFlagsEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CreateFeatureFlagsEnvironment`:\n%s\n", responseContent)
}
