// Trigger tests from CI/CD pipelines returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SyntheticsCITestBody{
		Tests: []datadog.SyntheticsCITest{
			{
				BasicAuth: &datadog.SyntheticsBasicAuth{
					SyntheticsBasicAuthWeb: &datadog.SyntheticsBasicAuthWeb{
						Password: "PaSSw0RD!",
						Type:     datadog.SYNTHETICSBASICAUTHWEBTYPE_WEB.Ptr(),
						Username: "my_username",
					}},
				DeviceIds: []datadog.SyntheticsDeviceID{
					datadog.SYNTHETICSDEVICEID_LAPTOP_LARGE,
				},
				Locations: []string{
					"aws:eu-west-3",
				},
				Metadata: &datadog.SyntheticsCIBatchMetadata{
					Ci: &datadog.SyntheticsCIBatchMetadataCI{
						Pipeline: &datadog.SyntheticsCIBatchMetadataPipeline{},
						Provider: &datadog.SyntheticsCIBatchMetadataProvider{},
					},
					Git: &datadog.SyntheticsCIBatchMetadataGit{},
				},
				PublicId: "aaa-aaa-aaa",
				Retry:    &datadog.SyntheticsTestOptionsRetry{},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.TriggerCITests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerCITests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.TriggerCITests`:\n%s\n", responseContent)
}
