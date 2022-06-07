// Update a security filter returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityFilterUpdateRequest{
		Data: datadog.SecurityFilterUpdateData{
			Attributes: datadog.SecurityFilterUpdateAttributes{
				ExclusionFilters: []datadog.SecurityFilterExclusionFilter{},
				FilteredDataType: datadog.SECURITYFILTERFILTEREDDATATYPE_LOGS.Ptr(),
				IsEnabled:        datadog.PtrBool(true),
				Name:             datadog.PtrString("Custom security filter"),
				Query:            datadog.PtrString("service:api"),
				Version:          datadog.PtrInt32(1),
			},
			Type: datadog.SECURITYFILTERTYPE_SECURITY_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPlatformApi.UpdateSecurityFilter(ctx, "security_filter_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPlatformApi.UpdateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityPlatformApi.UpdateSecurityFilter`:\n%s\n", responseContent)
}
