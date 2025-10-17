// Get all custom attributes config of case type returns "OK" response

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
	// there is a valid "case_type" in the system
	CaseTypeID := os.Getenv("CASE_TYPE_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementAttributeApi(apiClient)
	resp, r, err := api.GetAllCustomAttributeConfigsByCaseType(ctx, CaseTypeID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementAttributeApi.GetAllCustomAttributeConfigsByCaseType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementAttributeApi.GetAllCustomAttributeConfigsByCaseType`:\n%s\n", responseContent)
}
