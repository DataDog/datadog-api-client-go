// Delete resources from a batch migration returns "No Content" response

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
	// there is a valid "batch_migration" in the system
	BatchMigrationDataID := uuid.MustParse(os.Getenv("BATCH_MIGRATION_DATA_ID"))

	// there is a valid "batch_migration_resource" in the system
	BatchMigrationResourceID := uuid.MustParse(os.Getenv("BATCH_MIGRATION_RESOURCE_ID"))

	body := datadogV2.SecurityMonitoringBatchMigrationResourcesDeleteRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationResourcesDeleteData{
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONRESOURCESTYPE_BATCH_MIGRATION_RESOURCES,
			Attributes: datadogV2.SecurityMonitoringBatchMigrationResourcesDeleteAttributes{
				ResourceIds: []uuid.UUID{
					BatchMigrationResourceID,
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteSecurityMonitoringBatchMigrationResources", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.DeleteSecurityMonitoringBatchMigrationResources(ctx, BatchMigrationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.DeleteSecurityMonitoringBatchMigrationResources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
