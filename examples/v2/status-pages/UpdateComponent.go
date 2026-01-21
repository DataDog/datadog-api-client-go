// Update component returns "OK" response

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
	StatusPageDataAttributesComponents0ID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ATTRIBUTES_COMPONENTS_0_ID"))
	StatusPageDataID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ID"))

	body := datadogV2.PatchComponentRequest{
		Data: &datadogV2.PatchComponentRequestData{
			Attributes: &datadogV2.PatchComponentRequestDataAttributes{
				Name: datadog.PtrString("Logs Indexing"),
			},
			Id:   &StatusPageDataAttributesComponents0ID,
			Type: datadogV2.STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.UpdateComponent(ctx, StatusPageDataID, StatusPageDataAttributesComponents0ID, body, *datadogV2.NewUpdateComponentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.UpdateComponent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.UpdateComponent`:\n%s\n", responseContent)
}
