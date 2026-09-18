// Delete a Terraform backend sync configuration returns "No Content" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTerraformStateFilesApi(apiClient)
	r, err := api.DeleteTerraformBackendSyncConfig(ctx, "9007199254740993")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformStateFilesApi.DeleteTerraformBackendSyncConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
