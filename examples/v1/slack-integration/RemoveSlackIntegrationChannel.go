// Remove a Slack integration channel returns "The channel was removed successfully." response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SlackIntegrationApi(apiClient)
	r, err := api.RemoveSlackIntegrationChannel(ctx, "account_name", "channel_name")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.RemoveSlackIntegrationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
