// Delete a custom destination returns "OK" response

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
	// there is a valid "logs_custom_destination" in the system
	LogsCustomDestinationDataID := os.Getenv("LOGS_CUSTOM_DESTINATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLogsCustomDestination", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.DeleteLogsCustomDestination(ctx, LogsCustomDestinationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.DeleteLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.DeleteLogsCustomDestination`:\n%s\n", responseContent)
}
