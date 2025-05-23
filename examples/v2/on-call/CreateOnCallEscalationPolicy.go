// Create On-Call escalation policy returns "Created" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	// there is a valid "schedule" in the system
	ScheduleDataID := os.Getenv("SCHEDULE_DATA_ID")

	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	body := datadogV2.EscalationPolicyCreateRequest{
		Data: datadogV2.EscalationPolicyCreateRequestData{
			Attributes: datadogV2.EscalationPolicyCreateRequestDataAttributes{
				Name:                   "Example-On-Call",
				ResolvePageOnPolicyEnd: datadog.PtrBool(true),
				Retries:                datadog.PtrInt64(2),
				Steps: []datadogV2.EscalationPolicyCreateRequestDataAttributesStepsItems{
					{
						Assignment:           datadogV2.ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_DEFAULT.Ptr(),
						EscalateAfterSeconds: datadog.PtrInt64(3600),
						Targets: []datadogV2.EscalationPolicyStepTarget{
							{
								Id:   datadog.PtrString(UserDataID),
								Type: datadogV2.ESCALATIONPOLICYSTEPTARGETTYPE_USERS.Ptr(),
							},
							{
								Id:   datadog.PtrString(ScheduleDataID),
								Type: datadogV2.ESCALATIONPOLICYSTEPTARGETTYPE_SCHEDULES.Ptr(),
							},
							{
								Id:   datadog.PtrString(DdTeamDataID),
								Type: datadogV2.ESCALATIONPOLICYSTEPTARGETTYPE_TEAMS.Ptr(),
							},
						},
					},
					{
						Assignment:           datadogV2.ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_ROUND_ROBIN.Ptr(),
						EscalateAfterSeconds: datadog.PtrInt64(3600),
						Targets: []datadogV2.EscalationPolicyStepTarget{
							{
								Id:   datadog.PtrString(DdTeamDataID),
								Type: datadogV2.ESCALATIONPOLICYSTEPTARGETTYPE_TEAMS.Ptr(),
							},
						},
					},
				},
			},
			Relationships: &datadogV2.EscalationPolicyCreateRequestDataRelationships{
				Teams: &datadogV2.DataRelationshipsTeams{
					Data: []datadogV2.DataRelationshipsTeamsDataItems{
						{
							Id:   DdTeamDataID,
							Type: datadogV2.DATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
						},
					},
				},
			},
			Type: datadogV2.ESCALATIONPOLICYCREATEREQUESTDATATYPE_POLICIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateOnCallEscalationPolicy(ctx, body, *datadogV2.NewCreateOnCallEscalationPolicyOptionalParameters().WithInclude("steps.targets"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateOnCallEscalationPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateOnCallEscalationPolicy`:\n%s\n", responseContent)
}
