// Update an event email address returns "OK" response

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
	body := datadogV2.EventEmailAddressUpdateRequest{
		Data: datadogV2.EventEmailAddressUpdateData{
			Attributes: datadogV2.EventEmailAddressUpdateAttributes{
				AlertType:   datadogV2.EVENTEMAILADDRESSALERTTYPE_INFO,
				Description: *datadog.NewNullableString(datadog.PtrString("Updated description for the email address.")),
				NotifyHandles: []string{
					"@slack-my-channel",
				},
				Tags: []string{
					"env:production",
					"team:my-team",
				},
			},
			Type: datadogV2.EVENTEMAILADDRESSRESOURCETYPE_EVENT_EMAILS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateEventEmailAddress", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, r, err := api.UpdateEventEmailAddress(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000001"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.UpdateEventEmailAddress`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.UpdateEventEmailAddress`:\n%s\n", responseContent)
}
