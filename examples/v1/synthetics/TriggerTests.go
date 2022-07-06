// Trigger Synthetics tests returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestPublicID := os.Getenv("SYNTHETICS_API_TEST_PUBLIC_ID")

	body := datadog.SyntheticsTriggerBody{
		Tests: []datadog.SyntheticsTriggerTest{
			{
				PublicId: SyntheticsAPITestPublicID,
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SyntheticsApi(apiClient)
	resp, r, err := api.TriggerTests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.TriggerTests`:\n%s\n", responseContent)
}
