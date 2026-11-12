// List Agent Health issues returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListAgentHealthIssues", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentHealthApi(apiClient)
	resp, r, err := api.ListAgentHealthIssues(ctx, *datadogV2.NewListAgentHealthIssuesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentHealthApi.ListAgentHealthIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentHealthApi.ListAgentHealthIssues`:\n%s\n", responseContent)
}
