// Add resources to a batch migration returns "Created" response

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

	body := datadogV2.SecurityMonitoringBatchMigrationResourcesCreateRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationResourcesCreateData{
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONRESOURCESTYPE_BATCH_MIGRATION_RESOURCES,
			Attributes: datadogV2.SecurityMonitoringBatchMigrationResourcesCreateAttributes{
				Resources: []datadogV2.SecurityMonitoringBatchMigrationResourceCreateItem{
					{
						Name:      datadog.PtrString("Linux sudo abuse"),
						InputType: datadogV2.SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_SPL,
						Input:     `index=linux process=sudo "authentication failure" | stats count by user, host`,
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityMonitoringBatchMigrationResources", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringBatchMigrationResources(ctx, BatchMigrationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringBatchMigrationResources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringBatchMigrationResources`:\n%s\n", responseContent)
}
