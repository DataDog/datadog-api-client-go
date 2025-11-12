// Update an On-Call email for a user returns "OK" response

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

	// there is a valid "oncall_email" in the system
	OncallEmailDataAttributesAddress := os.Getenv("ONCALL_EMAIL_DATA_ATTRIBUTES_ADDRESS")
	OncallEmailDataID := os.Getenv("ONCALL_EMAIL_DATA_ID")

	body := datadogV2.EmailUpdateRequest{
		Data: &datadogV2.EmailData{
			Id: datadog.PtrString(OncallEmailDataID),
			Attributes: &datadogV2.EmailAttributes{
				Active:  datadog.PtrBool(true),
				Address: datadog.PtrString(OncallEmailDataAttributesAddress),
				Alias:   datadog.PtrString("Example-On-Call-alias"),
				Formats: []datadogV2.EmailFormatType{
					datadogV2.EMAILFORMATTYPE_HTML,
				},
			},
			Type: datadogV2.EMAILTYPE_EMAILS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.UpdateUserEmailNotificationChannel(ctx, UserDataID, OncallEmailDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.UpdateUserEmailNotificationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.UpdateUserEmailNotificationChannel`:\n%s\n", responseContent)
}
