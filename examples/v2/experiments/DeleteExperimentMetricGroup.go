// Delete experiment metric group returns "No Content" response

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

	// there is a valid "experiment_metric_group" in the system
	ExperimentMetricGroupDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_GROUP_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.DeleteExperimentMetricGroup(ctx, ExperimentDataID, ExperimentMetricGroupDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.DeleteExperimentMetricGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
