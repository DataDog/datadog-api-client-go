// Update status page returns "OK" response

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

	body := datadogV2.PatchStatusPageRequest{
		Data: &datadogV2.PatchStatusPageRequestData{
			Attributes: &datadogV2.PatchStatusPageRequestDataAttributes{
				Name: datadog.PtrString("[DD Integration Tests] 5e2fd69be33e79aa"),
			},
			Id:   datadog.PtrString(StatusPageDataID),
			Type: datadogV2.STATUSPAGEDATATYPE_STATUS_PAGES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateStatusPage(ctx, StatusPageDataID, body, *datadogV2.NewUpdateStatusPageOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateStatusPage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateStatusPage`:\n%s\n", responseContent)
}
