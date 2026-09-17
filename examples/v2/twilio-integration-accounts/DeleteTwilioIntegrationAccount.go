// Delete a Twilio integration account returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteTwilioIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTwilioIntegrationAccountsApi(apiClient)
	r, err := api.DeleteTwilioIntegrationAccount(ctx, "account_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwilioIntegrationAccountsApi.DeleteTwilioIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
