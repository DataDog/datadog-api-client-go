// Add a test to a Synthetics downtime returns "OK" response

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
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.AddTestToSyntheticsDowntime(ctx, "00000000-0000-0000-0000-000000000001", "abc-def-123")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.AddTestToSyntheticsDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.AddTestToSyntheticsDowntime`:\n%s\n", responseContent)
}
