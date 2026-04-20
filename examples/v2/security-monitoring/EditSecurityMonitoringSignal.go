// Update security signal triage state or assignee returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalUpdateRequest{
		Data: datadogV2.SecurityMonitoringSignalUpdateData{
			Attributes: datadogV2.SecurityMonitoringSignalUpdateAttributes{
				ArchiveReason: datadogV2.SECURITYMONITORINGSIGNALARCHIVEREASON_NONE.Ptr(),
				Assignee: &datadogV2.SecurityMonitoringTriageUser{
					Uuid: "773b045d-ccf8-4808-bd3b-955ef6a8c940",
				},
				State: datadogV2.SECURITYMONITORINGSIGNALSTATE_OPEN.Ptr(),
			},
			Type: datadogV2.SECURITYMONITORINGSIGNALMETADATATYPE_SIGNAL_METADATA.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringSignal(ctx, "signal_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignal`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignal`:\n%s\n", responseContent)
}
