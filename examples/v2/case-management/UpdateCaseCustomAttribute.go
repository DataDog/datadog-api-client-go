// Update case custom attribute returns "OK" response

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

	body := datadogV2.CaseUpdateCustomAttributeRequest{
		Data: datadogV2.CaseUpdateCustomAttribute{
			Attributes: datadogV2.CustomAttributeValue{
				Type:    datadogV2.CUSTOMATTRIBUTETYPE_TEXT,
				IsMulti: true,
				Value: datadogV2.CustomAttributeValuesUnion{
					CustomAttributeMultiStringValue: &[]string{
						"Abba",
						"The Cure",
					}},
			},
			Type: datadogV2.CASERESOURCETYPE_CASE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.UpdateCaseCustomAttribute(ctx, CaseWithTypeID, CustomAttributeAttributesKey, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdateCaseCustomAttribute`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.UpdateCaseCustomAttribute`:\n%s\n", responseContent)
}
