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
	configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
		OrgUUID: os.Getenv("DD_TEST_ORG_UUID"),
		ProviderAuth: &datadog.GCPAuth{
			// Minting via `gcloud` requires impersonating a service account:
			// user identity tokens cannot set the custom `datadog/<org-uuid>`
			// audience. Alternatively, pass a pre-minted identity token via
			// GCPAuth.IdentityToken or the DD_GCP_IDENTITY_TOKEN environment
			// variable to skip the `gcloud` invocation entirely.
			ImpersonateServiceAccount: os.Getenv("DD_TEST_GCP_SA_EMAIL"),
		},
		Provider: datadog.ProviderGCP,
	}
	apiClient := datadog.NewAPIClient(configuration)

	// Make example API call using the DelegatedTokenConfig
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.ListTeams(ctx, *datadogV2.NewListTeamsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeams`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeams`:\n%s\n", responseContent)
}
