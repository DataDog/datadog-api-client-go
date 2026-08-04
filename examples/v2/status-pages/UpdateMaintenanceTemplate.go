// Update maintenance template returns "OK" response

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
	body := datadogV2.PatchMaintenanceTemplateRequest{
		Data: &datadogV2.PatchMaintenanceTemplateRequestData{
			Attributes: &datadogV2.PatchMaintenanceTemplateRequestDataAttributes{
				ComponentIds: []string{},
			},
			Id:   "",
			Type: datadogV2.PATCHMAINTENANCETEMPLATEREQUESTDATATYPE_MAINTENANCE_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateMaintenanceTemplate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewUpdateMaintenanceTemplateOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateMaintenanceTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateMaintenanceTemplate`:\n%s\n", responseContent)
}
