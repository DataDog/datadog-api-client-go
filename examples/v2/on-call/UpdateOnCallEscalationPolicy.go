// Update On-Call escalation policy returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "escalation_policy" in the system
	EscalationPolicyDataID := os.Getenv("ESCALATION_POLICY_DATA_ID")
	EscalationPolicyDataRelationshipsStepsData0ID := os.Getenv("ESCALATION_POLICY_DATA_RELATIONSHIPS_STEPS_DATA_0_ID")


	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")


	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")


	body := datadogV2.EscalationPolicyUpdateRequest{
Data: datadogV2.EscalationPolicyUpdateRequestData{
Attributes: datadogV2.EscalationPolicyUpdateRequestDataAttributes{
Name: "Example-On-Call-updated",
ResolvePageOnPolicyEnd: datadog.PtrBool(false),
Retries: datadog.PtrInt64(0),
Steps: []datadogV2.EscalationPolicyUpdateRequestDataAttributesStepsItems{
{
Assignment: datadogV2.ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_DEFAULT.Ptr(),
EscalateAfterSeconds: datadog.PtrInt64(3600),
Id: datadog.PtrString(EscalationPolicyDataRelationshipsStepsData0ID),
Targets: []datadogV2.EscalationPolicyStepTarget{
{
Id: datadog.PtrString(UserDataID),
Type: datadogV2.ESCALATIONPOLICYSTEPTARGETTYPE_USERS.Ptr(),
},
},
},
},
},
Id: EscalationPolicyDataID,
Relationships: &datadogV2.EscalationPolicyUpdateRequestDataRelationships{
Teams: &datadogV2.DataRelationshipsTeams{
Data: []datadogV2.DataRelationshipsTeamsDataItems{
{
Id: DdTeamDataID,
Type: datadogV2.DATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
},
},
},
},
Type: datadogV2.ESCALATIONPOLICYUPDATEREQUESTDATATYPE_POLICIES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.UpdateOnCallEscalationPolicy(ctx, EscalationPolicyDataID, body, *datadogV2.NewUpdateOnCallEscalationPolicyOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.UpdateOnCallEscalationPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.UpdateOnCallEscalationPolicy`:\n%s\n", responseContent)
}