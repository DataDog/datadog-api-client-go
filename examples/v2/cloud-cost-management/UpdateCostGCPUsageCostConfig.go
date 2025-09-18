// Update Google Cloud Usage Cost config returns "OK" response

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
	body := datadogV2.GCPUsageCostConfigPatchRequest{
		Data: datadogV2.GCPUsageCostConfigPatchData{
			Attributes: datadogV2.GCPUsageCostConfigPatchRequestAttributes{
				IsEnabled: true,
			},
			Type: datadogV2.GCPUSAGECOSTCONFIGPATCHREQUESTTYPE_GCP_USAGE_COST_CONFIG_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCostGCPUsageCostConfig(ctx, 100, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCostGCPUsageCostConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCostGCPUsageCostConfig`:\n%s\n", responseContent)
}
