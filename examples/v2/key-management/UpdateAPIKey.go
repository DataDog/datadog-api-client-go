// Edit an API key returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	// there is a valid "api_key" in the system
	APIKeyDataID := os.Getenv("API_KEY_DATA_ID")

	body := datadog.APIKeyUpdateRequest{
		Data: datadog.APIKeyUpdateData{
			Type: datadog.APIKEYSTYPE_API_KEYS,
			Id:   APIKeyDataID,
			Attributes: datadog.APIKeyUpdateAttributes{
				Name: "Example-Edit_an_API_key_returns_OK_response",
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewKeyManagementApi(apiClient)
	resp, r, err := api.UpdateAPIKey(ctx, APIKeyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateAPIKey`:\n%s\n", responseContent)
}
