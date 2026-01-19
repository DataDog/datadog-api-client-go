// Update Jira issue template returns "OK" response

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
	body := datadogV2.JiraIssueTemplateUpdateRequest{
		Data: datadogV2.JiraIssueTemplateUpdateRequestData{
			Attributes: datadogV2.JiraIssueTemplateUpdateRequestAttributes{
				Fields: map[string]interface{}{
					"description": "{'payload': 'Updated Description', 'type': 'json'}",
				},
				Name: datadog.PtrString("test_template_updated"),
			},
			Type: datadogV2.JIRAISSUETEMPLATETYPE_JIRA_ISSUE_TEMPLATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateJiraIssueTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewJiraIntegrationApi(apiClient)
	resp, r, err := api.UpdateJiraIssueTemplate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationApi.UpdateJiraIssueTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationApi.UpdateJiraIssueTemplate`:\n%s\n", responseContent)
}
