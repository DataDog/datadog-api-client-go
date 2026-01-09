// Update degradation returns "OK" response

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
	// there is a valid "status_page" in the system
	StatusPageDataID := os.Getenv("STATUS_PAGE_DATA_ID")

	// there is a valid "degradation" in the system
	DegradationDataID := os.Getenv("DEGRADATION_DATA_ID")

	body := datadogV2.PatchDegradationRequest{
		Data: &datadogV2.PatchDegradationRequestData{
			Attributes: &datadogV2.PatchDegradationRequestDataAttributes{
				Title: datadog.PtrString("5e2fd69be33e79aa"),
			},
			Id:   datadog.PtrString(DegradationDataID),
			Type: datadogV2.PATCHDEGRADATIONREQUESTDATATYPE_DEGRADATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateDegradation(ctx, StatusPageDataID, DegradationDataID, body, *datadogV2.NewUpdateDegradationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateDegradation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateDegradation`:\n%s\n", responseContent)
}
