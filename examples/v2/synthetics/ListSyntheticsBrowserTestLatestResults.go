// Get a browser test's latest results returns "OK" response

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
	resp, r, err := api.ListSyntheticsBrowserTestLatestResults(ctx, "public_id", *datadogV2.NewListSyntheticsBrowserTestLatestResultsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListSyntheticsBrowserTestLatestResults`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.ListSyntheticsBrowserTestLatestResults`:\n%s\n", responseContent)
}
