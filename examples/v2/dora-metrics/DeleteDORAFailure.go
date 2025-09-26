// Delete a failure event returns "Accepted" response

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
	api := datadogV2.NewDORAMetricsApi(apiClient)
	r, err := api.DeleteDORAFailure(ctx, "NO_VALUE")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DORAMetricsApi.DeleteDORAFailure`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
