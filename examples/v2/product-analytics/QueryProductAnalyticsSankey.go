// Compute Sankey flow analysis returns "OK" response

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
	body := datadogV2.ProductAnalyticsSankeyRequest{
		Data: datadogV2.ProductAnalyticsSankeyRequestData{
			Attributes: datadogV2.ProductAnalyticsSankeyRequestAttributes{
				DataSource: "product_analytics",
				Definition: datadogV2.ProductAnalyticsSankeyDefinition{
					EntriesPerStep: datadog.PtrInt64(5),
					NumberOfSteps:  datadog.PtrInt64(5),
					Source:         datadog.PtrString("/logs"),
					Target:         datadog.PtrString(""),
				},
				Search: datadogV2.ProductAnalyticsSankeySearch{
					JoinKeys: &datadogV2.ProductAnalyticsJoinKeys{
						Primary: datadog.PtrString("@session.id"),
					},
					Query: "@type:view",
				},
				Time: datadogV2.ProductAnalyticsSankeyTime{
					From: 1771232048460,
					To:   1771836848262,
				},
			},
			Type: datadogV2.PRODUCTANALYTICSSANKEYREQUESTTYPE_SANKEY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
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
