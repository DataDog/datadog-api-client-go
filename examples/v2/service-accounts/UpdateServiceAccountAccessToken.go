// Update an access token for a service account returns "OK" response

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
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	// there is a valid "service_account_access_token" for "service_account_user"
	ServiceAccountAccessTokenDataID := os.Getenv("SERVICE_ACCOUNT_ACCESS_TOKEN_DATA_ID")

	body := datadogV2.PersonalAccessTokenUpdateRequest{
		Data: datadogV2.PersonalAccessTokenUpdateData{
			Id:   ServiceAccountAccessTokenDataID,
			Type: datadogV2.PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
			Attributes: datadogV2.PersonalAccessTokenUpdateAttributes{
				Name: datadog.PtrString("My Personal Access Token-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	resp, r, err := api.UpdateServiceAccountAccessToken(ctx, ServiceAccountUserDataID, ServiceAccountAccessTokenDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccountAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccountAccessToken`:\n%s\n", responseContent)
}
