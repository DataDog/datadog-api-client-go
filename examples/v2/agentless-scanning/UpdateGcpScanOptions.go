// Update GCP scan options returns "OK" response

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
	body := datadogV2.GcpScanOptionsInputUpdate{
		Data: &datadogV2.GcpScanOptionsInputUpdateData{
			Id:   "api-spec-test",
			Type: datadogV2.GCPSCANOPTIONSINPUTUPDATEDATATYPE_GCP_SCAN_OPTIONS,
			Attributes: &datadogV2.GcpScanOptionsInputUpdateDataAttributes{
				VulnContainersOs: datadog.PtrBool(false),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.UpdateGcpScanOptions(ctx, "api-spec-test", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.UpdateGcpScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.UpdateGcpScanOptions`:\n%s\n", responseContent)
}
