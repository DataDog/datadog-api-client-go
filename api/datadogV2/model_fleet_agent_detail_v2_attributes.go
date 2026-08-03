// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentDetailV2Attributes Attributes for the v2 agent detail response.
type FleetAgentDetailV2Attributes struct {
	// Detailed information about a Datadog Agent.
	AgentInfos FleetAgentInfoDetailsV2 `json:"agent_infos"`
	// Configuration details for an agent, organized by configuration layer.
	ConfigurationFiles *FleetAgentConfigurationFilesV2 `json:"configuration_files,omitempty"`
	// Integrations organized by their status.
	Integrations *FleetIntegrationsByStatusV2 `json:"integrations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentDetailV2Attributes instantiates a new FleetAgentDetailV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentDetailV2Attributes(agentInfos FleetAgentInfoDetailsV2) *FleetAgentDetailV2Attributes {
	this := FleetAgentDetailV2Attributes{}
	this.AgentInfos = agentInfos
	return &this
}

// NewFleetAgentDetailV2AttributesWithDefaults instantiates a new FleetAgentDetailV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentDetailV2AttributesWithDefaults() *FleetAgentDetailV2Attributes {
	this := FleetAgentDetailV2Attributes{}
	return &this
}

// GetAgentInfos returns the AgentInfos field value.
func (o *FleetAgentDetailV2Attributes) GetAgentInfos() FleetAgentInfoDetailsV2 {
	if o == nil {
		var ret FleetAgentInfoDetailsV2
		return ret
	}
	return o.AgentInfos
}

// GetAgentInfosOk returns a tuple with the AgentInfos field value
// and a boolean to check if the value has been set.
func (o *FleetAgentDetailV2Attributes) GetAgentInfosOk() (*FleetAgentInfoDetailsV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentInfos, true
}

// SetAgentInfos sets field value.
func (o *FleetAgentDetailV2Attributes) SetAgentInfos(v FleetAgentInfoDetailsV2) {
	o.AgentInfos = v
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise.
func (o *FleetAgentDetailV2Attributes) GetConfigurationFiles() FleetAgentConfigurationFilesV2 {
	if o == nil || o.ConfigurationFiles == nil {
		var ret FleetAgentConfigurationFilesV2
		return ret
	}
	return *o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentDetailV2Attributes) GetConfigurationFilesOk() (*FleetAgentConfigurationFilesV2, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *FleetAgentDetailV2Attributes) HasConfigurationFiles() bool {
	return o != nil && o.ConfigurationFiles != nil
}

// SetConfigurationFiles gets a reference to the given FleetAgentConfigurationFilesV2 and assigns it to the ConfigurationFiles field.
func (o *FleetAgentDetailV2Attributes) SetConfigurationFiles(v FleetAgentConfigurationFilesV2) {
	o.ConfigurationFiles = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *FleetAgentDetailV2Attributes) GetIntegrations() FleetIntegrationsByStatusV2 {
	if o == nil || o.Integrations == nil {
		var ret FleetIntegrationsByStatusV2
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentDetailV2Attributes) GetIntegrationsOk() (*FleetIntegrationsByStatusV2, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *FleetAgentDetailV2Attributes) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given FleetIntegrationsByStatusV2 and assigns it to the Integrations field.
func (o *FleetAgentDetailV2Attributes) SetIntegrations(v FleetIntegrationsByStatusV2) {
	o.Integrations = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentDetailV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["agent_infos"] = o.AgentInfos
	if o.ConfigurationFiles != nil {
		toSerialize["configuration_files"] = o.ConfigurationFiles
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentDetailV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentInfos         *FleetAgentInfoDetailsV2        `json:"agent_infos"`
		ConfigurationFiles *FleetAgentConfigurationFilesV2 `json:"configuration_files,omitempty"`
		Integrations       *FleetIntegrationsByStatusV2    `json:"integrations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AgentInfos == nil {
		return fmt.Errorf("required field agent_infos missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_infos", "configuration_files", "integrations"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AgentInfos.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AgentInfos = *all.AgentInfos
	if all.ConfigurationFiles != nil && all.ConfigurationFiles.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigurationFiles = all.ConfigurationFiles
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
