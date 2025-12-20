// Search security findings returns "OK" response with pagination

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
	body := datadogV2.SecurityFindingsSearchRequest{
		Data: &datadogV2.SecurityFindingsSearchRequestData{
			Attributes: &datadogV2.SecurityFindingsSearchRequestDataAttributes{
				Filter: datadog.PtrString("@severity:(critical OR high)"),
				Page: &datadogV2.SecurityFindingsSearchRequestPage{
					Limit: datadog.PtrInt64(1),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SearchSecurityFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.SearchSecurityFindings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityFindings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.SearchSecurityFindings`:\n%s\n", responseContent)
}
