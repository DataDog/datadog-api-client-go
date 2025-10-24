// Create AWS scan options returns "Agentless scan options enabled successfully." response

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
	body := datadogV2.AwsScanOptionsCreateRequest{
		Data: datadogV2.AwsScanOptionsCreateData{
			Id:   "000000000003",
			Type: datadogV2.AWSSCANOPTIONSTYPE_AWS_SCAN_OPTIONS,
			Attributes: datadogV2.AwsScanOptionsCreateAttributes{
				Lambda:           true,
				SensitiveData:    false,
				VulnContainersOs: true,
				VulnHostOs:       true,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.CreateAwsScanOptions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.CreateAwsScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.CreateAwsScanOptions`:\n%s\n", responseContent)
}
