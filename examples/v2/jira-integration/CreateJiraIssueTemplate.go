// Create Jira issue template returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.JiraIssueTemplateCreateRequest{
		Data: &datadogV2.JiraIssueTemplateCreateRequestData{
			Attributes: &datadogV2.JiraIssueTemplateCreateRequestAttributes{
				Fields: map[string]interface{}{
					"description": "{'payload': 'Test', 'type': 'json'}",
				},
				IssueTypeId: datadog.PtrString("12730"),
				JiraAccount: &datadogV2.JiraIssueTemplateCreateRequestAttributesJiraAccount{
					Id: uuid.MustParse("80f16d40-1fba-486e-b1fc-983e6ca19bec"),
				},
				Name:      datadog.PtrString("test-template"),
				ProjectId: datadog.PtrString("10772"),
			},
			Type: datadogV2.JIRAISSUETEMPLATETYPE_JIRA_ISSUE_TEMPLATE.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateJiraIssueTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewJiraIntegrationApi(apiClient)
	resp, r, err := api.CreateJiraIssueTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationApi.CreateJiraIssueTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationApi.CreateJiraIssueTemplate`:\n%s\n", responseContent)
}
