// Get a synthetics monitor's details

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestMonitorID, _ := strconv.ParseInt(os.Getenv("SYNTHETICS_API_TEST_MONITOR_ID"), 10, 64)


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.GetMonitor(ctx, SyntheticsAPITestMonitorID, *datadogV1.NewGetMonitorOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitor`:\n%s\n", responseContent)
}