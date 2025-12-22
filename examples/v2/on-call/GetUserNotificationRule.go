// Get an On-Call notification rule for a user returns "OK" response

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

	// there is a valid "oncall_email_notification_rule" in the system
	OncallEmailNotificationRuleDataID := os.Getenv("ONCALL_EMAIL_NOTIFICATION_RULE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.GetUserNotificationRule(ctx, UserDataID, OncallEmailNotificationRuleDataID, *datadogV2.NewGetUserNotificationRuleOptionalParameters().WithInclude("channel"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.GetUserNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.GetUserNotificationRule`:\n%s\n", responseContent)
}
