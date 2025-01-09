// Deploy App returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "app" in the system
	AppDataID := os.Getenv("APP_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeployApp", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppsApi(apiClient)
	resp, r, err := api.DeployApp(ctx, AppDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.DeployApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppsApi.DeployApp`:\n%s\n", responseContent)
}
