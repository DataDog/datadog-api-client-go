// Update an incident todo returns "OK" response

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

	// the "incident" has an "incident_todo"
	IncidentTodoDataID := os.Getenv("INCIDENT_TODO_DATA_ID")

	body := datadogV2.IncidentTodoPatchRequest{
		Data: datadogV2.IncidentTodoPatchData{
			Attributes: datadogV2.IncidentTodoAttributes{
				Assignees: []datadogV2.IncidentTodoAssignee{
					datadogV2.IncidentTodoAssignee{
						IncidentTodoAssigneeHandle: datadog.PtrString("@test.user@test.com")},
				},
				Content:   "Restore lost data.",
				Completed: *datadog.NewNullableString(datadog.PtrString("2023-03-06T22:00:00.000000+00:00")),
				DueDate:   *datadog.NewNullableString(datadog.PtrString("2023-07-10T05:00:00.000000+00:00")),
			},
			Type: datadogV2.INCIDENTTODOTYPE_INCIDENT_TODOS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentTodo", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentTodo(ctx, IncidentDataID, IncidentTodoDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentTodo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentTodo`:\n%s\n", responseContent)
}
