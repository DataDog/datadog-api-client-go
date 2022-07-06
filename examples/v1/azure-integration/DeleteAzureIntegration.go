// Delete an Azure integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.AzureAccount{
		Automute:     datadog.PtrBool(true),
		ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
		ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
		Errors: []string{
			"*",
		},
		HostFilters:   datadog.PtrString("key:value,filter:example"),
		NewClientId:   datadog.PtrString("new1c7f6-1234-5678-9101-3fcbf464test"),
		NewTenantName: datadog.PtrString("new1c44-1234-5678-9101-cc00736ftest"),
		TenantName:    datadog.PtrString("testc44-1234-5678-9101-cc00736ftest"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.AzureIntegrationApi(apiClient)
	resp, r, err := api.DeleteAzureIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.DeleteAzureIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.DeleteAzureIntegration`:\n%s\n", responseContent)
}
