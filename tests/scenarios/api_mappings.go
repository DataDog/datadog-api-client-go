// This file is autogenerated. Do not modify directly.

package scenarios

import (
	"reflect"

	v1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

	v2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

var apiMappings = map[string]map[string]reflect.Value{
	"v1": {
		"IPRangesApi":                         reflect.ValueOf(v1.IPRangesApi),
		"KeyManagementApi":                    reflect.ValueOf(v1.KeyManagementApi),
		"ServiceChecksApi":                    reflect.ValueOf(v1.ServiceChecksApi),
		"UsageMeteringApi":                    reflect.ValueOf(v1.UsageMeteringApi),
		"DashboardsApi":                       reflect.ValueOf(v1.DashboardsApi),
		"DashboardListsApi":                   reflect.ValueOf(v1.DashboardListsApi),
		"MetricsApi":                          reflect.ValueOf(v1.MetricsApi),
		"DowntimesApi":                        reflect.ValueOf(v1.DowntimesApi),
		"EventsApi":                           reflect.ValueOf(v1.EventsApi),
		"SnapshotsApi":                        reflect.ValueOf(v1.SnapshotsApi),
		"HostsApi":                            reflect.ValueOf(v1.HostsApi),
		"AWSIntegrationApi":                   reflect.ValueOf(v1.AWSIntegrationApi),
		"AWSLogsIntegrationApi":               reflect.ValueOf(v1.AWSLogsIntegrationApi),
		"AzureIntegrationApi":                 reflect.ValueOf(v1.AzureIntegrationApi),
		"GCPIntegrationApi":                   reflect.ValueOf(v1.GCPIntegrationApi),
		"PagerDutyIntegrationApi":             reflect.ValueOf(v1.PagerDutyIntegrationApi),
		"SlackIntegrationApi":                 reflect.ValueOf(v1.SlackIntegrationApi),
		"WebhooksIntegrationApi":              reflect.ValueOf(v1.WebhooksIntegrationApi),
		"LogsApi":                             reflect.ValueOf(v1.LogsApi),
		"LogsIndexesApi":                      reflect.ValueOf(v1.LogsIndexesApi),
		"LogsPipelinesApi":                    reflect.ValueOf(v1.LogsPipelinesApi),
		"MonitorsApi":                         reflect.ValueOf(v1.MonitorsApi),
		"NotebooksApi":                        reflect.ValueOf(v1.NotebooksApi),
		"OrganizationsApi":                    reflect.ValueOf(v1.OrganizationsApi),
		"SecurityMonitoringApi":               reflect.ValueOf(v1.SecurityMonitoringApi),
		"ServiceLevelObjectivesApi":           reflect.ValueOf(v1.ServiceLevelObjectivesApi),
		"ServiceLevelObjectiveCorrectionsApi": reflect.ValueOf(v1.ServiceLevelObjectiveCorrectionsApi),
		"SyntheticsApi":                       reflect.ValueOf(v1.SyntheticsApi),
		"TagsApi":                             reflect.ValueOf(v1.TagsApi),
		"UsersApi":                            reflect.ValueOf(v1.UsersApi),
		"AuthenticationApi":                   reflect.ValueOf(v1.AuthenticationApi),
	},
	"v2": {
		"KeyManagementApi":         reflect.ValueOf(v2.KeyManagementApi),
		"AuditApi":                 reflect.ValueOf(v2.AuditApi),
		"AuthNMappingsApi":         reflect.ValueOf(v2.AuthNMappingsApi),
		"DashboardListsApi":        reflect.ValueOf(v2.DashboardListsApi),
		"IncidentsApi":             reflect.ValueOf(v2.IncidentsApi),
		"OpsgenieIntegrationApi":   reflect.ValueOf(v2.OpsgenieIntegrationApi),
		"LogsApi":                  reflect.ValueOf(v2.LogsApi),
		"LogsArchivesApi":          reflect.ValueOf(v2.LogsArchivesApi),
		"LogsMetricsApi":           reflect.ValueOf(v2.LogsMetricsApi),
		"MetricsApi":               reflect.ValueOf(v2.MetricsApi),
		"RolesApi":                 reflect.ValueOf(v2.RolesApi),
		"ProcessesApi":             reflect.ValueOf(v2.ProcessesApi),
		"RUMApi":                   reflect.ValueOf(v2.RUMApi),
		"OrganizationsApi":         reflect.ValueOf(v2.OrganizationsApi),
		"CloudWorkloadSecurityApi": reflect.ValueOf(v2.CloudWorkloadSecurityApi),
		"SecurityMonitoringApi":    reflect.ValueOf(v2.SecurityMonitoringApi),
		"UsersApi":                 reflect.ValueOf(v2.UsersApi),
		"ServiceAccountsApi":       reflect.ValueOf(v2.ServiceAccountsApi),
		"IncidentServicesApi":      reflect.ValueOf(v2.IncidentServicesApi),
		"IncidentTeamsApi":         reflect.ValueOf(v2.IncidentTeamsApi),
		"UsageMeteringApi":         reflect.ValueOf(v2.UsageMeteringApi),
	},
}
