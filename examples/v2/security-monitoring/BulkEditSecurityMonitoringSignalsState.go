// Bulk update triage state of security signals returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalsBulkStateUpdateRequest{
		Data: []datadogV2.SecurityMonitoringSignalsBulkStateUpdateData{
			{
				Attributes: datadogV2.SecurityMonitoringSignalStateUpdateAttributes{
					ArchiveReason: datadogV2.SECURITYMONITORINGSIGNALARCHIVEREASON_NONE.Ptr(),
					State:         datadogV2.SECURITYMONITORINGSIGNALSTATE_OPEN,
				},
				Id:   "AAAAAWgN8Xwgr1vKDQAAAABBV2dOOFh3ZzZobm1mWXJFYTR0OA",
				Type: datadogV2.SECURITYMONITORINGSIGNALTYPE_SIGNAL.Ptr(),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.BulkEditSecurityMonitoringSignalsState(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkEditSecurityMonitoringSignalsState`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkEditSecurityMonitoringSignalsState`:\n%s\n", responseContent)
}
