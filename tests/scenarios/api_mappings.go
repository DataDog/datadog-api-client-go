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
	    "ActionConnectionApi": reflect.ValueOf(datadogV2.NewActionConnectionApi),
	    "AgentlessScanningApi": reflect.ValueOf(datadogV2.NewAgentlessScanningApi),
	    "KeyManagementApi": reflect.ValueOf(datadogV2.NewKeyManagementApi),
	    "APIManagementApi": reflect.ValueOf(datadogV2.NewAPIManagementApi),
	    "SpansMetricsApi": reflect.ValueOf(datadogV2.NewSpansMetricsApi),
	    "APMRetentionFiltersApi": reflect.ValueOf(datadogV2.NewAPMRetentionFiltersApi),
	    "AppBuilderApi": reflect.ValueOf(datadogV2.NewAppBuilderApi),
	    "AuditApi": reflect.ValueOf(datadogV2.NewAuditApi),
	    "AuthNMappingsApi": reflect.ValueOf(datadogV2.NewAuthNMappingsApi),
	    "CaseManagementApi": reflect.ValueOf(datadogV2.NewCaseManagementApi),
	    "SoftwareCatalogApi": reflect.ValueOf(datadogV2.NewSoftwareCatalogApi),
	    "CIVisibilityPipelinesApi": reflect.ValueOf(datadogV2.NewCIVisibilityPipelinesApi),
	    "CIVisibilityTestsApi": reflect.ValueOf(datadogV2.NewCIVisibilityTestsApi),
	    "SecurityMonitoringApi": reflect.ValueOf(datadogV2.NewSecurityMonitoringApi),
	    "ContainerImagesApi": reflect.ValueOf(datadogV2.NewContainerImagesApi),
	    "ContainersApi": reflect.ValueOf(datadogV2.NewContainersApi),
	    "CloudCostManagementApi": reflect.ValueOf(datadogV2.NewCloudCostManagementApi),
	    "UsageMeteringApi": reflect.ValueOf(datadogV2.NewUsageMeteringApi),
	    "CSMAgentsApi": reflect.ValueOf(datadogV2.NewCSMAgentsApi),
	    "CSMCoverageAnalysisApi": reflect.ValueOf(datadogV2.NewCSMCoverageAnalysisApi),
	    "DashboardListsApi": reflect.ValueOf(datadogV2.NewDashboardListsApi),
	    "DatasetsApi": reflect.ValueOf(datadogV2.NewDatasetsApi),
	    "DataDeletionApi": reflect.ValueOf(datadogV2.NewDataDeletionApi),
	    "DomainAllowlistApi": reflect.ValueOf(datadogV2.NewDomainAllowlistApi),
	    "DORAMetricsApi": reflect.ValueOf(datadogV2.NewDORAMetricsApi),
	    "DowntimesApi": reflect.ValueOf(datadogV2.NewDowntimesApi),
	    "EventsApi": reflect.ValueOf(datadogV2.NewEventsApi),
	    "IncidentsApi": reflect.ValueOf(datadogV2.NewIncidentsApi),
	    "AWSIntegrationApi": reflect.ValueOf(datadogV2.NewAWSIntegrationApi),
	    "AWSLogsIntegrationApi": reflect.ValueOf(datadogV2.NewAWSLogsIntegrationApi),
	    "GCPIntegrationApi": reflect.ValueOf(datadogV2.NewGCPIntegrationApi),
	    "MicrosoftTeamsIntegrationApi": reflect.ValueOf(datadogV2.NewMicrosoftTeamsIntegrationApi),
	    "OpsgenieIntegrationApi": reflect.ValueOf(datadogV2.NewOpsgenieIntegrationApi),
	    "CloudflareIntegrationApi": reflect.ValueOf(datadogV2.NewCloudflareIntegrationApi),
	    "ConfluentCloudApi": reflect.ValueOf(datadogV2.NewConfluentCloudApi),
	    "FastlyIntegrationApi": reflect.ValueOf(datadogV2.NewFastlyIntegrationApi),
	    "OktaIntegrationApi": reflect.ValueOf(datadogV2.NewOktaIntegrationApi),
	    "IPAllowlistApi": reflect.ValueOf(datadogV2.NewIPAllowlistApi),
	    "LogsApi": reflect.ValueOf(datadogV2.NewLogsApi),
	    "LogsArchivesApi": reflect.ValueOf(datadogV2.NewLogsArchivesApi),
	    "LogsCustomDestinationsApi": reflect.ValueOf(datadogV2.NewLogsCustomDestinationsApi),
	    "LogsMetricsApi": reflect.ValueOf(datadogV2.NewLogsMetricsApi),
	    "MetricsApi": reflect.ValueOf(datadogV2.NewMetricsApi),
	    "MonitorsApi": reflect.ValueOf(datadogV2.NewMonitorsApi),
	    "NetworkDeviceMonitoringApi": reflect.ValueOf(datadogV2.NewNetworkDeviceMonitoringApi),
	    "CloudNetworkMonitoringApi": reflect.ValueOf(datadogV2.NewCloudNetworkMonitoringApi),
	    "OnCallApi": reflect.ValueOf(datadogV2.NewOnCallApi),
	    "OnCallPagingApi": reflect.ValueOf(datadogV2.NewOnCallPagingApi),
	    "OrganizationsApi": reflect.ValueOf(datadogV2.NewOrganizationsApi),
	    "RolesApi": reflect.ValueOf(datadogV2.NewRolesApi),
	    "PowerpackApi": reflect.ValueOf(datadogV2.NewPowerpackApi),
	    "ProcessesApi": reflect.ValueOf(datadogV2.NewProcessesApi),
	    "ApplicationSecurityApi": reflect.ValueOf(datadogV2.NewApplicationSecurityApi),
	    "CSMThreatsApi": reflect.ValueOf(datadogV2.NewCSMThreatsApi),
	    "ObservabilityPipelinesApi": reflect.ValueOf(datadogV2.NewObservabilityPipelinesApi),
	    "RestrictionPoliciesApi": reflect.ValueOf(datadogV2.NewRestrictionPoliciesApi),
	    "RUMApi": reflect.ValueOf(datadogV2.NewRUMApi),
	    "RumRetentionFiltersApi": reflect.ValueOf(datadogV2.NewRumRetentionFiltersApi),
	    "RumMetricsApi": reflect.ValueOf(datadogV2.NewRumMetricsApi),
	    "ServiceScorecardsApi": reflect.ValueOf(datadogV2.NewServiceScorecardsApi),
	    "SensitiveDataScannerApi": reflect.ValueOf(datadogV2.NewSensitiveDataScannerApi),
	    "ServiceAccountsApi": reflect.ValueOf(datadogV2.NewServiceAccountsApi),
	    "IncidentServicesApi": reflect.ValueOf(datadogV2.NewIncidentServicesApi),
	    "ServiceDefinitionApi": reflect.ValueOf(datadogV2.NewServiceDefinitionApi),
	    "ServiceLevelObjectivesApi": reflect.ValueOf(datadogV2.NewServiceLevelObjectivesApi),
	    "SpansApi": reflect.ValueOf(datadogV2.NewSpansApi),
	    "SyntheticsApi": reflect.ValueOf(datadogV2.NewSyntheticsApi),
	    "TeamsApi": reflect.ValueOf(datadogV2.NewTeamsApi),
	    "IncidentTeamsApi": reflect.ValueOf(datadogV2.NewIncidentTeamsApi),
	    "UsersApi": reflect.ValueOf(datadogV2.NewUsersApi),
	    "WorkflowAutomationApi": reflect.ValueOf(datadogV2.NewWorkflowAutomationApi),
	},
}