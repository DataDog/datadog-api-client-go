// Update a Splunk custom destination with a sourcetype returns "OK" response

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
	// there is a valid "custom_destination_splunk" in the system
	CustomDestinationSplunkDataID := os.Getenv("CUSTOM_DESTINATION_SPLUNK_DATA_ID")

	body := datadogV2.CustomDestinationUpdateRequest{
		Data: &datadogV2.CustomDestinationUpdateRequestDefinition{
			Attributes: &datadogV2.CustomDestinationUpdateRequestAttributes{
				ForwarderDestination: &datadogV2.CustomDestinationForwardDestination{
					CustomDestinationForwardDestinationSplunk: &datadogV2.CustomDestinationForwardDestinationSplunk{
						Type:        datadogV2.CUSTOMDESTINATIONFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC,
						Endpoint:    "https://example.com",
						AccessToken: "my-access-token",
						Sourcetype:  *datadog.NewNullableString(datadog.PtrString("new-sourcetype")),
					}},
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
			Id:   CustomDestinationSplunkDataID,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.UpdateLogsCustomDestination(ctx, CustomDestinationSplunkDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.UpdateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.UpdateLogsCustomDestination`:\n%s\n", responseContent)
}
