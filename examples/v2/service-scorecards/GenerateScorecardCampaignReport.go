// Generate campaign report returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.GenerateCampaignReportRequest{
		Data: datadogV2.GenerateCampaignReportRequestData{
			Attributes: datadogV2.GenerateCampaignReportRequestAttributes{
				Slack: datadogV2.SlackRoutingOptions{
					ChannelId:     datadog.PtrString("C024BDQ4N"),
					ChannelName:   datadog.PtrString("service-scorecards"),
					WorkspaceId:   datadog.PtrString("T024BDQ4N"),
					WorkspaceName: datadog.PtrString("datadog-workspace"),
				},
			},
			Type: datadogV2.GENERATECAMPAIGNREPORTREQUESTDATATYPE_CAMPAIGN_REPORT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GenerateScorecardCampaignReport", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	r, err := api.GenerateScorecardCampaignReport(ctx, "c10ODp0VCrrIpXmz", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.GenerateScorecardCampaignReport`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
