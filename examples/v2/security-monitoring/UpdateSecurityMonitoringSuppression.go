// Update a suppression rule returns "OK" response

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
	// there is a valid "suppression" in the system
	SuppressionDataID := os.Getenv("SUPPRESSION_DATA_ID")

	body := datadogV2.SecurityMonitoringSuppressionUpdateRequest{
		Data: datadogV2.SecurityMonitoringSuppressionUpdateData{
			Attributes: datadogV2.SecurityMonitoringSuppressionUpdateAttributes{
				SuppressionQuery: datadog.PtrString("env:staging status:low"),
			},
			Type: datadogV2.SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringSuppression(ctx, SuppressionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringSuppression`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringSuppression`:\n%s\n", responseContent)
}
