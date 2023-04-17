// Update a downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "downtime" in the system
	DowntimeID, _ := strconv.ParseInt(os.Getenv("DOWNTIME_ID"), 10, 64)

	body := datadogV1.Downtime{
		Message:                       *datadog.NewNullableString(datadog.PtrString("Example-Downtime-updated")),
		MuteFirstRecoveryNotification: datadog.PtrBool(true),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	resp, r, err := api.UpdateDowntime(ctx, DowntimeID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.UpdateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.UpdateDowntime`:\n%s\n", responseContent)
}
