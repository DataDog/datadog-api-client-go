// Update a dashboard list returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "dashboard_list" in the system
	DashboardListID, _ := strconv.ParseInt(os.Getenv("DASHBOARD_LIST_ID"), 10, 64)

	body := datadog.DashboardList{
		Name: "updated Example-Update_a_dashboard_list_returns_OK_response",
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DashboardListsApi(apiClient)
	resp, r, err := api.UpdateDashboardList(ctx, DashboardListID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardList`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.UpdateDashboardList`:\n%s\n", responseContent)
}
