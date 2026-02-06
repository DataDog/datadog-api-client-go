// Update client token returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.ClientTokenUpdateRequest{
		Hash: "1234567890abcdef1234567890abcdef123",
		Name: "Updated Client Token Name",
		OriginUrls: []string{
			"https://example.com",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v1.UpdateClientToken", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdateClientToken(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateClientToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateClientToken`:\n%s\n", responseContent)
}
