// Create an on-call event email address returns "Created" response

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
	body := datadogV2.OnCallEventEmailAddressCreateRequest{
		Data: datadogV2.OnCallEventEmailAddressCreateData{
			Attributes: datadogV2.OnCallEventEmailAddressCreateAttributes{
				AlertType:   datadogV2.EVENTEMAILADDRESSALERTTYPE_INFO.Ptr(),
				Description: datadog.PtrString("On-call email address for my team."),
				Format:      datadogV2.EVENTEMAILADDRESSFORMAT_JSON,
				Tags: []string{
					"env:production",
					"team:my-team",
				},
				TeamHandle: "my-team",
			},
			Type: datadogV2.EVENTEMAILADDRESSRESOURCETYPE_EVENT_EMAILS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOnCallEventEmailAddress", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEventsApi(apiClient)
	resp, r, err := api.CreateOnCallEventEmailAddress(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateOnCallEventEmailAddress`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.CreateOnCallEventEmailAddress`:\n%s\n", responseContent)
}
