// Get deleted dashboards returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewDashboardsApi(apiClient)
	resp, r, err := api.ListDashboards(ctx, *datadog.NewListDashboardsOptionalParameters().WithFilterDeleted(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.ListDashboards`:\n%s\n", responseContent)
}
