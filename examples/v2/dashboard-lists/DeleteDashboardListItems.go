// Delete items from a dashboard list returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.DashboardListDeleteItemsRequest{
		Dashboards: []datadog.DashboardListItemRequest{
			{
				Id:   "q5j-nti-fv6",
				Type: datadog.DASHBOARDTYPE_HOST_TIMEBOARD,
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DashboardListsApi(apiClient)
	resp, r, err := api.DeleteDashboardListItems(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardListItems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.DeleteDashboardListItems`:\n%s\n", responseContent)
}
