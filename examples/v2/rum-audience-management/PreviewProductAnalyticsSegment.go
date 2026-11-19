// Preview a segment returns "OK" response

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
	body := datadogV2.SegmentPreviewRequest{
		Data: datadogV2.SegmentPreviewData{
			Attributes: datadogV2.SegmentPreviewAttributes{
				DataQuery: datadogV2.SegmentDataQuery{
					EventPlatform: []datadogV2.SegmentEventPlatformBlock{
						{
							Facet: "@usr.id",
							From:  datadog.PtrString("-2592000000"),
							Query: datadog.PtrString("@type:view @view.name:/logs"),
							To:    datadog.PtrString("0"),
						},
					},
				},
				Limit: datadog.PtrInt64(50),
				UserStoreColumns: []string{
					"user_id",
					"user_email",
					"user_name",
				},
			},
			Type: datadogV2.SEGMENTPREVIEWREQUESTTYPE_SEGMENT_PREVIEW_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PreviewProductAnalyticsSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.PreviewProductAnalyticsSegment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.PreviewProductAnalyticsSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.PreviewProductAnalyticsSegment`:\n%s\n", responseContent)
}
