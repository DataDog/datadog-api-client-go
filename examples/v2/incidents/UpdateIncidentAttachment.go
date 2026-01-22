// Update incident attachment returns "OK" response

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

	// there is a valid "incident_attachment" in the system
	IncidentAttachmentDataID := os.Getenv("INCIDENT_ATTACHMENT_DATA_ID")

	body := datadogV2.PatchAttachmentRequest{
		Data: &datadogV2.PatchAttachmentRequestData{
			Attributes: &datadogV2.PatchAttachmentRequestDataAttributes{
				Attachment: &datadogV2.PatchAttachmentRequestDataAttributesAttachment{
					DocumentUrl: datadog.PtrString("https://app.datadoghq.com/notebook/124/Example-Incident"),
					Title:       datadog.PtrString("Example-Incident"),
				},
			},
			Id:   datadog.PtrString(IncidentAttachmentDataID),
			Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentAttachment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentAttachment(ctx, IncidentDataID, IncidentAttachmentDataID, body, *datadogV2.NewUpdateIncidentAttachmentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentAttachment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentAttachment`:\n%s\n", responseContent)
}
