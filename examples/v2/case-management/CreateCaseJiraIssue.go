// Create Jira issue for case returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.JiraIssueCreateRequest{
		Data: datadogV2.JiraIssueCreateData{
			Attributes: datadogV2.JiraIssueCreateAttributes{
				Fields:        map[string]interface{}{},
				IssueTypeId:   "10001",
				JiraAccountId: "1234",
				ProjectId:     "5678",
			},
			Type: datadogV2.JIRAISSUERESOURCETYPE_ISSUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCaseJiraIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.CreateCaseJiraIssue(ctx, "case_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCaseJiraIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
