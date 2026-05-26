// Delete an entity integration configuration returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteEntityIntegrationConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEntityIntegrationConfigsApi(apiClient)
	r, err := api.DeleteEntityIntegrationConfig(ctx, "github")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationConfigsApi.DeleteEntityIntegrationConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
