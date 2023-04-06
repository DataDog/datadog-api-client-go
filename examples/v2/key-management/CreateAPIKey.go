// Create an API key returns "Created" response

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
	body := datadogV2.APIKeyCreateRequest{
		Data: datadogV2.APIKeyCreateData{
			Type: datadogV2.APIKEYSTYPE_API_KEYS,
			Attributes: datadogV2.APIKeyCreateAttributes{
				Name: "Example-Key-Management",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.CreateAPIKey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.CreateAPIKey`:\n%s\n", responseContent)
}
