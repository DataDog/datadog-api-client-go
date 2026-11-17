// Create a segment returns "Created" response

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
	body := datadogV2.SegmentCreateRequest{
		Data: datadogV2.SegmentCreateData{
			Attributes: datadogV2.SegmentCreateAttributes{
				DataQuery: datadogV2.SegmentDataQuery{
					EventPlatform: []datadogV2.SegmentEventPlatformBlock{
						{
							Facet: "@usr.id",
							From:  datadog.PtrString("-2592000000"),
							Name:  datadog.PtrString("viewed_logs"),
							Query: datadog.PtrString("@type:view @view.name:/logs"),
							To:    datadog.PtrString("0"),
						},
					},
				},
				Description: datadog.PtrString("Users who viewed the logs page in the last 30 days."),
				Name:        "High-value users",
				Tags: []string{
					"team:growth",
				},
			},
			Type: datadogV2.SEGMENTRESOURCETYPE_SEGMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateProductAnalyticsSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumAudienceManagementApi(apiClient)
	resp, r, err := api.CreateProductAnalyticsSegment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAudienceManagementApi.CreateProductAnalyticsSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumAudienceManagementApi.CreateProductAnalyticsSegment`:\n%s\n", responseContent)
}
