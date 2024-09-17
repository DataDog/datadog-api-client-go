// Update a Slack integration channel returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	body := datadogV1.SlackIntegrationChannel{
Display: &datadogV1.SlackIntegrationChannelDisplay{
Message: datadog.PtrBool(true),
Notified: datadog.PtrBool(true),
Snapshot: datadog.PtrBool(true),
Tags: datadog.PtrBool(true),
},
Name: datadog.PtrString("#general"),
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSlackIntegrationApi(apiClient)
	resp, r, err := api.UpdateSlackIntegrationChannel(ctx, "account_name", "channel_name", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.UpdateSlackIntegrationChannel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SlackIntegrationApi.UpdateSlackIntegrationChannel`:\n%s\n", responseContent)
}