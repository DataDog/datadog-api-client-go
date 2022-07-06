// Add custom screenboard dashboard to an existing dashboard list returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "dashboard_list" in the system
	DashboardListID, _ := strconv.ParseInt(os.Getenv("DASHBOARD_LIST_ID"), 10, 64)

	// there is a valid "screenboard_dashboard" in the system
	ScreenboardDashboardID := os.Getenv("SCREENBOARD_DASHBOARD_ID")

	body := datadog.DashboardListAddItemsRequest{
		Dashboards: []datadog.DashboardListItemRequest{
			{
				Id:   ScreenboardDashboardID,
				Type: datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD,
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DashboardListsApi(apiClient)
	resp, r, err := api.CreateDashboardListItems(ctx, DashboardListID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.CreateDashboardListItems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.CreateDashboardListItems`:\n%s\n", responseContent)
}
