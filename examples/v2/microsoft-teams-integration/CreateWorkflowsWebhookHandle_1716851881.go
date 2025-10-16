// Create workflow webhook handle returns "CREATED" response

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
	body := datadogV2.MicrosoftTeamsCreateWorkflowsWebhookHandleRequest{
Data: datadogV2.MicrosoftTeamsWorkflowsWebhookHandleRequestData{
Attributes: datadogV2.MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes{
Name: "Example-Microsoft-Teams-Integration",
Url: "https://example.logic.azure.com/workflows/123",
},
Type: datadogV2.MICROSOFTTEAMSWORKFLOWSWEBHOOKHANDLETYPE_WORKFLOWS_WEBHOOK_HANDLE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.CreateWorkflowsWebhookHandle(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.CreateWorkflowsWebhookHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.CreateWorkflowsWebhookHandle`:\n%s\n", responseContent)
}