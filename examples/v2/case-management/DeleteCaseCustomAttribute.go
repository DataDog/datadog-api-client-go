// Delete custom attribute from case returns "OK" response

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
	// there is a valid "case" with a custom "case_type" in the system
	CaseWithTypeID := os.Getenv("CASE_WITH_TYPE_ID")

	// there is a valid "custom_attribute" in the system
	CustomAttributeAttributesKey := os.Getenv("CUSTOM_ATTRIBUTE_ATTRIBUTES_KEY")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.DeleteCaseCustomAttribute(ctx, CaseWithTypeID, CustomAttributeAttributesKey)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.DeleteCaseCustomAttribute`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.DeleteCaseCustomAttribute`:\n%s\n", responseContent)
}
