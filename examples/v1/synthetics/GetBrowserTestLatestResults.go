// Get a browser test's latest results summaries returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.GetBrowserTestLatestResults(ctx, "2yy-sem-mjh", *datadog.NewGetBrowserTestLatestResultsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestLatestResults`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetBrowserTestLatestResults`:\n%s\n", responseContent)
}
