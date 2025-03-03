// Update a RUM retention filter returns "Updated" response

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
	// there is a valid "rum_retention_filter" in the system
	RumRetentionFilterDataID := os.Getenv("RUM_RETENTION_FILTER_DATA_ID")

	body := datadogV2.RumRetentionFilterUpdateRequest{
		Data: datadogV2.RumRetentionFilterUpdateData{
			Id:   RumRetentionFilterDataID,
			Type: datadogV2.RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
			Attributes: datadogV2.RumRetentionFilterUpdateAttributes{
				Name:       datadog.PtrString("Test updating retention filter"),
				EventType:  datadogV2.RUMRETENTIONFILTEREVENTTYPE_VIEW.Ptr(),
				Query:      datadog.PtrString("view_query"),
				SampleRate: datadog.PtrInt64(100),
				Enabled:    datadog.PtrBool(true),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.UpdateRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", RumRetentionFilterDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.UpdateRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.UpdateRetentionFilter`:\n%s\n", responseContent)
}
