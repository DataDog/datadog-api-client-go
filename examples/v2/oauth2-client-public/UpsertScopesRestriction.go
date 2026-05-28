// Upsert an OAuth2 client scopes restriction returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.UpsertOAuthScopesRestrictionRequest{
		Data: datadogV2.UpsertOAuthScopesRestrictionData{
			Attributes: &datadogV2.UpsertOAuthScopesRestrictionDataAttributes{
				OidcScopes: []datadogV2.OAuthOidcScope{
					datadogV2.OAUTHOIDCSCOPE_OPENID,
					datadogV2.OAUTHOIDCSCOPE_EMAIL,
				},
				PermissionScopes: []string{
					"dashboards_read",
					"metrics_read",
				},
			},
			Type: datadogV2.UPSERTOAUTHSCOPESRESTRICTIONTYPE_UPSERT_SCOPES_RESTRICTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertScopesRestriction", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOAuth2ClientPublicApi(apiClient)
	resp, r, err := api.UpsertScopesRestriction(ctx, uuid.MustParse("fafa8e1c-36a5-11f0-a83d-da7ad0900001"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientPublicApi.UpsertScopesRestriction`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientPublicApi.UpsertScopesRestriction`:\n%s\n", responseContent)
}
