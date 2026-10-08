// Send AI tool user activity returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AIImpactUserActivityRequest{
		Data: []datadogV2.AIImpactUserActivityData{
			{
				Attributes: datadogV2.AIImpactUserActivityAttributes{
					Day:      "2026-05-26",
					IsActive: true,
					Models: []string{
						"claude-sonnet-4.5",
						"gpt-5",
					},
					Tools: []string{
						"Claude Code",
						"Cursor",
					},
					UserEmail: "user@example.com",
				},
				Type: datadogV2.AIIMPACTUSERACTIVITYTYPE_AI_IMPACT_USER_ACTIVITY,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIImpactApi(apiClient)
	r, err := api.CreateAIImpactUserActivity(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIImpactApi.CreateAIImpactUserActivity`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
