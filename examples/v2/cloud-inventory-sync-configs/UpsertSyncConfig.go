// Create or update a sync configuration returns "OK" response

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
	body := datadogV2.UpsertCloudInventorySyncConfigRequest{
		Data: datadogV2.UpsertCloudInventorySyncConfigRequestData{
			Attributes: datadogV2.UpsertCloudInventorySyncConfigRequestAttributes{
				Aws: &datadogV2.CloudInventorySyncConfigAWSRequestAttributes{
					AwsAccountId:            "123456789012",
					DestinationBucketName:   "my-inventory-bucket",
					DestinationBucketRegion: "us-east-1",
					DestinationPrefix:       datadog.PtrString("logs/"),
				},
				Azure: &datadogV2.CloudInventorySyncConfigAzureRequestAttributes{
					ClientId:       "11111111-1111-1111-1111-111111111111",
					Container:      "inventory-container",
					ResourceGroup:  "my-resource-group",
					StorageAccount: "mystorageaccount",
					SubscriptionId: "33333333-3333-3333-3333-333333333333",
					TenantId:       "22222222-2222-2222-2222-222222222222",
				},
				Gcp: &datadogV2.CloudInventorySyncConfigGCPRequestAttributes{
					DestinationBucketName: "my-inventory-reports",
					ProjectId:             "my-gcp-project",
					ServiceAccountEmail:   "reader@my-gcp-project.iam.gserviceaccount.com",
					SourceBucketName:      "my-monitored-bucket",
				},
			},
			Id:   datadogV2.CLOUDINVENTORYCLOUDPROVIDERID_AWS,
			Type: datadogV2.CLOUDINVENTORYCLOUDPROVIDERREQUESTTYPE_CLOUD_PROVIDER,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertSyncConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudInventorySyncConfigsApi(apiClient)
	resp, r, err := api.UpsertSyncConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInventorySyncConfigsApi.UpsertSyncConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudInventorySyncConfigsApi.UpsertSyncConfig`:\n%s\n", responseContent)
}
