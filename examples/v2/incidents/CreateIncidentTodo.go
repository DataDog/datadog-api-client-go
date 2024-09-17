// Create an incident todo returns "CREATED" response

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
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	body := datadogV2.IncidentTodoCreateRequest{
		Data: datadogV2.IncidentTodoCreateData{
			Attributes: datadogV2.IncidentTodoAttributes{
				Assignees: []datadogV2.IncidentTodoAssignee{
					datadogV2.IncidentTodoAssignee{
						IncidentTodoAssigneeHandle: datadog.PtrString("@test.user@test.com")},
				},
				Content: "Restore lost data.",
			},
			Type: datadogV2.INCIDENTTODOTYPE_INCIDENT_TODOS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentTodo", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentTodo(ctx, IncidentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentTodo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentTodo`:\n%s\n", responseContent)
}
