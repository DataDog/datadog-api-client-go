// Get an archive configuration returns "OK" response

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
	// there is a valid "rum_application" in the system
	RumApplicationDataID := os.Getenv("RUM_APPLICATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMManagedArchivesApi(apiClient)
	resp, r, err := api.GetRUMArchiveConfig(ctx, RumApplicationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMManagedArchivesApi.GetRUMArchiveConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMManagedArchivesApi.GetRUMArchiveConfig`:\n%s\n", responseContent)
}
