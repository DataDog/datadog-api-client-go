// Submit deflate metrics returns "Payload accepted" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.MetricsPayload{
		Series: []datadog.Series{
			{
				Metric: "system.load.1",
				Type:   datadog.PtrString("gauge"),
				Points: [][]*float64{
					{
						datadog.PtrFloat64(float64(time.Now().Unix())),
						datadog.PtrFloat64(1.1),
					},
				},
				Tags: []string{
					"test:ExampleSubmitdeflatemetricsreturnsPayloadacceptedresponse",
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.MetricsApi(apiClient)
	resp, r, err := api.SubmitMetrics(ctx, body, *datadog.NewSubmitMetricsOptionalParameters().WithContentEncoding(datadog.METRICCONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitMetrics`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.SubmitMetrics`:\n%s\n", responseContent)
}
