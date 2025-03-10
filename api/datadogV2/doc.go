// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

// List of APIs:
//   - [APIManagementApi.CreateOpenAPI]
//   - [APIManagementApi.DeleteOpenAPI]
//   - [APIManagementApi.GetOpenAPI]
//   - [APIManagementApi.ListAPIs]
//   - [APIManagementApi.UpdateOpenAPI]
//   - [APMRetentionFiltersApi.CreateApmRetentionFilter]
//   - [APMRetentionFiltersApi.DeleteApmRetentionFilter]
//   - [APMRetentionFiltersApi.GetApmRetentionFilter]
//   - [APMRetentionFiltersApi.ListApmRetentionFilters]
//   - [APMRetentionFiltersApi.ReorderApmRetentionFilters]
//   - [APMRetentionFiltersApi.UpdateApmRetentionFilter]
//   - [AWSIntegrationApi.CreateAWSAccount]
//   - [AWSIntegrationApi.CreateNewAWSExternalID]
//   - [AWSIntegrationApi.DeleteAWSAccount]
//   - [AWSIntegrationApi.GetAWSAccount]
//   - [AWSIntegrationApi.ListAWSAccounts]
//   - [AWSIntegrationApi.ListAWSNamespaces]
//   - [AWSIntegrationApi.UpdateAWSAccount]
//   - [AWSLogsIntegrationApi.ListAWSLogsServices]
//   - [ActionConnectionApi.CreateActionConnection]
//   - [ActionConnectionApi.DeleteActionConnection]
//   - [ActionConnectionApi.GetActionConnection]
//   - [ActionConnectionApi.UpdateActionConnection]
//   - [AgentlessScanningApi.CreateAwsOnDemandTask]
//   - [AgentlessScanningApi.CreateAwsScanOptions]
//   - [AgentlessScanningApi.DeleteAwsScanOptions]
//   - [AgentlessScanningApi.GetAwsOnDemandTask]
//   - [AgentlessScanningApi.ListAwsOnDemandTasks]
//   - [AgentlessScanningApi.ListAwsScanOptions]
//   - [AgentlessScanningApi.UpdateAwsScanOptions]
//   - [AppBuilderApi.CreateApp]
//   - [AppBuilderApi.DeleteApp]
//   - [AppBuilderApi.DeleteApps]
//   - [AppBuilderApi.GetApp]
//   - [AppBuilderApi.ListApps]
//   - [AppBuilderApi.PublishApp]
//   - [AppBuilderApi.UnpublishApp]
//   - [AppBuilderApi.UpdateApp]
//   - [ApplicationSecurityApi.CreateApplicationSecurityWafCustomRule]
//   - [ApplicationSecurityApi.CreateApplicationSecurityWafExclusionFilter]
//   - [ApplicationSecurityApi.DeleteApplicationSecurityWafCustomRule]
//   - [ApplicationSecurityApi.DeleteApplicationSecurityWafExclusionFilter]
//   - [ApplicationSecurityApi.GetApplicationSecurityWafCustomRule]
//   - [ApplicationSecurityApi.GetApplicationSecurityWafExclusionFilter]
//   - [ApplicationSecurityApi.ListApplicationSecurityWAFCustomRules]
//   - [ApplicationSecurityApi.ListApplicationSecurityWafExclusionFilters]
//   - [ApplicationSecurityApi.UpdateApplicationSecurityWafCustomRule]
//   - [ApplicationSecurityApi.UpdateApplicationSecurityWafExclusionFilter]
//   - [AuditApi.ListAuditLogs]
//   - [AuditApi.SearchAuditLogs]
//   - [AuthNMappingsApi.CreateAuthNMapping]
//   - [AuthNMappingsApi.DeleteAuthNMapping]
//   - [AuthNMappingsApi.GetAuthNMapping]
//   - [AuthNMappingsApi.ListAuthNMappings]
//   - [AuthNMappingsApi.UpdateAuthNMapping]
//   - [CIVisibilityPipelinesApi.AggregateCIAppPipelineEvents]
//   - [CIVisibilityPipelinesApi.CreateCIAppPipelineEvent]
//   - [CIVisibilityPipelinesApi.ListCIAppPipelineEvents]
//   - [CIVisibilityPipelinesApi.SearchCIAppPipelineEvents]
//   - [CIVisibilityTestsApi.AggregateCIAppTestEvents]
//   - [CIVisibilityTestsApi.ListCIAppTestEvents]
//   - [CIVisibilityTestsApi.SearchCIAppTestEvents]
//   - [CSMAgentsApi.ListAllCSMAgents]
//   - [CSMAgentsApi.ListAllCSMServerlessAgents]
//   - [CSMCoverageAnalysisApi.GetCSMCloudAccountsCoverageAnalysis]
//   - [CSMCoverageAnalysisApi.GetCSMHostsAndContainersCoverageAnalysis]
//   - [CSMCoverageAnalysisApi.GetCSMServerlessCoverageAnalysis]
//   - [CSMThreatsApi.CreateCSMThreatsAgentRule]
//   - [CSMThreatsApi.CreateCloudWorkloadSecurityAgentRule]
//   - [CSMThreatsApi.DeleteCSMThreatsAgentRule]
//   - [CSMThreatsApi.DeleteCloudWorkloadSecurityAgentRule]
//   - [CSMThreatsApi.DownloadCSMThreatsPolicy]
//   - [CSMThreatsApi.DownloadCloudWorkloadPolicyFile]
//   - [CSMThreatsApi.GetCSMThreatsAgentRule]
//   - [CSMThreatsApi.GetCloudWorkloadSecurityAgentRule]
//   - [CSMThreatsApi.ListCSMThreatsAgentRules]
//   - [CSMThreatsApi.ListCloudWorkloadSecurityAgentRules]
//   - [CSMThreatsApi.UpdateCSMThreatsAgentRule]
//   - [CSMThreatsApi.UpdateCloudWorkloadSecurityAgentRule]
//   - [CaseManagementApi.ArchiveCase]
//   - [CaseManagementApi.AssignCase]
//   - [CaseManagementApi.CreateCase]
//   - [CaseManagementApi.CreateProject]
//   - [CaseManagementApi.DeleteProject]
//   - [CaseManagementApi.GetCase]
//   - [CaseManagementApi.GetProject]
//   - [CaseManagementApi.GetProjects]
//   - [CaseManagementApi.SearchCases]
//   - [CaseManagementApi.UnarchiveCase]
//   - [CaseManagementApi.UnassignCase]
//   - [CaseManagementApi.UpdatePriority]
//   - [CaseManagementApi.UpdateStatus]
//   - [CloudCostManagementApi.CreateCostAWSCURConfig]
//   - [CloudCostManagementApi.CreateCostAzureUCConfigs]
//   - [CloudCostManagementApi.DeleteCostAWSCURConfig]
//   - [CloudCostManagementApi.DeleteCostAzureUCConfig]
//   - [CloudCostManagementApi.DeleteCustomCostsFile]
//   - [CloudCostManagementApi.GetCustomCostsFile]
//   - [CloudCostManagementApi.ListCostAWSCURConfigs]
//   - [CloudCostManagementApi.ListCostAzureUCConfigs]
//   - [CloudCostManagementApi.ListCustomCostsFiles]
//   - [CloudCostManagementApi.UpdateCostAWSCURConfig]
//   - [CloudCostManagementApi.UpdateCostAzureUCConfigs]
//   - [CloudCostManagementApi.UploadCustomCostsFile]
//   - [CloudflareIntegrationApi.CreateCloudflareAccount]
//   - [CloudflareIntegrationApi.DeleteCloudflareAccount]
//   - [CloudflareIntegrationApi.GetCloudflareAccount]
//   - [CloudflareIntegrationApi.ListCloudflareAccounts]
//   - [CloudflareIntegrationApi.UpdateCloudflareAccount]
//   - [ConfluentCloudApi.CreateConfluentAccount]
//   - [ConfluentCloudApi.CreateConfluentResource]
//   - [ConfluentCloudApi.DeleteConfluentAccount]
//   - [ConfluentCloudApi.DeleteConfluentResource]
//   - [ConfluentCloudApi.GetConfluentAccount]
//   - [ConfluentCloudApi.GetConfluentResource]
//   - [ConfluentCloudApi.ListConfluentAccount]
//   - [ConfluentCloudApi.ListConfluentResource]
//   - [ConfluentCloudApi.UpdateConfluentAccount]
//   - [ConfluentCloudApi.UpdateConfluentResource]
//   - [ContainerImagesApi.ListContainerImages]
//   - [ContainersApi.ListContainers]
//   - [DORAMetricsApi.CreateDORADeployment]
//   - [DORAMetricsApi.CreateDORAIncident]
//   - [DashboardListsApi.CreateDashboardListItems]
//   - [DashboardListsApi.DeleteDashboardListItems]
//   - [DashboardListsApi.GetDashboardListItems]
//   - [DashboardListsApi.UpdateDashboardListItems]
//   - [DataDeletionApi.CancelDataDeletionRequest]
//   - [DataDeletionApi.CreateDataDeletionRequest]
//   - [DataDeletionApi.GetDataDeletionRequests]
//   - [DomainAllowlistApi.GetDomainAllowlist]
//   - [DomainAllowlistApi.PatchDomainAllowlist]
//   - [DowntimesApi.CancelDowntime]
//   - [DowntimesApi.CreateDowntime]
//   - [DowntimesApi.GetDowntime]
//   - [DowntimesApi.ListDowntimes]
//   - [DowntimesApi.ListMonitorDowntimes]
//   - [DowntimesApi.UpdateDowntime]
//   - [EventsApi.CreateEvent]
//   - [EventsApi.ListEvents]
//   - [EventsApi.SearchEvents]
//   - [FastlyIntegrationApi.CreateFastlyAccount]
//   - [FastlyIntegrationApi.CreateFastlyService]
//   - [FastlyIntegrationApi.DeleteFastlyAccount]
//   - [FastlyIntegrationApi.DeleteFastlyService]
//   - [FastlyIntegrationApi.GetFastlyAccount]
//   - [FastlyIntegrationApi.GetFastlyService]
//   - [FastlyIntegrationApi.ListFastlyAccounts]
//   - [FastlyIntegrationApi.ListFastlyServices]
//   - [FastlyIntegrationApi.UpdateFastlyAccount]
//   - [FastlyIntegrationApi.UpdateFastlyService]
//   - [GCPIntegrationApi.CreateGCPSTSAccount]
//   - [GCPIntegrationApi.DeleteGCPSTSAccount]
//   - [GCPIntegrationApi.GetGCPSTSDelegate]
//   - [GCPIntegrationApi.ListGCPSTSAccounts]
//   - [GCPIntegrationApi.MakeGCPSTSDelegate]
//   - [GCPIntegrationApi.UpdateGCPSTSAccount]
//   - [IPAllowlistApi.GetIPAllowlist]
//   - [IPAllowlistApi.UpdateIPAllowlist]
//   - [IncidentServicesApi.CreateIncidentService]
//   - [IncidentServicesApi.DeleteIncidentService]
//   - [IncidentServicesApi.GetIncidentService]
//   - [IncidentServicesApi.ListIncidentServices]
//   - [IncidentServicesApi.UpdateIncidentService]
//   - [IncidentTeamsApi.CreateIncidentTeam]
//   - [IncidentTeamsApi.DeleteIncidentTeam]
//   - [IncidentTeamsApi.GetIncidentTeam]
//   - [IncidentTeamsApi.ListIncidentTeams]
//   - [IncidentTeamsApi.UpdateIncidentTeam]
//   - [IncidentsApi.CreateIncident]
//   - [IncidentsApi.CreateIncidentIntegration]
//   - [IncidentsApi.CreateIncidentTodo]
//   - [IncidentsApi.CreateIncidentType]
//   - [IncidentsApi.DeleteIncident]
//   - [IncidentsApi.DeleteIncidentIntegration]
//   - [IncidentsApi.DeleteIncidentTodo]
//   - [IncidentsApi.DeleteIncidentType]
//   - [IncidentsApi.GetIncident]
//   - [IncidentsApi.GetIncidentIntegration]
//   - [IncidentsApi.GetIncidentTodo]
//   - [IncidentsApi.GetIncidentType]
//   - [IncidentsApi.ListIncidentAttachments]
//   - [IncidentsApi.ListIncidentIntegrations]
//   - [IncidentsApi.ListIncidentTodos]
//   - [IncidentsApi.ListIncidentTypes]
//   - [IncidentsApi.ListIncidents]
//   - [IncidentsApi.SearchIncidents]
//   - [IncidentsApi.UpdateIncident]
//   - [IncidentsApi.UpdateIncidentAttachments]
//   - [IncidentsApi.UpdateIncidentIntegration]
//   - [IncidentsApi.UpdateIncidentTodo]
//   - [IncidentsApi.UpdateIncidentType]
//   - [KeyManagementApi.CreateAPIKey]
//   - [KeyManagementApi.CreateCurrentUserApplicationKey]
//   - [KeyManagementApi.DeleteAPIKey]
//   - [KeyManagementApi.DeleteApplicationKey]
//   - [KeyManagementApi.DeleteCurrentUserApplicationKey]
//   - [KeyManagementApi.GetAPIKey]
//   - [KeyManagementApi.GetApplicationKey]
//   - [KeyManagementApi.GetCurrentUserApplicationKey]
//   - [KeyManagementApi.ListAPIKeys]
//   - [KeyManagementApi.ListApplicationKeys]
//   - [KeyManagementApi.ListCurrentUserApplicationKeys]
//   - [KeyManagementApi.UpdateAPIKey]
//   - [KeyManagementApi.UpdateApplicationKey]
//   - [KeyManagementApi.UpdateCurrentUserApplicationKey]
//   - [LogsApi.AggregateLogs]
//   - [LogsApi.ListLogs]
//   - [LogsApi.ListLogsGet]
//   - [LogsApi.SubmitLog]
//   - [LogsArchivesApi.AddReadRoleToArchive]
//   - [LogsArchivesApi.CreateLogsArchive]
//   - [LogsArchivesApi.DeleteLogsArchive]
//   - [LogsArchivesApi.GetLogsArchive]
//   - [LogsArchivesApi.GetLogsArchiveOrder]
//   - [LogsArchivesApi.ListArchiveReadRoles]
//   - [LogsArchivesApi.ListLogsArchives]
//   - [LogsArchivesApi.RemoveRoleFromArchive]
//   - [LogsArchivesApi.UpdateLogsArchive]
//   - [LogsArchivesApi.UpdateLogsArchiveOrder]
//   - [LogsCustomDestinationsApi.CreateLogsCustomDestination]
//   - [LogsCustomDestinationsApi.DeleteLogsCustomDestination]
//   - [LogsCustomDestinationsApi.GetLogsCustomDestination]
//   - [LogsCustomDestinationsApi.ListLogsCustomDestinations]
//   - [LogsCustomDestinationsApi.UpdateLogsCustomDestination]
//   - [LogsMetricsApi.CreateLogsMetric]
//   - [LogsMetricsApi.DeleteLogsMetric]
//   - [LogsMetricsApi.GetLogsMetric]
//   - [LogsMetricsApi.ListLogsMetrics]
//   - [LogsMetricsApi.UpdateLogsMetric]
//   - [MetricsApi.CreateBulkTagsMetricsConfiguration]
//   - [MetricsApi.CreateTagConfiguration]
//   - [MetricsApi.DeleteBulkTagsMetricsConfiguration]
//   - [MetricsApi.DeleteTagConfiguration]
//   - [MetricsApi.EstimateMetricsOutputSeries]
//   - [MetricsApi.ListActiveMetricConfigurations]
//   - [MetricsApi.ListMetricAssets]
//   - [MetricsApi.ListTagConfigurationByName]
//   - [MetricsApi.ListTagConfigurations]
//   - [MetricsApi.ListTagsByMetricName]
//   - [MetricsApi.ListVolumesByMetricName]
//   - [MetricsApi.QueryScalarData]
//   - [MetricsApi.QueryTimeseriesData]
//   - [MetricsApi.SubmitMetrics]
//   - [MetricsApi.UpdateTagConfiguration]
//   - [MicrosoftTeamsIntegrationApi.CreateTenantBasedHandle]
//   - [MicrosoftTeamsIntegrationApi.CreateWorkflowsWebhookHandle]
//   - [MicrosoftTeamsIntegrationApi.DeleteTenantBasedHandle]
//   - [MicrosoftTeamsIntegrationApi.DeleteWorkflowsWebhookHandle]
//   - [MicrosoftTeamsIntegrationApi.GetChannelByName]
//   - [MicrosoftTeamsIntegrationApi.GetTenantBasedHandle]
//   - [MicrosoftTeamsIntegrationApi.GetWorkflowsWebhookHandle]
//   - [MicrosoftTeamsIntegrationApi.ListTenantBasedHandles]
//   - [MicrosoftTeamsIntegrationApi.ListWorkflowsWebhookHandles]
//   - [MicrosoftTeamsIntegrationApi.UpdateTenantBasedHandle]
//   - [MicrosoftTeamsIntegrationApi.UpdateWorkflowsWebhookHandle]
//   - [MonitorsApi.CreateMonitorConfigPolicy]
//   - [MonitorsApi.DeleteMonitorConfigPolicy]
//   - [MonitorsApi.GetMonitorConfigPolicy]
//   - [MonitorsApi.ListMonitorConfigPolicies]
//   - [MonitorsApi.UpdateMonitorConfigPolicy]
//   - [NetworkDeviceMonitoringApi.GetDevice]
//   - [NetworkDeviceMonitoringApi.GetInterfaces]
//   - [NetworkDeviceMonitoringApi.ListDeviceUserTags]
//   - [NetworkDeviceMonitoringApi.ListDevices]
//   - [NetworkDeviceMonitoringApi.UpdateDeviceUserTags]
//   - [ObservabilityPipelinesApi.CreatePipeline]
//   - [ObservabilityPipelinesApi.DeletePipeline]
//   - [ObservabilityPipelinesApi.GetPipeline]
//   - [ObservabilityPipelinesApi.ListPipelines]
//   - [ObservabilityPipelinesApi.UpdatePipeline]
//   - [OktaIntegrationApi.CreateOktaAccount]
//   - [OktaIntegrationApi.DeleteOktaAccount]
//   - [OktaIntegrationApi.GetOktaAccount]
//   - [OktaIntegrationApi.ListOktaAccounts]
//   - [OktaIntegrationApi.UpdateOktaAccount]
//   - [OpsgenieIntegrationApi.CreateOpsgenieService]
//   - [OpsgenieIntegrationApi.DeleteOpsgenieService]
//   - [OpsgenieIntegrationApi.GetOpsgenieService]
//   - [OpsgenieIntegrationApi.ListOpsgenieServices]
//   - [OpsgenieIntegrationApi.UpdateOpsgenieService]
//   - [OrganizationsApi.GetOrgConfig]
//   - [OrganizationsApi.ListOrgConfigs]
//   - [OrganizationsApi.UpdateOrgConfig]
//   - [OrganizationsApi.UploadIdPMetadata]
//   - [PowerpackApi.CreatePowerpack]
//   - [PowerpackApi.DeletePowerpack]
//   - [PowerpackApi.GetPowerpack]
//   - [PowerpackApi.ListPowerpacks]
//   - [PowerpackApi.UpdatePowerpack]
//   - [ProcessesApi.ListProcesses]
//   - [RUMApi.AggregateRUMEvents]
//   - [RUMApi.CreateRUMApplication]
//   - [RUMApi.DeleteRUMApplication]
//   - [RUMApi.GetRUMApplication]
//   - [RUMApi.GetRUMApplications]
//   - [RUMApi.ListRUMEvents]
//   - [RUMApi.SearchRUMEvents]
//   - [RUMApi.UpdateRUMApplication]
//   - [RestrictionPoliciesApi.DeleteRestrictionPolicy]
//   - [RestrictionPoliciesApi.GetRestrictionPolicy]
//   - [RestrictionPoliciesApi.UpdateRestrictionPolicy]
//   - [RolesApi.AddPermissionToRole]
//   - [RolesApi.AddUserToRole]
//   - [RolesApi.CloneRole]
//   - [RolesApi.CreateRole]
//   - [RolesApi.DeleteRole]
//   - [RolesApi.GetRole]
//   - [RolesApi.ListPermissions]
//   - [RolesApi.ListRolePermissions]
//   - [RolesApi.ListRoleUsers]
//   - [RolesApi.ListRoles]
//   - [RolesApi.RemovePermissionFromRole]
//   - [RolesApi.RemoveUserFromRole]
//   - [RolesApi.UpdateRole]
//   - [RumMetricsApi.CreateRumMetric]
//   - [RumMetricsApi.DeleteRumMetric]
//   - [RumMetricsApi.GetRumMetric]
//   - [RumMetricsApi.ListRumMetrics]
//   - [RumMetricsApi.UpdateRumMetric]
//   - [RumRetentionFiltersApi.CreateRetentionFilter]
//   - [RumRetentionFiltersApi.DeleteRetentionFilter]
//   - [RumRetentionFiltersApi.GetRetentionFilter]
//   - [RumRetentionFiltersApi.ListRetentionFilters]
//   - [RumRetentionFiltersApi.OrderRetentionFilters]
//   - [RumRetentionFiltersApi.UpdateRetentionFilter]
//   - [SecurityMonitoringApi.CancelHistoricalJob]
//   - [SecurityMonitoringApi.ConvertExistingSecurityMonitoringRule]
//   - [SecurityMonitoringApi.ConvertJobResultToSignal]
//   - [SecurityMonitoringApi.ConvertSecurityMonitoringRuleFromJSONToTerraform]
//   - [SecurityMonitoringApi.CreateSecurityFilter]
//   - [SecurityMonitoringApi.CreateSecurityMonitoringRule]
//   - [SecurityMonitoringApi.CreateSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.CreateSignalNotificationRule]
//   - [SecurityMonitoringApi.CreateVulnerabilityNotificationRule]
//   - [SecurityMonitoringApi.DeleteHistoricalJob]
//   - [SecurityMonitoringApi.DeleteSecurityFilter]
//   - [SecurityMonitoringApi.DeleteSecurityMonitoringRule]
//   - [SecurityMonitoringApi.DeleteSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.DeleteSignalNotificationRule]
//   - [SecurityMonitoringApi.DeleteVulnerabilityNotificationRule]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalState]
//   - [SecurityMonitoringApi.GetFinding]
//   - [SecurityMonitoringApi.GetHistoricalJob]
//   - [SecurityMonitoringApi.GetRuleVersionHistory]
//   - [SecurityMonitoringApi.GetSBOM]
//   - [SecurityMonitoringApi.GetSecurityFilter]
//   - [SecurityMonitoringApi.GetSecurityMonitoringRule]
//   - [SecurityMonitoringApi.GetSecurityMonitoringSignal]
//   - [SecurityMonitoringApi.GetSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.GetSignalNotificationRule]
//   - [SecurityMonitoringApi.GetSignalNotificationRules]
//   - [SecurityMonitoringApi.GetVulnerabilityNotificationRule]
//   - [SecurityMonitoringApi.GetVulnerabilityNotificationRules]
//   - [SecurityMonitoringApi.ListFindings]
//   - [SecurityMonitoringApi.ListHistoricalJobs]
//   - [SecurityMonitoringApi.ListSecurityFilters]
//   - [SecurityMonitoringApi.ListSecurityMonitoringRules]
//   - [SecurityMonitoringApi.ListSecurityMonitoringSignals]
//   - [SecurityMonitoringApi.ListSecurityMonitoringSuppressions]
//   - [SecurityMonitoringApi.ListVulnerabilities]
//   - [SecurityMonitoringApi.ListVulnerableAssets]
//   - [SecurityMonitoringApi.MuteFindings]
//   - [SecurityMonitoringApi.PatchSignalNotificationRule]
//   - [SecurityMonitoringApi.PatchVulnerabilityNotificationRule]
//   - [SecurityMonitoringApi.RunHistoricalJob]
//   - [SecurityMonitoringApi.SearchSecurityMonitoringSignals]
//   - [SecurityMonitoringApi.TestExistingSecurityMonitoringRule]
//   - [SecurityMonitoringApi.TestSecurityMonitoringRule]
//   - [SecurityMonitoringApi.UpdateSecurityFilter]
//   - [SecurityMonitoringApi.UpdateSecurityMonitoringRule]
//   - [SecurityMonitoringApi.UpdateSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.ValidateSecurityMonitoringRule]
//   - [SensitiveDataScannerApi.CreateScanningGroup]
//   - [SensitiveDataScannerApi.CreateScanningRule]
//   - [SensitiveDataScannerApi.DeleteScanningGroup]
//   - [SensitiveDataScannerApi.DeleteScanningRule]
//   - [SensitiveDataScannerApi.ListScanningGroups]
//   - [SensitiveDataScannerApi.ListStandardPatterns]
//   - [SensitiveDataScannerApi.ReorderScanningGroups]
//   - [SensitiveDataScannerApi.UpdateScanningGroup]
//   - [SensitiveDataScannerApi.UpdateScanningRule]
//   - [ServiceAccountsApi.CreateServiceAccount]
//   - [ServiceAccountsApi.CreateServiceAccountApplicationKey]
//   - [ServiceAccountsApi.DeleteServiceAccountApplicationKey]
//   - [ServiceAccountsApi.GetServiceAccountApplicationKey]
//   - [ServiceAccountsApi.ListServiceAccountApplicationKeys]
//   - [ServiceAccountsApi.UpdateServiceAccountApplicationKey]
//   - [ServiceDefinitionApi.CreateOrUpdateServiceDefinitions]
//   - [ServiceDefinitionApi.DeleteServiceDefinition]
//   - [ServiceDefinitionApi.GetServiceDefinition]
//   - [ServiceDefinitionApi.ListServiceDefinitions]
//   - [ServiceLevelObjectivesApi.CreateSLOReportJob]
//   - [ServiceLevelObjectivesApi.GetSLOReport]
//   - [ServiceLevelObjectivesApi.GetSLOReportJobStatus]
//   - [ServiceScorecardsApi.CreateScorecardOutcomesBatch]
//   - [ServiceScorecardsApi.CreateScorecardRule]
//   - [ServiceScorecardsApi.DeleteScorecardRule]
//   - [ServiceScorecardsApi.ListScorecardOutcomes]
//   - [ServiceScorecardsApi.ListScorecardRules]
//   - [ServiceScorecardsApi.UpdateScorecardRule]
//   - [SoftwareCatalogApi.DeleteCatalogEntity]
//   - [SoftwareCatalogApi.ListCatalogEntity]
//   - [SoftwareCatalogApi.UpsertCatalogEntity]
//   - [SpansApi.AggregateSpans]
//   - [SpansApi.ListSpans]
//   - [SpansApi.ListSpansGet]
//   - [SpansMetricsApi.CreateSpansMetric]
//   - [SpansMetricsApi.DeleteSpansMetric]
//   - [SpansMetricsApi.GetSpansMetric]
//   - [SpansMetricsApi.ListSpansMetrics]
//   - [SpansMetricsApi.UpdateSpansMetric]
//   - [SyntheticsApi.GetOnDemandConcurrencyCap]
//   - [SyntheticsApi.SetOnDemandConcurrencyCap]
//   - [TeamsApi.CreateTeam]
//   - [TeamsApi.CreateTeamLink]
//   - [TeamsApi.CreateTeamMembership]
//   - [TeamsApi.DeleteTeam]
//   - [TeamsApi.DeleteTeamLink]
//   - [TeamsApi.DeleteTeamMembership]
//   - [TeamsApi.GetTeam]
//   - [TeamsApi.GetTeamLink]
//   - [TeamsApi.GetTeamLinks]
//   - [TeamsApi.GetTeamMemberships]
//   - [TeamsApi.GetTeamPermissionSettings]
//   - [TeamsApi.GetUserMemberships]
//   - [TeamsApi.ListTeams]
//   - [TeamsApi.UpdateTeam]
//   - [TeamsApi.UpdateTeamLink]
//   - [TeamsApi.UpdateTeamMembership]
//   - [TeamsApi.UpdateTeamPermissionSetting]
//   - [UsageMeteringApi.GetActiveBillingDimensions]
//   - [UsageMeteringApi.GetBillingDimensionMapping]
//   - [UsageMeteringApi.GetCostByOrg]
//   - [UsageMeteringApi.GetEstimatedCostByOrg]
//   - [UsageMeteringApi.GetHistoricalCostByOrg]
//   - [UsageMeteringApi.GetHourlyUsage]
//   - [UsageMeteringApi.GetMonthlyCostAttribution]
//   - [UsageMeteringApi.GetProjectedCost]
//   - [UsageMeteringApi.GetUsageApplicationSecurityMonitoring]
//   - [UsageMeteringApi.GetUsageLambdaTracedInvocations]
//   - [UsageMeteringApi.GetUsageObservabilityPipelines]
//   - [UsersApi.CreateUser]
//   - [UsersApi.DisableUser]
//   - [UsersApi.GetInvitation]
//   - [UsersApi.GetUser]
//   - [UsersApi.ListUserOrganizations]
//   - [UsersApi.ListUserPermissions]
//   - [UsersApi.ListUsers]
//   - [UsersApi.SendInvitations]
//   - [UsersApi.UpdateUser]
//   - [WorkflowAutomationApi.CancelWorkflowInstance]
//   - [WorkflowAutomationApi.CreateWorkflow]
//   - [WorkflowAutomationApi.CreateWorkflowInstance]
//   - [WorkflowAutomationApi.DeleteWorkflow]
//   - [WorkflowAutomationApi.GetWorkflow]
//   - [WorkflowAutomationApi.GetWorkflowInstance]
//   - [WorkflowAutomationApi.ListWorkflowInstances]
//   - [WorkflowAutomationApi.UpdateWorkflow]
package datadogV2
