// Revoke client token returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.ClientTokenRevokeRequest{
		Hash: "1234567890abcdef1234567890abcdef123",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v1.RevokeClientToken", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewKeyManagementApi(apiClient)
	r, err := api.RevokeClientToken(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.RevokeClientToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
