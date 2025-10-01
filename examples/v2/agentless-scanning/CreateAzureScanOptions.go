// Create azure scan options returns "Created" response

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
	body := datadogV2.AzureScanOptions{
		Data: &datadogV2.AzureScanOptionsData{
			Attributes: &datadogV2.AzureScanOptionsDataAttributes{
				VulnContainersOs: datadog.PtrBool(true),
				VulnHostOs:       datadog.PtrBool(true),
			},
			Id:   "12345678-90ab-cdef-1234-567890abcdef",
			Type: datadogV2.AZURESCANOPTIONSDATATYPE_AZURE_SCAN_OPTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.CreateAzureScanOptions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.CreateAzureScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.CreateAzureScanOptions`:\n%s\n", responseContent)
}
