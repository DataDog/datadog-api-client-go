// Create a page from an incident returns "OK" response

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
	body := datadogV2.IncidentCreatePageFromIncidentRequest{
		Data: datadogV2.IncidentCreatePageFromIncidentData{
			Attributes: datadogV2.IncidentCreatePageAttributes{
				Description: *datadog.NewNullableString(datadog.PtrString("Page created for incident response")),
				Services: *datadog.NewNullableList(&[]string{
					"web-service",
					"api-service",
				}),
				Tags: *datadog.NewNullableList(&[]string{
					"urgent",
					"production",
				}),
				Target: datadogV2.IncidentPageTarget{
					Identifier: "team-handle",
					Type:       datadogV2.INCIDENTPAGETARGETTYPE_TEAM_HANDLE,
				},
				Title: "Incident Response Page",
			},
			Type: datadogV2.INCIDENTPAGETYPE_PAGE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreatePageFromIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreatePageFromIncident(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreatePageFromIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreatePageFromIncident`:\n%s\n", responseContent)
}
