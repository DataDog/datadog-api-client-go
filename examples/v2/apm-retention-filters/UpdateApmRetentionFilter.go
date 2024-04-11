// Update a retention filter returns "OK" response

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
	// there is a valid "retention_filter" in the system
	RetentionFilterDataID := os.Getenv("RETENTION_FILTER_DATA_ID")

	body := datadogV2.RetentionFilterUpdateRequest{
		Data: datadogV2.RetentionFilterUpdateData{
			Attributes: datadogV2.RetentionFilterUpdateAttributes{
				Name: "test",
				Rate: 0.9,
				Filter: datadogV2.SpansFilterCreate{
					Query: "@_top_level:1 test:service-demo",
				},
				Enabled:    true,
				FilterType: datadogV2.RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR,
			},
			Id:   "test-id",
			Type: datadogV2.APMRETENTIONFILTERTYPE_apm_retention_filter,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMRetentionFiltersApi(apiClient)
	resp, r, err := api.UpdateApmRetentionFilter(ctx, RetentionFilterDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMRetentionFiltersApi.UpdateApmRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APMRetentionFiltersApi.UpdateApmRetentionFilter`:\n%s\n", responseContent)
}
