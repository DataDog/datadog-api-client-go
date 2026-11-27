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
			Type: datadogV2.RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS,
			Attributes: datadogV2.RumExclusionFilterCreateAttributes{
				Name:      "Example-Rum-Retention-Filter",
				Enabled:   datadog.PtrBool(true),
				EventType: datadogV2.RUMEXCLUSIONFILTEREVENTTYPE_ERROR.Ptr(),
				Query:     datadog.PtrString("@error.message:*extension*"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.CreateExclusionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.CreateExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.CreateExclusionFilter`:\n%s\n", responseContent)
}
