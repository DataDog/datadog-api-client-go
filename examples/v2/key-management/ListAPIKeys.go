// Get all API keys returns "OK" response

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
	APIKeyDataAttributesName := os.Getenv("API_KEY_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.ListAPIKeys(ctx, *datadogV2.NewListAPIKeysOptionalParameters().WithFilter(APIKeyDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListAPIKeys`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.ListAPIKeys`:\n%s\n", responseContent)
}
