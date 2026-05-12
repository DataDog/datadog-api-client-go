// Create backfilled maintenance returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "status_page" in the system
	StatusPageDataAttributesComponents0Components0ID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ATTRIBUTES_COMPONENTS_0_COMPONENTS_0_ID"))
	StatusPageDataID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ID"))

	body := datadogV2.CreateBackfilledMaintenanceRequest{
		Data: &datadogV2.CreateBackfilledMaintenanceRequestData{
			Attributes: &datadogV2.CreateBackfilledMaintenanceRequestDataAttributes{
				Title: "Past Database Maintenance",
				Updates: []datadogV2.CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems{
					{
						ComponentsAffected: []datadogV2.CreateMaintenanceRequestDataAttributesComponentsAffectedItems{
							{
								Id:     StatusPageDataAttributesComponents0Components0ID,
								Status: datadogV2.PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_MAINTENANCE,
							},
						},
						Description: "Database maintenance is in progress.",
						StartedAt:   time.Now().Add(time.Hour * -1),
						Status:      datadogV2.CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_IN_PROGRESS,
					},
					{
						ComponentsAffected: []datadogV2.CreateMaintenanceRequestDataAttributesComponentsAffectedItems{
							{
								Id:     StatusPageDataAttributesComponents0Components0ID,
								Status: datadogV2.PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
							},
						},
						Description: "Database maintenance has been completed successfully.",
						StartedAt:   time.Now(),
						Status:      datadogV2.CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_COMPLETED,
					},
				},
			},
			Type: datadogV2.PATCHMAINTENANCEREQUESTDATATYPE_MAINTENANCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateBackfilledMaintenance(ctx, StatusPageDataID, body, *datadogV2.NewCreateBackfilledMaintenanceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateBackfilledMaintenance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateBackfilledMaintenance`:\n%s\n", responseContent)
}
