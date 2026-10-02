// Create experiment metric group from collection returns "Created" response

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

	// there is a valid "experiment_metric_collection_with_metric" in the system
	ExperimentMetricCollectionWithMetricDataID := uuid.MustParse(os.Getenv("EXPERIMENT_METRIC_COLLECTION_WITH_METRIC_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateExperimentMetricGroupFromCollection(ctx, ExperimentDataID, ExperimentMetricCollectionWithMetricDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateExperimentMetricGroupFromCollection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateExperimentMetricGroupFromCollection`:\n%s\n", responseContent)
}
