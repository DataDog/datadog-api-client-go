// Create Jira issue for security finding returns "Created" response

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
	body := datadogV2.CreateJiraIssueRequestArray{
		Data: []datadogV2.CreateJiraIssueRequestData{
			{
				Attributes: &datadogV2.CreateJiraIssueRequestDataAttributes{
					Title:       datadog.PtrString("A title"),
					Description: datadog.PtrString("A description"),
				},
				Relationships: &datadogV2.CreateJiraIssueRequestDataRelationships{
					Findings: datadogV2.Findings{
						Data: []datadogV2.FindingData{
							{
								Id:   "YmNlZmJhYTcyMDU5ZDk0ZDhiNjRmNGI0NDk4MDdiNzN-MDJlMjg0NzNmYzJiODY2MzJkNjU0OTI4NmVhZTUyY2U=",
								Type: datadogV2.FINDINGDATATYPE_FINDINGS,
							},
						},
					},
					Project: datadogV2.CaseManagementProject{
						Data: datadogV2.CaseManagementProjectData{
							Id:   "959a6f71-bac8-4027-b1d3-2264f569296f",
							Type: datadogV2.CASEMANAGEMENTPROJECTDATATYPE_PROJECTS,
						},
					},
				},
				Type: datadogV2.JIRAISSUESDATATYPE_JIRA_ISSUES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateJiraIssues(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateJiraIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateJiraIssues`:\n%s\n", responseContent)
}
