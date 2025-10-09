// Post GCP Scan Options returns "Agentless scan options enabled successfully." response

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
	body := datadogV2.GcpScanOptions{
		Data: &datadogV2.GcpScanOptionsData{
			Id:   "new-project",
			Type: datadogV2.GCPSCANOPTIONSDATATYPE_GCP_SCAN_OPTIONS,
			Attributes: &datadogV2.GcpScanOptionsDataAttributes{
				VulnHostOs:       datadog.PtrBool(true),
				VulnContainersOs: datadog.PtrBool(true),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.CreateGcpScanOptions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.CreateGcpScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.CreateGcpScanOptions`:\n%s\n", responseContent)
}
