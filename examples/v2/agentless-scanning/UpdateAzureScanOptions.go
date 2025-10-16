// Update azure scan options returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.AzureScanOptionsInputUpdate{
Data: &datadogV2.AzureScanOptionsInputUpdateData{
Id: "12345678-90ab-cdef-1234-567890abcdef",
Type: datadogV2.AZURESCANOPTIONSINPUTUPDATEDATATYPE_AZURE_SCAN_OPTIONS,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.UpdateAzureScanOptions(ctx, "12345678-90ab-cdef-1234-567890abcdef", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.UpdateAzureScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.UpdateAzureScanOptions`:\n%s\n", responseContent)
}