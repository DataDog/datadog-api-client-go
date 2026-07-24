// Delete a RUM operation strong link returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteRUMOperationStrongLink", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	r, err := api.DeleteRUMOperationStrongLink(ctx, "rum_operation_id", "feature_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.DeleteRUMOperationStrongLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
