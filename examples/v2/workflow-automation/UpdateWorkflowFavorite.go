// Update workflow favorite status returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.WorkflowFavoriteRequest{
		Data: datadogV2.WorkflowFavoriteRequestData{
			Attributes: datadogV2.WorkflowFavoriteRequestAttributes{
				Favorite: true,
			},
			Type: datadogV2.WORKFLOWFAVORITEREQUESTTYPE_WORKFLOW_FAVORITE_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateWorkflowFavorite", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	r, err := api.UpdateWorkflowFavorite(ctx, "workflow_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.UpdateWorkflowFavorite`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
