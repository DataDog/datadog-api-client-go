// Update a maintenance window returns "OK" response

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
	body := datadogV2.MaintenanceWindowUpdateRequest{
		Data: datadogV2.MaintenanceWindowUpdate{
			Attributes: &datadogV2.MaintenanceWindowUpdateAttributes{},
			Type:       datadogV2.MAINTENANCEWINDOWRESOURCETYPE_MAINTENANCE_WINDOW,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.UpdateMaintenanceWindow(ctx, "maintenance_window_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdateMaintenanceWindow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.UpdateMaintenanceWindow`:\n%s\n", responseContent)
}
