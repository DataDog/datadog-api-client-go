// Update an incident status page notice returns "OK" response

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
	body := datadogV2.IncidentStatusPageNoticeUpdateRequest{
		Data: datadogV2.IncidentStatusPageNoticeUpdateData{
			Attributes: datadogV2.IncidentStatusPageNoticeUpdateDataAttributes{
				Message: datadog.PtrString("The issue has been resolved."),
				Status:  datadog.PtrString("resolved"),
				Title:   datadog.PtrString("Service degradation resolved."),
			},
			Type: datadogV2.INCIDENTSTATUSPAGENOTICEINTEGRATIONTYPE_INCIDENT_INTEGRATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentStatusPageNotice", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentStatusPageNotice(ctx, "incident_id", uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewUpdateIncidentStatusPageNoticeOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentStatusPageNotice`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentStatusPageNotice`:\n%s\n", responseContent)
}
