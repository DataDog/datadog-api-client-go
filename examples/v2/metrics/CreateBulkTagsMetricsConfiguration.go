// Configure tags for multiple metrics returns "Accepted" response

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
	// there is a valid "user" in the system
	UserDataAttributesEmail := os.Getenv("USER_DATA_ATTRIBUTES_EMAIL")

	body := datadogV2.MetricBulkTagConfigCreateRequest{
		Data: datadogV2.MetricBulkTagConfigCreate{
			Attributes: &datadogV2.MetricBulkTagConfigCreateAttributes{
				Emails: []string{
					UserDataAttributesEmail,
				},
				Tags: []string{
					"test",
					"examplemetric",
				},
			},
			Id:   "system.load.1",
			Type: datadogV2.METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateBulkTagsMetricsConfiguration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateBulkTagsMetricsConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateBulkTagsMetricsConfiguration`:\n%s\n", responseContent)
}
