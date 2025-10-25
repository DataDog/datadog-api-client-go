// Create event email address returns "Created" response

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
	body := datadogV2.CreateEventEmailAddressRequest{
		Data: &datadogV2.CreateEventEmailAddressRequestData{
			Attributes: &datadogV2.CreateEventEmailAddressRequestDataAttributes{
				Format: "",
				NotifyHandles: []string{
					"",
				},
				Tags: []string{
					"",
				},
			},
			Type: datadogV2.EVENTEMAILSTYPE_EVENT_EMAILS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateEventEmailAddress", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, r, err := api.CreateEventEmailAddress(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEventEmailAddress`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.CreateEventEmailAddress`:\n%s\n", responseContent)
}
