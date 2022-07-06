// Cancel a downtime returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "downtime" in the system
	DowntimeID, _ := strconv.ParseInt(os.Getenv("DOWNTIME_ID"), 10, 64)

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DowntimesApi(apiClient)
	r, err := api.CancelDowntime(ctx, DowntimeID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CancelDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
