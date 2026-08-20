// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaConfigUpdateAttributes The RUM retention quota configuration properties to create or update.
type RumRetentionQuotaConfigUpdateAttributes struct {
	// The configuration used when `mode` is `adaptive`.
	Adaptive *RumRetentionQuotaAdaptiveConfig `json:"adaptive,omitempty"`
	// The configuration used when `mode` is `custom`.
	Custom *RumRetentionQuotaCustomConfig `json:"custom,omitempty"`
	// The retention quota mode. `custom` enforces a fixed session limit, while
	// `adaptive` dynamically adjusts retention.
	Mode RumRetentionQuotaMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionQuotaConfigUpdateAttributes instantiates a new RumRetentionQuotaConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionQuotaConfigUpdateAttributes(mode RumRetentionQuotaMode) *RumRetentionQuotaConfigUpdateAttributes {
	this := RumRetentionQuotaConfigUpdateAttributes{}
	this.Mode = mode
	return &this
}

// NewRumRetentionQuotaConfigUpdateAttributesWithDefaults instantiates a new RumRetentionQuotaConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionQuotaConfigUpdateAttributesWithDefaults() *RumRetentionQuotaConfigUpdateAttributes {
	this := RumRetentionQuotaConfigUpdateAttributes{}
	return &this
}

// GetAdaptive returns the Adaptive field value if set, zero value otherwise.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetAdaptive() RumRetentionQuotaAdaptiveConfig {
	if o == nil || o.Adaptive == nil {
		var ret RumRetentionQuotaAdaptiveConfig
		return ret
	}
	return *o.Adaptive
}

// GetAdaptiveOk returns a tuple with the Adaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetAdaptiveOk() (*RumRetentionQuotaAdaptiveConfig, bool) {
	if o == nil || o.Adaptive == nil {
		return nil, false
	}
	return o.Adaptive, true
}

// HasAdaptive returns a boolean if a field has been set.
func (o *RumRetentionQuotaConfigUpdateAttributes) HasAdaptive() bool {
	return o != nil && o.Adaptive != nil
}

// SetAdaptive gets a reference to the given RumRetentionQuotaAdaptiveConfig and assigns it to the Adaptive field.
func (o *RumRetentionQuotaConfigUpdateAttributes) SetAdaptive(v RumRetentionQuotaAdaptiveConfig) {
	o.Adaptive = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetCustom() RumRetentionQuotaCustomConfig {
	if o == nil || o.Custom == nil {
		var ret RumRetentionQuotaCustomConfig
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetCustomOk() (*RumRetentionQuotaCustomConfig, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RumRetentionQuotaConfigUpdateAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given RumRetentionQuotaCustomConfig and assigns it to the Custom field.
func (o *RumRetentionQuotaConfigUpdateAttributes) SetCustom(v RumRetentionQuotaCustomConfig) {
	o.Custom = &v
}

// GetMode returns the Mode field value.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetMode() RumRetentionQuotaMode {
	if o == nil {
		var ret RumRetentionQuotaMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaConfigUpdateAttributes) GetModeOk() (*RumRetentionQuotaMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RumRetentionQuotaConfigUpdateAttributes) SetMode(v RumRetentionQuotaMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionQuotaConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Adaptive != nil {
		toSerialize["adaptive"] = o.Adaptive
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRetentionQuotaConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Adaptive *RumRetentionQuotaAdaptiveConfig `json:"adaptive,omitempty"`
		Custom   *RumRetentionQuotaCustomConfig   `json:"custom,omitempty"`
		Mode     *RumRetentionQuotaMode           `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"adaptive", "custom", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Adaptive != nil && all.Adaptive.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Adaptive = all.Adaptive
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
