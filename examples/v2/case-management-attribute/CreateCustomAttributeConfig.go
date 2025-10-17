// Create custom attribute config for a case type returns "CREATED" response

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

	body := datadogV2.CustomAttributeConfigCreateRequest{
		Data: datadogV2.CustomAttributeConfigCreate{
			Attributes: datadogV2.CustomAttributeConfigAttributesCreate{
				DisplayName: "AWS Region 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
				IsMulti:     true,
				Key:         "region_d9fe56bc9274fbb6",
				Type:        datadogV2.CUSTOMATTRIBUTETYPE_NUMBER,
			},
			Type: datadogV2.CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementAttributeApi(apiClient)
	resp, r, err := api.CreateCustomAttributeConfig(ctx, CaseTypeID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementAttributeApi.CreateCustomAttributeConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementAttributeApi.CreateCustomAttributeConfig`:\n%s\n", responseContent)
}
