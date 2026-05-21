// Create an incident Jira issue returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.IncidentJiraIssueRequest{
		Data: datadogV2.IncidentJiraIssueDataRequest{
			Attributes: datadogV2.IncidentJiraIssueDataAttributesRequest{
				AccountId:   "123456",
				IssueTypeId: "10001",
				ProjectId:   "10000",
				TemplateId:  datadog.PtrUUID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
			},
			Type: datadogV2.INCIDENTJIRAISSUETYPE_INCIDENT_JIRA_ISSUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentJiraIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentJiraIssue(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentJiraIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentJiraIssue`:\n%s\n", responseContent)
}
