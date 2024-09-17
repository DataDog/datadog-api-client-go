// Delete tags for multiple metrics returns "Accepted" response

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
	body := datadogV2.MetricBulkTagConfigDeleteRequest{
		Data: datadogV2.MetricBulkTagConfigDelete{
			Attributes: &datadogV2.MetricBulkTagConfigDeleteAttributes{
				Emails: []string{
					"sue@example.com",
					"bob@example.com",
				},
			},
			Id:   "kafka.lag",
			Type: datadogV2.METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.DeleteBulkTagsMetricsConfiguration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteBulkTagsMetricsConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.DeleteBulkTagsMetricsConfiguration`:\n%s\n", responseContent)
}
