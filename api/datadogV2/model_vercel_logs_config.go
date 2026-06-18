// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelLogsConfig Logs forwarding configuration for the Vercel integration.
type VercelLogsConfig struct {
	// Whether logs forwarding is enabled.
	Enabled bool `json:"enabled"`
	// List of Vercel deployment environments to forward telemetry from.
	Environments []VercelEnvironment `json:"environments"`
	// List of Vercel log sources to forward to Datadog.
	LogSources []VercelLogSource `json:"logSources"`
	// Percentage of logs to forward to Datadog, between 0 and 100.
	SamplingPercentage int32 `json:"samplingPercentage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVercelLogsConfig instantiates a new VercelLogsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVercelLogsConfig(enabled bool, environments []VercelEnvironment, logSources []VercelLogSource, samplingPercentage int32) *VercelLogsConfig {
	this := VercelLogsConfig{}
	this.Enabled = enabled
	this.Environments = environments
	this.LogSources = logSources
	this.SamplingPercentage = samplingPercentage
	return &this
}

// NewVercelLogsConfigWithDefaults instantiates a new VercelLogsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVercelLogsConfigWithDefaults() *VercelLogsConfig {
	this := VercelLogsConfig{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *VercelLogsConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *VercelLogsConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *VercelLogsConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironments returns the Environments field value.
func (o *VercelLogsConfig) GetEnvironments() []VercelEnvironment {
	if o == nil {
		var ret []VercelEnvironment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *VercelLogsConfig) GetEnvironmentsOk() (*[]VercelEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value.
func (o *VercelLogsConfig) SetEnvironments(v []VercelEnvironment) {
	o.Environments = v
}

// GetLogSources returns the LogSources field value.
func (o *VercelLogsConfig) GetLogSources() []VercelLogSource {
	if o == nil {
		var ret []VercelLogSource
		return ret
	}
	return o.LogSources
}

// GetLogSourcesOk returns a tuple with the LogSources field value
// and a boolean to check if the value has been set.
func (o *VercelLogsConfig) GetLogSourcesOk() (*[]VercelLogSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogSources, true
}

// SetLogSources sets field value.
func (o *VercelLogsConfig) SetLogSources(v []VercelLogSource) {
	o.LogSources = v
}

// GetSamplingPercentage returns the SamplingPercentage field value.
func (o *VercelLogsConfig) GetSamplingPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SamplingPercentage
}

// GetSamplingPercentageOk returns a tuple with the SamplingPercentage field value
// and a boolean to check if the value has been set.
func (o *VercelLogsConfig) GetSamplingPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamplingPercentage, true
}

// SetSamplingPercentage sets field value.
func (o *VercelLogsConfig) SetSamplingPercentage(v int32) {
	o.SamplingPercentage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VercelLogsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["environments"] = o.Environments
	toSerialize["logSources"] = o.LogSources
	toSerialize["samplingPercentage"] = o.SamplingPercentage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VercelLogsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled            *bool                `json:"enabled"`
		Environments       *[]VercelEnvironment `json:"environments"`
		LogSources         *[]VercelLogSource   `json:"logSources"`
		SamplingPercentage *int32               `json:"samplingPercentage"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Environments == nil {
		return fmt.Errorf("required field environments missing")
	}
	if all.LogSources == nil {
		return fmt.Errorf("required field logSources missing")
	}
	if all.SamplingPercentage == nil {
		return fmt.Errorf("required field samplingPercentage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "environments", "logSources", "samplingPercentage"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Environments = *all.Environments
	o.LogSources = *all.LogSources
	o.SamplingPercentage = *all.SamplingPercentage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
