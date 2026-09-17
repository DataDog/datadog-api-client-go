// List analytics events returns "OK" response

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
	body := datadogV2.ProductAnalyticsAnalyticsListRequest{
		Data: datadogV2.ProductAnalyticsAnalyticsListRequestData{
			Attributes: datadogV2.ProductAnalyticsAnalyticsListRequestAttributes{
				From: 1771232048460,
				Query: datadogV2.ProductAnalyticsAnalyticsListQuery{
					Columns: []string{
						"@view.name",
					},
					Limit: datadog.PtrInt64(100),
					Query: datadogV2.ProductAnalyticsBaseQuery{
						ProductAnalyticsEventQuery: &datadogV2.ProductAnalyticsEventQuery{
							DataSource: datadogV2.PRODUCTANALYTICSEVENTQUERYDATASOURCE_PRODUCT_ANALYTICS,
							Search: datadogV2.ProductAnalyticsEventSearch{
								Query: datadog.PtrString("@type:view"),
							},
						}},
				},
				To: 1771836848262,
			},
			Type: datadogV2.PRODUCTANALYTICSANALYTICSLISTREQUESTTYPE_FORMULA_ANALYTICS_EXTENDED_LIST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsList", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsList(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsList`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsList`:\n%s\n", responseContent)
}
