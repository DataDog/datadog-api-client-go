// Get an API test's latest results summaries returns "OK" response

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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SyntheticsApi(apiClient)
	resp, r, err := api.GetAPITestLatestResults(ctx, "public_id", *datadog.NewGetAPITestLatestResultsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestLatestResults`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetAPITestLatestResults`:\n%s\n", responseContent)
}
