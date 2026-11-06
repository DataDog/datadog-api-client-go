// Create a Custom OIDC provider returns "Created" response

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
	body := datadogV2.OIDCCloudAuthProviderCreateRequest{
		Data: datadogV2.OIDCCloudAuthProviderCreateData{
			Type: datadogV2.OIDCCLOUDAUTHPROVIDERTYPE_CUSTOM_OIDC_PROVIDER_CONFIG,
			Attributes: datadogV2.OIDCCloudAuthProviderCreateAttributes{
				ProviderName:  "oidc-de20b7e944abfac3",
				IssuerUrl:     "https://oidc-de20b7e944abfac3.example.com",
				DiscoveryType: datadogV2.OIDCCLOUDAUTHDISCOVERYTYPE_OIDC_DISCOVERY,
				Audiences: []string{
					"sdk-test",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOIDCCloudAuthProvider", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateOIDCCloudAuthProvider(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateOIDCCloudAuthProvider`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateOIDCCloudAuthProvider`:\n%s\n", responseContent)
}
