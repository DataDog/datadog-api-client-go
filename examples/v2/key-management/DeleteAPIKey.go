// Delete an API key returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "api_key" in the system
	APIKeyDataID := os.Getenv("API_KEY_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewKeyManagementApi(apiClient)
	r, err := api.DeleteAPIKey(ctx, APIKeyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
