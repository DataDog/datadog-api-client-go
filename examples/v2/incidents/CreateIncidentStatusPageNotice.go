// Publish an incident status page notice returns "OK" response

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
	body := datadogV2.IncidentStatusPageNoticeCreateRequest{
		Data: datadogV2.IncidentStatusPageNoticeCreateData{
			Attributes: datadogV2.IncidentStatusPageNoticeCreateDataAttributes{
				Components: map[string]string{
					"component_1": "degraded_performance",
				},
				Message: datadog.PtrString("We are investigating reports of elevated error rates."),
				Status:  datadog.PtrString("investigating"),
				Title:   datadog.PtrString("Service degradation detected."),
			},
			Type: datadogV2.INCIDENTSTATUSPAGENOTICEINTEGRATIONTYPE_INCIDENT_INTEGRATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentStatusPageNotice", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentStatusPageNotice(ctx, "incident_id", uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewCreateIncidentStatusPageNoticeOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentStatusPageNotice`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentStatusPageNotice`:\n%s\n", responseContent)
}
