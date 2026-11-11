// Update a GitHub Costs integration account returns "OK" response

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
	body := datadogV2.GitHubCostsIntegrationAccountUpdateRequest{
		Data: datadogV2.GitHubCostsIntegrationAccountUpdateData{
			Attributes: datadogV2.GitHubCostsIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.GitHubCostsIntegrationAccountAuthenticationUpdate{
					GitHubCostsIntegrationAccountBearerTokenAuthUpdate: &datadogV2.GitHubCostsIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.GITHUBCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Name: datadog.PtrString("github-costs-acme"),
				Settings: &datadogV2.GitHubCostsIntegrationAccountSettingsUpdate{
					EnterpriseName: datadog.PtrString("acme"),
					Tags:           datadog.PtrString("env:prod,team:finops"),
				},
			},
			Id:   "4b7c2e1a-6f3d-4e8b-9a52-1d0c8f7e3b64",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateGitHubCostsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGitHubCostsIntegrationApi(apiClient)
	resp, r, err := api.UpdateGitHubCostsIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubCostsIntegrationApi.UpdateGitHubCostsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GitHubCostsIntegrationApi.UpdateGitHubCostsIntegrationAccount`:\n%s\n", responseContent)
}
