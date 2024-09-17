// Update a tag configuration returns "OK" response

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
	// there is a valid "metric_tag_configuration" in the system
	MetricTagConfigurationDataID := os.Getenv("METRIC_TAG_CONFIGURATION_DATA_ID")

	body := datadogV2.MetricTagConfigurationUpdateRequest{
		Data: datadogV2.MetricTagConfigurationUpdateData{
			Type: datadogV2.METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS,
			Id:   MetricTagConfigurationDataID,
			Attributes: &datadogV2.MetricTagConfigurationUpdateAttributes{
				Tags: []string{
					"app",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateTagConfiguration(ctx, MetricTagConfigurationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateTagConfiguration`:\n%s\n", responseContent)
}
