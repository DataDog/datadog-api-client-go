// Create a security filter returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	body := datadog.SecurityFilterCreateRequest{
		Data: datadog.SecurityFilterCreateData{
			Attributes: datadog.SecurityFilterCreateAttributes{
				ExclusionFilters: []datadog.SecurityFilterExclusionFilter{
					{
						Name:  "Exclude staging",
						Query: "source:staging",
					},
				},
				FilteredDataType: datadog.SECURITYFILTERFILTEREDDATATYPE_LOGS,
				IsEnabled:        true,
				Name:             "Example-Create_a_security_filter_returns_OK_response",
				Query:            "service:ExampleCreateasecurityfilterreturnsOKresponse",
			},
			Type: datadog.SECURITYFILTERTYPE_SECURITY_FILTERS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFilter(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFilter`:\n%s\n", responseContent)
}
