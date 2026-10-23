// List experiment protocols returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "experiment_protocol" in the system
	ExperimentProtocolDataAttributesName := os.Getenv("EXPERIMENT_PROTOCOL_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.ListExperimentProtocols(ctx, *datadogV2.NewListExperimentProtocolsOptionalParameters().WithFilterStatus([]datadogV2.ExperimentsPublicProtocolResponseDataAttributesStatus{
		datadogV2.EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_DRAFT,
	}).WithFilterQuery(ExperimentProtocolDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.ListExperimentProtocols`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.ListExperimentProtocols`:\n%s\n", responseContent)
}
