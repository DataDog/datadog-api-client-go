// Attach security findings to a Linear issue returns "OK" response

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
	body := datadogV2.AttachLinearIssueRequest{
		Data: datadogV2.AttachLinearIssueRequestData{
			Attributes: datadogV2.AttachLinearIssueRequestDataAttributes{
				LinearIssueUrl: "https://linear.app/your-workspace/issue/ENG-123",
			},
			Relationships: datadogV2.AttachLinearIssueRequestDataRelationships{
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
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AttachLinearIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.AttachLinearIssue(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.AttachLinearIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.AttachLinearIssue`:\n%s\n", responseContent)
}
