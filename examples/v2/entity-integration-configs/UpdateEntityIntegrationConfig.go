// Create or update entity integration configuration returns "OK" response

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
	body := datadogV2.EntityIntegrationConfigRequest{
		Data: datadogV2.EntityIntegrationConfigRequestData{
			Attributes: datadogV2.EntityIntegrationConfigRequestAttributes{
				Config: map[string]interface{}{
					"enabled_repos": "[{'github_org_name': 'myorg', 'hostname': 'github.com', 'repo_name': 'myrepo'}]",
				},
			},
			Type: datadogV2.ENTITYINTEGRATIONCONFIGREQUESTTYPE_ENTITY_INTEGRATION_CONFIG_REQUESTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateEntityIntegrationConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEntityIntegrationConfigsApi(apiClient)
	resp, r, err := api.UpdateEntityIntegrationConfig(ctx, "github", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationConfigsApi.UpdateEntityIntegrationConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EntityIntegrationConfigsApi.UpdateEntityIntegrationConfig`:\n%s\n", responseContent)
}
