// Update a personal access token returns "OK" response

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
	// there is a valid "personal_access_token" in the system
	PersonalAccessTokenDataID := os.Getenv("PERSONAL_ACCESS_TOKEN_DATA_ID")

	body := datadogV2.PersonalAccessTokenUpdateRequest{
		Data: datadogV2.PersonalAccessTokenUpdateData{
			Type: datadogV2.PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
			Id:   PersonalAccessTokenDataID,
			Attributes: datadogV2.PersonalAccessTokenUpdateAttributes{
				Name: datadog.PtrString("Example-Key-Management-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdatePersonalAccessToken(ctx, PersonalAccessTokenDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdatePersonalAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdatePersonalAccessToken`:\n%s\n", responseContent)
}
