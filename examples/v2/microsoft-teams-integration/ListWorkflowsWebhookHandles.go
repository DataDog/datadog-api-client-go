// Get all Workflows webhook handles returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMicrosoftTeamsIntegrationApi(apiClient)
	resp, r, err := api.ListWorkflowsWebhookHandles(ctx, *datadogV2.NewListWorkflowsWebhookHandlesOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftTeamsIntegrationApi.ListWorkflowsWebhookHandles`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MicrosoftTeamsIntegrationApi.ListWorkflowsWebhookHandles`:\n%s\n", responseContent)
}