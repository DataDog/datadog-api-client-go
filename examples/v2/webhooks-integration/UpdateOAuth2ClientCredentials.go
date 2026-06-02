// Update an OAuth2 client credentials auth method returns "OK" response

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
	body := datadogV2.WebhooksOAuth2ClientCredentialsUpdateRequest{
		Data: datadogV2.WebhooksOAuth2ClientCredentialsUpdateData{
			Attributes: datadogV2.WebhooksOAuth2ClientCredentialsUpdateAttributes{
				AccessTokenUrl: datadog.PtrString("https://example.com/oauth/token"),
				Audience:       *datadog.NewNullableString(datadog.PtrString("https://api.example.com")),
				ClientId:       datadog.PtrString("my-client-id"),
				ClientSecret:   datadog.PtrString("my-client-secret"),
				Name:           datadog.PtrString("my-oauth2-auth"),
				Scope:          *datadog.NewNullableString(datadog.PtrString("read:webhooks write:webhooks")),
			},
			Type: datadogV2.WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.UpdateOAuth2ClientCredentials(ctx, "auth_method_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.UpdateOAuth2ClientCredentials`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.UpdateOAuth2ClientCredentials`:\n%s\n", responseContent)
}
