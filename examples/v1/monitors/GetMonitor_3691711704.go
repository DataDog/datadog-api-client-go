// Get a synthetics monitor's details

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
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestMonitorID, _ := strconv.ParseInt(os.Getenv("SYNTHETICS_API_TEST_MONITOR_ID"), 10, 64)

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.MonitorsApi(apiClient)
	resp, r, err := api.GetMonitor(ctx, SyntheticsAPITestMonitorID, *datadog.NewGetMonitorOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitor`:\n%s\n", responseContent)
}
