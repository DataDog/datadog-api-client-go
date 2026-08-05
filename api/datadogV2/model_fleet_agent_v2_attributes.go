// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentV2Attributes Attributes of a Datadog Agent in the v2 list response.
type FleetAgentV2Attributes struct {
	// The Datadog Agent version.
	AgentVersion *string `json:"agent_version,omitempty"`
	// The name of the API key used by the agent, if available and not redacted.
	ApiKeyName *string `json:"api_key_name,omitempty"`
	// The UUID of the API key used by the agent.
	ApiKeyUuid *string `json:"api_key_uuid,omitempty"`
	// The cloud provider where the agent is running.
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// The Kubernetes cluster name, if the agent runs in a cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
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
	// Unix timestamp when the agent was first seen.
	FirstSeenAt *int64 `json:"first_seen_at,omitempty"`
	// Identifiers of fleet policies applied to the agent.
	FleetPolicies []string `json:"fleet_policies,omitempty"`
	// The hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// Number of instrumentation errors on the agent. Absent from the response when the count is zero.
	InstrumentationErrorCounts *int64 `json:"instrumentation_error_counts,omitempty"`
	// The single-step instrumentation status of the Agent.
	InstrumentationStatus *FleetAgentV2AttributesInstrumentationStatus `json:"instrumentation_status,omitempty"`
	// Names of integrations configured on the agent.
	Integrations []string `json:"integrations,omitempty"`
	// IP addresses of the agent host.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Whether single-step instrumentation is enabled on the agent.
	IsSingleStepInstrumentationEnabled *bool `json:"is_single_step_instrumentation_enabled,omitempty"`
	// Unix timestamp of the last agent restart.
	LastRestartAt *int64 `json:"last_restart_at,omitempty"`
	// The operating system of the host.
	Os *string `json:"os,omitempty"`
	// OpenTelemetry collector deployment types associated with the agent.
	OtelCollectorDeploymentTypes []string `json:"otel_collector_deployment_types,omitempty"`
	// OpenTelemetry collector distributions associated with the agent.
	OtelCollectorDistributions []string `json:"otel_collector_distributions,omitempty"`
	// All OpenTelemetry collector versions associated with the agent.
	OtelCollectorVersions []string `json:"otel_collector_versions,omitempty"`
	// OpenTelemetry resource attributes reported by the agent.
	OtelResourceAttributes []string `json:"otel_resource_attributes,omitempty"`
	// The Kubernetes pod name, if the agent runs as a pod.
	PodName *string `json:"pod_name,omitempty"`
	// The remote agent management status.
	RemoteAgentManagement *string `json:"remote_agent_management,omitempty"`
	// The remote configuration connection status of the agent.
	RemoteConfigStatus *string `json:"remote_config_status,omitempty"`
	// Services running on the agent.
	Services []string `json:"services,omitempty"`
	// Tags associated with the agent. Returned as an empty array when the agent has no tags.
	Tags []FleetAgentAttributesTagsItems `json:"tags,omitempty"`
	// The team associated with the agent.
	Team *string `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentV2Attributes instantiates a new FleetAgentV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentV2Attributes() *FleetAgentV2Attributes {
	this := FleetAgentV2Attributes{}
	return &this
}

// NewFleetAgentV2AttributesWithDefaults instantiates a new FleetAgentV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentV2AttributesWithDefaults() *FleetAgentV2Attributes {
	this := FleetAgentV2Attributes{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *FleetAgentV2Attributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetApiKeyName returns the ApiKeyName field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetApiKeyName() string {
	if o == nil || o.ApiKeyName == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyName
}

// GetApiKeyNameOk returns a tuple with the ApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetApiKeyNameOk() (*string, bool) {
	if o == nil || o.ApiKeyName == nil {
		return nil, false
	}
	return o.ApiKeyName, true
}

// HasApiKeyName returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasApiKeyName() bool {
	return o != nil && o.ApiKeyName != nil
}

// SetApiKeyName gets a reference to the given string and assigns it to the ApiKeyName field.
func (o *FleetAgentV2Attributes) SetApiKeyName(v string) {
	o.ApiKeyName = &v
}

// GetApiKeyUuid returns the ApiKeyUuid field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetApiKeyUuid() string {
	if o == nil || o.ApiKeyUuid == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyUuid
}

// GetApiKeyUuidOk returns a tuple with the ApiKeyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetApiKeyUuidOk() (*string, bool) {
	if o == nil || o.ApiKeyUuid == nil {
		return nil, false
	}
	return o.ApiKeyUuid, true
}

// HasApiKeyUuid returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasApiKeyUuid() bool {
	return o != nil && o.ApiKeyUuid != nil
}

// SetApiKeyUuid gets a reference to the given string and assigns it to the ApiKeyUuid field.
func (o *FleetAgentV2Attributes) SetApiKeyUuid(v string) {
	o.ApiKeyUuid = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *FleetAgentV2Attributes) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *FleetAgentV2Attributes) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatadogDataCenter returns the DatadogDataCenter field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetDatadogDataCenter() string {
	if o == nil || o.DatadogDataCenter == nil {
		var ret string
		return ret
	}
	return *o.DatadogDataCenter
}

// GetDatadogDataCenterOk returns a tuple with the DatadogDataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetDatadogDataCenterOk() (*string, bool) {
	if o == nil || o.DatadogDataCenter == nil {
		return nil, false
	}
	return o.DatadogDataCenter, true
}

// HasDatadogDataCenter returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasDatadogDataCenter() bool {
	return o != nil && o.DatadogDataCenter != nil
}

// SetDatadogDataCenter gets a reference to the given string and assigns it to the DatadogDataCenter field.
func (o *FleetAgentV2Attributes) SetDatadogDataCenter(v string) {
	o.DatadogDataCenter = &v
}

// GetEcsFargateClusterName returns the EcsFargateClusterName field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetEcsFargateClusterName() string {
	if o == nil || o.EcsFargateClusterName == nil {
		var ret string
		return ret
	}
	return *o.EcsFargateClusterName
}

// GetEcsFargateClusterNameOk returns a tuple with the EcsFargateClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetEcsFargateClusterNameOk() (*string, bool) {
	if o == nil || o.EcsFargateClusterName == nil {
		return nil, false
	}
	return o.EcsFargateClusterName, true
}

// HasEcsFargateClusterName returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasEcsFargateClusterName() bool {
	return o != nil && o.EcsFargateClusterName != nil
}

// SetEcsFargateClusterName gets a reference to the given string and assigns it to the EcsFargateClusterName field.
func (o *FleetAgentV2Attributes) SetEcsFargateClusterName(v string) {
	o.EcsFargateClusterName = &v
}

// GetEcsFargateTaskArn returns the EcsFargateTaskArn field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetEcsFargateTaskArn() string {
	if o == nil || o.EcsFargateTaskArn == nil {
		var ret string
		return ret
	}
	return *o.EcsFargateTaskArn
}

// GetEcsFargateTaskArnOk returns a tuple with the EcsFargateTaskArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetEcsFargateTaskArnOk() (*string, bool) {
	if o == nil || o.EcsFargateTaskArn == nil {
		return nil, false
	}
	return o.EcsFargateTaskArn, true
}

// HasEcsFargateTaskArn returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasEcsFargateTaskArn() bool {
	return o != nil && o.EcsFargateTaskArn != nil
}

// SetEcsFargateTaskArn gets a reference to the given string and assigns it to the EcsFargateTaskArn field.
func (o *FleetAgentV2Attributes) SetEcsFargateTaskArn(v string) {
	o.EcsFargateTaskArn = &v
}

// GetEnabledProducts returns the EnabledProducts field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetEnabledProducts() []string {
	if o == nil || o.EnabledProducts == nil {
		var ret []string
		return ret
	}
	return o.EnabledProducts
}

// GetEnabledProductsOk returns a tuple with the EnabledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetEnabledProductsOk() (*[]string, bool) {
	if o == nil || o.EnabledProducts == nil {
		return nil, false
	}
	return &o.EnabledProducts, true
}

// HasEnabledProducts returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasEnabledProducts() bool {
	return o != nil && o.EnabledProducts != nil
}

// SetEnabledProducts gets a reference to the given []string and assigns it to the EnabledProducts field.
func (o *FleetAgentV2Attributes) SetEnabledProducts(v []string) {
	o.EnabledProducts = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetEnv() []string {
	if o == nil || o.Env == nil {
		var ret []string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetEnvOk() (*[]string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return &o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given []string and assigns it to the Env field.
func (o *FleetAgentV2Attributes) SetEnv(v []string) {
	o.Env = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetFirstSeenAt() int64 {
	if o == nil || o.FirstSeenAt == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetFirstSeenAtOk() (*int64, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given int64 and assigns it to the FirstSeenAt field.
func (o *FleetAgentV2Attributes) SetFirstSeenAt(v int64) {
	o.FirstSeenAt = &v
}

// GetFleetPolicies returns the FleetPolicies field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetFleetPolicies() []string {
	if o == nil || o.FleetPolicies == nil {
		var ret []string
		return ret
	}
	return o.FleetPolicies
}

// GetFleetPoliciesOk returns a tuple with the FleetPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetFleetPoliciesOk() (*[]string, bool) {
	if o == nil || o.FleetPolicies == nil {
		return nil, false
	}
	return &o.FleetPolicies, true
}

// HasFleetPolicies returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasFleetPolicies() bool {
	return o != nil && o.FleetPolicies != nil
}

// SetFleetPolicies gets a reference to the given []string and assigns it to the FleetPolicies field.
func (o *FleetAgentV2Attributes) SetFleetPolicies(v []string) {
	o.FleetPolicies = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetAgentV2Attributes) SetHostname(v string) {
	o.Hostname = &v
}

// GetInstrumentationErrorCounts returns the InstrumentationErrorCounts field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetInstrumentationErrorCounts() int64 {
	if o == nil || o.InstrumentationErrorCounts == nil {
		var ret int64
		return ret
	}
	return *o.InstrumentationErrorCounts
}

// GetInstrumentationErrorCountsOk returns a tuple with the InstrumentationErrorCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetInstrumentationErrorCountsOk() (*int64, bool) {
	if o == nil || o.InstrumentationErrorCounts == nil {
		return nil, false
	}
	return o.InstrumentationErrorCounts, true
}

// HasInstrumentationErrorCounts returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasInstrumentationErrorCounts() bool {
	return o != nil && o.InstrumentationErrorCounts != nil
}

// SetInstrumentationErrorCounts gets a reference to the given int64 and assigns it to the InstrumentationErrorCounts field.
func (o *FleetAgentV2Attributes) SetInstrumentationErrorCounts(v int64) {
	o.InstrumentationErrorCounts = &v
}

// GetInstrumentationStatus returns the InstrumentationStatus field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetInstrumentationStatus() FleetAgentV2AttributesInstrumentationStatus {
	if o == nil || o.InstrumentationStatus == nil {
		var ret FleetAgentV2AttributesInstrumentationStatus
		return ret
	}
	return *o.InstrumentationStatus
}

// GetInstrumentationStatusOk returns a tuple with the InstrumentationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetInstrumentationStatusOk() (*FleetAgentV2AttributesInstrumentationStatus, bool) {
	if o == nil || o.InstrumentationStatus == nil {
		return nil, false
	}
	return o.InstrumentationStatus, true
}

// HasInstrumentationStatus returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasInstrumentationStatus() bool {
	return o != nil && o.InstrumentationStatus != nil
}

// SetInstrumentationStatus gets a reference to the given FleetAgentV2AttributesInstrumentationStatus and assigns it to the InstrumentationStatus field.
func (o *FleetAgentV2Attributes) SetInstrumentationStatus(v FleetAgentV2AttributesInstrumentationStatus) {
	o.InstrumentationStatus = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetIntegrations() []string {
	if o == nil || o.Integrations == nil {
		var ret []string
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetIntegrationsOk() (*[]string, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return &o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given []string and assigns it to the Integrations field.
func (o *FleetAgentV2Attributes) SetIntegrations(v []string) {
	o.Integrations = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasIpAddresses() bool {
	return o != nil && o.IpAddresses != nil
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *FleetAgentV2Attributes) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetIsSingleStepInstrumentationEnabled returns the IsSingleStepInstrumentationEnabled field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetIsSingleStepInstrumentationEnabled() bool {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSingleStepInstrumentationEnabled
}

// GetIsSingleStepInstrumentationEnabledOk returns a tuple with the IsSingleStepInstrumentationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetIsSingleStepInstrumentationEnabledOk() (*bool, bool) {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		return nil, false
	}
	return o.IsSingleStepInstrumentationEnabled, true
}

// HasIsSingleStepInstrumentationEnabled returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasIsSingleStepInstrumentationEnabled() bool {
	return o != nil && o.IsSingleStepInstrumentationEnabled != nil
}

// SetIsSingleStepInstrumentationEnabled gets a reference to the given bool and assigns it to the IsSingleStepInstrumentationEnabled field.
func (o *FleetAgentV2Attributes) SetIsSingleStepInstrumentationEnabled(v bool) {
	o.IsSingleStepInstrumentationEnabled = &v
}

// GetLastRestartAt returns the LastRestartAt field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetLastRestartAt() int64 {
	if o == nil || o.LastRestartAt == nil {
		var ret int64
		return ret
	}
	return *o.LastRestartAt
}

// GetLastRestartAtOk returns a tuple with the LastRestartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetLastRestartAtOk() (*int64, bool) {
	if o == nil || o.LastRestartAt == nil {
		return nil, false
	}
	return o.LastRestartAt, true
}

// HasLastRestartAt returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasLastRestartAt() bool {
	return o != nil && o.LastRestartAt != nil
}

// SetLastRestartAt gets a reference to the given int64 and assigns it to the LastRestartAt field.
func (o *FleetAgentV2Attributes) SetLastRestartAt(v int64) {
	o.LastRestartAt = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *FleetAgentV2Attributes) SetOs(v string) {
	o.Os = &v
}

// GetOtelCollectorDeploymentTypes returns the OtelCollectorDeploymentTypes field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetOtelCollectorDeploymentTypes() []string {
	if o == nil || o.OtelCollectorDeploymentTypes == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorDeploymentTypes
}

// GetOtelCollectorDeploymentTypesOk returns a tuple with the OtelCollectorDeploymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetOtelCollectorDeploymentTypesOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorDeploymentTypes == nil {
		return nil, false
	}
	return &o.OtelCollectorDeploymentTypes, true
}

// HasOtelCollectorDeploymentTypes returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasOtelCollectorDeploymentTypes() bool {
	return o != nil && o.OtelCollectorDeploymentTypes != nil
}

// SetOtelCollectorDeploymentTypes gets a reference to the given []string and assigns it to the OtelCollectorDeploymentTypes field.
func (o *FleetAgentV2Attributes) SetOtelCollectorDeploymentTypes(v []string) {
	o.OtelCollectorDeploymentTypes = v
}

// GetOtelCollectorDistributions returns the OtelCollectorDistributions field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetOtelCollectorDistributions() []string {
	if o == nil || o.OtelCollectorDistributions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorDistributions
}

// GetOtelCollectorDistributionsOk returns a tuple with the OtelCollectorDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetOtelCollectorDistributionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorDistributions == nil {
		return nil, false
	}
	return &o.OtelCollectorDistributions, true
}

// HasOtelCollectorDistributions returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasOtelCollectorDistributions() bool {
	return o != nil && o.OtelCollectorDistributions != nil
}

// SetOtelCollectorDistributions gets a reference to the given []string and assigns it to the OtelCollectorDistributions field.
func (o *FleetAgentV2Attributes) SetOtelCollectorDistributions(v []string) {
	o.OtelCollectorDistributions = v
}

// GetOtelCollectorVersions returns the OtelCollectorVersions field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetOtelCollectorVersions() []string {
	if o == nil || o.OtelCollectorVersions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorVersions
}

// GetOtelCollectorVersionsOk returns a tuple with the OtelCollectorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetOtelCollectorVersionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorVersions == nil {
		return nil, false
	}
	return &o.OtelCollectorVersions, true
}

// HasOtelCollectorVersions returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasOtelCollectorVersions() bool {
	return o != nil && o.OtelCollectorVersions != nil
}

// SetOtelCollectorVersions gets a reference to the given []string and assigns it to the OtelCollectorVersions field.
func (o *FleetAgentV2Attributes) SetOtelCollectorVersions(v []string) {
	o.OtelCollectorVersions = v
}

// GetOtelResourceAttributes returns the OtelResourceAttributes field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetOtelResourceAttributes() []string {
	if o == nil || o.OtelResourceAttributes == nil {
		var ret []string
		return ret
	}
	return o.OtelResourceAttributes
}

// GetOtelResourceAttributesOk returns a tuple with the OtelResourceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetOtelResourceAttributesOk() (*[]string, bool) {
	if o == nil || o.OtelResourceAttributes == nil {
		return nil, false
	}
	return &o.OtelResourceAttributes, true
}

// HasOtelResourceAttributes returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasOtelResourceAttributes() bool {
	return o != nil && o.OtelResourceAttributes != nil
}

// SetOtelResourceAttributes gets a reference to the given []string and assigns it to the OtelResourceAttributes field.
func (o *FleetAgentV2Attributes) SetOtelResourceAttributes(v []string) {
	o.OtelResourceAttributes = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *FleetAgentV2Attributes) SetPodName(v string) {
	o.PodName = &v
}

// GetRemoteAgentManagement returns the RemoteAgentManagement field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetRemoteAgentManagement() string {
	if o == nil || o.RemoteAgentManagement == nil {
		var ret string
		return ret
	}
	return *o.RemoteAgentManagement
}

// GetRemoteAgentManagementOk returns a tuple with the RemoteAgentManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetRemoteAgentManagementOk() (*string, bool) {
	if o == nil || o.RemoteAgentManagement == nil {
		return nil, false
	}
	return o.RemoteAgentManagement, true
}

// HasRemoteAgentManagement returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasRemoteAgentManagement() bool {
	return o != nil && o.RemoteAgentManagement != nil
}

// SetRemoteAgentManagement gets a reference to the given string and assigns it to the RemoteAgentManagement field.
func (o *FleetAgentV2Attributes) SetRemoteAgentManagement(v string) {
	o.RemoteAgentManagement = &v
}

// GetRemoteConfigStatus returns the RemoteConfigStatus field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetRemoteConfigStatus() string {
	if o == nil || o.RemoteConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigStatus
}

// GetRemoteConfigStatusOk returns a tuple with the RemoteConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetRemoteConfigStatusOk() (*string, bool) {
	if o == nil || o.RemoteConfigStatus == nil {
		return nil, false
	}
	return o.RemoteConfigStatus, true
}

// HasRemoteConfigStatus returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasRemoteConfigStatus() bool {
	return o != nil && o.RemoteConfigStatus != nil
}

// SetRemoteConfigStatus gets a reference to the given string and assigns it to the RemoteConfigStatus field.
func (o *FleetAgentV2Attributes) SetRemoteConfigStatus(v string) {
	o.RemoteConfigStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FleetAgentV2Attributes) SetServices(v []string) {
	o.Services = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetTags() []FleetAgentAttributesTagsItems {
	if o == nil || o.Tags == nil {
		var ret []FleetAgentAttributesTagsItems
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetTagsOk() (*[]FleetAgentAttributesTagsItems, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []FleetAgentAttributesTagsItems and assigns it to the Tags field.
func (o *FleetAgentV2Attributes) SetTags(v []FleetAgentAttributesTagsItems) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *FleetAgentV2Attributes) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentV2Attributes) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *FleetAgentV2Attributes) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *FleetAgentV2Attributes) SetTeam(v string) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	if o.FleetPolicies != nil {
		toSerialize["fleet_policies"] = o.FleetPolicies
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.InstrumentationErrorCounts != nil {
		toSerialize["instrumentation_error_counts"] = o.InstrumentationErrorCounts
	}
	if o.InstrumentationStatus != nil {
		toSerialize["instrumentation_status"] = o.InstrumentationStatus
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
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
	if o.OtelCollectorDeploymentTypes != nil {
		toSerialize["otel_collector_deployment_types"] = o.OtelCollectorDeploymentTypes
	}
	if o.OtelCollectorDistributions != nil {
		toSerialize["otel_collector_distributions"] = o.OtelCollectorDistributions
	}
	if o.OtelCollectorVersions != nil {
		toSerialize["otel_collector_versions"] = o.OtelCollectorVersions
	}
	if o.OtelResourceAttributes != nil {
		toSerialize["otel_resource_attributes"] = o.OtelResourceAttributes
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
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
func (o *FleetAgentV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersion                       *string                                      `json:"agent_version,omitempty"`
		ApiKeyName                         *string                                      `json:"api_key_name,omitempty"`
		ApiKeyUuid                         *string                                      `json:"api_key_uuid,omitempty"`
		CloudProvider                      *string                                      `json:"cloud_provider,omitempty"`
		ClusterName                        *string                                      `json:"cluster_name,omitempty"`
		DatadogDataCenter                  *string                                      `json:"datadog_data_center,omitempty"`
		EcsFargateClusterName              *string                                      `json:"ecs_fargate_cluster_name,omitempty"`
		EcsFargateTaskArn                  *string                                      `json:"ecs_fargate_task_arn,omitempty"`
		EnabledProducts                    []string                                     `json:"enabled_products,omitempty"`
		Env                                []string                                     `json:"env,omitempty"`
		FirstSeenAt                        *int64                                       `json:"first_seen_at,omitempty"`
		FleetPolicies                      []string                                     `json:"fleet_policies,omitempty"`
		Hostname                           *string                                      `json:"hostname,omitempty"`
		InstrumentationErrorCounts         *int64                                       `json:"instrumentation_error_counts,omitempty"`
		InstrumentationStatus              *FleetAgentV2AttributesInstrumentationStatus `json:"instrumentation_status,omitempty"`
		Integrations                       []string                                     `json:"integrations,omitempty"`
		IpAddresses                        []string                                     `json:"ip_addresses,omitempty"`
		IsSingleStepInstrumentationEnabled *bool                                        `json:"is_single_step_instrumentation_enabled,omitempty"`
		LastRestartAt                      *int64                                       `json:"last_restart_at,omitempty"`
		Os                                 *string                                      `json:"os,omitempty"`
		OtelCollectorDeploymentTypes       []string                                     `json:"otel_collector_deployment_types,omitempty"`
		OtelCollectorDistributions         []string                                     `json:"otel_collector_distributions,omitempty"`
		OtelCollectorVersions              []string                                     `json:"otel_collector_versions,omitempty"`
		OtelResourceAttributes             []string                                     `json:"otel_resource_attributes,omitempty"`
		PodName                            *string                                      `json:"pod_name,omitempty"`
		RemoteAgentManagement              *string                                      `json:"remote_agent_management,omitempty"`
		RemoteConfigStatus                 *string                                      `json:"remote_config_status,omitempty"`
		Services                           []string                                     `json:"services,omitempty"`
		Tags                               []FleetAgentAttributesTagsItems              `json:"tags,omitempty"`
		Team                               *string                                      `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_version", "api_key_name", "api_key_uuid", "cloud_provider", "cluster_name", "datadog_data_center", "ecs_fargate_cluster_name", "ecs_fargate_task_arn", "enabled_products", "env", "first_seen_at", "fleet_policies", "hostname", "instrumentation_error_counts", "instrumentation_status", "integrations", "ip_addresses", "is_single_step_instrumentation_enabled", "last_restart_at", "os", "otel_collector_deployment_types", "otel_collector_distributions", "otel_collector_versions", "otel_resource_attributes", "pod_name", "remote_agent_management", "remote_config_status", "services", "tags", "team"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgentVersion = all.AgentVersion
	o.ApiKeyName = all.ApiKeyName
	o.ApiKeyUuid = all.ApiKeyUuid
	o.CloudProvider = all.CloudProvider
	o.ClusterName = all.ClusterName
	o.DatadogDataCenter = all.DatadogDataCenter
	o.EcsFargateClusterName = all.EcsFargateClusterName
	o.EcsFargateTaskArn = all.EcsFargateTaskArn
	o.EnabledProducts = all.EnabledProducts
	o.Env = all.Env
	o.FirstSeenAt = all.FirstSeenAt
	o.FleetPolicies = all.FleetPolicies
	o.Hostname = all.Hostname
	o.InstrumentationErrorCounts = all.InstrumentationErrorCounts
	if all.InstrumentationStatus != nil && !all.InstrumentationStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.InstrumentationStatus = all.InstrumentationStatus
	}
	o.Integrations = all.Integrations
	o.IpAddresses = all.IpAddresses
	o.IsSingleStepInstrumentationEnabled = all.IsSingleStepInstrumentationEnabled
	o.LastRestartAt = all.LastRestartAt
	o.Os = all.Os
	o.OtelCollectorDeploymentTypes = all.OtelCollectorDeploymentTypes
	o.OtelCollectorDistributions = all.OtelCollectorDistributions
	o.OtelCollectorVersions = all.OtelCollectorVersions
	o.OtelResourceAttributes = all.OtelResourceAttributes
	o.PodName = all.PodName
	o.RemoteAgentManagement = all.RemoteAgentManagement
	o.RemoteConfigStatus = all.RemoteConfigStatus
	o.Services = all.Services
	o.Tags = all.Tags
	o.Team = all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
