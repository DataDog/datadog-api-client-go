// Update a campaign returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UpdateCampaignRequest{
		Data: datadogV2.UpdateCampaignRequestData{
			Attributes: datadogV2.UpdateCampaignRequestAttributes{
				Description: datadog.PtrString("Campaign to improve security posture for Q1 2024."),
				DueDate:     datadog.PtrTime(time.Date(2024, 3, 31, 23, 59, 59, 0, time.UTC)),
				EntityScope: datadog.PtrString("kind:service AND team:platform"),
				Guidance:    datadog.PtrString("Please ensure all services pass the security requirements."),
				Key:         datadog.PtrString("q1-security-2024"),
				Name:        "Q1 Security Campaign",
				OwnerId:     "550e8400-e29b-41d4-a716-446655440000",
				RuleIds: []string{
					"q8MQxk8TCqrHnWkx",
					"r9NRyl9UDrsIoXly",
				},
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Status:    "in_progress",
			},
			Type: datadogV2.CAMPAIGNTYPE_CAMPAIGN,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateScorecardCampaign", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceScorecardsApi(apiClient)
	resp, r, err := api.UpdateScorecardCampaign(ctx, "c10ODp0VCrrIpXmz", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceScorecardsApi.UpdateScorecardCampaign`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceScorecardsApi.UpdateScorecardCampaign`:\n%s\n", responseContent)
}
