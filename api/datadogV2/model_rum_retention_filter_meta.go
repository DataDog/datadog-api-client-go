// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterMeta Metadata about a RUM retention filter.
type RumRetentionFilterMeta struct {
	// Which attributes of the retention filter can be edited. `complete`: all attributes can be edited. `limited`: only the sample rates can be edited; the name, event type, and query are fixed and the filter cannot be reordered (for example, default filters). `unknown`: the edition mode is not set.
	EditionMode *RumRetentionFilterMetaEditionMode `json:"edition_mode,omitempty"`
	// The source of the last change to a RUM retention filter.
	Source *RumRetentionFilterMetaSource `json:"source,omitempty"`
	// The timestamp of the last update to the retention filter, in milliseconds since the Unix epoch.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The handle of the user who last updated the retention filter.
	UpdatedByHandle *string `json:"updated_by_handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionFilterMeta instantiates a new RumRetentionFilterMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionFilterMeta() *RumRetentionFilterMeta {
	this := RumRetentionFilterMeta{}
	return &this
}

// NewRumRetentionFilterMetaWithDefaults instantiates a new RumRetentionFilterMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionFilterMetaWithDefaults() *RumRetentionFilterMeta {
	this := RumRetentionFilterMeta{}
	return &this
}

// GetEditionMode returns the EditionMode field value if set, zero value otherwise.
func (o *RumRetentionFilterMeta) GetEditionMode() RumRetentionFilterMetaEditionMode {
	if o == nil || o.EditionMode == nil {
		var ret RumRetentionFilterMetaEditionMode
		return ret
	}
	return *o.EditionMode
}

// GetEditionModeOk returns a tuple with the EditionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterMeta) GetEditionModeOk() (*RumRetentionFilterMetaEditionMode, bool) {
	if o == nil || o.EditionMode == nil {
		return nil, false
	}
	return o.EditionMode, true
}

// HasEditionMode returns a boolean if a field has been set.
func (o *RumRetentionFilterMeta) HasEditionMode() bool {
	return o != nil && o.EditionMode != nil
}

// SetEditionMode gets a reference to the given RumRetentionFilterMetaEditionMode and assigns it to the EditionMode field.
func (o *RumRetentionFilterMeta) SetEditionMode(v RumRetentionFilterMetaEditionMode) {
	o.EditionMode = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RumRetentionFilterMeta) GetSource() RumRetentionFilterMetaSource {
	if o == nil || o.Source == nil {
		var ret RumRetentionFilterMetaSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterMeta) GetSourceOk() (*RumRetentionFilterMetaSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RumRetentionFilterMeta) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given RumRetentionFilterMetaSource and assigns it to the Source field.
func (o *RumRetentionFilterMeta) SetSource(v RumRetentionFilterMetaSource) {
	o.Source = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RumRetentionFilterMeta) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterMeta) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RumRetentionFilterMeta) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *RumRetentionFilterMeta) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value if set, zero value otherwise.
func (o *RumRetentionFilterMeta) GetUpdatedByHandle() string {
	if o == nil || o.UpdatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterMeta) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil || o.UpdatedByHandle == nil {
		return nil, false
	}
	return o.UpdatedByHandle, true
}

// HasUpdatedByHandle returns a boolean if a field has been set.
func (o *RumRetentionFilterMeta) HasUpdatedByHandle() bool {
	return o != nil && o.UpdatedByHandle != nil
}

// SetUpdatedByHandle gets a reference to the given string and assigns it to the UpdatedByHandle field.
func (o *RumRetentionFilterMeta) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionFilterMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EditionMode != nil {
		toSerialize["edition_mode"] = o.EditionMode
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
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
func (o *RumRetentionFilterMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EditionMode     *RumRetentionFilterMetaEditionMode `json:"edition_mode,omitempty"`
		Source          *RumRetentionFilterMetaSource      `json:"source,omitempty"`
		UpdatedAt       *int64                             `json:"updated_at,omitempty"`
		UpdatedByHandle *string                            `json:"updated_by_handle,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"edition_mode", "source", "updated_at", "updated_by_handle"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EditionMode != nil && !all.EditionMode.IsValid() {
		hasInvalidField = true
	} else {
		o.EditionMode = all.EditionMode
	}
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedByHandle = all.UpdatedByHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
