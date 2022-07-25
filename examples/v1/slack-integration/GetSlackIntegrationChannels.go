// Get all channels in a Slack integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSlackIntegrationApi(apiClient)
	resp, r, err := api.GetSlackIntegrationChannels(ctx, "account_name")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.GetSlackIntegrationChannels`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SlackIntegrationApi.GetSlackIntegrationChannels`:\n%s\n", responseContent)
}
