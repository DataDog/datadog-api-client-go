// Create Cloud Cost Management Azure configs returns "OK" response

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
	body := datadogV2.AzureUCConfigPostRequest{
		Data: datadogV2.AzureUCConfigPostData{
			Attributes: datadogV2.AzureUCConfigPostRequestAttributes{
				AccountId: "1234abcd-1234-abcd-1234-1234abcd1234",
				ActualBillConfig: datadogV2.BillConfig{
					ExportName:       "dd-actual-export",
					ExportPath:       "dd-export-path",
					StorageAccount:   "dd-storage-account",
					StorageContainer: "dd-storage-container",
				},
				AmortizedBillConfig: datadogV2.BillConfig{
					ExportName:       "dd-actual-export",
					ExportPath:       "dd-export-path",
					StorageAccount:   "dd-storage-account",
					StorageContainer: "dd-storage-container",
				},
				ClientId:  "1234abcd-1234-abcd-1234-1234abcd1234",
				IsEnabled: datadog.PtrBool(true),
				Scope:     "subscriptions/1234abcd-1234-abcd-1234-1234abcd1234",
			},
			Type: datadogV2.AZUREUCCONFIGPOSTREQUESTTYPE_AZURE_UC_CONFIG_POST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateCostAzureUCConfigs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateCostAzureUCConfigs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateCostAzureUCConfigs`:\n%s\n", responseContent)
}
