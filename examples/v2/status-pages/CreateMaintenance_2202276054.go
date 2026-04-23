// Create maintenance returns "Created" response

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

	body := datadogV2.CreateMaintenanceRequest{
		Data: &datadogV2.CreateMaintenanceRequestData{
			Attributes: datadogV2.CreateMaintenanceRequestDataAttributes{
				Title:                 "API Maintenance",
				ScheduledDescription:  "We will be performing maintenance on the API to improve performance.",
				InProgressDescription: "We are currently performing maintenance on the API to improve performance.",
				CompletedDescription:  "We have completed maintenance on the API to improve performance.",
				StartDate:             time.Now().Add(time.Hour * 1),
				CompletedDate:         time.Now().Add(time.Hour * 2),
				ComponentsAffected: []datadogV2.CreateMaintenanceRequestDataAttributesComponentsAffectedItems{
					{
						Id:     StatusPageDataAttributesComponents0Components0ID,
						Status: datadogV2.PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
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
	resp, r, err := api.CreateMaintenance(ctx, StatusPageDataID, body, *datadogV2.NewCreateMaintenanceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateMaintenance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateMaintenance`:\n%s\n", responseContent)
}
