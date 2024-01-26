// Delete an API returns "API deleted successfully" response

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
	// there is a valid "managed_api" in the system
	ManagedAPIDataID := uuid.MustParse(os.Getenv("MANAGED_API_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteOpenAPI", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPIManagementApi(apiClient)
	r, err := api.DeleteOpenAPI(ctx, ManagedAPIDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIManagementApi.DeleteOpenAPI`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
