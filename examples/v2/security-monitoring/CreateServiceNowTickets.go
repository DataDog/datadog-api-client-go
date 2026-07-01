// Create ServiceNow tickets for security findings returns "Created" response

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
	body := datadogV2.CreateServiceNowTicketRequestArray{
		Data: []datadogV2.CreateServiceNowTicketRequestData{
			{
				Attributes: &datadogV2.CreateServiceNowTicketRequestDataAttributes{
					AssigneeId:  datadog.PtrString("f315bdaf-9ee7-4808-a9c1-99c15bf0f4d0"),
					Description: datadog.PtrString("A description of the ServiceNow ticket."),
					Priority:    datadogV2.CASEPRIORITY_NOT_DEFINED.Ptr(),
					Title:       datadog.PtrString("A title for the ServiceNow ticket."),
				},
				Relationships: datadogV2.CreateServiceNowTicketRequestDataRelationships{
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
				Type: datadogV2.SERVICENOWTICKETSDATATYPE_SERVICENOW_TICKETS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateServiceNowTickets(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateServiceNowTickets`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateServiceNowTickets`:\n%s\n", responseContent)
}
