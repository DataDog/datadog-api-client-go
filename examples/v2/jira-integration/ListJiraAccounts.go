// List Jira accounts returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListJiraAccounts", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewJiraIntegrationApi(apiClient)
	resp, r, err := api.ListJiraAccounts(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationApi.ListJiraAccounts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationApi.ListJiraAccounts`:\n%s\n", responseContent)
}
