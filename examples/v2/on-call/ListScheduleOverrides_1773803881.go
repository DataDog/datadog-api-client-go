// List On-Call schedule overrides returns "OK" response with pagination

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, _ := api.ListScheduleOverridesWithPagination(ctx, "3653d3c6-0c75-11ea-ad28-fb5701eabc7d", time.Date(2024, 1, 7, 2, 53, 1, 0, time.UTC), time.Date(2024, 1, 14, 2, 53, 1, 0, time.UTC), *datadogV2.NewListScheduleOverridesOptionalParameters())

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.ListScheduleOverrides`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
