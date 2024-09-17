// Update case priority returns "OK" response

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
	// there is a valid "case" in the system
	CaseID := os.Getenv("CASE_ID")

	body := datadogV2.CaseUpdatePriorityRequest{
		Data: datadogV2.CaseUpdatePriority{
			Attributes: datadogV2.CaseUpdatePriorityAttributes{
				Priority: datadogV2.CASEPRIORITY_P3,
			},
			Type: datadogV2.CASERESOURCETYPE_CASE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.UpdatePriority(ctx, CaseID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdatePriority`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.UpdatePriority`:\n%s\n", responseContent)
}
