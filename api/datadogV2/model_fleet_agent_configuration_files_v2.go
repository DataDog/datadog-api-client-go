// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentConfigurationFilesV2 Configuration details for an agent, organized by configuration layer.
type FleetAgentConfigurationFilesV2 struct {
	// Configuration information organized by layers.
	AgentConfiguration *FleetConfigurationLayer `json:"agent_configuration,omitempty"`
	// Configuration information organized by layers.
	ApplicationMonitoringConfiguration *FleetConfigurationLayer `json:"application_monitoring_configuration,omitempty"`
	// The unique agent key identifier.
	DatadogAgentKey *string `json:"datadog_agent_key,omitempty"`
	// Configuration for OpenTelemetry collectors associated with the agent. Present only when the agent has associated OpenTelemetry collectors.
	OtelCollectorsConfiguration []FleetOtelCollectorConfigurationV2 `json:"otel_collectors_configuration,omitempty"`
	// Configuration information organized by layers.
	SecurityAgentConfiguration *FleetConfigurationLayer `json:"security_agent_configuration,omitempty"`
	// Configuration information organized by layers.
	SystemProbeConfiguration *FleetConfigurationLayer `json:"system_probe_configuration,omitempty"`
	// The configuration version.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentConfigurationFilesV2 instantiates a new FleetAgentConfigurationFilesV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentConfigurationFilesV2() *FleetAgentConfigurationFilesV2 {
	this := FleetAgentConfigurationFilesV2{}
	return &this
}

// NewFleetAgentConfigurationFilesV2WithDefaults instantiates a new FleetAgentConfigurationFilesV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentConfigurationFilesV2WithDefaults() *FleetAgentConfigurationFilesV2 {
	this := FleetAgentConfigurationFilesV2{}
	return &this
}

// GetAgentConfiguration returns the AgentConfiguration field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetAgentConfiguration() FleetConfigurationLayer {
	if o == nil || o.AgentConfiguration == nil {
		var ret FleetConfigurationLayer
		return ret
	}
	return *o.AgentConfiguration
}

// GetAgentConfigurationOk returns a tuple with the AgentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetAgentConfigurationOk() (*FleetConfigurationLayer, bool) {
	if o == nil || o.AgentConfiguration == nil {
		return nil, false
	}
	return o.AgentConfiguration, true
}

// HasAgentConfiguration returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasAgentConfiguration() bool {
	return o != nil && o.AgentConfiguration != nil
}

// SetAgentConfiguration gets a reference to the given FleetConfigurationLayer and assigns it to the AgentConfiguration field.
func (o *FleetAgentConfigurationFilesV2) SetAgentConfiguration(v FleetConfigurationLayer) {
	o.AgentConfiguration = &v
}

// GetApplicationMonitoringConfiguration returns the ApplicationMonitoringConfiguration field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetApplicationMonitoringConfiguration() FleetConfigurationLayer {
	if o == nil || o.ApplicationMonitoringConfiguration == nil {
		var ret FleetConfigurationLayer
		return ret
	}
	return *o.ApplicationMonitoringConfiguration
}

// GetApplicationMonitoringConfigurationOk returns a tuple with the ApplicationMonitoringConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetApplicationMonitoringConfigurationOk() (*FleetConfigurationLayer, bool) {
	if o == nil || o.ApplicationMonitoringConfiguration == nil {
		return nil, false
	}
	return o.ApplicationMonitoringConfiguration, true
}

// HasApplicationMonitoringConfiguration returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasApplicationMonitoringConfiguration() bool {
	return o != nil && o.ApplicationMonitoringConfiguration != nil
}

// SetApplicationMonitoringConfiguration gets a reference to the given FleetConfigurationLayer and assigns it to the ApplicationMonitoringConfiguration field.
func (o *FleetAgentConfigurationFilesV2) SetApplicationMonitoringConfiguration(v FleetConfigurationLayer) {
	o.ApplicationMonitoringConfiguration = &v
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil || o.DatadogAgentKey == nil {
		return nil, false
	}
	return o.DatadogAgentKey, true
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey != nil
}

// SetDatadogAgentKey gets a reference to the given string and assigns it to the DatadogAgentKey field.
func (o *FleetAgentConfigurationFilesV2) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey = &v
}

// GetOtelCollectorsConfiguration returns the OtelCollectorsConfiguration field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetOtelCollectorsConfiguration() []FleetOtelCollectorConfigurationV2 {
	if o == nil || o.OtelCollectorsConfiguration == nil {
		var ret []FleetOtelCollectorConfigurationV2
		return ret
	}
	return o.OtelCollectorsConfiguration
}

// GetOtelCollectorsConfigurationOk returns a tuple with the OtelCollectorsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetOtelCollectorsConfigurationOk() (*[]FleetOtelCollectorConfigurationV2, bool) {
	if o == nil || o.OtelCollectorsConfiguration == nil {
		return nil, false
	}
	return &o.OtelCollectorsConfiguration, true
}

// HasOtelCollectorsConfiguration returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasOtelCollectorsConfiguration() bool {
	return o != nil && o.OtelCollectorsConfiguration != nil
}

// SetOtelCollectorsConfiguration gets a reference to the given []FleetOtelCollectorConfigurationV2 and assigns it to the OtelCollectorsConfiguration field.
func (o *FleetAgentConfigurationFilesV2) SetOtelCollectorsConfiguration(v []FleetOtelCollectorConfigurationV2) {
	o.OtelCollectorsConfiguration = v
}

// GetSecurityAgentConfiguration returns the SecurityAgentConfiguration field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetSecurityAgentConfiguration() FleetConfigurationLayer {
	if o == nil || o.SecurityAgentConfiguration == nil {
		var ret FleetConfigurationLayer
		return ret
	}
	return *o.SecurityAgentConfiguration
}

// GetSecurityAgentConfigurationOk returns a tuple with the SecurityAgentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetSecurityAgentConfigurationOk() (*FleetConfigurationLayer, bool) {
	if o == nil || o.SecurityAgentConfiguration == nil {
		return nil, false
	}
	return o.SecurityAgentConfiguration, true
}

// HasSecurityAgentConfiguration returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasSecurityAgentConfiguration() bool {
	return o != nil && o.SecurityAgentConfiguration != nil
}

// SetSecurityAgentConfiguration gets a reference to the given FleetConfigurationLayer and assigns it to the SecurityAgentConfiguration field.
func (o *FleetAgentConfigurationFilesV2) SetSecurityAgentConfiguration(v FleetConfigurationLayer) {
	o.SecurityAgentConfiguration = &v
}

// GetSystemProbeConfiguration returns the SystemProbeConfiguration field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetSystemProbeConfiguration() FleetConfigurationLayer {
	if o == nil || o.SystemProbeConfiguration == nil {
		var ret FleetConfigurationLayer
		return ret
	}
	return *o.SystemProbeConfiguration
}

// GetSystemProbeConfigurationOk returns a tuple with the SystemProbeConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetSystemProbeConfigurationOk() (*FleetConfigurationLayer, bool) {
	if o == nil || o.SystemProbeConfiguration == nil {
		return nil, false
	}
	return o.SystemProbeConfiguration, true
}

// HasSystemProbeConfiguration returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasSystemProbeConfiguration() bool {
	return o != nil && o.SystemProbeConfiguration != nil
}

// SetSystemProbeConfiguration gets a reference to the given FleetConfigurationLayer and assigns it to the SystemProbeConfiguration field.
func (o *FleetAgentConfigurationFilesV2) SetSystemProbeConfiguration(v FleetConfigurationLayer) {
	o.SystemProbeConfiguration = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FleetAgentConfigurationFilesV2) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentConfigurationFilesV2) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FleetAgentConfigurationFilesV2) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FleetAgentConfigurationFilesV2) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentConfigurationFilesV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentConfiguration != nil {
		toSerialize["agent_configuration"] = o.AgentConfiguration
	}
	if o.ApplicationMonitoringConfiguration != nil {
		toSerialize["application_monitoring_configuration"] = o.ApplicationMonitoringConfiguration
	}
	if o.DatadogAgentKey != nil {
		toSerialize["datadog_agent_key"] = o.DatadogAgentKey
	}
	if o.OtelCollectorsConfiguration != nil {
		toSerialize["otel_collectors_configuration"] = o.OtelCollectorsConfiguration
	}
	if o.SecurityAgentConfiguration != nil {
		toSerialize["security_agent_configuration"] = o.SecurityAgentConfiguration
	}
	if o.SystemProbeConfiguration != nil {
		toSerialize["system_probe_configuration"] = o.SystemProbeConfiguration
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentConfigurationFilesV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentConfiguration                 *FleetConfigurationLayer            `json:"agent_configuration,omitempty"`
		ApplicationMonitoringConfiguration *FleetConfigurationLayer            `json:"application_monitoring_configuration,omitempty"`
		DatadogAgentKey                    *string                             `json:"datadog_agent_key,omitempty"`
		OtelCollectorsConfiguration        []FleetOtelCollectorConfigurationV2 `json:"otel_collectors_configuration,omitempty"`
		SecurityAgentConfiguration         *FleetConfigurationLayer            `json:"security_agent_configuration,omitempty"`
		SystemProbeConfiguration           *FleetConfigurationLayer            `json:"system_probe_configuration,omitempty"`
		Version                            *string                             `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_configuration", "application_monitoring_configuration", "datadog_agent_key", "otel_collectors_configuration", "security_agent_configuration", "system_probe_configuration", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AgentConfiguration != nil && all.AgentConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AgentConfiguration = all.AgentConfiguration
	if all.ApplicationMonitoringConfiguration != nil && all.ApplicationMonitoringConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApplicationMonitoringConfiguration = all.ApplicationMonitoringConfiguration
	o.DatadogAgentKey = all.DatadogAgentKey
	o.OtelCollectorsConfiguration = all.OtelCollectorsConfiguration
	if all.SecurityAgentConfiguration != nil && all.SecurityAgentConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SecurityAgentConfiguration = all.SecurityAgentConfiguration
	if all.SystemProbeConfiguration != nil && all.SystemProbeConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SystemProbeConfiguration = all.SystemProbeConfiguration
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
