// Delete a tag configuration returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "metric_tag_configuration" in the system
	MetricTagConfigurationDataID := os.Getenv("METRIC_TAG_CONFIGURATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("DeleteTagConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.MetricsApi.DeleteTagConfiguration(ctx, MetricTagConfigurationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
