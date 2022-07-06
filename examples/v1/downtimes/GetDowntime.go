// Get a downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
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
	resp, r, err := api.GetDowntime(ctx, DowntimeID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.GetDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.GetDowntime`:\n%s\n", responseContent)
}
