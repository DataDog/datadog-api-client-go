// Compute a Sankey diagram returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.ProductAnalyticsSankeyRequest{
		Data: datadogV2.ProductAnalyticsSankeyRequestData{
			Attributes: datadogV2.ProductAnalyticsSankeyRequestAttributes{
				Definition: datadogV2.ProductAnalyticsSankeyDefinition{
					EntriesPerStep: datadog.PtrInt64(10),
					NumberOfSteps:  datadog.PtrInt64(3),
					Source:         "@view.name",
					Target:         "@view.name",
				},
				Search: datadogV2.ProductAnalyticsSankeySearch{
					AudienceFilters: &datadogV2.ProductAnalyticsAudienceFilters{
						Accounts: []datadogV2.ProductAnalyticsAudienceAccountSubquery{
							{
								Name: "",
							},
						},
						Formula: datadog.PtrString("u"),
						Segments: []datadogV2.ProductAnalyticsAudienceSegmentSubquery{
							{
								Name:      "",
								SegmentId: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
							},
						},
						Users: []datadogV2.ProductAnalyticsAudienceUserSubquery{
							{
								Name:  "u",
								Query: datadog.PtrString("*"),
							},
						},
					},
					JoinKeys: &datadogV2.ProductAnalyticsJoinKeys{
						Primary:   datadog.PtrString("@session.id"),
						Secondary: []string{},
					},
					Query: datadog.PtrString("@type:view"),
				},
				Time: datadogV2.ProductAnalyticsSankeyTime{
					From: 1756425600000,
					To:   1756857600000,
				},
			},
			Type: datadogV2.PRODUCTANALYTICSSANKEYREQUESTTYPE_SANKEY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsSankey", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsSankey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsSankey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsSankey`:\n%s\n", responseContent)
}
