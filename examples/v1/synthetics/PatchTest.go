// Patch a Synthetic test returns "OK" response

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

	body := datadogV1.SyntheticsPatchTestBody{
		Data: []datadogV1.SyntheticsPatchTestOperation{
			{
				Op:    datadogV1.SYNTHETICSPATCHTESTOPERATIONNAME_REPLACE.Ptr(),
				Path:  datadog.PtrString("/name"),
				Value: "New test name",
			},
			{
				Op:   datadogV1.SYNTHETICSPATCHTESTOPERATIONNAME_REMOVE.Ptr(),
				Path: datadog.PtrString("/config/assertions/0"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.PatchTest(ctx, SyntheticsAPITestPublicID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.PatchTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.PatchTest`:\n%s\n", responseContent)
}
