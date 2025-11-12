// Delete an On-Call email for a user returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	// there is a valid "oncall_email" in the system
	OncallEmailDataID := os.Getenv("ONCALL_EMAIL_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	r, err := api.DeleteUserEmailNotificationChannel(ctx, UserDataID, OncallEmailDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.DeleteUserEmailNotificationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
