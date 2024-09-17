// Create, update, and delete incident attachments returns "OK" response

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
	body := datadogV2.IncidentAttachmentUpdateRequest{
		Data: []datadogV2.IncidentAttachmentUpdateData{
			{
				Attributes: &datadogV2.IncidentAttachmentUpdateAttributes{
					IncidentAttachmentPostmortemAttributes: &datadogV2.IncidentAttachmentPostmortemAttributes{
						Attachment: datadogV2.IncidentAttachmentsPostmortemAttributesAttachmentObject{
							DocumentUrl: "https://app.datadoghq.com/notebook/123",
							Title:       "Postmortem IR-123",
						},
						AttachmentType: datadogV2.INCIDENTATTACHMENTPOSTMORTEMATTACHMENTTYPE_POSTMORTEM,
					}},
				Id:   datadog.PtrString("00000000-abcd-0002-0000-000000000000"),
				Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
			},
			{
				Attributes: &datadogV2.IncidentAttachmentUpdateAttributes{
					IncidentAttachmentLinkAttributes: &datadogV2.IncidentAttachmentLinkAttributes{
						Attachment: datadogV2.IncidentAttachmentLinkAttributesAttachmentObject{
							DocumentUrl: "https://www.example.com/webstore-failure-runbook",
							Title:       "Runbook for webstore service failures",
						},
						AttachmentType: datadogV2.INCIDENTATTACHMENTLINKATTACHMENTTYPE_LINK,
					}},
				Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
			},
			{
				Id:   datadog.PtrString("00000000-abcd-0003-0000-000000000000"),
				Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentAttachments", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentAttachments(ctx, "incident_id", body, *datadogV2.NewUpdateIncidentAttachmentsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentAttachments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentAttachments`:\n%s\n", responseContent)
}
