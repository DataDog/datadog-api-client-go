// Create a change request returns "Created" response

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
	body := datadogV2.ChangeRequestCreateRequest{
		Data: datadogV2.ChangeRequestCreateData{
			Attributes: datadogV2.ChangeRequestCreateAttributes{
				ChangeRequestLinkedIncidentUuid:     datadog.PtrString("00000000-0000-0000-0000-000000000000"),
				ChangeRequestMaintenanceWindowQuery: datadog.PtrString(""),
				ChangeRequestPlan:                   datadog.PtrString("1. Deploy to staging 2. Run tests 3. Deploy to production"),
				ChangeRequestRisk:                   datadogV2.CHANGEREQUESTRISKLEVEL_LOW.Ptr(),
				ChangeRequestType:                   datadogV2.CHANGEREQUESTCHANGETYPE_NORMAL.Ptr(),
				Description:                         datadog.PtrString("Deploying new payment service v2.1"),
				EndDate:                             datadog.PtrTime(time.Date(2024, 1, 2, 15, 0, 0, 0, time.UTC)),
				ProjectId:                           datadog.PtrString("d4bbe1af-f36e-42f1-87c1-493ca35c320e"),
				RequestedTeams: []string{
					"team-handle-1",
				},
				StartDate: datadog.PtrTime(time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC)),
				Title:     "Deploy new payment service",
			},
			Type: datadogV2.CHANGEREQUESTRESOURCETYPE_CHANGE_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateChangeRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewChangeManagementApi(apiClient)
	resp, r, err := api.CreateChangeRequest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeManagementApi.CreateChangeRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ChangeManagementApi.CreateChangeRequest`:\n%s\n", responseContent)
}
