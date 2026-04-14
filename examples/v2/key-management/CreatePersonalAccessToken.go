// Create a personal access token returns "Created" response

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
			Type: datadogV2.PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
			Attributes: datadogV2.PersonalAccessTokenCreateAttributes{
				Name: "Example-Key-Management",
				Scopes: []string{
					"dashboards_read",
				},
				ExpiresAt: time.Now().AddDate(0, 0, 365),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
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
