// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterUpdateData The permanent retention filter properties to update.
type RumPermanentRetentionFilterUpdateData struct {
	// The attributes of a permanent retention filter that can be updated.
	Attributes RumPermanentRetentionFilterUpdateAttributes `json:"attributes"`
	// Permanent retention filter ID.
	Id string `json:"id"`
	// The resource type. The value must be `permanent_retention_filters`.
	Type RumPermanentRetentionFilterType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentRetentionFilterUpdateData instantiates a new RumPermanentRetentionFilterUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentRetentionFilterUpdateData(attributes RumPermanentRetentionFilterUpdateAttributes, id string, typeVar RumPermanentRetentionFilterType) *RumPermanentRetentionFilterUpdateData {
	this := RumPermanentRetentionFilterUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRumPermanentRetentionFilterUpdateDataWithDefaults instantiates a new RumPermanentRetentionFilterUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentRetentionFilterUpdateDataWithDefaults() *RumPermanentRetentionFilterUpdateData {
	this := RumPermanentRetentionFilterUpdateData{}
	var typeVar RumPermanentRetentionFilterType = RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RumPermanentRetentionFilterUpdateData) GetAttributes() RumPermanentRetentionFilterUpdateAttributes {
	if o == nil {
		var ret RumPermanentRetentionFilterUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterUpdateData) GetAttributesOk() (*RumPermanentRetentionFilterUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RumPermanentRetentionFilterUpdateData) SetAttributes(v RumPermanentRetentionFilterUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *RumPermanentRetentionFilterUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RumPermanentRetentionFilterUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *RumPermanentRetentionFilterUpdateData) GetType() RumPermanentRetentionFilterType {
	if o == nil {
		var ret RumPermanentRetentionFilterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterUpdateData) GetTypeOk() (*RumPermanentRetentionFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RumPermanentRetentionFilterUpdateData) SetType(v RumPermanentRetentionFilterType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentRetentionFilterUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumPermanentRetentionFilterUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RumPermanentRetentionFilterUpdateAttributes `json:"attributes"`
		Id         *string                                      `json:"id"`
		Type       *RumPermanentRetentionFilterType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
