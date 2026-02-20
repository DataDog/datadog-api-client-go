// Edit a test suite returns "OK" response

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
	body := datadogV2.SuiteCreateEditRequest{
		Data: datadogV2.SuiteCreateEdit{
			Attributes: datadogV2.SyntheticsSuite{
				Message: datadog.PtrString("Notification message"),
				Name:    "Example suite name",
				Options: datadogV2.SyntheticsSuiteOptions{},
				Tags: []string{
					"env:production",
				},
				Tests: []datadogV2.SyntheticsSuiteTest{
					{
						AlertingCriticality: datadogV2.SYNTHETICSSUITETESTALERTINGCRITICALITY_CRITICAL.Ptr(),
						PublicId:            "",
					},
				},
				Type: datadogV2.SYNTHETICSSUITETYPE_SUITE,
			},
			Type: datadogV2.SYNTHETICSSUITETYPES_SUITES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.EditSyntheticsSuite(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditSyntheticsSuite`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.EditSyntheticsSuite`:\n%s\n", responseContent)
}
