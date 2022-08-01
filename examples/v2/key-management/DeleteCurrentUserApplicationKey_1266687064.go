// Delete an Application key owned by current user returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	// there is a valid "application_key" in the system
	ApplicationKeyDataID := os.Getenv("APPLICATION_KEY_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewKeyManagementApi(apiClient)
	r, err := api.DeleteCurrentUserApplicationKey(ctx, ApplicationKeyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteCurrentUserApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
