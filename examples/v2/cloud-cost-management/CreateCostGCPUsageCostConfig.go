// Create Cloud Cost Management GCP Usage Cost config returns "OK" response

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
	body := datadogV2.GCPUsageCostConfigPostRequest{
		Data: datadogV2.GCPUsageCostConfigPostData{
			Attributes: datadogV2.GCPUsageCostConfigPostRequestAttributes{
				BillingAccountId:  "123456_A123BC_12AB34",
				BucketName:        "dd-cost-bucket",
				ExportDatasetName: "billing",
				ExportPrefix:      datadog.PtrString("datadog_cloud_cost_usage_export"),
				ExportProjectName: "dd-cloud-cost-report",
				ServiceAccount:    "dd-ccm-gcp-integration@my-environment.iam.gserviceaccount.com",
			},
			Type: datadogV2.GCPUSAGECOSTCONFIGPOSTREQUESTTYPE_GCP_USAGE_COST_CONFIG_POST_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateCostGCPUsageCostConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateCostGCPUsageCostConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateCostGCPUsageCostConfig`:\n%s\n", responseContent)
}
