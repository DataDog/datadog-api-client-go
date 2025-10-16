// Update Azure integration host filters returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	body := datadogV1.AzureAccount{
AppServicePlanFilters: datadog.PtrString("key:value,filter:example"),
Automute: datadog.PtrBool(true),
ClientId: datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
ClientSecret: datadog.PtrString("TestingRh2nx664kUy5dIApvM54T4AtO"),
ContainerAppFilters: datadog.PtrString("key:value,filter:example"),
CspmEnabled: datadog.PtrBool(true),
CustomMetricsEnabled: datadog.PtrBool(true),
Errors: []string{
"*",
},
HostFilters: datadog.PtrString("key:value,filter:example"),
MetricsEnabled: datadog.PtrBool(true),
MetricsEnabledDefault: datadog.PtrBool(true),
NewClientId: datadog.PtrString("new1c7f6-1234-5678-9101-3fcbf464test"),
NewTenantName: datadog.PtrString("new1c44-1234-5678-9101-cc00736ftest"),
ResourceCollectionEnabled: datadog.PtrBool(true),
ResourceProviderConfigs: []datadogV1.ResourceProviderConfig{
{
MetricsEnabled: datadog.PtrBool(true),
Namespace: datadog.PtrString("Microsoft.Compute"),
},
},
TenantName: datadog.PtrString("testc44-1234-5678-9101-cc00736ftest"),
UsageMetricsEnabled: datadog.PtrBool(true),
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAzureIntegrationApi(apiClient)
	resp, r, err := api.UpdateAzureHostFilters(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureHostFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.UpdateAzureHostFilters`:\n%s\n", responseContent)
}