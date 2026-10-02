// Start experiment returns "The experiment was started." response

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
	// there is a valid "configured_experiment" in the system
	ConfiguredExperimentDataID := uuid.MustParse(os.Getenv("CONFIGURED_EXPERIMENT_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.StartExperiment(ctx, ConfiguredExperimentDataID, *datadogV2.NewStartExperimentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.StartExperiment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
