// Update Workflows webhook handle returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.MicrosoftTeamsUpdateWorkflowsWebhookHandleRequest{
Data: datadogV2.MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData{
Attributes: datadogV2.MicrosoftTeamsWorkflowsWebhookHandleAttributes{
Name: datadog.PtrString("fake-handle-name"),
Url: datadog.PtrString("https://fake.url.com"),
},
Type: datadogV2.MICROSOFTTEAMSWORKFLOWSWEBHOOKHANDLETYPE_WORKFLOWS_WEBHOOK_HANDLE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.UpdateWorkflowsWebhookHandle(ctx, "handle_id", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.UpdateWorkflowsWebhookHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.UpdateWorkflowsWebhookHandle`:\n%s\n", responseContent)
}