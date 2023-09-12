// Update a custom destination returns "OK" response

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
	// there is a valid "logs_custom_destination" in the system
	LogsCustomDestinationDataID := os.Getenv("LOGS_CUSTOM_DESTINATION_DATA_ID")

	body := datadogV2.CustomDestinationUpdatePayload{
		Data: datadogV2.CustomDestinationWithId{
			Id: LogsCustomDestinationDataID,
			Attributes: datadogV2.CustomDestinationAttributes{
				Name:    "my-destination--updated-name",
				Version: 0,
				ForwarderDestination: datadogV2.CustomDestinationForwarderDestination{
					HttpDestination: &datadogV2.HttpDestination{
						Type:     datadogV2.HTTPDESTINATIONTYPE_HTTP,
						Endpoint: "https://example.org/my-intake",
						Auth: datadogV2.HttpDestinationAuth{
							HttpDestinationBasicAuth: &datadogV2.HttpDestinationBasicAuth{
								Type:     datadogV2.HTTPDESTINATIONBASICAUTHTYPE_BASIC,
								Username: "username",
								Password: "password",
							}},
					}},
			},
			Type: "custom_destination",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLogsCustomDestination", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.UpdateLogsCustomDestination(ctx, LogsCustomDestinationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.UpdateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.UpdateLogsCustomDestination`:\n%s\n", responseContent)
}
