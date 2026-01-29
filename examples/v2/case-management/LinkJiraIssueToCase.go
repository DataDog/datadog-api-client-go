// Link existing Jira issue to case returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.JiraIssueLinkRequest{
		Data: datadogV2.JiraIssueLinkData{
			Attributes: datadogV2.JiraIssueLinkAttributes{
				JiraIssueUrl: "https://jira.example.com/browse/PROJ-123",
			},
			Type: datadogV2.JIRAISSUERESOURCETYPE_ISSUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.LinkJiraIssueToCase", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.LinkJiraIssueToCase(ctx, "case_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.LinkJiraIssueToCase`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
