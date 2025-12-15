// Create case for security finding returns "Created" response

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
	body := datadogV2.CreateCaseRequestArray{
		Data: []datadogV2.CreateCaseRequestData{
			{
				Attributes: &datadogV2.CreateCaseRequestDataAttributes{
					Title:       datadog.PtrString("A title"),
					Description: datadog.PtrString("A description"),
				},
				Relationships: &datadogV2.CreateCaseRequestDataRelationships{
					Findings: datadogV2.Findings{
						Data: []datadogV2.FindingData{
							{
								Id:   "YjdhNDM3N2QyNTFjYmUwYTY3NDdhMTg0YTk2Yjg5MDl-ZjNmMzAwOTFkZDNhNGQzYzI0MzgxNTk4MjRjZmE2NzE=",
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
				Type: datadogV2.CASEDATATYPE_CASES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateCases(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateCases`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateCases`:\n%s\n", responseContent)
}
