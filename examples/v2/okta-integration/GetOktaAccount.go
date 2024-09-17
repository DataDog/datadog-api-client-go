// Get Okta account returns "OK" response

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
	// there is a valid "okta_account" in the system
	OktaAccountDataID := os.Getenv("OKTA_ACCOUNT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOktaIntegrationApi(apiClient)
	resp, r, err := api.GetOktaAccount(ctx, OktaAccountDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaIntegrationApi.GetOktaAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OktaIntegrationApi.GetOktaAccount`:\n%s\n", responseContent)
}
