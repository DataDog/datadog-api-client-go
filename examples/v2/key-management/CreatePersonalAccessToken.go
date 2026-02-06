// Create personal access token returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.PersonalAccessTokenCreateRequest{
		Data: datadogV2.PersonalAccessTokenCreateData{
			Attributes: datadogV2.PersonalAccessTokenCreateAttributes{
				ExpiresAt: time.Date(2025, 3, 15, 10, 30, 0, 0, time.UTC),
				Name:      "Example Personal Access Token",
				Scopes: []string{
					"dashboards_read",
					"monitors_read",
				},
			},
			Type: datadogV2.PERSONALACCESSTOKENTYPE_PERSONAL_ACCESS_TOKENS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreatePersonalAccessToken", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.CreatePersonalAccessToken(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreatePersonalAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.CreatePersonalAccessToken`:\n%s\n", responseContent)
}
