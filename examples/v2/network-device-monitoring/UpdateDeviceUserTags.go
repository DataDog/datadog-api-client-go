// Update the tags for a device returns "OK" response

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
	body := datadogV2.ListTagsResponse{
		Data: &datadogV2.ListTagsResponseData{
			Attributes: &datadogV2.ListTagsResponseDataAttributes{
				Tags: []string{
					"tag:test",
					"tag:testbis",
				},
			},
			Id:   datadog.PtrString("default_device"),
			Type: datadog.PtrString("tags"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewNetworkDeviceMonitoringApi(apiClient)
	resp, r, err := api.UpdateDeviceUserTags(ctx, "default_device", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceMonitoringApi.UpdateDeviceUserTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceMonitoringApi.UpdateDeviceUserTags`:\n%s\n", responseContent)
}
