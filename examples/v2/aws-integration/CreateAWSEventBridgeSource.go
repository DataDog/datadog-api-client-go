// Create an Amazon EventBridge source returns "Amazon EventBridge source created." response

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
	body := datadogV2.AWSEventBridgeCreateRequest{
		Data: datadogV2.AWSEventBridgeCreateRequestData{
			Attributes: datadogV2.AWSEventBridgeCreateRequestAttributes{
				AccountId:          "123456789012",
				CreateEventBus:     datadog.PtrBool(true),
				EventGeneratorName: "app-alerts",
				Region:             "us-east-1",
			},
			Type: datadogV2.AWSEVENTBRIDGETYPE_EVENT_BRIDGE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateAWSEventBridgeSource(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSEventBridgeSource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateAWSEventBridgeSource`:\n%s\n", responseContent)
}
