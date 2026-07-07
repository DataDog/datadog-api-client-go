// Create Linear issues for security findings returns "Created" response

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
	body := datadogV2.CreateLinearIssueRequestArray{
		Data: []datadogV2.CreateLinearIssueRequestData{
			{
				Attributes: &datadogV2.CreateLinearIssueRequestDataAttributes{
					AssigneeId:  datadog.PtrString("f315bdaf-9ee7-4808-a9c1-99c15bf0f4d0"),
					Description: datadog.PtrString("A description of the Linear issue."),
					LabelIds: []string{
						"a1b2c3d4-5e6f-7a8b-9c0d-1e2f3a4b5c6d",
					},
					LinearProjectId: datadog.PtrString("d4c3b2a1-6f5e-8b7a-0d9c-2f1e4a3b6c5d"),
					Priority:        datadogV2.CASEPRIORITY_NOT_DEFINED.Ptr(),
					Title:           datadog.PtrString("A title for the Linear issue."),
				},
				Relationships: &datadogV2.CreateLinearIssueRequestDataRelationships{
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
				Type: datadogV2.LINEARISSUESDATATYPE_LINEAR_ISSUES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLinearIssues", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateLinearIssues(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateLinearIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateLinearIssues`:\n%s\n", responseContent)
}
