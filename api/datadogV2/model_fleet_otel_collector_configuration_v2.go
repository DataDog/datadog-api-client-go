// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetOtelCollectorConfigurationV2 Configuration for a single OpenTelemetry collector associated with the agent.
type FleetOtelCollectorConfigurationV2 struct {
	// The unique identifier of the OpenTelemetry collector.
	CollectorId *string `json:"collector_id,omitempty"`
	// The final compiled configuration of the OpenTelemetry collector.
	CompiledConfiguration *string `json:"compiled_configuration,omitempty"`
	// The distribution of the OpenTelemetry collector.
	Distribution *string `json:"distribution,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetOtelCollectorConfigurationV2 instantiates a new FleetOtelCollectorConfigurationV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetOtelCollectorConfigurationV2() *FleetOtelCollectorConfigurationV2 {
	this := FleetOtelCollectorConfigurationV2{}
	return &this
}

// NewFleetOtelCollectorConfigurationV2WithDefaults instantiates a new FleetOtelCollectorConfigurationV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetOtelCollectorConfigurationV2WithDefaults() *FleetOtelCollectorConfigurationV2 {
	this := FleetOtelCollectorConfigurationV2{}
	return &this
}

// GetCollectorId returns the CollectorId field value if set, zero value otherwise.
func (o *FleetOtelCollectorConfigurationV2) GetCollectorId() string {
	if o == nil || o.CollectorId == nil {
		var ret string
		return ret
	}
	return *o.CollectorId
}

// GetCollectorIdOk returns a tuple with the CollectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetOtelCollectorConfigurationV2) GetCollectorIdOk() (*string, bool) {
	if o == nil || o.CollectorId == nil {
		return nil, false
	}
	return o.CollectorId, true
}

// HasCollectorId returns a boolean if a field has been set.
func (o *FleetOtelCollectorConfigurationV2) HasCollectorId() bool {
	return o != nil && o.CollectorId != nil
}

// SetCollectorId gets a reference to the given string and assigns it to the CollectorId field.
func (o *FleetOtelCollectorConfigurationV2) SetCollectorId(v string) {
	o.CollectorId = &v
}

// GetCompiledConfiguration returns the CompiledConfiguration field value if set, zero value otherwise.
func (o *FleetOtelCollectorConfigurationV2) GetCompiledConfiguration() string {
	if o == nil || o.CompiledConfiguration == nil {
		var ret string
		return ret
	}
	return *o.CompiledConfiguration
}

// GetCompiledConfigurationOk returns a tuple with the CompiledConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetOtelCollectorConfigurationV2) GetCompiledConfigurationOk() (*string, bool) {
	if o == nil || o.CompiledConfiguration == nil {
		return nil, false
	}
	return o.CompiledConfiguration, true
}

// HasCompiledConfiguration returns a boolean if a field has been set.
func (o *FleetOtelCollectorConfigurationV2) HasCompiledConfiguration() bool {
	return o != nil && o.CompiledConfiguration != nil
}

// SetCompiledConfiguration gets a reference to the given string and assigns it to the CompiledConfiguration field.
func (o *FleetOtelCollectorConfigurationV2) SetCompiledConfiguration(v string) {
	o.CompiledConfiguration = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *FleetOtelCollectorConfigurationV2) GetDistribution() string {
	if o == nil || o.Distribution == nil {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetOtelCollectorConfigurationV2) GetDistributionOk() (*string, bool) {
	if o == nil || o.Distribution == nil {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *FleetOtelCollectorConfigurationV2) HasDistribution() bool {
	return o != nil && o.Distribution != nil
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *FleetOtelCollectorConfigurationV2) SetDistribution(v string) {
	o.Distribution = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetOtelCollectorConfigurationV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CollectorId != nil {
		toSerialize["collector_id"] = o.CollectorId
	}
	if o.CompiledConfiguration != nil {
		toSerialize["compiled_configuration"] = o.CompiledConfiguration
	}
	if o.Distribution != nil {
		toSerialize["distribution"] = o.Distribution
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetOtelCollectorConfigurationV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectorId           *string `json:"collector_id,omitempty"`
		CompiledConfiguration *string `json:"compiled_configuration,omitempty"`
		Distribution          *string `json:"distribution,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"collector_id", "compiled_configuration", "distribution"})
	} else {
		return err
	}
	o.CollectorId = all.CollectorId
	o.CompiledConfiguration = all.CompiledConfiguration
	o.Distribution = all.Distribution

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
