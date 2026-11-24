// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaConfigUpdateAttributes The RUM retention quota configuration properties to create or update.
type RUMRetentionQuotaConfigUpdateAttributes struct {
	// The configuration used when `mode` is `custom`.
	Custom *RUMRetentionQuotaCustomConfig `json:"custom,omitempty"`
	// The retention quota mode. `custom` enforces a fixed session limit.
	// `custom` is the only supported mode.
	Mode RUMRetentionQuotaMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRetentionQuotaConfigUpdateAttributes instantiates a new RUMRetentionQuotaConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRetentionQuotaConfigUpdateAttributes(mode RUMRetentionQuotaMode) *RUMRetentionQuotaConfigUpdateAttributes {
	this := RUMRetentionQuotaConfigUpdateAttributes{}
	this.Mode = mode
	return &this
}

// NewRUMRetentionQuotaConfigUpdateAttributesWithDefaults instantiates a new RUMRetentionQuotaConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRetentionQuotaConfigUpdateAttributesWithDefaults() *RUMRetentionQuotaConfigUpdateAttributes {
	this := RUMRetentionQuotaConfigUpdateAttributes{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RUMRetentionQuotaConfigUpdateAttributes) GetCustom() RUMRetentionQuotaCustomConfig {
	if o == nil || o.Custom == nil {
		var ret RUMRetentionQuotaCustomConfig
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigUpdateAttributes) GetCustomOk() (*RUMRetentionQuotaCustomConfig, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RUMRetentionQuotaConfigUpdateAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given RUMRetentionQuotaCustomConfig and assigns it to the Custom field.
func (o *RUMRetentionQuotaConfigUpdateAttributes) SetCustom(v RUMRetentionQuotaCustomConfig) {
	o.Custom = &v
}

// GetMode returns the Mode field value.
func (o *RUMRetentionQuotaConfigUpdateAttributes) GetMode() RUMRetentionQuotaMode {
	if o == nil {
		var ret RUMRetentionQuotaMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigUpdateAttributes) GetModeOk() (*RUMRetentionQuotaMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RUMRetentionQuotaConfigUpdateAttributes) SetMode(v RUMRetentionQuotaMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRetentionQuotaConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *RUMRetentionQuotaConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Custom *RUMRetentionQuotaCustomConfig `json:"custom,omitempty"`
		Mode   *RUMRetentionQuotaMode         `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
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
