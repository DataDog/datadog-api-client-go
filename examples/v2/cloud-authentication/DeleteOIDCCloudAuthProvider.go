// Delete a Custom OIDC provider returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "oidc_cloud_auth_provider" in the system
	OidcCloudAuthProviderDataID := uuid.MustParse(os.Getenv("OIDC_CLOUD_AUTH_PROVIDER_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteOIDCCloudAuthProvider", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	r, err := api.DeleteOIDCCloudAuthProvider(ctx, OidcCloudAuthProviderDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.DeleteOIDCCloudAuthProvider`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
