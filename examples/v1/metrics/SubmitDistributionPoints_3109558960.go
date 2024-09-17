// Submit deflate distribution points returns "Payload accepted" response

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
	body := datadogV1.DistributionPointsPayload{
		Series: []datadogV1.DistributionPointsSeries{
			{
				Metric: "system.load.1.dist",
				Points: [][]datadogV1.DistributionPointItem{
					{
						{DistributionPointTimestamp: datadog.PtrFloat64(float64(time.Now().Unix()))},
						{DistributionPointData: &[]float64{
							1.0,
							2.0,
						}},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMetricsApi(apiClient)
	resp, r, err := api.SubmitDistributionPoints(ctx, body, *datadogV1.NewSubmitDistributionPointsOptionalParameters().WithContentEncoding(datadogV1.DISTRIBUTIONPOINTSCONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitDistributionPoints`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.SubmitDistributionPoints`:\n%s\n", responseContent)
}
