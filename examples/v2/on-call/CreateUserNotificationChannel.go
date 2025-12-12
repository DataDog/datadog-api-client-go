// Create an On-Call notification channel for a user returns "Created" response

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

	body := datadogV2.CreateUserNotificationChannelRequest{
		Data: datadogV2.CreateNotificationChannelData{
			Attributes: &datadogV2.CreateNotificationChannelAttributes{
				Config: &datadogV2.CreateNotificationChannelConfig{
					CreateEmailNotificationChannelConfig: &datadogV2.CreateEmailNotificationChannelConfig{
						Address: "foo@bar.com",
						Formats: []datadogV2.NotificationChannelEmailFormatType{
							datadogV2.NOTIFICATIONCHANNELEMAILFORMATTYPE_HTML,
						},
						Type: datadogV2.NOTIFICATIONCHANNELEMAILCONFIGTYPE_EMAIL,
					}},
			},
			Type: datadogV2.NOTIFICATIONCHANNELTYPE_NOTIFICATION_CHANNELS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateUserNotificationChannel(ctx, UserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateUserNotificationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateUserNotificationChannel`:\n%s\n", responseContent)
}
