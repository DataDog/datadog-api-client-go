// Revoke an access token for a service account returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	// there is a valid "service_account_access_token" for "service_account_user"
	ServiceAccountAccessTokenDataID := os.Getenv("SERVICE_ACCOUNT_ACCESS_TOKEN_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	r, err := api.RevokeServiceAccountAccessToken(ctx, ServiceAccountUserDataID, ServiceAccountAccessTokenDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.RevokeServiceAccountAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
