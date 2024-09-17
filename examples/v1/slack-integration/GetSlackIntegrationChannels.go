// Get all channels in a Slack integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSlackIntegrationApi(apiClient)
	resp, r, err := api.GetSlackIntegrationChannels(ctx, "account_name")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.GetSlackIntegrationChannels`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SlackIntegrationApi.GetSlackIntegrationChannels`:\n%s\n", responseContent)
}
