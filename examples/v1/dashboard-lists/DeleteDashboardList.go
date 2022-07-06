// Delete a dashboard list returns "OK" response

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

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DashboardListsApi(apiClient)
	resp, r, err := api.DeleteDashboardList(ctx, DashboardListID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardList`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.DeleteDashboardList`:\n%s\n", responseContent)
}
