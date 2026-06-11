// List user authorizations for a client returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgAuthorizedClientsApi(apiClient)
	resp, r, err := api.ListOrgAuthorizedClientUserAuthorizations(ctx, "00000000-0000-0000-0000-000000000001", *datadogV2.NewListOrgAuthorizedClientUserAuthorizationsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAuthorizedClientsApi.ListOrgAuthorizedClientUserAuthorizations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgAuthorizedClientsApi.ListOrgAuthorizedClientUserAuthorizations`:\n%s\n", responseContent)
}
