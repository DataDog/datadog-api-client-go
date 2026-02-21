// Update personal access token returns "OK" response

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
	body := datadogV2.PersonalAccessTokenUpdateRequest{
		Data: datadogV2.PersonalAccessTokenUpdateData{
			Attributes: datadogV2.PersonalAccessTokenUpdateAttributes{
				Name: datadog.PtrString("Updated Personal Access Token Name"),
				Scopes: []string{
					"dashboards_read",
					"dashboards_write",
				},
			},
			Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Type: datadogV2.PERSONALACCESSTOKENTYPE_PERSONAL_ACCESS_TOKENS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdatePersonalAccessToken", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdatePersonalAccessToken(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdatePersonalAccessToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdatePersonalAccessToken`:\n%s\n", responseContent)
}
