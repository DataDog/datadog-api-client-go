// Create a cost recommendation notification rule returns "Created" response

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
	body := datadogV2.CostRecommendationNotificationRuleCreateRequest{
		Data: datadogV2.CostRecommendationNotificationRuleCreateData{
			Attributes: datadogV2.CostRecommendationNotificationRuleCreateAttributes{
				ChannelId:             datadog.PtrString("C0456ABCD"),
				DestinationType:       datadogV2.COSTRECOMMENDATIONNOTIFICATIONRULEDESTINATIONTYPE_SLACK.Ptr(),
				Enabled:               datadog.PtrBool(true),
				Frequency:             datadogV2.COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_WEEKLY,
				IntervalExecutionDate: datadog.PtrString("2026-11-02"),
				IntervalExecutionTime: datadog.PtrString("09:30"),
				MentionTargets: []string{
					"U01ABCDE",
				},
				MessageCustomization: &datadogV2.CostRecommendationNotificationRuleMessageCustomization{
					Actions: &datadogV2.CostRecommendationNotificationRuleMessageActions{
						ViewInCcm: &datadogV2.CostRecommendationNotificationRuleMessageAction{
							Bulk:   datadog.PtrBool(true),
							PerRec: datadog.PtrBool(true),
						},
					},
					Header: &datadogV2.CostRecommendationNotificationRuleMessageCustomizationHeader{
						CustomText:       datadog.PtrString(" recommendations to review this week."),
						NotificationName: datadog.PtrBool(true),
						TotalSavings:     datadog.PtrBool(true),
					},
					IncludeTags: []string{
						"team",
					},
					Mode: datadogV2.COSTRECOMMENDATIONNOTIFICATIONRULEMESSAGECUSTOMIZATIONMODE_STRUCTURED.Ptr(),
					PerRec: &datadogV2.CostRecommendationNotificationRuleMessageCustomizationPerRec{
						Effort: datadog.PtrBool(true),
						Risk:   datadog.PtrBool(true),
						Why:    datadog.PtrBool(true),
					},
				},
				Name:                "Production recommendations",
				ScopeQuery:          datadog.PtrString("service:api AND env:prod"),
				Timezone:            "America/New_York",
				WeeklyExecutionDay:  datadog.PtrString("monday"),
				WeeklyExecutionTime: datadog.PtrString("09:30"),
				WorkspaceId:         datadog.PtrString("T0123ABCD"),
			},
			Type: datadogV2.COSTRECOMMENDATIONNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCostRecommendationNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateCostRecommendationNotificationRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateCostRecommendationNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateCostRecommendationNotificationRule`:\n%s\n", responseContent)
}
