// Update a downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "downtime" in the system
	DowntimeID, _ := strconv.ParseInt(os.Getenv("DOWNTIME_ID"), 10, 64)

	body := datadog.Downtime{
		Message:                       datadog.PtrString("Example-Update_a_downtime_returns_OK_response-updated"),
		MuteFirstRecoveryNotification: datadog.PtrBool(true),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DowntimesApi.UpdateDowntime(ctx, DowntimeID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.UpdateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.UpdateDowntime`:\n%s\n", responseContent)
}
