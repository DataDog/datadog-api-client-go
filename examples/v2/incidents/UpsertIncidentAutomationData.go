// Create or update incident automation data returns "OK" response

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
	body := datadogV2.IncidentAutomationDataRequest{
		Data: datadogV2.IncidentAutomationDataDataRequest{
			Attributes: datadogV2.IncidentAutomationDataAttributesRequest{
				Value: "completed",
			},
			Type: datadogV2.INCIDENTAUTOMATIONDATATYPE_INCIDENTS_AUTOMATION_DATA,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertIncidentAutomationData", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpsertIncidentAutomationData(ctx, "incident_id", "key", body, *datadogV2.NewUpsertIncidentAutomationDataOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpsertIncidentAutomationData`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpsertIncidentAutomationData`:\n%s\n", responseContent)
}
