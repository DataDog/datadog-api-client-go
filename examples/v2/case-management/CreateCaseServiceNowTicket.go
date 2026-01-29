// Create ServiceNow ticket for case returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ServiceNowTicketCreateRequest{
		Data: datadogV2.ServiceNowTicketCreateData{
			Attributes: datadogV2.ServiceNowTicketCreateAttributes{
				AssignmentGroup: datadog.PtrString("IT Support"),
				InstanceName:    "my-instance",
			},
			Type: datadogV2.SERVICENOWTICKETRESOURCETYPE_TICKETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCaseServiceNowTicket", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.CreateCaseServiceNowTicket(ctx, "case_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCaseServiceNowTicket`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
