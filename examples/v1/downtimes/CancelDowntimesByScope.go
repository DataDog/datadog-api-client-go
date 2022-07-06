// Cancel downtimes by scope returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "downtime" in the system
	DowntimeScope0 := os.Getenv("DOWNTIME_SCOPE_0")

	body := datadog.CancelDowntimesByScopeRequest{
		Scope: DowntimeScope0,
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DowntimesApi(apiClient)
	resp, r, err := api.CancelDowntimesByScope(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CancelDowntimesByScope`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CancelDowntimesByScope`:\n%s\n", responseContent)
}
