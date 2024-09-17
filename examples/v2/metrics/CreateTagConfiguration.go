// Create a tag configuration returns "Created" response

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
	body := datadogV2.MetricTagConfigurationCreateRequest{
		Data: datadogV2.MetricTagConfigurationCreateData{
			Type: datadogV2.METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS,
			Id:   "ExampleMetric",
			Attributes: &datadogV2.MetricTagConfigurationCreateAttributes{
				Tags: []string{
					"app",
					"datacenter",
				},
				MetricType: datadogV2.METRICTAGCONFIGURATIONMETRICTYPES_GAUGE,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateTagConfiguration(ctx, "ExampleMetric", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateTagConfiguration`:\n%s\n", responseContent)
}
