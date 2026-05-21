// Create an incident timeline entry returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.IncidentTimelineEntryRequest{
		Data: datadogV2.IncidentTimelineEntryDataRequest{
			Attributes: datadogV2.IncidentTimelineEntryDataAttributesRequest{
				CellType: datadogV2.INCIDENTTIMELINECELLTYPE_MARKDOWN,
				Content: datadogV2.IncidentTimelineEntryContent{
					Message: datadog.PtrString("Investigating the issue."),
				},
				DisplayTime: datadog.PtrTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				Important:   datadog.PtrBool(false),
			},
			Type: datadogV2.INCIDENTTIMELINEENTRYTYPE_INCIDENT_TIMELINE_CELLS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentTimelineEntry", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentTimelineEntry(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentTimelineEntry`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentTimelineEntry`:\n%s\n", responseContent)
}
