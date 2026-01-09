// Get component returns "OK" response

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
	StatusPageDataAttributesComponents0ID := os.Getenv("STATUS_PAGE_DATA_ATTRIBUTES_COMPONENTS_0_ID")
	StatusPageDataID := os.Getenv("STATUS_PAGE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.GetComponent(ctx, StatusPageDataID, StatusPageDataAttributesComponents0ID, *datadogV2.NewGetComponentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.GetComponent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.GetComponent`:\n%s\n", responseContent)
}
