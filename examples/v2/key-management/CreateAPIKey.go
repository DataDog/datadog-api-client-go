// Create an API key returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.APIKeyCreateRequest{
		Data: datadog.APIKeyCreateData{
			Type: datadog.APIKEYSTYPE_API_KEYS,
			Attributes: datadog.APIKeyCreateAttributes{
				Name: "Example-Create_an_API_key_returns_Created_response",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementApi.CreateAPIKey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.CreateAPIKey`:\n%s\n", responseContent)
}
