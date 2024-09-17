// Delete a pipeline returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)
	r, err := api.DeleteLogsPipeline(ctx, "pipeline_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.DeleteLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
