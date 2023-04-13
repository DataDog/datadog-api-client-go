// Create an incident attachment returns "OK" response

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

	body := datadogV2.IncidentAttachmentUpdateRequest{
		Data: []datadogV2.IncidentAttachmentUpdateData{
			{
				Type: datadogV2.INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
				Attributes: &datadogV2.IncidentAttachmentUpdateAttributes{
					IncidentAttachmentLinkAttributes: &datadogV2.IncidentAttachmentLinkAttributes{
						AttachmentType: datadogV2.INCIDENTATTACHMENTLINKATTACHMENTTYPE_LINK,
						Attachment: datadogV2.IncidentAttachmentLinkAttributesAttachmentObject{
							DocumentUrl: "https://www.example.com/doc",
							Title:       "Example-Incident",
						},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentAttachments", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentAttachments(ctx, IncidentDataID, body, *datadogV2.NewUpdateIncidentAttachmentsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentAttachments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentAttachments`:\n%s\n", responseContent)
}
