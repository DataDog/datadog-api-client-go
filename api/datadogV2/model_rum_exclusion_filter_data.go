// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMExclusionFilterData An exclusion filter.
type RUMExclusionFilterData struct {
	// The attributes of an exclusion filter.
	Attributes *RUMExclusionFilterAttributes `json:"attributes,omitempty"`
	// The ID of the exclusion filter.
	Id string `json:"id"`
	// Metadata about the exclusion filter.
	Meta *RUMExclusionFilterMeta `json:"meta,omitempty"`
	// The resource type. The value must be `exclusion_filters`.
	Type RUMExclusionFilterType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMExclusionFilterData instantiates a new RUMExclusionFilterData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMExclusionFilterData(id string, typeVar RUMExclusionFilterType) *RUMExclusionFilterData {
	this := RUMExclusionFilterData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRUMExclusionFilterDataWithDefaults instantiates a new RUMExclusionFilterData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMExclusionFilterDataWithDefaults() *RUMExclusionFilterData {
	this := RUMExclusionFilterData{}
	var typeVar RUMExclusionFilterType = RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RUMExclusionFilterData) GetAttributes() RUMExclusionFilterAttributes {
	if o == nil || o.Attributes == nil {
		var ret RUMExclusionFilterAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterData) GetAttributesOk() (*RUMExclusionFilterAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RUMExclusionFilterData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RUMExclusionFilterAttributes and assigns it to the Attributes field.
func (o *RUMExclusionFilterData) SetAttributes(v RUMExclusionFilterAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *RUMExclusionFilterData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RUMExclusionFilterData) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RUMExclusionFilterData) GetMeta() RUMExclusionFilterMeta {
	if o == nil || o.Meta == nil {
		var ret RUMExclusionFilterMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterData) GetMetaOk() (*RUMExclusionFilterMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RUMExclusionFilterData) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given RUMExclusionFilterMeta and assigns it to the Meta field.
func (o *RUMExclusionFilterData) SetMeta(v RUMExclusionFilterMeta) {
	o.Meta = &v
}

// GetType returns the Type field value.
func (o *RUMExclusionFilterData) GetType() RUMExclusionFilterType {
	if o == nil {
		var ret RUMExclusionFilterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterData) GetTypeOk() (*RUMExclusionFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMExclusionFilterData) SetType(v RUMExclusionFilterType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMExclusionFilterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMExclusionFilterData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RUMExclusionFilterAttributes `json:"attributes,omitempty"`
		Id         *string                       `json:"id"`
		Meta       *RUMExclusionFilterMeta       `json:"meta,omitempty"`
		Type       *RUMExclusionFilterType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
