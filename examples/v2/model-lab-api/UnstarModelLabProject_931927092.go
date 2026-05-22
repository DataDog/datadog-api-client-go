// Unstar a Model Lab project returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.UnstarModelLabProject", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewModelLabAPIApi(apiClient)
	r, err := api.UnstarModelLabProject(ctx, 2387)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelLabAPIApi.UnstarModelLabProject`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
