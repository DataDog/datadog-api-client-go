// Get all dashboards returns "OK" response with pagination

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, _ := api.ListDashboardsWithPagination(ctx, *datadogV1.NewListDashboardsOptionalParameters().WithCount(2), )

        for paginationResult := range resp {
            if paginationResult.Error != nil {
                fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards`: %v\n", paginationResult.Error)
            }
            responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
            fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}