// Schedule maintenance returns "Created" response

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
	body := datadogV2.CreateMaintenanceRequest{
		Data: &datadogV2.CreateMaintenanceRequestData{
			Attributes: datadogV2.CreateMaintenanceRequestDataAttributes{
				CompletedDate:        datadog.PtrTime(time.Date(2026, 2, 18, 19, 51, 13, 332360, time.UTC)),
				CompletedDescription: datadog.PtrString("We have completed maintenance on the API to improve performance."),
				ComponentsAffected: []datadogV2.CreateMaintenanceRequestDataAttributesComponentsAffectedItems{
					{
						Id:     uuid.MustParse("1234abcd-12ab-34cd-56ef-123456abcdef"),
						Status: datadogV2.PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
					},
				},
				InProgressDescription: datadog.PtrString("We are currently performing maintenance on the API to improve performance."),
				ScheduledDescription:  datadog.PtrString("We will be performing maintenance on the API to improve performance."),
				StartDate:             datadog.PtrTime(time.Date(2026, 2, 18, 19, 21, 13, 332360, time.UTC)),
				Title:                 "API Maintenance",
			},
			Type: datadogV2.PATCHMAINTENANCEREQUESTDATATYPE_MAINTENANCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateMaintenance(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewCreateMaintenanceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateMaintenance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateMaintenance`:\n%s\n", responseContent)
}
