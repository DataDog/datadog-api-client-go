// Create an On-Call notification rule for a user returns "Created" response

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

	// there is a valid "oncall_email_notification_channel" in the system
	OncallEmailNotificationChannelDataID := os.Getenv("ONCALL_EMAIL_NOTIFICATION_CHANNEL_DATA_ID")

	body := datadogV2.CreateOnCallNotificationRuleRequest{
		Data: datadogV2.CreateOnCallNotificationRuleRequestData{
			Attributes: &datadogV2.OnCallNotificationRuleRequestAttributes{
				Category:     datadogV2.ONCALLNOTIFICATIONRULECATEGORY_HIGH_URGENCY.Ptr(),
				DelayMinutes: datadog.PtrInt64(0),
			},
			Relationships: &datadogV2.OnCallNotificationRuleRelationships{
				Channel: &datadogV2.OnCallNotificationRuleChannelRelationship{
					Data: datadogV2.OnCallNotificationRuleChannelRelationshipData{
						Id:   datadog.PtrString(OncallEmailNotificationChannelDataID),
						Type: datadogV2.NOTIFICATIONCHANNELTYPE_NOTIFICATION_CHANNELS.Ptr(),
					},
				},
			},
			Type: datadogV2.ONCALLNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateUserNotificationRule(ctx, UserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateUserNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateUserNotificationRule`:\n%s\n", responseContent)
}
