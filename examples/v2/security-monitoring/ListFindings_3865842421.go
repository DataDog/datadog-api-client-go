// List findings returns "OK" response with pagination

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
	configuration.SetUnstableOperationEnabled("v2.ListFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, _ := api.ListFindingsWithPagination(ctx, *datadogV2.NewListFindingsOptionalParameters())

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListFindings`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
