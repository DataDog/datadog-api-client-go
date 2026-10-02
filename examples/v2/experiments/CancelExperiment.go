// Cancel experiment returns "The experiment was canceled and unlinked from its feature flag allocations." response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "experiment" in the system
	ExperimentDataID := uuid.MustParse(os.Getenv("EXPERIMENT_DATA_ID"))

	body := datadogV2.ExperimentsCancelExperimentV2Request{
		Data: datadogV2.ExperimentsCancelExperimentV2RequestData{
			Type: datadogV2.EXPERIMENTSCANCELEXPERIMENTV2REQUESTDATATYPE_CANCEL_EXPERIMENT_REQUEST,
			Attributes: datadogV2.ExperimentsCancelExperimentV2RequestDataAttributes{
				Reason: "Cancel the test experiment",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.CancelExperiment(ctx, ExperimentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CancelExperiment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
