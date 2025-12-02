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
				Type:       datadogV2.JIRAISSUESDATATYPE_JIRA_ISSUES,
				Attributes: &datadogV2.CreateJiraIssueRequestDataAttributes{},
				Relationships: &datadogV2.CreateJiraIssueRequestDataRelationships{
					Case: datadogV2.CreateJiraIssueRequestDataRelationshipsCase{
						Data: datadogV2.CreateJiraIssueRequestDataRelationshipsCaseData{
							Type: datadogV2.CASEDATATYPE_CASES,
							Id:   "53e242c6-a7d6-46ad-9680-b8d14753f716",
						},
					},
				},
			},
			{
				Type:       datadogV2.JIRAISSUESDATATYPE_JIRA_ISSUES,
				Attributes: &datadogV2.CreateJiraIssueRequestDataAttributes{},
				Relationships: &datadogV2.CreateJiraIssueRequestDataRelationships{
					Case: datadogV2.CreateJiraIssueRequestDataRelationshipsCase{
						Data: datadogV2.CreateJiraIssueRequestDataRelationshipsCaseData{
							Type: datadogV2.CASEDATATYPE_CASES,
							Id:   "195772b2-1f53-41d2-b81e-48c8e6c21d33",
						},
					},
				},
			},
		},
		Included: []datadogV2.CreateJiraIssueRequestArrayIncluded{
			datadogV2.CreateJiraIssueRequestArrayIncluded{
				CreateCaseRequestData: &datadogV2.CreateCaseRequestData{
					Type: datadogV2.CASEDATATYPE_CASES,
					Attributes: &datadogV2.CreateCaseRequestDataAttributes{
						Title:       datadog.PtrString("A title"),
						Description: datadog.PtrString("A description"),
					},
					Relationships: &datadogV2.CreateCaseRequestDataRelationships{
						Project: datadogV2.CaseManagementProject{
							Data: datadogV2.CaseManagementProjectData{
								Type: datadogV2.CASEMANAGEMENTPROJECTDATATYPE_PROJECTS,
								Id:   "959a6f71-bac8-4027-b1d3-2264f569296f",
							},
						},
						Findings: datadogV2.Findings{
							Data: []datadogV2.FindingData{
								{
									Type: datadogV2.FINDINGDATATYPE_FINDINGS,
									Id:   "OTQ3NjJkMmYwMTIzMzMxNTc1Y2Q4MTA5NWU0NTBmMDl-ZjE3NjMxZWVkYzBjZGI1NDY2NWY2OGQxZDk4MDY4MmI=",
								},
							},
						},
					},
					Id: datadog.PtrString("53e242c6-a7d6-46ad-9680-b8d14753f716"),
				}},
			datadogV2.CreateJiraIssueRequestArrayIncluded{
				CreateCaseRequestData: &datadogV2.CreateCaseRequestData{
					Type: datadogV2.CASEDATATYPE_CASES,
					Attributes: &datadogV2.CreateCaseRequestDataAttributes{
						Title:       datadog.PtrString("A title"),
						Description: datadog.PtrString("A description"),
					},
					Relationships: &datadogV2.CreateCaseRequestDataRelationships{
						Project: datadogV2.CaseManagementProject{
							Data: datadogV2.CaseManagementProjectData{
								Type: datadogV2.CASEMANAGEMENTPROJECTDATATYPE_PROJECTS,
								Id:   "959a6f71-bac8-4027-b1d3-2264f569296f",
							},
						},
						Findings: datadogV2.Findings{
							Data: []datadogV2.FindingData{
								{
									Type: datadogV2.FINDINGDATATYPE_FINDINGS,
									Id:   "MTNjN2ZmYWMzMDIxYmU1ZDFiZDRjNWUwN2I1NzVmY2F-YTA3MzllMTUzNWM3NmEyZjdiNzEzOWM5YmViZTMzOGM=",
								},
							},
						},
					},
					Id: datadog.PtrString("195772b2-1f53-41d2-b81e-48c8e6c21d33"),
				}},
			datadogV2.CreateJiraIssueRequestArrayIncluded{
				CaseManagementProjectData: &datadogV2.CaseManagementProjectData{
					Type: datadogV2.CASEMANAGEMENTPROJECTDATATYPE_PROJECTS,
					Id:   "959a6f71-bac8-4027-b1d3-2264f569296f",
				}},
			datadogV2.CreateJiraIssueRequestArrayIncluded{
				FindingData: &datadogV2.FindingData{
					Type: datadogV2.FINDINGDATATYPE_FINDINGS,
					Id:   "OTQ3NjJkMmYwMTIzMzMxNTc1Y2Q4MTA5NWU0NTBmMDl-ZjE3NjMxZWVkYzBjZGI1NDY2NWY2OGQxZDk4MDY4MmI=",
				}},
			datadogV2.CreateJiraIssueRequestArrayIncluded{
				FindingData: &datadogV2.FindingData{
					Type: datadogV2.FINDINGDATATYPE_FINDINGS,
					Id:   "MTNjN2ZmYWMzMDIxYmU1ZDFiZDRjNWUwN2I1NzVmY2F-YTA3MzllMTUzNWM3NmEyZjdiNzEzOWM5YmViZTMzOGM=",
				}},
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
