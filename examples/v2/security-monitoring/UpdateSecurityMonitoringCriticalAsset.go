// Update a critical asset returns "OK" response

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
	// there is a valid "critical_asset" in the system
	CriticalAssetDataID := os.Getenv("CRITICAL_ASSET_DATA_ID")

	body := datadogV2.SecurityMonitoringCriticalAssetUpdateRequest{
		Data: datadogV2.SecurityMonitoringCriticalAssetUpdateData{
			Type: datadogV2.SECURITYMONITORINGCRITICALASSETTYPE_CRITICAL_ASSETS,
			Attributes: datadogV2.SecurityMonitoringCriticalAssetUpdateAttributes{
				Enabled:   datadog.PtrBool(false),
				Query:     datadog.PtrString("no:alert"),
				RuleQuery: datadog.PtrString("type:(log_detection OR signal_correlation OR workload_security OR application_security) ruleId:djg-ktx-ipq"),
				Severity:  datadogV2.SECURITYMONITORINGCRITICALASSETSEVERITY_DECREASE.Ptr(),
				Tags: []string{
					"env:production",
				},
				Version: datadog.PtrInt32(1),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringCriticalAsset(ctx, CriticalAssetDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringCriticalAsset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringCriticalAsset`:\n%s\n", responseContent)
}
