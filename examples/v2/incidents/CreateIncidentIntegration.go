// Create an incident integration metadata returns "CREATED" response

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

	body := datadogV2.IncidentIntegrationMetadataCreateRequest{
		Data: datadogV2.IncidentIntegrationMetadataCreateData{
			Attributes: datadogV2.IncidentIntegrationMetadataAttributes{
				IncidentId:      datadog.PtrString(IncidentDataID),
				IntegrationType: 1,
				Metadata: datadogV2.IncidentIntegrationMetadataMetadata{
					SlackIntegrationMetadata: &datadogV2.SlackIntegrationMetadata{
						Channels: []datadogV2.SlackIntegrationMetadataChannelItem{
							{
								ChannelId:   "C0123456789",
								ChannelName: "#new-channel",
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
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentIntegration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentIntegration(ctx, IncidentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentIntegration`:\n%s\n", responseContent)
}
