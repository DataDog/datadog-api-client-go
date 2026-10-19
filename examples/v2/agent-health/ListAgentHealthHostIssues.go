// List active Agent Health issues returns "OK" response

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
	api := datadogV2.NewAgentHealthApi(apiClient)
	resp, r, err := api.ListAgentHealthHostIssues(ctx, "123e4567-e89b-12d3-a456-426614174000", *datadogV2.NewListAgentHealthHostIssuesOptionalParameters().WithAgentType(datadogV2.AGENTHEALTHAGENTTYPE_NODE))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentHealthApi.ListAgentHealthHostIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentHealthApi.ListAgentHealthHostIssues`:\n%s\n", responseContent)
}
