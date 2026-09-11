// Update a Snowflake integration account returns "OK" response

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
	body := datadogV2.SnowflakeIntegrationAccountUpdateRequest{
		Data: datadogV2.SnowflakeIntegrationAccountUpdateData{
			Attributes: datadogV2.SnowflakeIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.SnowflakeIntegrationAccountAuthenticationRequest{
					AuthType: datadogV2.SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY,
					PrivateKey: `-----BEGIN PRIVATE KEY-----
MIIE...
-----END PRIVATE KEY-----`,
					PrivateKeyName:       "my-rsa-key",
					PrivateKeyPassphrase: datadog.PtrString("your-private-key-passphrase"),
				},
				Dataflows: &datadogV2.SnowflakeIntegrationDataflowsRequest{
					SnowflakeAccountUsageMetrics: &datadogV2.SnowflakeAccountUsageMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeAccountUsageMetricsIntegrationDataflowSettingsRequest{
							AccountUsageMetricsAggregateLast24h: datadog.PtrBool(false),
						},
					},
					SnowflakeCloudCostMetrics: &datadogV2.SnowflakeCloudCostMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest{
							QueryTags: datadog.PtrString("env,team,cost_center"),
						},
					},
					SnowflakeDataObservabilityQualityMonitoring: &datadogV2.SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsRequest{
							DoTableCrawlerCron:          datadog.PtrString("0 */6 * * *"),
							SyncSnowflakeSystemDatabase: datadog.PtrBool(true),
						},
					},
					SnowflakeEventTableLogs: &datadogV2.SnowflakeEventTableLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeEventTableLogsIntegrationDataflowSettingsRequest{
							EventTableEventsEnabled:     datadog.PtrBool(true),
							EventTableLogsEnabled:       datadog.PtrBool(true),
							EventTableLogsIntervalMin:   datadog.PtrInt64(15),
							EventTableSpanEventsEnabled: datadog.PtrBool(false),
							EventTableSpansEnabled:      datadog.PtrBool(false),
						},
					},
					SnowflakeOrganizationUsageMetrics: &datadogV2.SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest{
							OrganizationUsageMetricsAggregateLast24h: datadog.PtrBool(false),
						},
					},
					SnowflakeQueryHistoryLogs: &datadogV2.SnowflakeQueryHistoryLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest{
							JoinQueryHistoryWithAccessHistoryEnabled: datadog.PtrBool(true),
							QueryHistoryLogsIntervalMin:              datadog.PtrInt64(15),
						},
					},
					SnowflakeSecurityLogs: &datadogV2.SnowflakeSecurityLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeSecurityLogsIntegrationDataflowSettingsRequest{
							SecurityLogsIntervalMin: datadog.PtrInt64(60),
						},
					},
					SnowflakeTaskHistoryLogs: &datadogV2.SnowflakeTaskHistoryLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest{
							TaskHistoryLogsIntervalMin: datadog.PtrInt64(30),
						},
					},
				},
				Name: datadog.PtrString("prod-snowflake"),
				Settings: &datadogV2.SnowflakeIntegrationAccountSettingsUpdate{
					SnowflakeAccountIdentifier: datadog.PtrString("myorg-myaccount"),
					Username:                   datadog.PtrString("datadog_user"),
				},
			},
			Id:   "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSnowflakeIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSnowflakeIntegrationApi(apiClient)
	resp, r, err := api.UpdateSnowflakeIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnowflakeIntegrationApi.UpdateSnowflakeIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SnowflakeIntegrationApi.UpdateSnowflakeIntegrationAccount`:\n%s\n", responseContent)
}
