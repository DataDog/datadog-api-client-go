// Render an incident template returns "OK" response

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
	body := datadogV2.IncidentRenderTemplateRequest{
		Data: datadogV2.IncidentRenderTemplateDataRequest{
			Attributes: datadogV2.IncidentRenderTemplateDataAttributesRequest{
				Content:           "Incident INC-123 is SEV-1.",
				DatetimeFormat:    datadog.PtrString("2006-01-02T15:04:05Z07:00"),
				Timezone:          datadog.PtrString("America/New_York"),
				ValidateLinks:     *datadog.NewNullableBool(datadog.PtrBool(false)),
				ValidateVariables: *datadog.NewNullableBool(datadog.PtrBool(false)),
			},
			Type: datadogV2.INCIDENTRENDEREDTEMPLATETYPE_RENDERED_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RenderIncidentTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.RenderIncidentTemplate(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.RenderIncidentTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.RenderIncidentTemplate`:\n%s\n", responseContent)
}
