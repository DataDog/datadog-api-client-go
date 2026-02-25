// Compute timeseries analytics returns "OK" response

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
	body := datadogV2.ProductAnalyticsAnalyticsRequest{
		Data: datadogV2.ProductAnalyticsAnalyticsRequestData{
			Attributes: datadogV2.ProductAnalyticsAnalyticsRequestAttributes{
				From: 1771232048460,
				Query: datadogV2.ProductAnalyticsAnalyticsQuery{
					Compute: datadogV2.ProductAnalyticsCompute{
						Aggregation: "count",
					},
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
			Type: datadogV2.PRODUCTANALYTICSANALYTICSREQUESTTYPE_FORMULA_ANALYTICS_EXTENDED_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsTimeseries(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsTimeseries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsTimeseries`:\n%s\n", responseContent)
}
