// Update an Azure integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureIntegrationApi.UpdateAzureIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.UpdateAzureIntegration`:\n%s\n", responseContent)
}
