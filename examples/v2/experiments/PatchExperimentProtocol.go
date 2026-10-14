// Update experiment protocol returns "OK" response

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
	// there is a valid "experiment_protocol" in the system
	ExperimentProtocolDataID := uuid.MustParse(os.Getenv("EXPERIMENT_PROTOCOL_DATA_ID"))

	body := datadogV2.ExperimentsPatchExperimentProtocolRequest{
		Data: datadogV2.ExperimentsPatchExperimentProtocolRequestData{
			Type: datadogV2.EXPERIMENTSPUBLICPROTOCOLRESPONSEDATATYPE_PROTOCOLS,
			Id:   datadog.PtrUUID(ExperimentProtocolDataID),
			Attributes: &datadogV2.ExperimentsPatchExperimentProtocolRequestDataAttributes{
				Description: *datadog.NewNullableString(datadog.PtrString("Updated protocol description.")),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.PatchExperimentProtocol(ctx, ExperimentProtocolDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.PatchExperimentProtocol`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.PatchExperimentProtocol`:\n%s\n", responseContent)
}
