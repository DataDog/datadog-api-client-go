// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterData A permanent retention filter.
type RumPermanentRetentionFilterData struct {
	// The attributes of a permanent retention filter.
	Attributes *RumPermanentRetentionFilterAttributes `json:"attributes,omitempty"`
	// Permanent retention filter ID.
	Id *string `json:"id,omitempty"`
	// Metadata about the permanent retention filter.
	Meta *RumPermanentRetentionFilterMeta `json:"meta,omitempty"`
	// The resource type. The value must be `permanent_retention_filters`.
	Type *RumPermanentRetentionFilterType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentRetentionFilterData instantiates a new RumPermanentRetentionFilterData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentRetentionFilterData() *RumPermanentRetentionFilterData {
	this := RumPermanentRetentionFilterData{}
	var typeVar RumPermanentRetentionFilterType = RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS
	this.Type = &typeVar
	return &this
}

// NewRumPermanentRetentionFilterDataWithDefaults instantiates a new RumPermanentRetentionFilterData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentRetentionFilterDataWithDefaults() *RumPermanentRetentionFilterData {
	this := RumPermanentRetentionFilterData{}
	var typeVar RumPermanentRetentionFilterType = RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterData) GetAttributes() RumPermanentRetentionFilterAttributes {
	if o == nil || o.Attributes == nil {
		var ret RumPermanentRetentionFilterAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterData) GetAttributesOk() (*RumPermanentRetentionFilterAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RumPermanentRetentionFilterAttributes and assigns it to the Attributes field.
func (o *RumPermanentRetentionFilterData) SetAttributes(v RumPermanentRetentionFilterAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RumPermanentRetentionFilterData) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterData) GetMeta() RumPermanentRetentionFilterMeta {
	if o == nil || o.Meta == nil {
		var ret RumPermanentRetentionFilterMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterData) GetMetaOk() (*RumPermanentRetentionFilterMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterData) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given RumPermanentRetentionFilterMeta and assigns it to the Meta field.
func (o *RumPermanentRetentionFilterData) SetMeta(v RumPermanentRetentionFilterMeta) {
	o.Meta = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterData) GetType() RumPermanentRetentionFilterType {
	if o == nil || o.Type == nil {
		var ret RumPermanentRetentionFilterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterData) GetTypeOk() (*RumPermanentRetentionFilterType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RumPermanentRetentionFilterType and assigns it to the Type field.
func (o *RumPermanentRetentionFilterData) SetType(v RumPermanentRetentionFilterType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentRetentionFilterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumPermanentRetentionFilterData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RumPermanentRetentionFilterAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Meta       *RumPermanentRetentionFilterMeta       `json:"meta,omitempty"`
		Type       *RumPermanentRetentionFilterType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
