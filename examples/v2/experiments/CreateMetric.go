// Create metric returns "Created" response

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
	body := datadogV2.ExperimentsCreateMetricV2Request{
		Data: datadogV2.ExperimentsCreateMetricV2RequestData{
			Type: datadogV2.METRICTYPE_METRICS,
			Attributes: datadogV2.ExperimentsCreateMetricV2RequestDataAttributes{
				ExperimentsCreateMetricNumeratorAttributes: &datadogV2.ExperimentsCreateMetricNumeratorAttributes{
					Name:           "ex-14bb9543f523edde",
					DataSourceType: datadogV2.EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_DATADOG,
					DesiredChange:  datadogV2.EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDESIREDCHANGE_METRIC_INCREASES,
					NumeratorAggregation: datadogV2.ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation{
						ExperimentsDatadogMetricAggregationInput: &datadogV2.ExperimentsDatadogMetricAggregationInput{
							Operation: "sum",
							DatadogMetricMeasure: datadogV2.ExperimentsDatadogMetricMeasureInput{
								Name:          "ex-14bb9543f523edde view duration",
								SourceType:    "PRODUCT_ANALYTICS",
								SourceSubtype: "RUM_VIEWS",
								ColumnType:    "double",
								ColumnName:    datadog.PtrString("@view.time_spent"),
							},
						}},
				}},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateMetric`:\n%s\n", responseContent)
}
