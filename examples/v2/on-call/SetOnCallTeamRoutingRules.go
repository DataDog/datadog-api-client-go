// Set On-Call team routing rules returns "OK" response

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
	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	// there is a valid "escalation_policy" in the system
	EscalationPolicyDataID := os.Getenv("ESCALATION_POLICY_DATA_ID")

	body := datadogV2.TeamRoutingRulesRequest{
		Data: &datadogV2.TeamRoutingRulesRequestData{
			Attributes: &datadogV2.TeamRoutingRulesRequestDataAttributes{
				Rules: []datadogV2.TeamRoutingRulesRequestRule{
					{
						Actions: []datadogV2.RoutingRuleAction{
							datadogV2.RoutingRuleAction{
								RoutingRuleEscalationPolicyAction: &datadogV2.RoutingRuleEscalationPolicyAction{
									Type:     datadogV2.ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY,
									PolicyId: EscalationPolicyDataID,
									Urgency:  datadogV2.URGENCY_LOW.Ptr(),
								}},
						},
						Query: datadog.PtrString("tags.service:time_restrictions"),
						TimeRestriction: &datadogV2.TimeRestrictions{
							TimeZone: "Europe/Paris",
							Restrictions: []datadogV2.TimeRestriction{
								{
									EndDay:    datadogV2.WEEKDAY_MONDAY.Ptr(),
									EndTime:   datadog.PtrString("17:00:00"),
									StartDay:  datadogV2.WEEKDAY_MONDAY.Ptr(),
									StartTime: datadog.PtrString("09:00:00"),
								},
								{
									EndDay:    datadogV2.WEEKDAY_TUESDAY.Ptr(),
									EndTime:   datadog.PtrString("17:00:00"),
									StartDay:  datadogV2.WEEKDAY_TUESDAY.Ptr(),
									StartTime: datadog.PtrString("09:00:00"),
								},
							},
						},
					},
					{
						Actions: []datadogV2.RoutingRuleAction{
							datadogV2.RoutingRuleAction{
								RoutingRuleEscalationPolicyAction: &datadogV2.RoutingRuleEscalationPolicyAction{
									Type:              datadogV2.ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY,
									PolicyId:          EscalationPolicyDataID,
									Urgency:           datadogV2.URGENCY_LOW.Ptr(),
									AckTimeoutMinutes: datadog.PtrInt64(30),
									SupportHours: &datadogV2.RoutingRuleEscalationPolicyActionSupportHours{
										TimeZone: "Europe/Paris",
										Restrictions: []datadogV2.TimeRestriction{
											{
												EndDay:    datadogV2.WEEKDAY_WEDNESDAY.Ptr(),
												EndTime:   datadog.PtrString("17:00:00"),
												StartDay:  datadogV2.WEEKDAY_WEDNESDAY.Ptr(),
												StartTime: datadog.PtrString("09:00:00"),
											},
											{
												EndDay:    datadogV2.WEEKDAY_THURSDAY.Ptr(),
												EndTime:   datadog.PtrString("17:00:00"),
												StartDay:  datadogV2.WEEKDAY_THURSDAY.Ptr(),
												StartTime: datadog.PtrString("09:00:00"),
											},
										},
									},
								}},
						},
						Query: datadog.PtrString("tags.service:support_hours_and_acknowledgment_timeout"),
					},
					{
						PolicyId: datadog.PtrString(EscalationPolicyDataID),
						Query:    datadog.PtrString("tags.service:legacy_policy_definition"),
						Urgency:  datadogV2.URGENCY_LOW.Ptr(),
					},
					{
						Actions: []datadogV2.RoutingRuleAction{
							datadogV2.RoutingRuleAction{
								RoutingRuleEscalationPolicyAction: &datadogV2.RoutingRuleEscalationPolicyAction{
									Type:     datadogV2.ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY,
									PolicyId: EscalationPolicyDataID,
									Urgency:  datadogV2.URGENCY_LOW.Ptr(),
								}},
						},
						Query: datadog.PtrString(""),
					},
				},
			},
			Id:   datadog.PtrString(DdTeamDataID),
			Type: datadogV2.TEAMROUTINGRULESREQUESTDATATYPE_TEAM_ROUTING_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.SetOnCallTeamRoutingRules(ctx, DdTeamDataID, body, *datadogV2.NewSetOnCallTeamRoutingRulesOptionalParameters().WithInclude("rules"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.SetOnCallTeamRoutingRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.SetOnCallTeamRoutingRules`:\n%s\n", responseContent)
}
