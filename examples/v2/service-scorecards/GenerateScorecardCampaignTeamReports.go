// Generate team-specific campaign reports returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.GenerateCampaignTeamReportsRequest{
		Data: datadogV2.GenerateCampaignTeamReportsRequestData{
			Attributes: datadogV2.GenerateCampaignTeamReportsRequestAttributes{
				EntityOwners: []datadogV2.EntityOwnerDestination{
					{
						Slack: datadogV2.SlackRoutingOptions{
							ChannelId:   datadog.PtrString("C024BDQ4N"),
							WorkspaceId: datadog.PtrString("T024BDQ4N"),
						},
						TeamId: "550e8400-e29b-41d4-a716-446655440000",
					},
				},
			},
			Type: datadogV2.GENERATECAMPAIGNTEAMREPORTSREQUESTDATATYPE_CAMPAIGN_TEAM_REPORT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GenerateScorecardCampaignTeamReports", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	r, err := api.GenerateScorecardCampaignTeamReports(ctx, "c10ODp0VCrrIpXmz", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.GenerateScorecardCampaignTeamReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
