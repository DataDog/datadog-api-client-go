// Delete custom attributes config returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "case_type" in the system
	CaseTypeID := os.Getenv("CASE_TYPE_ID")

	// there is a valid "custom_attribute" in the system
	CustomAttributeID := os.Getenv("CUSTOM_ATTRIBUTE_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementAttributeApi(apiClient)
	r, err := api.DeleteCustomAttributeConfig(ctx, CaseTypeID, CustomAttributeID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementAttributeApi.DeleteCustomAttributeConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
