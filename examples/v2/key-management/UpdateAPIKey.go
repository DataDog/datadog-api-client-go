// Edit an API key returns "OK" response

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
	// there is a valid "api_key" in the system
	APIKeyDataID := os.Getenv("API_KEY_DATA_ID")

	body := datadogV2.APIKeyUpdateRequest{
		Data: datadogV2.APIKeyUpdateData{
			Type: datadogV2.APIKEYSTYPE_API_KEYS,
			Id:   APIKeyDataID,
			Attributes: datadogV2.APIKeyUpdateAttributes{
				Name: "Example-Key-Management",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdateAPIKey(ctx, APIKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateAPIKey`:\n%s\n", responseContent)
}
