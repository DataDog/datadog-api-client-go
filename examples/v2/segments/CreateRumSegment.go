// Create rum segment returns "Segment created successfully" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.Segment{
		Data: &datadogV2.SegmentData{
			Attributes: &datadogV2.SegmentDataAttributes{
				CreatedAt: datadog.PtrTime(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)),
				CreatedBy: &datadogV2.SegmentDataSource{
					Handle: "",
					Id:     "",
					Uuid:   "",
				},
				DataQuery: datadogV2.SegmentDataAttributesDataQuery{
					EventPlatform: []datadogV2.SegmentDataAttributesDataQueryEventPlatformItems{
						{
							Facet: "@usr.id",
							From:  datadog.PtrString("2025-08-01"),
							Name:  datadog.PtrString("high_value_users"),
							Query: datadog.PtrString("@type:view @view.name:/logs @usr.session_duration:>300000"),
							To:    datadog.PtrString("2025-09-01"),
						},
					},
				},
				Description: datadog.PtrString("Users who frequently visit logs and have high session duration"),
				ModifiedAt:  datadog.PtrTime(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)),
				ModifiedBy: &datadogV2.SegmentDataSource{
					Handle: "",
					Id:     "",
					Uuid:   "",
				},
				Name:   "High-Value Users",
				OrgId:  datadog.PtrInt64(123456),
				Source: datadog.PtrInt64(0),
				Tags: []string{
					"high-value",
					"logs",
					"active",
				},
				Version: datadog.PtrInt64(1),
			},
			Id:   datadog.PtrString("segment-12345"),
			Type: datadogV2.SEGMENTDATATYPE_SEGMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRumSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSegmentsApi(apiClient)
	resp, r, err := api.CreateRumSegment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.CreateRumSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.CreateRumSegment`:\n%s\n", responseContent)
}
