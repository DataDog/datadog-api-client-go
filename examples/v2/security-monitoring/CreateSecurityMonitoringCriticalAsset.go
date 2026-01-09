// Create a critical asset returns "OK" response

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
	body := datadogV2.SecurityMonitoringCriticalAssetCreateRequest{
		Data: datadogV2.SecurityMonitoringCriticalAssetCreateData{
			Type: datadogV2.SECURITYMONITORINGCRITICALASSETTYPE_CRITICAL_ASSETS,
			Attributes: datadogV2.SecurityMonitoringCriticalAssetCreateAttributes{
				Query:     "host:examplesecuritymonitoring",
				RuleQuery: "type:(log_detection OR signal_correlation OR workload_security OR application_security) source:cloudtrail",
				Severity:  datadogV2.SECURITYMONITORINGCRITICALASSETSEVERITY_DECREASE,
				Tags: []string{
					"team:security",
					"env:test",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringCriticalAsset(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringCriticalAsset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringCriticalAsset`:\n%s\n", responseContent)
}
