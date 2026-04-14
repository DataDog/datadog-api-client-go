// Revoke a personal access token returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "personal_access_token" in the system
	PersonalAccessTokenDataID := os.Getenv("PERSONAL_ACCESS_TOKEN_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	r, err := api.RevokePersonalAccessToken(ctx, PersonalAccessTokenDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.RevokePersonalAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
