// Create experiment metric group returns "Created" response

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
	// there is a valid "experiment" in the system
	ExperimentDataID := uuid.MustParse(os.Getenv("EXPERIMENT_DATA_ID"))

	// there is a valid "experiment_metric" in the system
	ExperimentMetricDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_DATA_ID"))

	body := datadogV2.ExperimentsCreateExperimentMetricGroupV2Request{
		Data: datadogV2.ExperimentsCreateExperimentMetricGroupV2RequestData{
			Type: datadogV2.EXPERIMENTSPATCHEXPERIMENTMETRICGROUPV2REQUESTDATATYPE_EXPERIMENT_METRIC_GROUPS,
			Attributes: datadogV2.ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes{
				Name: "ex-14bb9543f523edde",
				Metrics: []datadogV2.ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems{
					{
						MetricId: ExperimentMetricDataID,
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateExperimentMetricGroup(ctx, ExperimentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateExperimentMetricGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateExperimentMetricGroup`:\n%s\n", responseContent)
}
