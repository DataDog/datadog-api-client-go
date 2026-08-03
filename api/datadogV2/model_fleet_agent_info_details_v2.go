// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentInfoDetailsV2 Detailed information about a Datadog Agent.
type FleetAgentInfoDetailsV2 struct {
	// The currently active agent in the high-availability group.
	ActiveHaAgent *string `json:"active_ha_agent,omitempty"`
	// The Datadog Agent version.
	AgentVersion *string `json:"agent_version,omitempty"`
	// The API key name (if available and not redacted).
	ApiKeyName *string `json:"api_key_name,omitempty"`
	// The API key UUID.
	ApiKeyUuid *string `json:"api_key_uuid,omitempty"`
	// The cloud provider where the agent is running.
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// Kubernetes cluster name (if applicable).
	ClusterName *string `json:"cluster_name,omitempty"`
	// The configuration identifier applied to the agent.
	ConfigId *string `json:"config_id,omitempty"`
	// The unique agent key identifier.
	DatadogAgentKey *string `json:"datadog_agent_key,omitempty"`
	// The Datadog data center the agent reports to.
	DatadogDataCenter *string `json:"datadog_data_center,omitempty"`
	// The ECS Fargate cluster name, if the agent runs in an ECS Fargate environment.
	EcsFargateClusterName *string `json:"ecs_fargate_cluster_name,omitempty"`
	// The ECS Fargate task ARN, if the agent runs in an ECS Fargate environment.
	EcsFargateTaskArn *string `json:"ecs_fargate_task_arn,omitempty"`
	// Datadog products enabled on the agent.
	EnabledProducts []string `json:"enabled_products,omitempty"`
	// Environments the agent is reporting from.
	Env []string `json:"env,omitempty"`
	// Timestamp when the agent was first seen.
	FirstSeenAt *int64 `json:"first_seen_at,omitempty"`
	// Hosts participating in the agent's high-availability group.
	HaAgentHosts []string `json:"ha_agent_hosts,omitempty"`
	// The high-availability state of the agent.
	HaAgentState *string `json:"ha_agent_state,omitempty"`
	// The hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// Alternative hostname list for the agent.
	HostnameAliases []string `json:"hostname_aliases,omitempty"`
	// The version of the installer used.
	InstallMethodInstallerVersion *string `json:"install_method_installer_version,omitempty"`
	// The tool used to install the agent.
	InstallMethodTool *string `json:"install_method_tool,omitempty"`
	// IP addresses of the agent.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Whether single-step instrumentation is enabled.
	IsSingleStepInstrumentationEnabled *bool `json:"is_single_step_instrumentation_enabled,omitempty"`
	// Timestamp of the last agent restart.
	LastRestartAt *int64 `json:"last_restart_at,omitempty"`
	// The operating system.
	Os *string `json:"os,omitempty"`
	// The operating system version.
	OsVersion *string `json:"os_version,omitempty"`
	// OpenTelemetry collector deployment types associated with the agent.
	OtelCollectorDeploymentTypes []string `json:"otel_collector_deployment_types,omitempty"`
	// OpenTelemetry collector distributions associated with the agent.
	OtelCollectorDistributions []string `json:"otel_collector_distributions,omitempty"`
	// OpenTelemetry collector version (if applicable).
	OtelCollectorVersion *string `json:"otel_collector_version,omitempty"`
	// List of OpenTelemetry collector versions (if applicable).
	OtelCollectorVersions []string `json:"otel_collector_versions,omitempty"`
	// OpenTelemetry collectors associated with the agent (if applicable).
	OtelCollectors []map[string]interface{} `json:"otel_collectors,omitempty"`
	// OpenTelemetry resource attributes reported by the agent.
	OtelResourceAttributes []string `json:"otel_resource_attributes,omitempty"`
	// Kubernetes pod name (if applicable).
	PodName *string `json:"pod_name,omitempty"`
	// The preferred active agent in the high-availability group.
	PreferredHaActiveAgent *string `json:"preferred_ha_active_agent,omitempty"`
	// The Python version used by the agent.
	PythonVersion *string `json:"python_version,omitempty"`
	// Regions where the agent is running.
	Region []string `json:"region,omitempty"`
	// Remote agent management status.
	RemoteAgentManagement *string `json:"remote_agent_management,omitempty"`
	// Remote configuration status.
	RemoteConfigStatus *string `json:"remote_config_status,omitempty"`
	// Services running on the agent.
	Services []string `json:"services,omitempty"`
	// Whether the agent supports remote agent upgrade.
	SupportAgentUpgrade *bool `json:"support_agent_upgrade,omitempty"`
	// Tags associated with the agent.
	Tags []string `json:"tags,omitempty"`
	// Team associated with the agent.
	Team *string `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentInfoDetailsV2 instantiates a new FleetAgentInfoDetailsV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentInfoDetailsV2() *FleetAgentInfoDetailsV2 {
	this := FleetAgentInfoDetailsV2{}
	return &this
}

// NewFleetAgentInfoDetailsV2WithDefaults instantiates a new FleetAgentInfoDetailsV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentInfoDetailsV2WithDefaults() *FleetAgentInfoDetailsV2 {
	this := FleetAgentInfoDetailsV2{}
	return &this
}

// GetActiveHaAgent returns the ActiveHaAgent field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetActiveHaAgent() string {
	if o == nil || o.ActiveHaAgent == nil {
		var ret string
		return ret
	}
	return *o.ActiveHaAgent
}

// GetActiveHaAgentOk returns a tuple with the ActiveHaAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetActiveHaAgentOk() (*string, bool) {
	if o == nil || o.ActiveHaAgent == nil {
		return nil, false
	}
	return o.ActiveHaAgent, true
}

// HasActiveHaAgent returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasActiveHaAgent() bool {
	return o != nil && o.ActiveHaAgent != nil
}

// SetActiveHaAgent gets a reference to the given string and assigns it to the ActiveHaAgent field.
func (o *FleetAgentInfoDetailsV2) SetActiveHaAgent(v string) {
	o.ActiveHaAgent = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *FleetAgentInfoDetailsV2) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetApiKeyName returns the ApiKeyName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetApiKeyName() string {
	if o == nil || o.ApiKeyName == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyName
}

// GetApiKeyNameOk returns a tuple with the ApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetApiKeyNameOk() (*string, bool) {
	if o == nil || o.ApiKeyName == nil {
		return nil, false
	}
	return o.ApiKeyName, true
}

// HasApiKeyName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasApiKeyName() bool {
	return o != nil && o.ApiKeyName != nil
}

// SetApiKeyName gets a reference to the given string and assigns it to the ApiKeyName field.
func (o *FleetAgentInfoDetailsV2) SetApiKeyName(v string) {
	o.ApiKeyName = &v
}

// GetApiKeyUuid returns the ApiKeyUuid field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetApiKeyUuid() string {
	if o == nil || o.ApiKeyUuid == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyUuid
}

// GetApiKeyUuidOk returns a tuple with the ApiKeyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetApiKeyUuidOk() (*string, bool) {
	if o == nil || o.ApiKeyUuid == nil {
		return nil, false
	}
	return o.ApiKeyUuid, true
}

// HasApiKeyUuid returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasApiKeyUuid() bool {
	return o != nil && o.ApiKeyUuid != nil
}

// SetApiKeyUuid gets a reference to the given string and assigns it to the ApiKeyUuid field.
func (o *FleetAgentInfoDetailsV2) SetApiKeyUuid(v string) {
	o.ApiKeyUuid = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *FleetAgentInfoDetailsV2) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *FleetAgentInfoDetailsV2) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetConfigId() string {
	if o == nil || o.ConfigId == nil {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetConfigIdOk() (*string, bool) {
	if o == nil || o.ConfigId == nil {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasConfigId() bool {
	return o != nil && o.ConfigId != nil
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *FleetAgentInfoDetailsV2) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil || o.DatadogAgentKey == nil {
		return nil, false
	}
	return o.DatadogAgentKey, true
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey != nil
}

// SetDatadogAgentKey gets a reference to the given string and assigns it to the DatadogAgentKey field.
func (o *FleetAgentInfoDetailsV2) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey = &v
}

// GetDatadogDataCenter returns the DatadogDataCenter field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetDatadogDataCenter() string {
	if o == nil || o.DatadogDataCenter == nil {
		var ret string
		return ret
	}
	return *o.DatadogDataCenter
}

// GetDatadogDataCenterOk returns a tuple with the DatadogDataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetDatadogDataCenterOk() (*string, bool) {
	if o == nil || o.DatadogDataCenter == nil {
		return nil, false
	}
	return o.DatadogDataCenter, true
}

// HasDatadogDataCenter returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasDatadogDataCenter() bool {
	return o != nil && o.DatadogDataCenter != nil
}

// SetDatadogDataCenter gets a reference to the given string and assigns it to the DatadogDataCenter field.
func (o *FleetAgentInfoDetailsV2) SetDatadogDataCenter(v string) {
	o.DatadogDataCenter = &v
}

// GetEcsFargateClusterName returns the EcsFargateClusterName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetEcsFargateClusterName() string {
	if o == nil || o.EcsFargateClusterName == nil {
		var ret string
		return ret
	}
	return *o.EcsFargateClusterName
}

// GetEcsFargateClusterNameOk returns a tuple with the EcsFargateClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetEcsFargateClusterNameOk() (*string, bool) {
	if o == nil || o.EcsFargateClusterName == nil {
		return nil, false
	}
	return o.EcsFargateClusterName, true
}

// HasEcsFargateClusterName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasEcsFargateClusterName() bool {
	return o != nil && o.EcsFargateClusterName != nil
}

// SetEcsFargateClusterName gets a reference to the given string and assigns it to the EcsFargateClusterName field.
func (o *FleetAgentInfoDetailsV2) SetEcsFargateClusterName(v string) {
	o.EcsFargateClusterName = &v
}

// GetEcsFargateTaskArn returns the EcsFargateTaskArn field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetEcsFargateTaskArn() string {
	if o == nil || o.EcsFargateTaskArn == nil {
		var ret string
		return ret
	}
	return *o.EcsFargateTaskArn
}

// GetEcsFargateTaskArnOk returns a tuple with the EcsFargateTaskArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetEcsFargateTaskArnOk() (*string, bool) {
	if o == nil || o.EcsFargateTaskArn == nil {
		return nil, false
	}
	return o.EcsFargateTaskArn, true
}

// HasEcsFargateTaskArn returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasEcsFargateTaskArn() bool {
	return o != nil && o.EcsFargateTaskArn != nil
}

// SetEcsFargateTaskArn gets a reference to the given string and assigns it to the EcsFargateTaskArn field.
func (o *FleetAgentInfoDetailsV2) SetEcsFargateTaskArn(v string) {
	o.EcsFargateTaskArn = &v
}

// GetEnabledProducts returns the EnabledProducts field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetEnabledProducts() []string {
	if o == nil || o.EnabledProducts == nil {
		var ret []string
		return ret
	}
	return o.EnabledProducts
}

// GetEnabledProductsOk returns a tuple with the EnabledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetEnabledProductsOk() (*[]string, bool) {
	if o == nil || o.EnabledProducts == nil {
		return nil, false
	}
	return &o.EnabledProducts, true
}

// HasEnabledProducts returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasEnabledProducts() bool {
	return o != nil && o.EnabledProducts != nil
}

// SetEnabledProducts gets a reference to the given []string and assigns it to the EnabledProducts field.
func (o *FleetAgentInfoDetailsV2) SetEnabledProducts(v []string) {
	o.EnabledProducts = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetEnv() []string {
	if o == nil || o.Env == nil {
		var ret []string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetEnvOk() (*[]string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return &o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given []string and assigns it to the Env field.
func (o *FleetAgentInfoDetailsV2) SetEnv(v []string) {
	o.Env = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetFirstSeenAt() int64 {
	if o == nil || o.FirstSeenAt == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetFirstSeenAtOk() (*int64, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given int64 and assigns it to the FirstSeenAt field.
func (o *FleetAgentInfoDetailsV2) SetFirstSeenAt(v int64) {
	o.FirstSeenAt = &v
}

// GetHaAgentHosts returns the HaAgentHosts field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetHaAgentHosts() []string {
	if o == nil || o.HaAgentHosts == nil {
		var ret []string
		return ret
	}
	return o.HaAgentHosts
}

// GetHaAgentHostsOk returns a tuple with the HaAgentHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetHaAgentHostsOk() (*[]string, bool) {
	if o == nil || o.HaAgentHosts == nil {
		return nil, false
	}
	return &o.HaAgentHosts, true
}

// HasHaAgentHosts returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasHaAgentHosts() bool {
	return o != nil && o.HaAgentHosts != nil
}

// SetHaAgentHosts gets a reference to the given []string and assigns it to the HaAgentHosts field.
func (o *FleetAgentInfoDetailsV2) SetHaAgentHosts(v []string) {
	o.HaAgentHosts = v
}

// GetHaAgentState returns the HaAgentState field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetHaAgentState() string {
	if o == nil || o.HaAgentState == nil {
		var ret string
		return ret
	}
	return *o.HaAgentState
}

// GetHaAgentStateOk returns a tuple with the HaAgentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetHaAgentStateOk() (*string, bool) {
	if o == nil || o.HaAgentState == nil {
		return nil, false
	}
	return o.HaAgentState, true
}

// HasHaAgentState returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasHaAgentState() bool {
	return o != nil && o.HaAgentState != nil
}

// SetHaAgentState gets a reference to the given string and assigns it to the HaAgentState field.
func (o *FleetAgentInfoDetailsV2) SetHaAgentState(v string) {
	o.HaAgentState = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetAgentInfoDetailsV2) SetHostname(v string) {
	o.Hostname = &v
}

// GetHostnameAliases returns the HostnameAliases field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetHostnameAliases() []string {
	if o == nil || o.HostnameAliases == nil {
		var ret []string
		return ret
	}
	return o.HostnameAliases
}

// GetHostnameAliasesOk returns a tuple with the HostnameAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetHostnameAliasesOk() (*[]string, bool) {
	if o == nil || o.HostnameAliases == nil {
		return nil, false
	}
	return &o.HostnameAliases, true
}

// HasHostnameAliases returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasHostnameAliases() bool {
	return o != nil && o.HostnameAliases != nil
}

// SetHostnameAliases gets a reference to the given []string and assigns it to the HostnameAliases field.
func (o *FleetAgentInfoDetailsV2) SetHostnameAliases(v []string) {
	o.HostnameAliases = v
}

// GetInstallMethodInstallerVersion returns the InstallMethodInstallerVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetInstallMethodInstallerVersion() string {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodInstallerVersion
}

// GetInstallMethodInstallerVersionOk returns a tuple with the InstallMethodInstallerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetInstallMethodInstallerVersionOk() (*string, bool) {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		return nil, false
	}
	return o.InstallMethodInstallerVersion, true
}

// HasInstallMethodInstallerVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasInstallMethodInstallerVersion() bool {
	return o != nil && o.InstallMethodInstallerVersion != nil
}

// SetInstallMethodInstallerVersion gets a reference to the given string and assigns it to the InstallMethodInstallerVersion field.
func (o *FleetAgentInfoDetailsV2) SetInstallMethodInstallerVersion(v string) {
	o.InstallMethodInstallerVersion = &v
}

// GetInstallMethodTool returns the InstallMethodTool field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetInstallMethodTool() string {
	if o == nil || o.InstallMethodTool == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodTool
}

// GetInstallMethodToolOk returns a tuple with the InstallMethodTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetInstallMethodToolOk() (*string, bool) {
	if o == nil || o.InstallMethodTool == nil {
		return nil, false
	}
	return o.InstallMethodTool, true
}

// HasInstallMethodTool returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasInstallMethodTool() bool {
	return o != nil && o.InstallMethodTool != nil
}

// SetInstallMethodTool gets a reference to the given string and assigns it to the InstallMethodTool field.
func (o *FleetAgentInfoDetailsV2) SetInstallMethodTool(v string) {
	o.InstallMethodTool = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasIpAddresses() bool {
	return o != nil && o.IpAddresses != nil
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *FleetAgentInfoDetailsV2) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetIsSingleStepInstrumentationEnabled returns the IsSingleStepInstrumentationEnabled field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetIsSingleStepInstrumentationEnabled() bool {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSingleStepInstrumentationEnabled
}

// GetIsSingleStepInstrumentationEnabledOk returns a tuple with the IsSingleStepInstrumentationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetIsSingleStepInstrumentationEnabledOk() (*bool, bool) {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		return nil, false
	}
	return o.IsSingleStepInstrumentationEnabled, true
}

// HasIsSingleStepInstrumentationEnabled returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasIsSingleStepInstrumentationEnabled() bool {
	return o != nil && o.IsSingleStepInstrumentationEnabled != nil
}

// SetIsSingleStepInstrumentationEnabled gets a reference to the given bool and assigns it to the IsSingleStepInstrumentationEnabled field.
func (o *FleetAgentInfoDetailsV2) SetIsSingleStepInstrumentationEnabled(v bool) {
	o.IsSingleStepInstrumentationEnabled = &v
}

// GetLastRestartAt returns the LastRestartAt field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetLastRestartAt() int64 {
	if o == nil || o.LastRestartAt == nil {
		var ret int64
		return ret
	}
	return *o.LastRestartAt
}

// GetLastRestartAtOk returns a tuple with the LastRestartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetLastRestartAtOk() (*int64, bool) {
	if o == nil || o.LastRestartAt == nil {
		return nil, false
	}
	return o.LastRestartAt, true
}

// HasLastRestartAt returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasLastRestartAt() bool {
	return o != nil && o.LastRestartAt != nil
}

// SetLastRestartAt gets a reference to the given int64 and assigns it to the LastRestartAt field.
func (o *FleetAgentInfoDetailsV2) SetLastRestartAt(v int64) {
	o.LastRestartAt = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *FleetAgentInfoDetailsV2) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *FleetAgentInfoDetailsV2) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetOtelCollectorDeploymentTypes returns the OtelCollectorDeploymentTypes field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorDeploymentTypes() []string {
	if o == nil || o.OtelCollectorDeploymentTypes == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorDeploymentTypes
}

// GetOtelCollectorDeploymentTypesOk returns a tuple with the OtelCollectorDeploymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorDeploymentTypesOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorDeploymentTypes == nil {
		return nil, false
	}
	return &o.OtelCollectorDeploymentTypes, true
}

// HasOtelCollectorDeploymentTypes returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelCollectorDeploymentTypes() bool {
	return o != nil && o.OtelCollectorDeploymentTypes != nil
}

// SetOtelCollectorDeploymentTypes gets a reference to the given []string and assigns it to the OtelCollectorDeploymentTypes field.
func (o *FleetAgentInfoDetailsV2) SetOtelCollectorDeploymentTypes(v []string) {
	o.OtelCollectorDeploymentTypes = v
}

// GetOtelCollectorDistributions returns the OtelCollectorDistributions field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorDistributions() []string {
	if o == nil || o.OtelCollectorDistributions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorDistributions
}

// GetOtelCollectorDistributionsOk returns a tuple with the OtelCollectorDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorDistributionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorDistributions == nil {
		return nil, false
	}
	return &o.OtelCollectorDistributions, true
}

// HasOtelCollectorDistributions returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelCollectorDistributions() bool {
	return o != nil && o.OtelCollectorDistributions != nil
}

// SetOtelCollectorDistributions gets a reference to the given []string and assigns it to the OtelCollectorDistributions field.
func (o *FleetAgentInfoDetailsV2) SetOtelCollectorDistributions(v []string) {
	o.OtelCollectorDistributions = v
}

// GetOtelCollectorVersion returns the OtelCollectorVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorVersion() string {
	if o == nil || o.OtelCollectorVersion == nil {
		var ret string
		return ret
	}
	return *o.OtelCollectorVersion
}

// GetOtelCollectorVersionOk returns a tuple with the OtelCollectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorVersionOk() (*string, bool) {
	if o == nil || o.OtelCollectorVersion == nil {
		return nil, false
	}
	return o.OtelCollectorVersion, true
}

// HasOtelCollectorVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelCollectorVersion() bool {
	return o != nil && o.OtelCollectorVersion != nil
}

// SetOtelCollectorVersion gets a reference to the given string and assigns it to the OtelCollectorVersion field.
func (o *FleetAgentInfoDetailsV2) SetOtelCollectorVersion(v string) {
	o.OtelCollectorVersion = &v
}

// GetOtelCollectorVersions returns the OtelCollectorVersions field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorVersions() []string {
	if o == nil || o.OtelCollectorVersions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorVersions
}

// GetOtelCollectorVersionsOk returns a tuple with the OtelCollectorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorVersionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorVersions == nil {
		return nil, false
	}
	return &o.OtelCollectorVersions, true
}

// HasOtelCollectorVersions returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelCollectorVersions() bool {
	return o != nil && o.OtelCollectorVersions != nil
}

// SetOtelCollectorVersions gets a reference to the given []string and assigns it to the OtelCollectorVersions field.
func (o *FleetAgentInfoDetailsV2) SetOtelCollectorVersions(v []string) {
	o.OtelCollectorVersions = v
}

// GetOtelCollectors returns the OtelCollectors field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectors() []map[string]interface{} {
	if o == nil || o.OtelCollectors == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.OtelCollectors
}

// GetOtelCollectorsOk returns a tuple with the OtelCollectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelCollectorsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.OtelCollectors == nil {
		return nil, false
	}
	return &o.OtelCollectors, true
}

// HasOtelCollectors returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelCollectors() bool {
	return o != nil && o.OtelCollectors != nil
}

// SetOtelCollectors gets a reference to the given []map[string]interface{} and assigns it to the OtelCollectors field.
func (o *FleetAgentInfoDetailsV2) SetOtelCollectors(v []map[string]interface{}) {
	o.OtelCollectors = v
}

// GetOtelResourceAttributes returns the OtelResourceAttributes field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetOtelResourceAttributes() []string {
	if o == nil || o.OtelResourceAttributes == nil {
		var ret []string
		return ret
	}
	return o.OtelResourceAttributes
}

// GetOtelResourceAttributesOk returns a tuple with the OtelResourceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetOtelResourceAttributesOk() (*[]string, bool) {
	if o == nil || o.OtelResourceAttributes == nil {
		return nil, false
	}
	return &o.OtelResourceAttributes, true
}

// HasOtelResourceAttributes returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasOtelResourceAttributes() bool {
	return o != nil && o.OtelResourceAttributes != nil
}

// SetOtelResourceAttributes gets a reference to the given []string and assigns it to the OtelResourceAttributes field.
func (o *FleetAgentInfoDetailsV2) SetOtelResourceAttributes(v []string) {
	o.OtelResourceAttributes = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *FleetAgentInfoDetailsV2) SetPodName(v string) {
	o.PodName = &v
}

// GetPreferredHaActiveAgent returns the PreferredHaActiveAgent field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetPreferredHaActiveAgent() string {
	if o == nil || o.PreferredHaActiveAgent == nil {
		var ret string
		return ret
	}
	return *o.PreferredHaActiveAgent
}

// GetPreferredHaActiveAgentOk returns a tuple with the PreferredHaActiveAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetPreferredHaActiveAgentOk() (*string, bool) {
	if o == nil || o.PreferredHaActiveAgent == nil {
		return nil, false
	}
	return o.PreferredHaActiveAgent, true
}

// HasPreferredHaActiveAgent returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasPreferredHaActiveAgent() bool {
	return o != nil && o.PreferredHaActiveAgent != nil
}

// SetPreferredHaActiveAgent gets a reference to the given string and assigns it to the PreferredHaActiveAgent field.
func (o *FleetAgentInfoDetailsV2) SetPreferredHaActiveAgent(v string) {
	o.PreferredHaActiveAgent = &v
}

// GetPythonVersion returns the PythonVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetPythonVersion() string {
	if o == nil || o.PythonVersion == nil {
		var ret string
		return ret
	}
	return *o.PythonVersion
}

// GetPythonVersionOk returns a tuple with the PythonVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetPythonVersionOk() (*string, bool) {
	if o == nil || o.PythonVersion == nil {
		return nil, false
	}
	return o.PythonVersion, true
}

// HasPythonVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasPythonVersion() bool {
	return o != nil && o.PythonVersion != nil
}

// SetPythonVersion gets a reference to the given string and assigns it to the PythonVersion field.
func (o *FleetAgentInfoDetailsV2) SetPythonVersion(v string) {
	o.PythonVersion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetRegion() []string {
	if o == nil || o.Region == nil {
		var ret []string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetRegionOk() (*[]string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return &o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given []string and assigns it to the Region field.
func (o *FleetAgentInfoDetailsV2) SetRegion(v []string) {
	o.Region = v
}

// GetRemoteAgentManagement returns the RemoteAgentManagement field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetRemoteAgentManagement() string {
	if o == nil || o.RemoteAgentManagement == nil {
		var ret string
		return ret
	}
	return *o.RemoteAgentManagement
}

// GetRemoteAgentManagementOk returns a tuple with the RemoteAgentManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetRemoteAgentManagementOk() (*string, bool) {
	if o == nil || o.RemoteAgentManagement == nil {
		return nil, false
	}
	return o.RemoteAgentManagement, true
}

// HasRemoteAgentManagement returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasRemoteAgentManagement() bool {
	return o != nil && o.RemoteAgentManagement != nil
}

// SetRemoteAgentManagement gets a reference to the given string and assigns it to the RemoteAgentManagement field.
func (o *FleetAgentInfoDetailsV2) SetRemoteAgentManagement(v string) {
	o.RemoteAgentManagement = &v
}

// GetRemoteConfigStatus returns the RemoteConfigStatus field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetRemoteConfigStatus() string {
	if o == nil || o.RemoteConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigStatus
}

// GetRemoteConfigStatusOk returns a tuple with the RemoteConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetRemoteConfigStatusOk() (*string, bool) {
	if o == nil || o.RemoteConfigStatus == nil {
		return nil, false
	}
	return o.RemoteConfigStatus, true
}

// HasRemoteConfigStatus returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasRemoteConfigStatus() bool {
	return o != nil && o.RemoteConfigStatus != nil
}

// SetRemoteConfigStatus gets a reference to the given string and assigns it to the RemoteConfigStatus field.
func (o *FleetAgentInfoDetailsV2) SetRemoteConfigStatus(v string) {
	o.RemoteConfigStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FleetAgentInfoDetailsV2) SetServices(v []string) {
	o.Services = v
}

// GetSupportAgentUpgrade returns the SupportAgentUpgrade field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetSupportAgentUpgrade() bool {
	if o == nil || o.SupportAgentUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.SupportAgentUpgrade
}

// GetSupportAgentUpgradeOk returns a tuple with the SupportAgentUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetSupportAgentUpgradeOk() (*bool, bool) {
	if o == nil || o.SupportAgentUpgrade == nil {
		return nil, false
	}
	return o.SupportAgentUpgrade, true
}

// HasSupportAgentUpgrade returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasSupportAgentUpgrade() bool {
	return o != nil && o.SupportAgentUpgrade != nil
}

// SetSupportAgentUpgrade gets a reference to the given bool and assigns it to the SupportAgentUpgrade field.
func (o *FleetAgentInfoDetailsV2) SetSupportAgentUpgrade(v bool) {
	o.SupportAgentUpgrade = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FleetAgentInfoDetailsV2) SetTags(v []string) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *FleetAgentInfoDetailsV2) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetailsV2) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *FleetAgentInfoDetailsV2) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *FleetAgentInfoDetailsV2) SetTeam(v string) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentInfoDetailsV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ActiveHaAgent != nil {
		toSerialize["active_ha_agent"] = o.ActiveHaAgent
	}
	if o.AgentVersion != nil {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if o.ApiKeyName != nil {
		toSerialize["api_key_name"] = o.ApiKeyName
	}
	if o.ApiKeyUuid != nil {
		toSerialize["api_key_uuid"] = o.ApiKeyUuid
	}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.ConfigId != nil {
		toSerialize["config_id"] = o.ConfigId
	}
	if o.DatadogAgentKey != nil {
		toSerialize["datadog_agent_key"] = o.DatadogAgentKey
	}
	if o.DatadogDataCenter != nil {
		toSerialize["datadog_data_center"] = o.DatadogDataCenter
	}
	if o.EcsFargateClusterName != nil {
		toSerialize["ecs_fargate_cluster_name"] = o.EcsFargateClusterName
	}
	if o.EcsFargateTaskArn != nil {
		toSerialize["ecs_fargate_task_arn"] = o.EcsFargateTaskArn
	}
	if o.EnabledProducts != nil {
		toSerialize["enabled_products"] = o.EnabledProducts
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FirstSeenAt != nil {
		toSerialize["first_seen_at"] = o.FirstSeenAt
	}
	if o.HaAgentHosts != nil {
		toSerialize["ha_agent_hosts"] = o.HaAgentHosts
	}
	if o.HaAgentState != nil {
		toSerialize["ha_agent_state"] = o.HaAgentState
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.HostnameAliases != nil {
		toSerialize["hostname_aliases"] = o.HostnameAliases
	}
	if o.InstallMethodInstallerVersion != nil {
		toSerialize["install_method_installer_version"] = o.InstallMethodInstallerVersion
	}
	if o.InstallMethodTool != nil {
		toSerialize["install_method_tool"] = o.InstallMethodTool
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.IsSingleStepInstrumentationEnabled != nil {
		toSerialize["is_single_step_instrumentation_enabled"] = o.IsSingleStepInstrumentationEnabled
	}
	if o.LastRestartAt != nil {
		toSerialize["last_restart_at"] = o.LastRestartAt
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.OtelCollectorDeploymentTypes != nil {
		toSerialize["otel_collector_deployment_types"] = o.OtelCollectorDeploymentTypes
	}
	if o.OtelCollectorDistributions != nil {
		toSerialize["otel_collector_distributions"] = o.OtelCollectorDistributions
	}
	if o.OtelCollectorVersion != nil {
		toSerialize["otel_collector_version"] = o.OtelCollectorVersion
	}
	if o.OtelCollectorVersions != nil {
		toSerialize["otel_collector_versions"] = o.OtelCollectorVersions
	}
	if o.OtelCollectors != nil {
		toSerialize["otel_collectors"] = o.OtelCollectors
	}
	if o.OtelResourceAttributes != nil {
		toSerialize["otel_resource_attributes"] = o.OtelResourceAttributes
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
	}
	if o.PreferredHaActiveAgent != nil {
		toSerialize["preferred_ha_active_agent"] = o.PreferredHaActiveAgent
	}
	if o.PythonVersion != nil {
		toSerialize["python_version"] = o.PythonVersion
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RemoteAgentManagement != nil {
		toSerialize["remote_agent_management"] = o.RemoteAgentManagement
	}
	if o.RemoteConfigStatus != nil {
		toSerialize["remote_config_status"] = o.RemoteConfigStatus
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.SupportAgentUpgrade != nil {
		toSerialize["support_agent_upgrade"] = o.SupportAgentUpgrade
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentInfoDetailsV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActiveHaAgent                      *string                  `json:"active_ha_agent,omitempty"`
		AgentVersion                       *string                  `json:"agent_version,omitempty"`
		ApiKeyName                         *string                  `json:"api_key_name,omitempty"`
		ApiKeyUuid                         *string                  `json:"api_key_uuid,omitempty"`
		CloudProvider                      *string                  `json:"cloud_provider,omitempty"`
		ClusterName                        *string                  `json:"cluster_name,omitempty"`
		ConfigId                           *string                  `json:"config_id,omitempty"`
		DatadogAgentKey                    *string                  `json:"datadog_agent_key,omitempty"`
		DatadogDataCenter                  *string                  `json:"datadog_data_center,omitempty"`
		EcsFargateClusterName              *string                  `json:"ecs_fargate_cluster_name,omitempty"`
		EcsFargateTaskArn                  *string                  `json:"ecs_fargate_task_arn,omitempty"`
		EnabledProducts                    []string                 `json:"enabled_products,omitempty"`
		Env                                []string                 `json:"env,omitempty"`
		FirstSeenAt                        *int64                   `json:"first_seen_at,omitempty"`
		HaAgentHosts                       []string                 `json:"ha_agent_hosts,omitempty"`
		HaAgentState                       *string                  `json:"ha_agent_state,omitempty"`
		Hostname                           *string                  `json:"hostname,omitempty"`
		HostnameAliases                    []string                 `json:"hostname_aliases,omitempty"`
		InstallMethodInstallerVersion      *string                  `json:"install_method_installer_version,omitempty"`
		InstallMethodTool                  *string                  `json:"install_method_tool,omitempty"`
		IpAddresses                        []string                 `json:"ip_addresses,omitempty"`
		IsSingleStepInstrumentationEnabled *bool                    `json:"is_single_step_instrumentation_enabled,omitempty"`
		LastRestartAt                      *int64                   `json:"last_restart_at,omitempty"`
		Os                                 *string                  `json:"os,omitempty"`
		OsVersion                          *string                  `json:"os_version,omitempty"`
		OtelCollectorDeploymentTypes       []string                 `json:"otel_collector_deployment_types,omitempty"`
		OtelCollectorDistributions         []string                 `json:"otel_collector_distributions,omitempty"`
		OtelCollectorVersion               *string                  `json:"otel_collector_version,omitempty"`
		OtelCollectorVersions              []string                 `json:"otel_collector_versions,omitempty"`
		OtelCollectors                     []map[string]interface{} `json:"otel_collectors,omitempty"`
		OtelResourceAttributes             []string                 `json:"otel_resource_attributes,omitempty"`
		PodName                            *string                  `json:"pod_name,omitempty"`
		PreferredHaActiveAgent             *string                  `json:"preferred_ha_active_agent,omitempty"`
		PythonVersion                      *string                  `json:"python_version,omitempty"`
		Region                             []string                 `json:"region,omitempty"`
		RemoteAgentManagement              *string                  `json:"remote_agent_management,omitempty"`
		RemoteConfigStatus                 *string                  `json:"remote_config_status,omitempty"`
		Services                           []string                 `json:"services,omitempty"`
		SupportAgentUpgrade                *bool                    `json:"support_agent_upgrade,omitempty"`
		Tags                               []string                 `json:"tags,omitempty"`
		Team                               *string                  `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active_ha_agent", "agent_version", "api_key_name", "api_key_uuid", "cloud_provider", "cluster_name", "config_id", "datadog_agent_key", "datadog_data_center", "ecs_fargate_cluster_name", "ecs_fargate_task_arn", "enabled_products", "env", "first_seen_at", "ha_agent_hosts", "ha_agent_state", "hostname", "hostname_aliases", "install_method_installer_version", "install_method_tool", "ip_addresses", "is_single_step_instrumentation_enabled", "last_restart_at", "os", "os_version", "otel_collector_deployment_types", "otel_collector_distributions", "otel_collector_version", "otel_collector_versions", "otel_collectors", "otel_resource_attributes", "pod_name", "preferred_ha_active_agent", "python_version", "region", "remote_agent_management", "remote_config_status", "services", "support_agent_upgrade", "tags", "team"})
	} else {
		return err
	}
	o.ActiveHaAgent = all.ActiveHaAgent
	o.AgentVersion = all.AgentVersion
	o.ApiKeyName = all.ApiKeyName
	o.ApiKeyUuid = all.ApiKeyUuid
	o.CloudProvider = all.CloudProvider
	o.ClusterName = all.ClusterName
	o.ConfigId = all.ConfigId
	o.DatadogAgentKey = all.DatadogAgentKey
	o.DatadogDataCenter = all.DatadogDataCenter
	o.EcsFargateClusterName = all.EcsFargateClusterName
	o.EcsFargateTaskArn = all.EcsFargateTaskArn
	o.EnabledProducts = all.EnabledProducts
	o.Env = all.Env
	o.FirstSeenAt = all.FirstSeenAt
	o.HaAgentHosts = all.HaAgentHosts
	o.HaAgentState = all.HaAgentState
	o.Hostname = all.Hostname
	o.HostnameAliases = all.HostnameAliases
	o.InstallMethodInstallerVersion = all.InstallMethodInstallerVersion
	o.InstallMethodTool = all.InstallMethodTool
	o.IpAddresses = all.IpAddresses
	o.IsSingleStepInstrumentationEnabled = all.IsSingleStepInstrumentationEnabled
	o.LastRestartAt = all.LastRestartAt
	o.Os = all.Os
	o.OsVersion = all.OsVersion
	o.OtelCollectorDeploymentTypes = all.OtelCollectorDeploymentTypes
	o.OtelCollectorDistributions = all.OtelCollectorDistributions
	o.OtelCollectorVersion = all.OtelCollectorVersion
	o.OtelCollectorVersions = all.OtelCollectorVersions
	o.OtelCollectors = all.OtelCollectors
	o.OtelResourceAttributes = all.OtelResourceAttributes
	o.PodName = all.PodName
	o.PreferredHaActiveAgent = all.PreferredHaActiveAgent
	o.PythonVersion = all.PythonVersion
	o.Region = all.Region
	o.RemoteAgentManagement = all.RemoteAgentManagement
	o.RemoteConfigStatus = all.RemoteConfigStatus
	o.Services = all.Services
	o.SupportAgentUpgrade = all.SupportAgentUpgrade
	o.Tags = all.Tags
	o.Team = all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
