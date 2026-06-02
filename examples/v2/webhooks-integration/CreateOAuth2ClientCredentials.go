// Create an OAuth2 client credentials auth method returns "CREATED" response

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
	body := datadogV2.WebhooksOAuth2ClientCredentialsCreateRequest{
		Data: datadogV2.WebhooksOAuth2ClientCredentialsCreateData{
			Attributes: datadogV2.WebhooksOAuth2ClientCredentialsCreateAttributes{
				AccessTokenUrl: "https://example.com/oauth/token",
				Audience:       *datadog.NewNullableString(datadog.PtrString("https://api.example.com")),
				ClientId:       "my-client-id",
				ClientSecret:   "my-client-secret",
				Name:           "my-oauth2-auth",
				Scope:          *datadog.NewNullableString(datadog.PtrString("read:webhooks write:webhooks")),
			},
			Type: datadogV2.WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.CreateOAuth2ClientCredentials(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateOAuth2ClientCredentials`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.CreateOAuth2ClientCredentials`:\n%s\n", responseContent)
}
