// Attach security findings to a ServiceNow ticket returns "OK" response

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
	body := datadogV2.AttachServiceNowTicketRequest{
		Data: datadogV2.AttachServiceNowTicketRequestData{
			Attributes: datadogV2.AttachServiceNowTicketRequestDataAttributes{
				ServicenowTicketUrl: "https://example.service-now.com/now/nav/ui/classic/params/target/incident.do?sys_id=abcdef0123456789abcdef0123456789",
			},
			Relationships: datadogV2.AttachServiceNowTicketRequestDataRelationships{
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
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AttachServiceNowTicket", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.AttachServiceNowTicket(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.AttachServiceNowTicket`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.AttachServiceNowTicket`:\n%s\n", responseContent)
}
