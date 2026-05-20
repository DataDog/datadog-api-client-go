// Update custom attribute config returns "OK" response

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
	body := datadogV2.CustomAttributeConfigUpdateRequest{
		Data: datadogV2.CustomAttributeConfigUpdate{
			Attributes: &datadogV2.CustomAttributeConfigUpdateAttributes{
				Description: datadog.PtrString("Updated description."),
				DisplayName: datadog.PtrString("AWS Region"),
				Type:        datadogV2.CUSTOMATTRIBUTETYPE_NUMBER.Ptr(),
				TypeData: &datadogV2.CustomAttributeTypeData{
					Options: []datadogV2.CustomAttributeSelectOption{
						{
							Value: "us-east-1",
						},
					},
				},
			},
			Type: datadogV2.CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCustomAttributeConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementAttributeApi(apiClient)
	resp, r, err := api.UpdateCustomAttributeConfig(ctx, "case_type_id", "custom_attribute_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementAttributeApi.UpdateCustomAttributeConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementAttributeApi.UpdateCustomAttributeConfig`:\n%s\n", responseContent)
}
