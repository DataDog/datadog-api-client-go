// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyConfigDataAttributesRegionsConfig Region configuration for an OCI tenancy, indicating which regions are available, enabled, or disabled for data collection.
type TenancyConfigDataAttributesRegionsConfig struct {
	// List of OCI regions available for data collection in the tenancy. Returns `null` when no regions are set.
	Available datadog.NullableList[string] `json:"available,omitempty"`
	// List of OCI regions explicitly disabled for data collection. Returns `null` when no regions are set.
	Disabled datadog.NullableList[string] `json:"disabled,omitempty"`
	// List of OCI regions enabled for data collection. Returns `null` when no regions are set.
	Enabled datadog.NullableList[string] `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyConfigDataAttributesRegionsConfig instantiates a new TenancyConfigDataAttributesRegionsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyConfigDataAttributesRegionsConfig() *TenancyConfigDataAttributesRegionsConfig {
	this := TenancyConfigDataAttributesRegionsConfig{}
	return &this
}

// NewTenancyConfigDataAttributesRegionsConfigWithDefaults instantiates a new TenancyConfigDataAttributesRegionsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyConfigDataAttributesRegionsConfigWithDefaults() *TenancyConfigDataAttributesRegionsConfig {
	this := TenancyConfigDataAttributesRegionsConfig{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenancyConfigDataAttributesRegionsConfig) GetAvailable() []string {
	if o == nil || o.Available.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Available.Get()
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TenancyConfigDataAttributesRegionsConfig) GetAvailableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Available.Get(), o.Available.IsSet()
}

// HasAvailable returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesRegionsConfig) HasAvailable() bool {
	return o != nil && o.Available.IsSet()
}

// SetAvailable gets a reference to the given datadog.NullableList[string] and assigns it to the Available field.
func (o *TenancyConfigDataAttributesRegionsConfig) SetAvailable(v []string) {
	o.Available.Set(&v)
}

// SetAvailableNil sets the value for Available to be an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) SetAvailableNil() {
	o.Available.Set(nil)
}

// UnsetAvailable ensures that no value is present for Available, not even an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) UnsetAvailable() {
	o.Available.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenancyConfigDataAttributesRegionsConfig) GetDisabled() []string {
	if o == nil || o.Disabled.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TenancyConfigDataAttributesRegionsConfig) GetDisabledOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesRegionsConfig) HasDisabled() bool {
	return o != nil && o.Disabled.IsSet()
}

// SetDisabled gets a reference to the given datadog.NullableList[string] and assigns it to the Disabled field.
func (o *TenancyConfigDataAttributesRegionsConfig) SetDisabled(v []string) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenancyConfigDataAttributesRegionsConfig) GetEnabled() []string {
	if o == nil || o.Enabled.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TenancyConfigDataAttributesRegionsConfig) GetEnabledOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesRegionsConfig) HasEnabled() bool {
	return o != nil && o.Enabled.IsSet()
}

// SetEnabled gets a reference to the given datadog.NullableList[string] and assigns it to the Enabled field.
func (o *TenancyConfigDataAttributesRegionsConfig) SetEnabled(v []string) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil.
func (o *TenancyConfigDataAttributesRegionsConfig) UnsetEnabled() {
	o.Enabled.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyConfigDataAttributesRegionsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Available.IsSet() {
		toSerialize["available"] = o.Available.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TenancyConfigDataAttributesRegionsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Available datadog.NullableList[string] `json:"available,omitempty"`
		Disabled  datadog.NullableList[string] `json:"disabled,omitempty"`
		Enabled   datadog.NullableList[string] `json:"enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available", "disabled", "enabled"})
	} else {
		return err
	}
	o.Available = all.Available
	o.Disabled = all.Disabled
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
