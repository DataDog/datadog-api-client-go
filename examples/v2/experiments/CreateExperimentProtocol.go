// Create experiment protocol returns "Created" response

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
	body := datadogV2.ExperimentsCreateExperimentProtocolRequest{
		Data: datadogV2.ExperimentsCreateExperimentProtocolRequestData{
			Type: datadogV2.EXPERIMENTSPUBLICPROTOCOLRESPONSEDATATYPE_PROTOCOLS,
			Attributes: datadogV2.ExperimentsCreateExperimentProtocolRequestDataAttributes{
				Name:        "ex-protocol-14bb9543f523edde",
				Description: *datadog.NewNullableString(datadog.PtrString("Protocol created by the API test.")),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateExperimentProtocol(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateExperimentProtocol`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateExperimentProtocol`:\n%s\n", responseContent)
}
