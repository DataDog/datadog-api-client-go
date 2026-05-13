// Update an AI workflow returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.UpdateAIWorkflowRequest{
		CompletedAt: datadog.PtrTime(time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)),
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
		GroupingLogic:                 datadog.PtrString("by-team"),
		MaxNumberOfEntitiesPerSession: datadog.PtrInt64(10),
		Prompt:                        datadog.PtrString("Updated prompt for the dependency upgrade"),
		Repository:                    datadog.PtrString("DataDog/datadog-agent"),
		WorkflowName:                  datadog.PtrString("Updated workflow name"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAIWorkflow", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDependencyManagementApi(apiClient)
	resp, r, err := api.UpdateAIWorkflow(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyManagementApi.UpdateAIWorkflow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DependencyManagementApi.UpdateAIWorkflow`:\n%s\n", responseContent)
}
