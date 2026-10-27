// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchListResponseMeta Metadata of the Archive Searches response.
type ArchiveSearchListResponseMeta struct {
	// Whether more Archive Searches exist beyond the ones returned.
	Truncated *bool `json:"truncated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchListResponseMeta instantiates a new ArchiveSearchListResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchListResponseMeta() *ArchiveSearchListResponseMeta {
	this := ArchiveSearchListResponseMeta{}
	return &this
}

// NewArchiveSearchListResponseMetaWithDefaults instantiates a new ArchiveSearchListResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchListResponseMetaWithDefaults() *ArchiveSearchListResponseMeta {
	this := ArchiveSearchListResponseMeta{}
	return &this
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *ArchiveSearchListResponseMeta) GetTruncated() bool {
	if o == nil || o.Truncated == nil {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchListResponseMeta) GetTruncatedOk() (*bool, bool) {
	if o == nil || o.Truncated == nil {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *ArchiveSearchListResponseMeta) HasTruncated() bool {
	return o != nil && o.Truncated != nil
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *ArchiveSearchListResponseMeta) SetTruncated(v bool) {
	o.Truncated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchListResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Truncated != nil {
		toSerialize["truncated"] = o.Truncated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchListResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Truncated *bool `json:"truncated,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"truncated"})
	} else {
		return err
	}
	o.Truncated = all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
