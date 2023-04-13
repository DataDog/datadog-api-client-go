// Submit deflate metrics returns "Payload accepted" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.MetricsPayload{
		Series: []datadogV1.Series{
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
					"test:ExampleMetric",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMetricsApi(apiClient)
	resp, r, err := api.SubmitMetrics(ctx, body, *datadogV1.NewSubmitMetricsOptionalParameters().WithContentEncoding(datadogV1.METRICCONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitMetrics`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.SubmitMetrics`:\n%s\n", responseContent)
}
