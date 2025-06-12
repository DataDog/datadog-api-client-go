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
	// Set your site with the `DD_SITE` environment variable
	ctx := context.WithValue(
		datadog.NewDefaultContext(context.Background()),
		datadog.ContextDelegatedToken,
		&datadog.DelegatedTokenConfig{
			DelegatedTokenCredentials: datadog.DelegatedTokenCredentials{
				OrgUUID: os.Getenv("DD_TEST_ORG_UUID"),
			},
			ProviderAuth: &datadog.AWSAuth{},
			Provider:     datadog.ProviderAWS,
		},
	)

	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	// Make example API call using the Public Token Config
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.ListTeams(ctx, *datadogV2.NewListTeamsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeams`:\n%s\n", responseContent)
}
