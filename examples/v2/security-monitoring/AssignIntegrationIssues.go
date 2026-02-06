// Assign or un-assign Jira issues to security findings returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.IntegrationAssignmentRequest{
		Data: datadogV2.IntegrationAssignmentDataRequest{
			Attributes: datadogV2.IntegrationAssignmentDataAttributesRequest{
				Action: datadogV2.INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTACTION_ASSIGN,
				Assignment: datadogV2.IntegrationAssignmentDataAttributesRequestAssignment{
					Jira: map[string][]string{
						"https://jira.example.com/browse/SEC-123": []string{
							"MDBjMzdhYzgyNGZkZGJiZmY0OGNmYjNiMWQ2ODY0YmR-OTc0YjMzNjM1Y2UyODA2YTEyNWQxYmNkZjhmODllNzg=",
						},
					},
				},
				Type: datadogV2.INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTTYPE_FINDINGS,
			},
			Id:   "some_id",
			Type: datadogV2.INTEGRATIONASSIGNMENTTYPE_ISSUE_ASSIGNMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AssignIntegrationIssues", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.AssignIntegrationIssues(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.AssignIntegrationIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
