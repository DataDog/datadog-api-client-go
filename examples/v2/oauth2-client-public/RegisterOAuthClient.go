// Register an OAuth2 client returns "Created" response

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
	body := datadogV2.OAuthClientRegistrationRequest{
		ClientName: "Example MCP Client",
		ClientUri:  datadog.PtrString("https://example.com"),
		GrantTypes: []datadogV2.OAuthClientRegistrationGrantType{
			datadogV2.OAUTHCLIENTREGISTRATIONGRANTTYPE_AUTHORIZATION_CODE,
			datadogV2.OAUTHCLIENTREGISTRATIONGRANTTYPE_REFRESH_TOKEN,
		},
		JwksUri:   datadog.PtrString("https://example.com/.well-known/jwks.json"),
		LogoUri:   datadog.PtrString("https://example.com/logo.png"),
		PolicyUri: datadog.PtrString("https://example.com/privacy"),
		RedirectUris: []string{
			"https://example.com/oauth/callback",
		},
		ResponseTypes: []datadogV2.OAuthClientRegistrationResponseType{
			datadogV2.OAUTHCLIENTREGISTRATIONRESPONSETYPE_CODE,
		},
		Scope:                   datadog.PtrString("openid profile"),
		TokenEndpointAuthMethod: datadog.PtrString("none"),
		TosUri:                  datadog.PtrString("https://example.com/tos"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RegisterOAuthClient", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOAuth2ClientPublicApi(apiClient)
	resp, r, err := api.RegisterOAuthClient(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientPublicApi.RegisterOAuthClient`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientPublicApi.RegisterOAuthClient`:\n%s\n", responseContent)
}
