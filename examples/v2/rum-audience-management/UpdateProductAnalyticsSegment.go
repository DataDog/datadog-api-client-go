// Update a segment returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SegmentUpdateRequest{
		Data: datadogV2.SegmentUpdateData{
			Attributes: datadogV2.SegmentUpdateAttributes{
				Description: datadog.PtrString("Users who viewed the logs page in the last 7 days."),
				Name:        datadog.PtrString("High-value users"),
			},
			Id:   "8a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d",
			Type: datadogV2.SEGMENTRESOURCETYPE_SEGMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateProductAnalyticsSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	r, err := api.UpdateProductAnalyticsSegment(ctx, "8a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.UpdateProductAnalyticsSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
