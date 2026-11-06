// Create a Custom OIDC persona mapping returns "Created" response

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
	// there is a valid "oidc_cloud_auth_provider" in the system
	OidcCloudAuthProviderDataID := uuid.MustParse(os.Getenv("OIDC_CLOUD_AUTH_PROVIDER_DATA_ID"))

	// there is a valid "service_account_user_with_role" in the system
	ServiceAccountUserWithRoleDataID := os.Getenv("SERVICE_ACCOUNT_USER_WITH_ROLE_DATA_ID")

	body := datadogV2.OIDCCloudAuthPersonaMappingCreateRequest{
		Data: datadogV2.OIDCCloudAuthPersonaMappingCreateData{
			Type: datadogV2.OIDCCLOUDAUTHPERSONAMAPPINGTYPE_CUSTOM_OIDC_IDENTITY_MAPPING,
			Attributes: datadogV2.OIDCCloudAuthPersonaMappingCreateAttributes{
				ProviderId:  OidcCloudAuthProviderDataID,
				AccountUuid: ServiceAccountUserWithRoleDataID,
				ClaimMatchers: []datadogV2.OIDCCloudAuthClaimMatcherCreate{
					{
						Claim: "sub",
						Mode:  datadogV2.OIDCCLOUDAUTHCLAIMMATCHMODE_EXACT.Ptr(),
						Values: []string{
							"system:serviceaccount:controllers:example-controller",
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOIDCCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.CreateOIDCCloudAuthPersonaMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.CreateOIDCCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.CreateOIDCCloudAuthPersonaMapping`:\n%s\n", responseContent)
}
