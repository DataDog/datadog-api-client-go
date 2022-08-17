// Change the triage state of a security signal returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalStateUpdateRequest{
		Data: datadogV2.SecurityMonitoringSignalStateUpdateData{
			Attributes: datadogV2.SecurityMonitoringSignalStateUpdateAttributes{
				ArchiveReason: datadogV2.SECURITYMONITORINGSIGNALARCHIVEREASON_NONE.Ptr(),
				State:         datadogV2.SECURITYMONITORINGSIGNALSTATE_OPEN,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringSignalState(ctx, "AQAAAYG1bl5K4HuUewAAAABBWUcxYmw1S0FBQmt2RmhRN0V4ZUVnQUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignalState`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignalState`:\n%s\n", responseContent)
}
