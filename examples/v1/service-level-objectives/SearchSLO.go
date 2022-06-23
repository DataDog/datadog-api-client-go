// Search for SLOs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "slo" in the system
	SloData0Name := os.Getenv("SLO_DATA_0_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("SearchSLO", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectivesApi.SearchSLO(ctx, *datadog.NewSearchSLOOptionalParameters().WithQuery(SloData0Name).WithPageSize(20).WithPageNumber(0))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.SearchSLO`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.SearchSLO`:\n%s\n", responseContent)
}
