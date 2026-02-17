// Update the tags for an interface returns "OK" response

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
	body := datadogV2.ListInterfaceTagsResponse{
		Data: &datadogV2.ListInterfaceTagsResponseData{
			Attributes: &datadogV2.ListTagsResponseDataAttributes{
				Tags: []string{
					"tag:test",
					"tag:testbis",
				},
			},
			Id:   datadog.PtrString("example:1.2.3.4:1"),
			Type: datadog.PtrString("tags"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewNetworkDeviceMonitoringApi(apiClient)
	resp, r, err := api.UpdateInterfaceUserTags(ctx, "example:1.2.3.4:1", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceMonitoringApi.UpdateInterfaceUserTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceMonitoringApi.UpdateInterfaceUserTags`:\n%s\n", responseContent)
}
