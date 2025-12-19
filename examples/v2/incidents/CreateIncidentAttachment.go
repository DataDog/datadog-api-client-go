// Create incident attachment returns "Created" response

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
	body := datadogV2.CreateAttachmentRequest{
		Data: &datadogV2.CreateAttachmentRequestData{
			Attributes: &datadogV2.CreateAttachmentRequestDataAttributes{
				Attachment: &datadogV2.CreateAttachmentRequestDataAttributesAttachment{
					DocumentUrl: datadog.PtrString("https://app.datadoghq.com/notebook/123/Postmortem-IR-123"),
					Title:       datadog.PtrString("Postmortem-IR-123"),
				},
				AttachmentType: datadogV2.ATTACHMENTDATAATTRIBUTESATTACHMENTTYPE_POSTMORTEM.Ptr(),
			},
			Id:   datadog.PtrString("00000000-0000-0000-0000-000000000000"),
			Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentAttachment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentAttachment(ctx, "incident_id", body, *datadogV2.NewCreateIncidentAttachmentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentAttachment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentAttachment`:\n%s\n", responseContent)
}
