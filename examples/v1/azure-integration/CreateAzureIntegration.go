// Create an Azure integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.AzureAccount{
		AppServicePlanFilters: datadog.PtrString("key:value,filter:example"),
		Automute:              datadog.PtrBool(true),
		ClientId:              datadog.PtrString("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"),
		ClientSecret:          datadog.PtrString("TestingRh2nx664kUy5dIApvM54T4AtO"),
		ContainerAppFilters:   datadog.PtrString("key:value,filter:example"),
		CspmEnabled:           datadog.PtrBool(true),
		CustomMetricsEnabled:  datadog.PtrBool(true),
		Errors: []string{
			"*",
		},
		HostFilters:               datadog.PtrString("key:value,filter:example"),
		NewClientId:               datadog.PtrString("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"),
		NewTenantName:             datadog.PtrString("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"),
		ResourceCollectionEnabled: datadog.PtrBool(true),
		TenantName:                datadog.PtrString("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAzureIntegrationApi(apiClient)
	resp, r, err := api.CreateAzureIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.CreateAzureIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.CreateAzureIntegration`:\n%s\n", responseContent)
}
