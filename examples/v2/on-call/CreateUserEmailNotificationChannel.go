// Create an On-Call email for a user returns "Created" response

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

	body := datadogV2.EmailCreateRequest{
		Data: &datadogV2.EmailData{
			Attributes: &datadogV2.EmailAttributes{
				Active:  datadog.PtrBool(true),
				Address: datadog.PtrString("john.doe@datadoghq.com"),
				Alias:   datadog.PtrString(""),
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
	resp, r, err := api.CreateUserEmailNotificationChannel(ctx, UserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateUserEmailNotificationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateUserEmailNotificationChannel`:\n%s\n", responseContent)
}
