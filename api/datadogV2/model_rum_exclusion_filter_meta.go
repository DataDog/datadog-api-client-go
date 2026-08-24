// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumExclusionFilterMeta Metadata about the exclusion filter.
type RumExclusionFilterMeta struct {
	// Unix epoch (in milliseconds) when the exclusion filter was last enabled.
	EnabledAt *int64 `json:"enabled_at,omitempty"`
	// Unix epoch (in milliseconds) of the last update.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Handle of the user who last updated the exclusion filter.
	UpdatedByHandle *string `json:"updated_by_handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumExclusionFilterMeta instantiates a new RumExclusionFilterMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumExclusionFilterMeta() *RumExclusionFilterMeta {
	this := RumExclusionFilterMeta{}
	return &this
}

// NewRumExclusionFilterMetaWithDefaults instantiates a new RumExclusionFilterMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumExclusionFilterMetaWithDefaults() *RumExclusionFilterMeta {
	this := RumExclusionFilterMeta{}
	return &this
}

// GetEnabledAt returns the EnabledAt field value if set, zero value otherwise.
func (o *RumExclusionFilterMeta) GetEnabledAt() int64 {
	if o == nil || o.EnabledAt == nil {
		var ret int64
		return ret
	}
	return *o.EnabledAt
}

// GetEnabledAtOk returns a tuple with the EnabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumExclusionFilterMeta) GetEnabledAtOk() (*int64, bool) {
	if o == nil || o.EnabledAt == nil {
		return nil, false
	}
	return o.EnabledAt, true
}

// HasEnabledAt returns a boolean if a field has been set.
func (o *RumExclusionFilterMeta) HasEnabledAt() bool {
	return o != nil && o.EnabledAt != nil
}

// SetEnabledAt gets a reference to the given int64 and assigns it to the EnabledAt field.
func (o *RumExclusionFilterMeta) SetEnabledAt(v int64) {
	o.EnabledAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RumExclusionFilterMeta) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumExclusionFilterMeta) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RumExclusionFilterMeta) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *RumExclusionFilterMeta) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value if set, zero value otherwise.
func (o *RumExclusionFilterMeta) GetUpdatedByHandle() string {
	if o == nil || o.UpdatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumExclusionFilterMeta) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil || o.UpdatedByHandle == nil {
		return nil, false
	}
	return o.UpdatedByHandle, true
}

// HasUpdatedByHandle returns a boolean if a field has been set.
func (o *RumExclusionFilterMeta) HasUpdatedByHandle() bool {
	return o != nil && o.UpdatedByHandle != nil
}

// SetUpdatedByHandle gets a reference to the given string and assigns it to the UpdatedByHandle field.
func (o *RumExclusionFilterMeta) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumExclusionFilterMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnabledAt != nil {
		toSerialize["enabled_at"] = o.EnabledAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedByHandle != nil {
		toSerialize["updated_by_handle"] = o.UpdatedByHandle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumExclusionFilterMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnabledAt       *int64  `json:"enabled_at,omitempty"`
		UpdatedAt       *int64  `json:"updated_at,omitempty"`
		UpdatedByHandle *string `json:"updated_by_handle,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled_at", "updated_at", "updated_by_handle"})
	} else {
		return err
	}
	o.EnabledAt = all.EnabledAt
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedByHandle = all.UpdatedByHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
