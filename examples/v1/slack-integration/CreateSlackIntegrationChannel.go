// Create a Slack integration channel returns "OK" response

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
	api := datadog.NewSlackIntegrationApi(apiClient)
	resp, r, err := api.CreateSlackIntegrationChannel(ctx, "account_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.CreateSlackIntegrationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SlackIntegrationApi.CreateSlackIntegrationChannel`:\n%s\n", responseContent)
}
