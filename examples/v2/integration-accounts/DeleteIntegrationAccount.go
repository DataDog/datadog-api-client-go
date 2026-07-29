// Delete an integration account returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	r, err := api.DeleteIntegrationAccount(ctx, datadogV2.INTEGRATIONACCOUNTINTEGRATIONID_TWILIO, datadogV2.INTEGRATIONACCOUNTINTERFACEID_TWILIO, "account_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.DeleteIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
