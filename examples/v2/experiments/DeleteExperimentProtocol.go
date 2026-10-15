// Delete experiment protocol returns "No Content" response

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
	// there is a valid "experiment_protocol" in the system
	ExperimentProtocolDataAttributesUpdatedAt := os.Getenv("EXPERIMENT_PROTOCOL_DATA_ATTRIBUTES_UPDATED_AT")
	ExperimentProtocolDataID := uuid.MustParse(os.Getenv("EXPERIMENT_PROTOCOL_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	r, err := api.DeleteExperimentProtocol(ctx, ExperimentProtocolDataID, *datadogV2.NewDeleteExperimentProtocolOptionalParameters().WithExpectedUpdatedAt(ExperimentProtocolDataAttributesUpdatedAt))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.DeleteExperimentProtocol`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
