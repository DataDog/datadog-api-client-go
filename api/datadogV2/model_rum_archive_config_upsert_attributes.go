// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveConfigUpsertAttributes Archive configuration to save. When enabled is true, retention_days is required and must be 30, 90, or 180. When disabled, retention_days can be omitted.
type RumArchiveConfigUpsertAttributes struct {
	// Whether archiving is enabled.
	Enabled bool `json:"enabled"`
	// Retention in days. When enabled is true, required and restricted to 30, 90, or 180. When enabled is false, this field can be omitted.
	RetentionDays *int32 `json:"retention_days,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumArchiveConfigUpsertAttributes instantiates a new RumArchiveConfigUpsertAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumArchiveConfigUpsertAttributes(enabled bool) *RumArchiveConfigUpsertAttributes {
	this := RumArchiveConfigUpsertAttributes{}
	this.Enabled = enabled
	return &this
}

// NewRumArchiveConfigUpsertAttributesWithDefaults instantiates a new RumArchiveConfigUpsertAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumArchiveConfigUpsertAttributesWithDefaults() *RumArchiveConfigUpsertAttributes {
	this := RumArchiveConfigUpsertAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *RumArchiveConfigUpsertAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigUpsertAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RumArchiveConfigUpsertAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRetentionDays returns the RetentionDays field value if set, zero value otherwise.
func (o *RumArchiveConfigUpsertAttributes) GetRetentionDays() int32 {
	if o == nil || o.RetentionDays == nil {
		var ret int32
		return ret
	}
	return *o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigUpsertAttributes) GetRetentionDaysOk() (*int32, bool) {
	if o == nil || o.RetentionDays == nil {
		return nil, false
	}
	return o.RetentionDays, true
}

// HasRetentionDays returns a boolean if a field has been set.
func (o *RumArchiveConfigUpsertAttributes) HasRetentionDays() bool {
	return o != nil && o.RetentionDays != nil
}

// SetRetentionDays gets a reference to the given int32 and assigns it to the RetentionDays field.
func (o *RumArchiveConfigUpsertAttributes) SetRetentionDays(v int32) {
	o.RetentionDays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumArchiveConfigUpsertAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.RetentionDays != nil {
		toSerialize["retention_days"] = o.RetentionDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumArchiveConfigUpsertAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool  `json:"enabled"`
		RetentionDays *int32 `json:"retention_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "retention_days"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.RetentionDays = all.RetentionDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
