// Update a schedule returns "OK" response

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
	body := datadogV2.FleetSchedulePatchRequest{
		Data: datadogV2.FleetSchedulePatch{
			Attributes: &datadogV2.FleetSchedulePatchAttributes{
				Name:  datadog.PtrString("Weekly Production Agent Updates"),
				Query: datadog.PtrString("env:prod AND service:web"),
				Rule: &datadogV2.FleetScheduleRecurrenceRule{
					DaysOfWeek: []string{
						"Mon",
						"Wed",
						"Fri",
					},
					MaintenanceWindowDuration: 1200,
					StartMaintenanceWindow:    "02:00",
					Timezone:                  "America/New_York",
				},
				Status:          datadogV2.FLEETSCHEDULESTATUS_ACTIVE.Ptr(),
				VersionToLatest: datadog.PtrInt64(0),
			},
			Type: datadogV2.FLEETSCHEDULERESOURCETYPE_SCHEDULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFleetSchedule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.UpdateFleetSchedule(ctx, "id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.UpdateFleetSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.UpdateFleetSchedule`:\n%s\n", responseContent)
}
