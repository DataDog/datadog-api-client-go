// Configure tags for multiple metrics returns "Accepted" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "user" in the system
	UserDataAttributesEmail := os.Getenv("USER_DATA_ATTRIBUTES_EMAIL")

	body := datadog.MetricBulkTagConfigCreateRequest{
		Data: datadog.MetricBulkTagConfigCreate{
			Attributes: &datadog.MetricBulkTagConfigCreateAttributes{
				Emails: []string{
					UserDataAttributesEmail,
				},
				Tags: []string{
					"test",
					"exampleconfiguretagsformultiplemetricsreturnsacceptedresponse",
				},
			},
			Id:   "system.load.1",
			Type: datadog.METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.CreateBulkTagsMetricsConfiguration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateBulkTagsMetricsConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateBulkTagsMetricsConfiguration`:\n%s\n", responseContent)
}
