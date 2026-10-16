// Publish a batch migration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.SecurityMonitoringBatchMigrationPublishRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationPublishData{
			Attributes: datadogV2.SecurityMonitoringBatchMigrationPublishAttributes{
				ResourceIds: []uuid.UUID{
					uuid.MustParse("c968cf83-56c2-4d54-a993-6c86dfb83527"),
				},
			},
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONPUBLISHTYPE_BATCH_MIGRATION_PUBLISH,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.PublishSecurityMonitoringBatchMigration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.PublishSecurityMonitoringBatchMigration(ctx, uuid.MustParse("7cdb3a75-64fb-46ce-9445-41969e78fe8e"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.PublishSecurityMonitoringBatchMigration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.PublishSecurityMonitoringBatchMigration`:\n%s\n", responseContent)
}
