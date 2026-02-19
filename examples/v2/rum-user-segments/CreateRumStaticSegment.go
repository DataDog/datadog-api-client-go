// Create a static RUM segment returns "Created" response

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
	body := datadogV2.RumStaticSegmentCreateRequest{
		Data: datadogV2.RumStaticSegmentCreateData{
			Attributes: datadogV2.RumStaticSegmentCreateAttributes{
				Description: "Users from a specific journey.",
				JourneyQueryObject: datadogV2.RumStaticSegmentJourneyQueryObject{
					Nodes: []datadogV2.RumStaticSegmentJourneyNode{
						{
							Filters: []datadogV2.RumStaticSegmentJourneyFilter{
								{
									Attribute: "@type",
									Value:     "view",
								},
							},
						},
					},
				},
				Name: "My Static Segment",
				Tags: []string{
					"team:frontend",
				},
			},
			Type: datadogV2.RUMSTATICSEGMENTREQUESTTYPE_CREATE_STATIC_SEGMENT_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRumStaticSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMUserSegmentsApi(apiClient)
	resp, r, err := api.CreateRumStaticSegment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMUserSegmentsApi.CreateRumStaticSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMUserSegmentsApi.CreateRumStaticSegment`:\n%s\n", responseContent)
}
