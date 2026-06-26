// Update governance control notification settings returns "OK" response

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
	body := datadogV2.ControlNotificationSettingsUpdateRequest{
		Data: datadogV2.ControlNotificationSettingsUpdateData{
			Attributes: &datadogV2.ControlNotificationSettingsUpdateAttributes{
				EventSettings: []datadogV2.ControlNotificationEventSetting{
					{
						Enabled:   true,
						EventType: "new_detection",
						Targets: []datadogV2.ControlNotificationTarget{
							{
								Handle: "#governance-alerts",
								Type:   datadogV2.CONTROLNOTIFICATIONTARGETTYPE_SLACK,
							},
						},
					},
				},
			},
			Type: datadogV2.CONTROLNOTIFICATIONSETTINGSRESOURCETYPE_CONTROL_NOTIFICATION_SETTINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateGovernanceControlNotificationSettings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceControlsApi(apiClient)
	resp, r, err := api.UpdateGovernanceControlNotificationSettings(ctx, "detection_type", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceControlsApi.UpdateGovernanceControlNotificationSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceControlsApi.UpdateGovernanceControlNotificationSettings`:\n%s\n", responseContent)
}
