// Create a retention filter returns "OK" response

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
	body := datadogV2.RetentionFilterCreateRequest{
		Data: datadogV2.RetentionFilterCreateData{
			Attributes: datadogV2.RetentionFilterCreateAttributes{
				Enabled: true,
				Filter: datadogV2.SpansFilterCreate{
					Query: "@http.status_code:200 service:my-service",
				},
				FilterType: datadogV2.RETENTIONFILTERTYPE_SPANS_SAMPLING_PROCESSOR,
				Name:       "my retention filter",
				Rate:       1.0,
			},
			Type: datadogV2.APMRETENTIONFILTERTYPE_apm_retention_filter,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMRetentionFiltersApi(apiClient)
	resp, r, err := api.CreateApmRetentionFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMRetentionFiltersApi.CreateApmRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APMRetentionFiltersApi.CreateApmRetentionFilter`:\n%s\n", responseContent)
}
