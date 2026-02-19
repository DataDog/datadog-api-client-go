// Create a RUM segment returns "Created" response

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
	body := datadogV2.RumSegmentCreateRequest{
		Data: datadogV2.RumSegmentCreateData{
			Attributes: datadogV2.RumSegmentCreateAttributes{
				DataQuery: datadogV2.RumSegmentDataQuery{
					Combination: datadog.PtrString("(logs && apm_home) && NOT(apm_trace)"),
					EventPlatforms: []datadogV2.RumSegmentEventPlatform{
						{
							Facet: "@usr.id",
							From:  datadog.PtrInt64(1709888355000),
							Name:  "logs",
							Query: "@type:view @view.url_path:/logs",
							To:    datadog.PtrInt64(1710493155000),
						},
					},
					Journeys: []datadogV2.RumSegmentJourney{
						{
							ConversionType: datadog.PtrString("any"),
							GroupBy:        datadog.PtrString("@usr.id"),
							Name:           datadog.PtrString("my_journey"),
							Search:         datadog.PtrString("@type:view"),
						},
					},
					ReferenceTables: []datadogV2.RumSegmentReferenceTable{
						{
							Columns: []datadogV2.RumSegmentReferenceTableColumn{
								{
									Name: "user_id",
								},
							},
							FilterQuery: datadog.PtrString(""),
							JoinCondition: datadogV2.RumSegmentReferenceTableJoinCondition{
								ColumnName: "user_id",
								Facet:      "@usr.id",
							},
							Name:      "my_ref_table",
							TableName: "my_table",
						},
					},
					Static: []datadogV2.RumSegmentStaticEntry{
						{
							Id:        "static-list-1",
							Name:      "My Static List",
							UserCount: 500,
						},
					},
					Templates: []datadogV2.RumSegmentTemplateInstance{
						{
							From: datadog.PtrInt64(1709888355000),
							Parameters: map[string]string{
								"threshold": "5",
							},
							TemplateId: "stickiness-v1",
							To:         datadog.PtrInt64(1710493155000),
						},
					},
				},
				Description: datadog.PtrString("Users who visited the homepage."),
				Name:        "My Segment",
				Tags: []string{
					"team:frontend",
				},
			},
			Type: datadogV2.RUMSEGMENTRESOURCETYPE_SEGMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRumSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMUserSegmentsApi(apiClient)
	resp, r, err := api.CreateRumSegment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMUserSegmentsApi.CreateRumSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMUserSegmentsApi.CreateRumSegment`:\n%s\n", responseContent)
}
