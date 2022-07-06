// Submit deflate distribution points returns "Payload accepted" response

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
	body := datadog.DistributionPointsPayload{
		Series: []datadog.DistributionPointsSeries{
			{
				Metric: "system.load.1.dist",
				Points: [][]datadog.DistributionPointItem{
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.MetricsApi(apiClient)
	resp, r, err := api.SubmitDistributionPoints(ctx, body, *datadog.NewSubmitDistributionPointsOptionalParameters().WithContentEncoding(datadog.DISTRIBUTIONPOINTSCONTENTENCODING_DEFLATE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitDistributionPoints`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.SubmitDistributionPoints`:\n%s\n", responseContent)
}
