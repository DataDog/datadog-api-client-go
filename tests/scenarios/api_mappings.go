// This file is autogenerated. Do not modify directly.

package scenarios

import (
    "reflect"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
    "github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

var apiMappings = map[string]map[string]reflect.Value{
	"v1": {
	    "IPRangesApi": reflect.ValueOf(datadogV1.NewIPRangesApi),
	    "KeyManagementApi": reflect.ValueOf(datadogV1.NewKeyManagementApi),
	    "ServiceChecksApi": reflect.ValueOf(datadogV1.NewServiceChecksApi),
	    "UsageMeteringApi": reflect.ValueOf(datadogV1.NewUsageMeteringApi),
	    "DashboardsApi": reflect.ValueOf(datadogV1.NewDashboardsApi),
	    "DashboardListsApi": reflect.ValueOf(datadogV1.NewDashboardListsApi),
	    "MetricsApi": reflect.ValueOf(datadogV1.NewMetricsApi),
	    "DowntimesApi": reflect.ValueOf(datadogV1.NewDowntimesApi),
	    "EventsApi": reflect.ValueOf(datadogV1.NewEventsApi),
	    "SnapshotsApi": reflect.ValueOf(datadogV1.NewSnapshotsApi),
	    "HostsApi": reflect.ValueOf(datadogV1.NewHostsApi),
	    "AWSIntegrationApi": reflect.ValueOf(datadogV1.NewAWSIntegrationApi),
	    "AWSLogsIntegrationApi": reflect.ValueOf(datadogV1.NewAWSLogsIntegrationApi),
	    "AzureIntegrationApi": reflect.ValueOf(datadogV1.NewAzureIntegrationApi),
	    "GCPIntegrationApi": reflect.ValueOf(datadogV1.NewGCPIntegrationApi),
	    "PagerDutyIntegrationApi": reflect.ValueOf(datadogV1.NewPagerDutyIntegrationApi),
	    "SlackIntegrationApi": reflect.ValueOf(datadogV1.NewSlackIntegrationApi),
	    "WebhooksIntegrationApi": reflect.ValueOf(datadogV1.NewWebhooksIntegrationApi),
	    "LogsApi": reflect.ValueOf(datadogV1.NewLogsApi),
	    "LogsIndexesApi": reflect.ValueOf(datadogV1.NewLogsIndexesApi),
	    "LogsPipelinesApi": reflect.ValueOf(datadogV1.NewLogsPipelinesApi),
	    "MonitorsApi": reflect.ValueOf(datadogV1.NewMonitorsApi),
	    "NotebooksApi": reflect.ValueOf(datadogV1.NewNotebooksApi),
	    "OrganizationsApi": reflect.ValueOf(datadogV1.NewOrganizationsApi),
	    "SecurityMonitoringApi": reflect.ValueOf(datadogV1.NewSecurityMonitoringApi),
	    "ServiceLevelObjectivesApi": reflect.ValueOf(datadogV1.NewServiceLevelObjectivesApi),
	    "ServiceLevelObjectiveCorrectionsApi": reflect.ValueOf(datadogV1.NewServiceLevelObjectiveCorrectionsApi),
	    "SyntheticsApi": reflect.ValueOf(datadogV1.NewSyntheticsApi),
	    "TagsApi": reflect.ValueOf(datadogV1.NewTagsApi),
	    "UsersApi": reflect.ValueOf(datadogV1.NewUsersApi),
	    "AuthenticationApi": reflect.ValueOf(datadogV1.NewAuthenticationApi),
	},
	"v2": {
	    "KeyManagementApi": reflect.ValueOf(datadogV2.NewKeyManagementApi),
	    "APIManagementApi": reflect.ValueOf(datadogV2.NewAPIManagementApi),
	    "SpansMetricsApi": reflect.ValueOf(datadogV2.NewSpansMetricsApi),
	    "APMRetentionFiltersApi": reflect.ValueOf(datadogV2.NewAPMRetentionFiltersApi),
	    "AuditApi": reflect.ValueOf(datadogV2.NewAuditApi),
	    "AuthNMappingsApi": reflect.ValueOf(datadogV2.NewAuthNMappingsApi),
	    "CaseManagementApi": reflect.ValueOf(datadogV2.NewCaseManagementApi),
	    "CIVisibilityPipelinesApi": reflect.ValueOf(datadogV2.NewCIVisibilityPipelinesApi),
	    "CIVisibilityTestsApi": reflect.ValueOf(datadogV2.NewCIVisibilityTestsApi),
	    "ContainerImagesApi": reflect.ValueOf(datadogV2.NewContainerImagesApi),
	    "ContainersApi": reflect.ValueOf(datadogV2.NewContainersApi),
	    "CloudCostManagementApi": reflect.ValueOf(datadogV2.NewCloudCostManagementApi),
	    "UsageMeteringApi": reflect.ValueOf(datadogV2.NewUsageMeteringApi),
	    "DashboardListsApi": reflect.ValueOf(datadogV2.NewDashboardListsApi),
	    "DORAMetricsApi": reflect.ValueOf(datadogV2.NewDORAMetricsApi),
	    "DowntimesApi": reflect.ValueOf(datadogV2.NewDowntimesApi),
	    "EventsApi": reflect.ValueOf(datadogV2.NewEventsApi),
	    "IncidentsApi": reflect.ValueOf(datadogV2.NewIncidentsApi),
	    "GCPIntegrationApi": reflect.ValueOf(datadogV2.NewGCPIntegrationApi),
	    "OpsgenieIntegrationApi": reflect.ValueOf(datadogV2.NewOpsgenieIntegrationApi),
	    "CloudflareIntegrationApi": reflect.ValueOf(datadogV2.NewCloudflareIntegrationApi),
	    "ConfluentCloudApi": reflect.ValueOf(datadogV2.NewConfluentCloudApi),
	    "FastlyIntegrationApi": reflect.ValueOf(datadogV2.NewFastlyIntegrationApi),
	    "OktaIntegrationApi": reflect.ValueOf(datadogV2.NewOktaIntegrationApi),
	    "IPAllowlistApi": reflect.ValueOf(datadogV2.NewIPAllowlistApi),
	    "LogsApi": reflect.ValueOf(datadogV2.NewLogsApi),
	    "LogsArchivesApi": reflect.ValueOf(datadogV2.NewLogsArchivesApi),
	    "LogsMetricsApi": reflect.ValueOf(datadogV2.NewLogsMetricsApi),
	    "MetricsApi": reflect.ValueOf(datadogV2.NewMetricsApi),
	    "MonitorsApi": reflect.ValueOf(datadogV2.NewMonitorsApi),
	    "RolesApi": reflect.ValueOf(datadogV2.NewRolesApi),
	    "SecurityMonitoringApi": reflect.ValueOf(datadogV2.NewSecurityMonitoringApi),
	    "PowerpackApi": reflect.ValueOf(datadogV2.NewPowerpackApi),
	    "ProcessesApi": reflect.ValueOf(datadogV2.NewProcessesApi),
	    "RestrictionPoliciesApi": reflect.ValueOf(datadogV2.NewRestrictionPoliciesApi),
	    "RUMApi": reflect.ValueOf(datadogV2.NewRUMApi),
	    "OrganizationsApi": reflect.ValueOf(datadogV2.NewOrganizationsApi),
	    "ServiceScorecardsApi": reflect.ValueOf(datadogV2.NewServiceScorecardsApi),
	    "CloudWorkloadSecurityApi": reflect.ValueOf(datadogV2.NewCloudWorkloadSecurityApi),
	    "SensitiveDataScannerApi": reflect.ValueOf(datadogV2.NewSensitiveDataScannerApi),
	    "ServiceAccountsApi": reflect.ValueOf(datadogV2.NewServiceAccountsApi),
	    "IncidentServicesApi": reflect.ValueOf(datadogV2.NewIncidentServicesApi),
	    "ServiceDefinitionApi": reflect.ValueOf(datadogV2.NewServiceDefinitionApi),
	    "SpansApi": reflect.ValueOf(datadogV2.NewSpansApi),
	    "SyntheticsApi": reflect.ValueOf(datadogV2.NewSyntheticsApi),
	    "TeamsApi": reflect.ValueOf(datadogV2.NewTeamsApi),
	    "IncidentTeamsApi": reflect.ValueOf(datadogV2.NewIncidentTeamsApi),
	    "UsersApi": reflect.ValueOf(datadogV2.NewUsersApi),
	},
}