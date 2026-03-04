// Update maintenance returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "status_page" in the system
	StatusPageDataID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ID"))

	// there is a valid "maintenance" in the system
	MaintenanceDataID := uuid.MustParse(os.Getenv("MAINTENANCE_DATA_ID"))

	body := datadogV2.PatchMaintenanceRequest{
		Data: &datadogV2.PatchMaintenanceRequestData{
			Attributes: datadogV2.PatchMaintenanceRequestDataAttributes{
				ScheduledDescription:  datadog.PtrString("We will be performing maintenance on the API to improve performance for 40 minutes."),
				InProgressDescription: datadog.PtrString("We are currently performing maintenance on the API to improve performance for 40 minutes."),
			},
			Id:   MaintenanceDataID,
			Type: datadogV2.PATCHMAINTENANCEREQUESTDATATYPE_MAINTENANCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateMaintenance(ctx, StatusPageDataID, MaintenanceDataID, body, *datadogV2.NewUpdateMaintenanceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateMaintenance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateMaintenance`:\n%s\n", responseContent)
}
