// Delete a powerpack returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "powerpack" in the system
	PowerpackDataID := os.Getenv("POWERPACK_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewPowerpackApi(apiClient)
	r, err := api.DeletePowerpack(ctx, PowerpackDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerpackApi.DeletePowerpack`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
