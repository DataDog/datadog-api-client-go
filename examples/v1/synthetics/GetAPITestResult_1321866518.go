// Get an API test result returns result with failure object

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
	// there is a "synthetics_api_test_with_wrong_dns" in the system
	SyntheticsAPITestWithWrongDNSPublicID := os.Getenv("SYNTHETICS_API_TEST_WITH_WRONG_DNS_PUBLIC_ID")

	// the "synthetics_api_test_with_wrong_dns" is triggered
	SyntheticsAPITestWithWrongDNSResultResults0ResultID := os.Getenv("SYNTHETICS_API_TEST_WITH_WRONG_DNS_RESULT_RESULTS_0_RESULT_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.GetAPITestResult(ctx, SyntheticsAPITestWithWrongDNSPublicID, SyntheticsAPITestWithWrongDNSResultResults0ResultID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestResult`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetAPITestResult`:\n%s\n", responseContent)
}
