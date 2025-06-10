// Delete an override returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "schedule" in the system
	ScheduleDataID := os.Getenv("SCHEDULE_DATA_ID")

	// there is a valid "override" in the system
	OverrideData0ID := os.Getenv("OVERRIDE_DATA_0_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	r, err := api.DeleteOnCallScheduleOverride(ctx, ScheduleDataID, OverrideData0ID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.DeleteOnCallScheduleOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
