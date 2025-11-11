// Delete a restriction query returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteRestrictionQuery", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsRestrictionQueriesApi(apiClient)
	r, err := api.DeleteRestrictionQuery(ctx, "restriction_query_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsRestrictionQueriesApi.DeleteRestrictionQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
