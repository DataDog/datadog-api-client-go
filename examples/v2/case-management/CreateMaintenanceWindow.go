// Create a maintenance window returns "Created" response

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
	body := datadogV2.MaintenanceWindowCreateRequest{
		Data: datadogV2.MaintenanceWindowCreate{
			Attributes: datadogV2.MaintenanceWindowCreateAttributes{
				EndAt:   time.Date(2026, 6, 1, 6, 0, 0, 0, time.UTC),
				Name:    "Weekly maintenance",
				Query:   "project:SEC",
				StartAt: time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			Type: datadogV2.MAINTENANCEWINDOWRESOURCETYPE_MAINTENANCE_WINDOW,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateMaintenanceWindow(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateMaintenanceWindow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateMaintenanceWindow`:\n%s\n", responseContent)
}
