// Create a security filter returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityFilterCreateRequest{
		Data: datadog.SecurityFilterCreateData{
			Attributes: datadog.SecurityFilterCreateAttributes{
				ExclusionFilters: []datadog.SecurityFilterExclusionFilter{
					datadog.SecurityFilterExclusionFilter{
						Name:  "Exclude staging",
						Query: "source:staging",
					},
				},
				FilteredDataType: datadog.SecurityFilterFilteredDataType("logs"),
				IsEnabled:        true,
				Name:             "Example-Create_a_security_filter_returns_OK_response",
				Query:            "service:ExampleCreateasecurityfilterreturnsOKresponse",
			},
			Type: datadog.SecurityFilterType("security_filters"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.CreateSecurityFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFilter`:\n%s\n", responseContent)
}
