// Export an LLM Observability dataset returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ExportLLMObsDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.ExportLLMObsDataset(ctx, "project_id", "dataset_id", *datadogV2.NewExportLLMObsDatasetOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.ExportLLMObsDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.ExportLLMObsDataset`:\n%s\n", resp)
}
