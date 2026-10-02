// Update experiment metric group returns "OK" response

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

	// there is a valid "experiment_metric_group" in the system
	ExperimentMetricGroupDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_GROUP_DATA_ID"))

	body := datadogV2.ExperimentsPatchExperimentMetricGroupV2Request{
		Data: datadogV2.ExperimentsPatchExperimentMetricGroupV2RequestData{
			Type: datadogV2.EXPERIMENTSPATCHEXPERIMENTMETRICGROUPV2REQUESTDATATYPE_EXPERIMENT_METRIC_GROUPS,
			Id:   datadog.PtrUUID(ExperimentMetricGroupDataID),
			Attributes: &datadogV2.ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes{
				Name: datadog.PtrString("ex-14bb9543f523edde updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.UpdateExperimentMetricGroup(ctx, ExperimentDataID, ExperimentMetricGroupDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.UpdateExperimentMetricGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.UpdateExperimentMetricGroup`:\n%s\n", responseContent)
}
