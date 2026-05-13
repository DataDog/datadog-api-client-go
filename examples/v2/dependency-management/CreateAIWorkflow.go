// Create an AI workflow returns "Created" response

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
	body := datadogV2.CreateAIWorkflowRequest{
		Entities: [][]datadogV2.Entity{
			{
				{
					EntityKind:      datadog.PtrString("service"),
					EntityName:      "my-service",
					EntityNamespace: datadog.PtrString("default"),
					EntityTeam:      datadog.PtrString("platform"),
				},
			},
		},
		FilteringLogic:                new(interface{}),
		GroupingLogic:                 "by-service",
		IdpCampaignId:                 "campaign-abc123",
		MaxNumberOfEntitiesPerSession: 5,
		Prompt:                        "Upgrade the lodash dependency to version 4.17.21",
		Repository:                    "DataDog/datadog-agent",
		User:                          "john.doe@example.com",
		WorkflowName:                  "Upgrade lodash to 4.17.21",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAIWorkflow", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDependencyManagementApi(apiClient)
	resp, r, err := api.CreateAIWorkflow(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyManagementApi.CreateAIWorkflow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DependencyManagementApi.CreateAIWorkflow`:\n%s\n", responseContent)
}
