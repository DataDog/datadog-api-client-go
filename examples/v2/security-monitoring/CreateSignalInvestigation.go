// Create a security signal investigation returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalInvestigationRequest{
		Data: datadogV2.SecurityMonitoringSignalInvestigationRequestData{
			Attributes: datadogV2.SecurityMonitoringSignalInvestigationRequestAttributes{
				Deployment: datadog.PtrString("live"),
				SignalId:   "AAAAAWgN8Xwgr1vKDQAAAABBV2dOOFh3ZzZobm1mWXJFYTR0OA",
			},
			Type: datadogV2.SECURITYMONITORINGSIGNALINVESTIGATIONTYPE_INVESTIGATION_SIGNAL,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSignalInvestigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSignalInvestigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSignalInvestigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSignalInvestigation`:\n%s\n", responseContent)
}
