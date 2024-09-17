// Get SLO report returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetSLOReport", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.GetSLOReport(ctx, "9fb2dc2a-ead0-11ee-a174-9fe3a9d7627f")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLOReport`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSLOReport`:\n%s\n", resp)
}
