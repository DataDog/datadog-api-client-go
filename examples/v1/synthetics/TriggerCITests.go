// Trigger tests from CI/CD pipelines returns "OK" response

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
	body := datadogV1.SyntheticsCITestBody{
		Tests: []datadogV1.SyntheticsCITest{
			{
				BasicAuth: &datadogV1.SyntheticsBasicAuth{
					SyntheticsBasicAuthWeb: &datadogV1.SyntheticsBasicAuthWeb{
						Password: "PaSSw0RD!",
						Type:     datadogV1.SYNTHETICSBASICAUTHWEBTYPE_WEB.Ptr(),
						Username: "my_username",
					}},
				DeviceIds: []datadogV1.SyntheticsDeviceID{
					datadogV1.SYNTHETICSDEVICEID_CHROME_LAPTOP_LARGE,
				},
				Locations: []string{
					"aws:eu-west-3",
				},
				Metadata: &datadogV1.SyntheticsCIBatchMetadata{
					Ci: &datadogV1.SyntheticsCIBatchMetadataCI{
						Pipeline: &datadogV1.SyntheticsCIBatchMetadataPipeline{},
						Provider: &datadogV1.SyntheticsCIBatchMetadataProvider{},
					},
					Git: &datadogV1.SyntheticsCIBatchMetadataGit{},
				},
				PublicId: "aaa-aaa-aaa",
				Retry:    &datadogV1.SyntheticsTestOptionsRetry{},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.TriggerCITests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerCITests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.TriggerCITests`:\n%s\n", responseContent)
}
