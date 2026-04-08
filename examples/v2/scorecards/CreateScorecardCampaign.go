// Create a new campaign returns "Created" response

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
	body := datadogV2.CreateCampaignRequest{
		Data: datadogV2.CreateCampaignRequestData{
			Attributes: datadogV2.CreateCampaignRequestAttributes{
				Description: datadog.PtrString("Campaign to improve security posture for Q1 2024."),
				DueDate:     datadog.PtrTime(time.Date(2024, 3, 31, 23, 59, 59, 0, time.UTC)),
				EntityScope: datadog.PtrString("kind:service AND team:platform"),
				Guidance:    datadog.PtrString("Please ensure all services pass the security requirements."),
				Key:         "q1-security-2024",
				Name:        "Q1 Security Campaign",
				OwnerId:     "550e8400-e29b-41d4-a716-446655440000",
				RuleIds: []string{
					"q8MQxk8TCqrHnWkx",
					"r9NRyl9UDrsIoXly",
				},
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Status:    datadogV2.CAMPAIGNSTATUS_IN_PROGRESS.Ptr(),
			},
			Type: datadogV2.CAMPAIGNTYPE_CAMPAIGN,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewScorecardsApi(apiClient)
	resp, r, err := api.CreateScorecardCampaign(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScorecardsApi.CreateScorecardCampaign`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ScorecardsApi.CreateScorecardCampaign`:\n%s\n", responseContent)
}
