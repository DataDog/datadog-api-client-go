// Attach security finding to a Jira issue returns "OK" response

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
	body := datadogV2.AttachJiraIssueRequest{
		Data: &datadogV2.AttachJiraIssueRequestData{
			Attributes: &datadogV2.AttachJiraIssueRequestDataAttributes{
				JiraIssueUrl: "https://datadoghq-sandbox-538.atlassian.net/browse/CSMSEC-105476",
			},
			Relationships: &datadogV2.AttachJiraIssueRequestDataRelationships{
				Findings: datadogV2.Findings{
					Data: []datadogV2.FindingData{
						{
							Id:   "OTQ3NjJkMmYwMTIzMzMxNTc1Y2Q4MTA5NWU0NTBmMDl-ZjE3NjMxZWVkYzBjZGI1NDY2NWY2OGQxZDk4MDY4MmI=",
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
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.AttachJiraIssue(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.AttachJiraIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.AttachJiraIssue`:\n%s\n", responseContent)
}
