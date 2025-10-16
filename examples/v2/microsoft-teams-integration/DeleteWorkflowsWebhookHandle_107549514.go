// Delete workflow webhook handle returns "OK" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "workflows_webhook_handle" in the system
	WorkflowsWebhookHandleDataID := os.Getenv("WORKFLOWS_WEBHOOK_HANDLE_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	r, err := api.DeleteWorkflowsWebhookHandle(ctx, WorkflowsWebhookHandleDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.DeleteWorkflowsWebhookHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}