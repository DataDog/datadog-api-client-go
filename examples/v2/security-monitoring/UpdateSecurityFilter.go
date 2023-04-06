// Update a security filter returns "OK" response

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
	// there is a valid "security_filter" in the system
	SecurityFilterDataID := os.Getenv("SECURITY_FILTER_DATA_ID")

	body := datadogV2.SecurityFilterUpdateRequest{
		Data: datadogV2.SecurityFilterUpdateData{
			Attributes: datadogV2.SecurityFilterUpdateAttributes{
				ExclusionFilters: []datadogV2.SecurityFilterExclusionFilter{},
				FilteredDataType: datadogV2.SECURITYFILTERFILTEREDDATATYPE_LOGS.Ptr(),
				IsEnabled:        datadog.PtrBool(true),
				Name:             datadog.PtrString("Example-Security-Monitoring"),
				Query:            datadog.PtrString("service:ExampleSecurityMonitoring"),
				Version:          datadog.PtrInt32(1),
			},
			Type: datadogV2.SECURITYFILTERTYPE_SECURITY_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityFilter(ctx, SecurityFilterDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityFilter`:\n%s\n", responseContent)
}
