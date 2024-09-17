// Add Okta account returns "OK" response

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
	body := datadogV2.OktaAccountRequest{
		Data: datadogV2.OktaAccount{
			Attributes: datadogV2.OktaAccountAttributes{
				AuthMethod:   "oauth",
				Domain:       "https://example.okta.com/",
				Name:         "exampleoktaintegration",
				ClientId:     datadog.PtrString("client_id"),
				ClientSecret: datadog.PtrString("client_secret"),
			},
			Id:   datadog.PtrString("f749daaf-682e-4208-a38d-c9b43162c609"),
			Type: datadogV2.OKTAACCOUNTTYPE_OKTA_ACCOUNTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOktaIntegrationApi(apiClient)
	resp, r, err := api.CreateOktaAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaIntegrationApi.CreateOktaAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OktaIntegrationApi.CreateOktaAccount`:\n%s\n", responseContent)
}
