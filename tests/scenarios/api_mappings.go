// This file is autogenerated. Do not modify directly.

package scenarios

import (
    "reflect"


    v1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

    v2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"

)

var apiMappings = map[string]map[string]reflect.Value{
	"v1": {
	    "IPRangesApi": reflect.ValueOf(v1.NewIPRangesApi),
	    "KeyManagementApi": reflect.ValueOf(v1.NewKeyManagementApi),
	    "ServiceChecksApi": reflect.ValueOf(v1.NewServiceChecksApi),
	    "UsageMeteringApi": reflect.ValueOf(v1.NewUsageMeteringApi),
	    "DashboardsApi": reflect.ValueOf(v1.NewDashboardsApi),
	    "DashboardListsApi": reflect.ValueOf(v1.NewDashboardListsApi),
	    "MetricsApi": reflect.ValueOf(v1.NewMetricsApi),
	    "DowntimesApi": reflect.ValueOf(v1.NewDowntimesApi),
	    "EventsApi": reflect.ValueOf(v1.NewEventsApi),
	    "SnapshotsApi": reflect.ValueOf(v1.NewSnapshotsApi),
	    "HostsApi": reflect.ValueOf(v1.NewHostsApi),
	    "AWSIntegrationApi": reflect.ValueOf(v1.NewAWSIntegrationApi),
	    "AWSLogsIntegrationApi": reflect.ValueOf(v1.NewAWSLogsIntegrationApi),
	    "AzureIntegrationApi": reflect.ValueOf(v1.NewAzureIntegrationApi),
	    "GCPIntegrationApi": reflect.ValueOf(v1.NewGCPIntegrationApi),
	    "PagerDutyIntegrationApi": reflect.ValueOf(v1.NewPagerDutyIntegrationApi),
	    "SlackIntegrationApi": reflect.ValueOf(v1.NewSlackIntegrationApi),
	    "WebhooksIntegrationApi": reflect.ValueOf(v1.NewWebhooksIntegrationApi),
	    "LogsApi": reflect.ValueOf(v1.NewLogsApi),
	    "LogsIndexesApi": reflect.ValueOf(v1.NewLogsIndexesApi),
	    "LogsPipelinesApi": reflect.ValueOf(v1.NewLogsPipelinesApi),
	    "MonitorsApi": reflect.ValueOf(v1.NewMonitorsApi),
	    "NotebooksApi": reflect.ValueOf(v1.NewNotebooksApi),
	    "OrganizationsApi": reflect.ValueOf(v1.NewOrganizationsApi),
	    "SecurityMonitoringApi": reflect.ValueOf(v1.NewSecurityMonitoringApi),
	    "ServiceLevelObjectivesApi": reflect.ValueOf(v1.NewServiceLevelObjectivesApi),
	    "ServiceLevelObjectiveCorrectionsApi": reflect.ValueOf(v1.NewServiceLevelObjectiveCorrectionsApi),
	    "SyntheticsApi": reflect.ValueOf(v1.NewSyntheticsApi),
	    "TagsApi": reflect.ValueOf(v1.NewTagsApi),
	    "UsersApi": reflect.ValueOf(v1.NewUsersApi),
	    "AuthenticationApi": reflect.ValueOf(v1.NewAuthenticationApi),
	},
	"v2": {
	    "KeyManagementApi": reflect.ValueOf(v2.NewKeyManagementApi),
	    "AuditApi": reflect.ValueOf(v2.NewAuditApi),
	    "AuthNMappingsApi": reflect.ValueOf(v2.NewAuthNMappingsApi),
	    "DashboardListsApi": reflect.ValueOf(v2.NewDashboardListsApi),
	    "IncidentsApi": reflect.ValueOf(v2.NewIncidentsApi),
	    "OpsgenieIntegrationApi": reflect.ValueOf(v2.NewOpsgenieIntegrationApi),
	    "LogsApi": reflect.ValueOf(v2.NewLogsApi),
	    "LogsArchivesApi": reflect.ValueOf(v2.NewLogsArchivesApi),
	    "LogsMetricsApi": reflect.ValueOf(v2.NewLogsMetricsApi),
	    "MetricsApi": reflect.ValueOf(v2.NewMetricsApi),
	    "RolesApi": reflect.ValueOf(v2.NewRolesApi),
	    "ProcessesApi": reflect.ValueOf(v2.NewProcessesApi),
	    "RUMApi": reflect.ValueOf(v2.NewRUMApi),
	    "OrganizationsApi": reflect.ValueOf(v2.NewOrganizationsApi),
	    "CloudWorkloadSecurityApi": reflect.ValueOf(v2.NewCloudWorkloadSecurityApi),
	    "SecurityMonitoringApi": reflect.ValueOf(v2.NewSecurityMonitoringApi),
	    "UsersApi": reflect.ValueOf(v2.NewUsersApi),
	    "ServiceAccountsApi": reflect.ValueOf(v2.NewServiceAccountsApi),
	    "IncidentServicesApi": reflect.ValueOf(v2.NewIncidentServicesApi),
	    "IncidentTeamsApi": reflect.ValueOf(v2.NewIncidentTeamsApi),
	    "UsageMeteringApi": reflect.ValueOf(v2.NewUsageMeteringApi),
	},
}