// Create Jira issues for security findings returns "Created" response

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
					AssigneeId:  datadog.PtrString("f315bdaf-9ee7-4808-a9c1-99c15bf0f4d0"),
					Description: datadog.PtrString("A description of the Jira issue."),
					Fields: map[string]interface{}{
						"key1": "value",
						"key2": "['value']",
						"key3": "{'key4': 'value'}",
					},
					Priority: datadogV2.CASEPRIORITY_NOT_DEFINED.Ptr(),
					Title:    datadog.PtrString("A title for the Jira issue."),
				},
				Relationships: &datadogV2.CreateJiraIssueRequestDataRelationships{
					Findings: datadogV2.Findings{
						Data: []datadogV2.FindingData{
							{
								Id:   "ZGVmLTAwcC1pZXJ-aS0wZjhjNjMyZDNmMzRlZTgzNw==",
								Type: datadogV2.FINDINGDATATYPE_FINDINGS,
							},
						},
					},
					Project: datadogV2.CaseManagementProject{
						Data: datadogV2.CaseManagementProjectData{
							Id:   "aeadc05e-98a8-11ec-ac2c-da7ad0900001",
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
	configuration.SetUnstableOperationEnabled("v2.CreateJiraIssues", true)
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
