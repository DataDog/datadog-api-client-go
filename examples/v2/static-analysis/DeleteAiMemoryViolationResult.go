// Delete an AI memory violation result returns "Successfully deleted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteAiMemoryViolationResult", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.DeleteAiMemoryViolationResult(ctx, "42")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.DeleteAiMemoryViolationResult`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
