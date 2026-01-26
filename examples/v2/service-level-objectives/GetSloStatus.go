// Get SLO status returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetSloStatus", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.GetSloStatus(ctx, "00000000-0000-0000-0000-000000000000", 1690901870, 1706803070, *datadogV2.NewGetSloStatusOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSloStatus`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSloStatus`:\n%s\n", responseContent)
}
