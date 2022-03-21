// List tags by metric name returns "Success" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "metric_tag_configuration" in the system
	MetricTagConfigurationDataID := os.Getenv("METRIC_TAG_CONFIGURATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.ListTagsByMetricName(ctx, MetricTagConfigurationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagsByMetricName`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListTagsByMetricName`:\n%s\n", responseContent)
}
