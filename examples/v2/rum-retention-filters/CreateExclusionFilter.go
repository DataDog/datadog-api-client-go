// Create a RUM exclusion filter returns "Created" response

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
	body := datadogV2.RumExclusionFilterCreateRequest{
		Data: datadogV2.RumExclusionFilterCreateData{
			Attributes: datadogV2.RumExclusionFilterCreateAttributes{
				Enabled:   datadog.PtrBool(true),
				EventType: datadogV2.RUMEXCLUSIONFILTEREVENTTYPE_ERROR.Ptr(),
				Name:      "Exclude noisy browser extension errors",
				Query:     datadog.PtrString("@error.message:*extension*"),
			},
			Type: datadogV2.RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateExclusionFilter", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.CreateExclusionFilter(ctx, "app_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.CreateExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.CreateExclusionFilter`:\n%s\n", responseContent)
}
