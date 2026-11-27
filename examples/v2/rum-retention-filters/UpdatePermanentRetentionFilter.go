// Update a permanent RUM retention filter returns "Updated" response

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
	// there is a valid "rum_application" in the system
	RumApplicationDataID := os.Getenv("RUM_APPLICATION_DATA_ID")

	body := datadogV2.RumPermanentRetentionFilterUpdateRequest{
		Data: datadogV2.RumPermanentRetentionFilterUpdateData{
			Id:   datadogV2.RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS,
			Type: datadogV2.RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS,
			Attributes: datadogV2.RumPermanentRetentionFilterUpdateAttributes{
				CrossProductSampling: &datadogV2.RumCrossProductSamplingUpdate{
					TraceEnabled:    datadog.PtrBool(true),
					TraceSampleRate: datadog.PtrFloat64(25.0),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.UpdatePermanentRetentionFilter(ctx, RumApplicationDataID, datadogV2.RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.UpdatePermanentRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.UpdatePermanentRetentionFilter`:\n%s\n", responseContent)
}
