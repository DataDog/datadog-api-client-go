// Create a batch migration returns "Created" response

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
	body := datadogV2.SecurityMonitoringBatchMigrationCreateRequest{
		Data: datadogV2.SecurityMonitoringBatchMigrationCreateData{
			Type: datadogV2.SECURITYMONITORINGBATCHMIGRATIONTYPE_BATCH_MIGRATION,
			Attributes: datadogV2.SecurityMonitoringBatchMigrationCreateAttributes{
				Name:        datadog.PtrString("Example-Security-Monitoring"),
				Description: datadog.PtrString("Rules imported from Splunk"),
				Tags: []string{
					"source:splunk",
				},
				Resources: []datadogV2.SecurityMonitoringBatchMigrationResourceCreateItem{
					{
						Name:        datadog.PtrString("Repeated Windows login failures"),
						Description: datadog.PtrString("Detects repeated failed Windows authentication attempts."),
						InputType:   datadogV2.SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_SPL,
						Input:       "index=wineventlog EventCode=4625",
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityMonitoringBatchMigration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringBatchMigration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringBatchMigration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringBatchMigration`:\n%s\n", responseContent)
}
