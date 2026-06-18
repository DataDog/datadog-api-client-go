// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelTraceConfig Tracing configuration for the Vercel integration.
type VercelTraceConfig struct {
	// Whether tracing is enabled.
	Enabled bool `json:"enabled"`
	// Whether the configuration uses the deprecated OpenTelemetry tracing setup.
	IsDeprecatedOtel bool `json:"isDeprecatedOtel"`
	// Percentage of traces to forward to Datadog, between 0 and 100.
	SamplingPercentage int32 `json:"samplingPercentage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVercelTraceConfig instantiates a new VercelTraceConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVercelTraceConfig(enabled bool, isDeprecatedOtel bool, samplingPercentage int32) *VercelTraceConfig {
	this := VercelTraceConfig{}
	this.Enabled = enabled
	this.IsDeprecatedOtel = isDeprecatedOtel
	this.SamplingPercentage = samplingPercentage
	return &this
}

// NewVercelTraceConfigWithDefaults instantiates a new VercelTraceConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVercelTraceConfigWithDefaults() *VercelTraceConfig {
	this := VercelTraceConfig{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *VercelTraceConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *VercelTraceConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *VercelTraceConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIsDeprecatedOtel returns the IsDeprecatedOtel field value.
func (o *VercelTraceConfig) GetIsDeprecatedOtel() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDeprecatedOtel
}

// GetIsDeprecatedOtelOk returns a tuple with the IsDeprecatedOtel field value
// and a boolean to check if the value has been set.
func (o *VercelTraceConfig) GetIsDeprecatedOtelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecatedOtel, true
}

// SetIsDeprecatedOtel sets field value.
func (o *VercelTraceConfig) SetIsDeprecatedOtel(v bool) {
	o.IsDeprecatedOtel = v
}

// GetSamplingPercentage returns the SamplingPercentage field value.
func (o *VercelTraceConfig) GetSamplingPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SamplingPercentage
}

// GetSamplingPercentageOk returns a tuple with the SamplingPercentage field value
// and a boolean to check if the value has been set.
func (o *VercelTraceConfig) GetSamplingPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamplingPercentage, true
}

// SetSamplingPercentage sets field value.
func (o *VercelTraceConfig) SetSamplingPercentage(v int32) {
	o.SamplingPercentage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VercelTraceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["isDeprecatedOtel"] = o.IsDeprecatedOtel
	toSerialize["samplingPercentage"] = o.SamplingPercentage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VercelTraceConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled            *bool  `json:"enabled"`
		IsDeprecatedOtel   *bool  `json:"isDeprecatedOtel"`
		SamplingPercentage *int32 `json:"samplingPercentage"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.IsDeprecatedOtel == nil {
		return fmt.Errorf("required field isDeprecatedOtel missing")
	}
	if all.SamplingPercentage == nil {
		return fmt.Errorf("required field samplingPercentage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "isDeprecatedOtel", "samplingPercentage"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.IsDeprecatedOtel = *all.IsDeprecatedOtel
	o.SamplingPercentage = *all.SamplingPercentage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
