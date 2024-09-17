// Update Okta account returns "OK" response

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

	body := datadogV2.OktaAccountUpdateRequest{
		Data: datadogV2.OktaAccountUpdateRequestData{
			Attributes: &datadogV2.OktaAccountUpdateRequestAttributes{
				AuthMethod:   "oauth",
				Domain:       "https://example.okta.com/",
				ClientId:     datadog.PtrString("client_id"),
				ClientSecret: datadog.PtrString("client_secret"),
			},
			Type: datadogV2.OKTAACCOUNTTYPE_OKTA_ACCOUNTS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOktaIntegrationApi(apiClient)
	resp, r, err := api.UpdateOktaAccount(ctx, OktaAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaIntegrationApi.UpdateOktaAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OktaIntegrationApi.UpdateOktaAccount`:\n%s\n", responseContent)
}
