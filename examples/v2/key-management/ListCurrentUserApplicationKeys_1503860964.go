// Get all Application keys owned by current user returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementApi.ListCurrentUserApplicationKeys(ctx, *datadog.NewListCurrentUserApplicationKeysOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListCurrentUserApplicationKeys`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.ListCurrentUserApplicationKeys`:\n%s\n", responseContent)
}
