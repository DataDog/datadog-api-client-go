// Update a batch migration resource returns "OK" response

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
	// there is a valid "batch_migration" in the system
	BatchMigrationDataID := uuid.MustParse(os.Getenv("BATCH_MIGRATION_DATA_ID"))

	// there is a valid "batch_migration_resource" in the system
	BatchMigrationResourceID := uuid.MustParse(os.Getenv("BATCH_MIGRATION_RESOURCE_ID"))

	body := datadogV2.SecurityMonitoringBatchMigrationResourceUpdateRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationResourceUpdateData{
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONRESOURCETYPE_BATCH_MIGRATION_RESOURCE,
			Attributes: datadogV2.SecurityMonitoringBatchMigrationResourceUpdateAttributes{
				Input: datadog.PtrString("index=wineventlog EventCode=4625 | stats count by user"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityMonitoringBatchMigrationResource", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringBatchMigrationResource(ctx, BatchMigrationDataID, BatchMigrationResourceID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringBatchMigrationResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringBatchMigrationResource`:\n%s\n", responseContent)
}
