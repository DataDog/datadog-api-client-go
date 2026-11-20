// List Claude Compliance integration accounts returns "OK" response

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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListClaudeComplianceIntegrationAccounts", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewClaudeComplianceIntegrationApi(apiClient)
	resp, r, err := api.ListClaudeComplianceIntegrationAccounts(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaudeComplianceIntegrationApi.ListClaudeComplianceIntegrationAccounts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ClaudeComplianceIntegrationApi.ListClaudeComplianceIntegrationAccounts`:\n%s\n", responseContent)
}
