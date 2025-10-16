// Get a list of metrics returns "Success" response with pagination

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, _ := api.ListTagConfigurationsWithPagination(ctx, *datadogV2.NewListTagConfigurationsOptionalParameters().WithPageSize(2), )

        for paginationResult := range resp {
            if paginationResult.Error != nil {
                fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations`: %v\n", paginationResult.Error)
            }
            responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
            fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}