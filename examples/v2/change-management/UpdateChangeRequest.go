// Update a change request returns "OK" response

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
	body := datadogV2.ChangeRequestUpdateRequest{
		Data: datadogV2.ChangeRequestUpdateData{
			Attributes: &datadogV2.ChangeRequestUpdateAttributes{
				ChangeRequestPlan: datadog.PtrString("Updated deployment plan"),
				ChangeRequestRisk: datadogV2.CHANGEREQUESTRISKLEVEL_LOW.Ptr(),
				ChangeRequestType: datadogV2.CHANGEREQUESTCHANGETYPE_NORMAL.Ptr(),
				EndDate:           datadog.PtrTime(time.Date(2024, 1, 2, 15, 0, 0, 0, time.UTC)),
				Id:                datadog.PtrString("CHM-1234"),
				StartDate:         datadog.PtrTime(time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC)),
			},
			Relationships: &datadogV2.ChangeRequestUpdateRelationships{
				ChangeRequestDecisions: &datadogV2.ChangeRequestDecisionsRelationship{
					Data: []datadogV2.ChangeRequestDecisionRelationshipData{
						{
							Id:   "decision-id-0",
							Type: datadogV2.CHANGEREQUESTDECISIONRESOURCETYPE_CHANGE_REQUEST_DECISION,
						},
					},
				},
			},
			Type: datadogV2.CHANGEREQUESTRESOURCETYPE_CHANGE_REQUEST,
		},
		Included: []datadogV2.ChangeRequestDecisionCreateItem{
			{
				Attributes: &datadogV2.ChangeRequestDecisionCreateAttributes{
					ChangeRequestStatus: datadogV2.CHANGEREQUESTDECISIONSTATUSTYPE_REQUESTED.Ptr(),
					RequestReason:       datadog.PtrString("Please review and approve this change"),
				},
				Id: "decision-id-0",
				Relationships: &datadogV2.ChangeRequestDecisionCreateRelationships{
					RequestedUser: &datadogV2.ChangeRequestUserRelationship{
						Data: *datadogV2.NewNullableChangeRequestUserRelationshipData(&datadogV2.ChangeRequestUserRelationshipData{
							Id:   "00000000-0000-0000-0000-000000000000",
							Type: "user",
						}),
					},
				},
				Type: datadogV2.CHANGEREQUESTDECISIONRESOURCETYPE_CHANGE_REQUEST_DECISION,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateChangeRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewChangeManagementApi(apiClient)
	resp, r, err := api.UpdateChangeRequest(ctx, "change_request_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeManagementApi.UpdateChangeRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ChangeManagementApi.UpdateChangeRequest`:\n%s\n", responseContent)
}
