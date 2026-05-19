// Remove insights from a case returns "OK" response

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
	body := datadogV2.CaseInsightsRequest{
		Data: datadogV2.CaseInsightsData{
			Attributes: datadogV2.CaseInsightsAttributes{
				Insights: []datadogV2.CaseInsight{
					{
						Ref:        "/monitors/12345?q=total",
						ResourceId: "12345",
						Type:       datadogV2.CASEINSIGHTTYPE_SECURITY_SIGNAL,
					},
				},
			},
			Type: datadogV2.CASERESOURCETYPE_CASE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RemoveCaseInsights", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.RemoveCaseInsights(ctx, "case_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.RemoveCaseInsights`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.RemoveCaseInsights`:\n%s\n", responseContent)
}
