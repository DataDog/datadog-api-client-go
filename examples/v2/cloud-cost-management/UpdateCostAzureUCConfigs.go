// Update Cloud Cost Management Azure config returns "OK" response

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
	body := datadogV2.AzureUCConfigPatchRequest{
		Data: datadogV2.AzureUCConfigPatchData{
			Attributes: datadogV2.AzureUCConfigPatchRequestAttributes{
				IsEnabled: true,
			},
			Type: datadogV2.AZUREUCCONFIGPATCHREQUESTTYPE_AZURE_UC_CONFIG_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCostAzureUCConfigs(ctx, "100", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCostAzureUCConfigs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCostAzureUCConfigs`:\n%s\n", responseContent)
}
