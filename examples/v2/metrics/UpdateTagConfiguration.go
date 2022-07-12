// Update a tag configuration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "metric_tag_configuration" in the system
	MetricTagConfigurationDataID := os.Getenv("METRIC_TAG_CONFIGURATION_DATA_ID")

	body := datadog.MetricTagConfigurationUpdateRequest{
		Data: datadog.MetricTagConfigurationUpdateData{
			Type: datadog.METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS,
			Id:   MetricTagConfigurationDataID,
			Attributes: &datadog.MetricTagConfigurationUpdateAttributes{
				Tags: []string{
					"app",
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateTagConfiguration(ctx, MetricTagConfigurationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateTagConfiguration`:\n%s\n", responseContent)
}
