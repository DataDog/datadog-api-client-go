// Delete a RUM operation returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteRUMOperation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMOperationsApi(apiClient)
	r, err := api.DeleteRUMOperation(ctx, "rum_operation_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMOperationsApi.DeleteRUMOperation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
