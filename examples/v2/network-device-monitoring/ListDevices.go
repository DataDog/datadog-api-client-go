// Get the list of devices returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewNetworkDeviceMonitoringApi(apiClient)
	resp, r, err := api.ListDevices(ctx, *datadogV2.NewListDevicesOptionalParameters().WithPageSize(1).WithPageNumber(0).WithFilterTag("device_namespace:default"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceMonitoringApi.ListDevices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceMonitoringApi.ListDevices`:\n%s\n", responseContent)
}
