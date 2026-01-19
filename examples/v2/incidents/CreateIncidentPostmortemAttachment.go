// Create postmortem attachment returns "Created" response

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
	body := datadogV2.PostmortemAttachmentRequest{
		Data: datadogV2.PostmortemAttachmentRequestData{
			Attributes: datadogV2.PostmortemAttachmentRequestAttributes{
				Cells: []datadogV2.PostmortemCell{
					{
						Attributes: &datadogV2.PostmortemCellAttributes{
							Definition: &datadogV2.PostmortemCellDefinition{
								Content: datadog.PtrString(`## Incident Summary
This incident was caused by...`),
							},
						},
						Id:   datadog.PtrString("cell-1"),
						Type: datadogV2.POSTMORTEMCELLTYPE_MARKDOWN.Ptr(),
					},
				},
				Content: datadog.PtrString(`# Incident Report - IR-123
[...]`),
				PostmortemTemplateId: datadog.PtrString("93645509-874e-45c4-adfa-623bfeaead89-123"),
				Title:                datadog.PtrString("Postmortem-IR-123"),
			},
			Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentPostmortemAttachment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentPostmortemAttachment(ctx, "00000000-0000-0000-0000-000000000000", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentPostmortemAttachment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentPostmortemAttachment`:\n%s\n", responseContent)
}
