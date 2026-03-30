// +build ignore

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
	pat := "ddpat_REPLACE_WITH_YOUR_PAT"

	// Start with a base context (no default API key/app key needed).
	ctx := context.Background()

	// Set the Bearer token — the client adds "Authorization: Bearer <token>" to every request.
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, pat)

	configuration := datadog.NewConfiguration()
	configuration.Debug = true // prints the raw HTTP request/response

	// Override the server URL to point at staging directly,
	// since "dd.datad0g.com" is not in the allowed enum for the {site} variable.
	configuration.Servers = datadog.ServerConfigurations{
		{
			URL:         "https://dd.datad0g.com",
			Description: "Staging",
		},
	}

	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsersApi(apiClient)

	resp, r, err := api.ListUsers(ctx, *datadogV2.NewListUsersOptionalParameters().WithPageSize(2))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(1)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`:\n%s\n", responseContent)
}
