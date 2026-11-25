// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterMeta Metadata about the retention filter.
type RUMRetentionFilterMeta struct {
	// The types of changes the retention filter accepts. `complete` filters can be updated and deleted.
	// `limited` filters are default retention filters that can't be deleted.
	EditionMode *RUMRetentionFilterEditionMode `json:"edition_mode,omitempty"`
	// The source of the last update to the retention filter. The value is `unknown` if the request doesn't specify a source.
	Source *RUMRetentionFilterMetaSource `json:"source,omitempty"`
	// The time of the last update, as a Unix timestamp in milliseconds.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The handle of the user who last updated the retention filter.
	UpdatedByHandle *string `json:"updated_by_handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRetentionFilterMeta instantiates a new RUMRetentionFilterMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRetentionFilterMeta() *RUMRetentionFilterMeta {
	this := RUMRetentionFilterMeta{}
	return &this
}

// NewRUMRetentionFilterMetaWithDefaults instantiates a new RUMRetentionFilterMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRetentionFilterMetaWithDefaults() *RUMRetentionFilterMeta {
	this := RUMRetentionFilterMeta{}
	return &this
}

// GetEditionMode returns the EditionMode field value if set, zero value otherwise.
func (o *RUMRetentionFilterMeta) GetEditionMode() RUMRetentionFilterEditionMode {
	if o == nil || o.EditionMode == nil {
		var ret RUMRetentionFilterEditionMode
		return ret
	}
	return *o.EditionMode
}

// GetEditionModeOk returns a tuple with the EditionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterMeta) GetEditionModeOk() (*RUMRetentionFilterEditionMode, bool) {
	if o == nil || o.EditionMode == nil {
		return nil, false
	}
	return o.EditionMode, true
}

// HasEditionMode returns a boolean if a field has been set.
func (o *RUMRetentionFilterMeta) HasEditionMode() bool {
	return o != nil && o.EditionMode != nil
}

// SetEditionMode gets a reference to the given RUMRetentionFilterEditionMode and assigns it to the EditionMode field.
func (o *RUMRetentionFilterMeta) SetEditionMode(v RUMRetentionFilterEditionMode) {
	o.EditionMode = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RUMRetentionFilterMeta) GetSource() RUMRetentionFilterMetaSource {
	if o == nil || o.Source == nil {
		var ret RUMRetentionFilterMetaSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterMeta) GetSourceOk() (*RUMRetentionFilterMetaSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RUMRetentionFilterMeta) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given RUMRetentionFilterMetaSource and assigns it to the Source field.
func (o *RUMRetentionFilterMeta) SetSource(v RUMRetentionFilterMetaSource) {
	o.Source = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RUMRetentionFilterMeta) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterMeta) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RUMRetentionFilterMeta) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *RUMRetentionFilterMeta) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value if set, zero value otherwise.
func (o *RUMRetentionFilterMeta) GetUpdatedByHandle() string {
	if o == nil || o.UpdatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterMeta) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil || o.UpdatedByHandle == nil {
		return nil, false
	}
	return o.UpdatedByHandle, true
}

// HasUpdatedByHandle returns a boolean if a field has been set.
func (o *RUMRetentionFilterMeta) HasUpdatedByHandle() bool {
	return o != nil && o.UpdatedByHandle != nil
}

// SetUpdatedByHandle gets a reference to the given string and assigns it to the UpdatedByHandle field.
func (o *RUMRetentionFilterMeta) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRetentionFilterMeta) MarshalJSON() ([]byte, error) {
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
func (o *RUMRetentionFilterMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EditionMode     *RUMRetentionFilterEditionMode `json:"edition_mode,omitempty"`
		Source          *RUMRetentionFilterMetaSource  `json:"source,omitempty"`
		UpdatedAt       *int64                         `json:"updated_at,omitempty"`
		UpdatedByHandle *string                        `json:"updated_by_handle,omitempty"`
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
