// Import an incident returns "CREATED" response

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
	body := datadogV2.IncidentImportRequest{
		Data: datadogV2.IncidentImportRequestData{
			Type: datadogV2.INCIDENTTYPE_INCIDENTS,
			Attributes: datadogV2.IncidentImportRequestAttributes{
				Title:      "Example-Incident",
				Visibility: datadogV2.INCIDENTIMPORTVISIBILITY_ORGANIZATION.Ptr(),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ImportIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.ImportIncident(ctx, body, *datadogV2.NewImportIncidentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ImportIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ImportIncident`:\n%s\n", responseContent)
}
