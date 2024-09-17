// Create a security filter returns "OK" response

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
	body := datadogV2.SecurityFilterCreateRequest{
		Data: datadogV2.SecurityFilterCreateData{
			Attributes: datadogV2.SecurityFilterCreateAttributes{
				ExclusionFilters: []datadogV2.SecurityFilterExclusionFilter{
					{
						Name:  "Exclude staging",
						Query: "source:staging",
					},
				},
				FilteredDataType: datadogV2.SECURITYFILTERFILTEREDDATATYPE_LOGS,
				IsEnabled:        true,
				Name:             "Example-Security-Monitoring",
				Query:            "service:ExampleSecurityMonitoring",
			},
			Type: datadogV2.SECURITYFILTERTYPE_SECURITY_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFilter`:\n%s\n", responseContent)
}
