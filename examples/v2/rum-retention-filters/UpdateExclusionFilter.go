// Update a RUM exclusion filter returns "Updated" response

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
	body := datadogV2.RumExclusionFilterUpdateRequest{
		Data: datadogV2.RumExclusionFilterUpdateData{
			Attributes: datadogV2.RumExclusionFilterUpdateAttributes{
				Enabled:   datadog.PtrBool(true),
				EventType: datadogV2.RUMEXCLUSIONFILTEREVENTTYPE_ERROR.Ptr(),
				Name:      datadog.PtrString("Exclude noisy browser extension errors"),
				Query:     datadog.PtrString("@error.message:*extension*"),
			},
			Id:   "051601eb-54a0-abc0-03f9-cc02efa18892",
			Type: datadogV2.RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.UpdateExclusionFilter(ctx, "app_id", "ef_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.UpdateExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.UpdateExclusionFilter`:\n%s\n", responseContent)
}
