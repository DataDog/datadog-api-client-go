// Update a hardcoded retention filter returns "Updated" response

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
	body := datadogV2.RumHardcodedRetentionFilterUpdateRequest{
		Data: datadogV2.RumHardcodedRetentionFilterUpdateData{
			Id:   "REPLACE.ME",
			Type: datadogV2.RUMHARDCODEDRETENTIONFILTERTYPE_HARDCODED_RETENTION_FILTERS,
			Attributes: datadogV2.RumHardcodedRetentionFilterUpdateAttributes{
				CrossProductSampling: &datadogV2.RumHardcodedCrossProductSamplingUpdate{
					SessionReplaySampleRate: datadog.PtrFloat64(50.0),
					SessionReplayEnabled:    datadog.PtrBool(true),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionFiltersHardcodedApi(apiClient)
	resp, r, err := api.UpdateHardcodedRetentionFilter(ctx, "Example-RUM-Retention-Filters-Hardcoded", "Example-RUM-Retention-Filters-Hardcoded", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionFiltersHardcodedApi.UpdateHardcodedRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRetentionFiltersHardcodedApi.UpdateHardcodedRetentionFilter`:\n%s\n", responseContent)
}
