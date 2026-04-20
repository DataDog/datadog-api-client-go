// Bulk update security signals returns "OK" response

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
	body := datadogV2.SecurityMonitoringSignalsBulkUpdateRequest{
		Data: []datadogV2.SecurityMonitoringSignalsBulkUpdateData{
			{
				Attributes: datadogV2.SecurityMonitoringSignalUpdateAttributes{
					ArchiveReason: datadogV2.SECURITYMONITORINGSIGNALARCHIVEREASON_NONE.Ptr(),
					Assignee: &datadogV2.SecurityMonitoringTriageUser{
						Uuid: "773b045d-ccf8-4808-bd3b-955ef6a8c940",
					},
					State: datadogV2.SECURITYMONITORINGSIGNALSTATE_OPEN.Ptr(),
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
	resp, r, err := api.BulkEditSecurityMonitoringSignals(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.BulkEditSecurityMonitoringSignals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.BulkEditSecurityMonitoringSignals`:\n%s\n", responseContent)
}
