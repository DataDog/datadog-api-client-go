// Delete a pipeline returns "OK" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "pipeline" in the system
	PipelineDataID := os.Getenv("PIPELINE_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeletePipeline", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewObservabilityPipelinesApi(apiClient)
	r, err := api.DeletePipeline(ctx, PipelineDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityPipelinesApi.DeletePipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}