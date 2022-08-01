// Create a tag configuration returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	body := datadog.MetricTagConfigurationCreateRequest{
		Data: datadog.MetricTagConfigurationCreateData{
			Type: datadog.METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS,
			Id:   "ExampleCreateatagconfigurationreturnsCreatedresponse",
			Attributes: &datadog.MetricTagConfigurationCreateAttributes{
				Tags: []string{
					"app",
					"datacenter",
				},
				MetricType: datadog.METRICTAGCONFIGURATIONMETRICTYPES_GAUGE,
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMetricsApi(apiClient)
	resp, r, err := api.CreateTagConfiguration(ctx, "ExampleCreateatagconfigurationreturnsCreatedresponse", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateTagConfiguration`:\n%s\n", responseContent)
}
