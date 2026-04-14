// Create an access token for a service account returns "Created" response

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

	body := datadogV2.ServiceAccountAccessTokenCreateRequest{
		Data: datadogV2.ServiceAccountAccessTokenCreateData{
			Type: datadogV2.PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
			Attributes: datadogV2.ServiceAccountAccessTokenCreateAttributes{
				Name: "Example-Service-Account",
				Scopes: []string{
					"dashboards_read",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	resp, r, err := api.CreateServiceAccountAccessToken(ctx, ServiceAccountUserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountAccessToken`:\n%s\n", responseContent)
}
