// Delete an Amazon EventBridge source returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.AWSEventBridgeDeleteRequest{
		AccountId:          datadog.PtrString("123456789012"),
		EventGeneratorName: datadog.PtrString("app-alerts-zyxw3210"),
		Region:             datadog.PtrString("us-east-1"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.DeleteAWSEventBridgeSource(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSEventBridgeSource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.DeleteAWSEventBridgeSource`:\n%s\n", responseContent)
}
