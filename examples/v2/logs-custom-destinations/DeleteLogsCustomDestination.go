// Delete a custom destination returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "custom_destination" in the system
	CustomDestinationDataID := os.Getenv("CUSTOM_DESTINATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	r, err := api.DeleteLogsCustomDestination(ctx, CustomDestinationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.DeleteLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
