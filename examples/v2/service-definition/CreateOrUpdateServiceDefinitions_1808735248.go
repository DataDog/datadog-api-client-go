// Create or update service definition using schema v2 returns "CREATED" response

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
		ServiceDefinitionV2: &datadogV2.ServiceDefinitionV2{
			Contacts: []datadogV2.ServiceDefinitionV2Contact{
				datadogV2.ServiceDefinitionV2Contact{
					ServiceDefinitionV2Email: &datadogV2.ServiceDefinitionV2Email{
						Contact: "contact@datadoghq.com",
						Name:    datadog.PtrString("Team Email"),
						Type:    datadogV2.SERVICEDEFINITIONV2EMAILTYPE_EMAIL,
					}},
			},
			DdService: "service-exampleservicedefinition",
			DdTeam:    datadog.PtrString("my-team"),
			Docs: []datadogV2.ServiceDefinitionV2Doc{
				{
					Name:     "Architecture",
					Provider: datadog.PtrString("google drive"),
					Url:      "https://gdrive/mydoc",
				},
			},
			Extensions: map[string]interface{}{
				"myorgextension": "extensionvalue",
			},
			Integrations: &datadogV2.ServiceDefinitionV2Integrations{
				Opsgenie: &datadogV2.ServiceDefinitionV2Opsgenie{
					Region:     datadogV2.SERVICEDEFINITIONV2OPSGENIEREGION_US.Ptr(),
					ServiceUrl: "https://my-org.opsgenie.com/service/123e4567-e89b-12d3-a456-426614174000",
				},
				Pagerduty: datadog.PtrString("https://my-org.pagerduty.com/service-directory/PMyService"),
			},
			Links: []datadogV2.ServiceDefinitionV2Link{
				{
					Name: "Runbook",
					Type: datadogV2.SERVICEDEFINITIONV2LINKTYPE_RUNBOOK,
					Url:  "https://my-runbook",
				},
			},
			Repos: []datadogV2.ServiceDefinitionV2Repo{
				{
					Name:     "Source Code",
					Provider: datadog.PtrString("GitHub"),
					Url:      "https://github.com/DataDog/schema",
				},
			},
			SchemaVersion: datadogV2.SERVICEDEFINITIONV2VERSION_V2,
			Tags: []string{
				"my:tag",
				"service:tag",
			},
			Team: datadog.PtrString("my-team"),
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
