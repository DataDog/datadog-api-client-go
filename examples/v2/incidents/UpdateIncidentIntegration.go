// Update an existing incident integration metadata returns "OK" response

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

	// the "incident" has an "incident_integration_metadata"
	IncidentIntegrationMetadataDataID := os.Getenv("INCIDENT_INTEGRATION_METADATA_DATA_ID")

	body := datadogV2.IncidentIntegrationMetadataPatchRequest{
		Data: datadogV2.IncidentIntegrationMetadataPatchData{
			Attributes: datadogV2.IncidentIntegrationMetadataAttributes{
				IncidentId:      datadog.PtrString(IncidentDataID),
				IntegrationType: 1,
				Metadata: datadogV2.IncidentIntegrationMetadataMetadata{
					SlackIntegrationMetadata: &datadogV2.SlackIntegrationMetadata{
						Channels: []datadogV2.SlackIntegrationMetadataChannelItem{
							{
								ChannelId:   "C0123456789",
								ChannelName: "#updated-channel-name",
								TeamId:      datadog.PtrString("T01234567"),
								RedirectUrl: "https://slack.com/app_redirect?channel=C0123456789&team=T01234567",
							},
						},
					}},
			},
			Type: datadogV2.INCIDENTINTEGRATIONMETADATATYPE_INCIDENT_INTEGRATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentIntegration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentIntegration(ctx, IncidentDataID, IncidentIntegrationMetadataDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentIntegration`:\n%s\n", responseContent)
}
