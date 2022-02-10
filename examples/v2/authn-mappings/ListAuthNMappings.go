// List all AuthN Mappings returns "OK" response

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
	resp, r, err := apiClient.AuthNMappingsApi.ListAuthNMappings(ctx, *datadog.NewListAuthNMappingsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.ListAuthNMappings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuthNMappingsApi.ListAuthNMappings`:\n%s\n", responseContent)
}
