// Delete a user authorization for a client returns "No Content" response

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
	r, err := api.DeleteOrgAuthorizedClientUserAuthorization(ctx, "00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000002")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAuthorizedClientsApi.DeleteOrgAuthorizedClientUserAuthorization`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
