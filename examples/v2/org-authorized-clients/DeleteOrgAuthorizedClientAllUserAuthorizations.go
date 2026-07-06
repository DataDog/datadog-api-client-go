// Delete a user's authorizations for a client returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgAuthorizedClientsApi(apiClient)
	r, err := api.DeleteOrgAuthorizedClientAllUserAuthorizations(ctx, "00000000-0000-0000-0000-000000000001", "00000000-0000-9999-0000-000000000001")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAuthorizedClientsApi.DeleteOrgAuthorizedClientAllUserAuthorizations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
