// Update a Claude Compliance integration account returns "OK" response

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
	body := datadogV2.ClaudeComplianceIntegrationAccountUpdateRequest{
		Data: datadogV2.ClaudeComplianceIntegrationAccountUpdateData{
			Attributes: datadogV2.ClaudeComplianceIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.ClaudeComplianceIntegrationAccountAuthenticationUpdate{
					ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate: &datadogV2.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.CLAUDECOMPLIANCEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Name: datadog.PtrString("claude-compliance-acme"),
			},
			Id:   "5f1a8d3c-7e2b-4f9a-b6c4-8d0e3a7f1b52",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateClaudeComplianceIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewClaudeComplianceIntegrationApi(apiClient)
	resp, r, err := api.UpdateClaudeComplianceIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaudeComplianceIntegrationApi.UpdateClaudeComplianceIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ClaudeComplianceIntegrationApi.UpdateClaudeComplianceIntegrationAccount`:\n%s\n", responseContent)
}
