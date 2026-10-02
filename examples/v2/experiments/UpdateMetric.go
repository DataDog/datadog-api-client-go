// Update metric returns "OK" response

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
	// there is a valid "experiment_metric" in the system
	ExperimentMetricDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_DATA_ID"))

	body := datadogV2.ExperimentsUpdateMetricV2Request{
		Data: datadogV2.ExperimentsUpdateMetricV2RequestData{
			Type: datadogV2.METRICTYPE_METRICS,
			Id:   datadog.PtrUUID(ExperimentMetricDataID),
			Attributes: &datadogV2.ExperimentsUpdateMetricV2RequestDataAttributes{
				Name: datadog.PtrString("ex-14bb9543f523edde updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateMetric(ctx, ExperimentMetricDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateMetric`:\n%s\n", responseContent)
}
