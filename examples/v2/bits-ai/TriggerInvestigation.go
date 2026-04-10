// Trigger a Bits AI investigation returns "OK" response

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
	body := datadogV2.TriggerInvestigationRequest{
		Data: datadogV2.TriggerInvestigationRequestData{
			Attributes: datadogV2.TriggerInvestigationRequestDataAttributes{
				Trigger: datadogV2.TriggerAttributes{
					MonitorAlertTrigger: datadogV2.MonitorAlertTriggerAttributes{
						EventId:   "1234567890123456789",
						EventTs:   1700000000000,
						MonitorId: 12345678,
					},
					Type: datadogV2.TRIGGERTYPE_MONITOR_ALERT_TRIGGER,
				},
			},
			Type: datadogV2.TRIGGERINVESTIGATIONREQUESTTYPE_TRIGGER_INVESTIGATION_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.TriggerInvestigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsAIApi(apiClient)
	resp, r, err := api.TriggerInvestigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsAIApi.TriggerInvestigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsAIApi.TriggerInvestigation`:\n%s\n", responseContent)
}
