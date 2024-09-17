// Create or update service definition returns "CREATED" response

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
	body := datadogV2.ServiceDefinitionsCreateRequest{
		ServiceDefinitionV2Dot2: &datadogV2.ServiceDefinitionV2Dot2{
			Application: datadog.PtrString("my-app"),
			CiPipelineFingerprints: []string{
				"j88xdEy0J5lc",
				"eZ7LMljCk8vo",
			},
			Contacts: []datadogV2.ServiceDefinitionV2Dot2Contact{
				{
					Contact: "https://teams.microsoft.com/myteam",
					Name:    datadog.PtrString("My team channel"),
					Type:    "slack",
				},
			},
			DdService:   "my-service",
			Description: datadog.PtrString("My service description"),
			Extensions: map[string]interface{}{
				"myorg/extension": "extensionValue",
			},
			Integrations: &datadogV2.ServiceDefinitionV2Dot2Integrations{
				Opsgenie: &datadogV2.ServiceDefinitionV2Dot2Opsgenie{
					Region:     datadogV2.SERVICEDEFINITIONV2DOT2OPSGENIEREGION_US.Ptr(),
					ServiceUrl: "https://my-org.opsgenie.com/service/123e4567-e89b-12d3-a456-426614174000",
				},
				Pagerduty: &datadogV2.ServiceDefinitionV2Dot2Pagerduty{
					ServiceUrl: datadog.PtrString("https://my-org.pagerduty.com/service-directory/PMyService"),
				},
			},
			Languages: []string{
				"dotnet",
				"go",
				"java",
				"js",
				"php",
				"python",
				"ruby",
				"c++",
			},
			Lifecycle: datadog.PtrString("sandbox"),
			Links: []datadogV2.ServiceDefinitionV2Dot2Link{
				{
					Name:     "Runbook",
					Provider: datadog.PtrString("Github"),
					Type:     "runbook",
					Url:      "https://my-runbook",
				},
			},
			SchemaVersion: datadogV2.SERVICEDEFINITIONV2DOT2VERSION_V2_2,
			Tags: []string{
				"my:tag",
				"service:tag",
			},
			Team: datadog.PtrString("my-team"),
			Tier: datadog.PtrString("High"),
			Type: datadogV2.SERVICEDEFINITIONV2DOT2TYPE_WEB.Ptr(),
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceDefinitionApi(apiClient)
	resp, r, err := api.CreateOrUpdateServiceDefinitions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceDefinitionApi.CreateOrUpdateServiceDefinitions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceDefinitionApi.CreateOrUpdateServiceDefinitions`:\n%s\n", responseContent)
}
