// Delete component returns "No Content" response

package main

import (
	"context"
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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	r, err := api.DeleteComponent(ctx, StatusPageDataID, StatusPageDataAttributesComponents0ID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.DeleteComponent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
