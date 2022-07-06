// Update a Slack integration channel returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SlackIntegrationChannel{
		Display: &datadog.SlackIntegrationChannelDisplay{
			Message:  datadog.PtrBool(true),
			Notified: datadog.PtrBool(true),
			Snapshot: datadog.PtrBool(true),
			Tags:     datadog.PtrBool(true),
		},
		Name: datadog.PtrString("#general"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SlackIntegrationApi(apiClient)
	resp, r, err := api.UpdateSlackIntegrationChannel(ctx, "account_name", "channel_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.UpdateSlackIntegrationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SlackIntegrationApi.UpdateSlackIntegrationChannel`:\n%s\n", responseContent)
}
