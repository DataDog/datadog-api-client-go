// Update a batch migration returns "OK" response

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

	body := datadogV2.SecurityMonitoringBatchMigrationUpdateRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationUpdateData{
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONTYPE_BATCH_MIGRATION,
			Attributes: datadogV2.SecurityMonitoringBatchMigrationUpdateAttributes{
				Name: datadog.PtrString("Example-Security-Monitoring-updated"),
				Tags: []string{
					"source:splunk",
					"team:secops",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityMonitoringBatchMigration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringBatchMigration(ctx, BatchMigrationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringBatchMigration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringBatchMigration`:\n%s\n", responseContent)
}
