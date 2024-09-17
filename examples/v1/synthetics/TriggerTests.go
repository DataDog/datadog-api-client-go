// Trigger Synthetic tests returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestPublicID := os.Getenv("SYNTHETICS_API_TEST_PUBLIC_ID")

	body := datadogV1.SyntheticsTriggerBody{
		Tests: []datadogV1.SyntheticsTriggerTest{
			{
				PublicId: SyntheticsAPITestPublicID,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.TriggerTests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.TriggerTests`:\n%s\n", responseContent)
}
