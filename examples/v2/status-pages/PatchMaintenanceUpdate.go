// Edit maintenance update returns "OK" response

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
	body := datadogV2.PatchMaintenanceUpdateRequest{
		Data: &datadogV2.PatchMaintenanceUpdateRequestData{
			Attributes: &datadogV2.PatchMaintenanceUpdateRequestDataAttributes{
				Description: datadog.PtrString("We have completed maintenance on the API to improve performance."),
			},
			Id:   "00000000-0000-0000-0000-000000000000",
			Type: datadogV2.PATCHMAINTENANCEUPDATEREQUESTDATATYPE_MAINTENANCE_UPDATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.PatchMaintenanceUpdate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.PatchMaintenanceUpdate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.PatchMaintenanceUpdate`:\n%s\n", responseContent)
}
