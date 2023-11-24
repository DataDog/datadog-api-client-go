// Delete Fastly account returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "fastly_account" in the system
	FastlyAccountDataID := os.Getenv("FASTLY_ACCOUNT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFastlyIntegrationApi(apiClient)
	r, err := api.DeleteFastlyAccount(ctx, FastlyAccountDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.DeleteFastlyAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
