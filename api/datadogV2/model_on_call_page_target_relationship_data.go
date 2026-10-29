// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageTargetRelationshipData Reference to the target of the Page.
type OnCallPageTargetRelationshipData struct {
	// The ID of the target.
	Id string `json:"id"`
	// The resource type of the Page target.
	Type OnCallPageTargetRelationshipDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageTargetRelationshipData instantiates a new OnCallPageTargetRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageTargetRelationshipData(id string, typeVar OnCallPageTargetRelationshipDataType) *OnCallPageTargetRelationshipData {
	this := OnCallPageTargetRelationshipData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOnCallPageTargetRelationshipDataWithDefaults instantiates a new OnCallPageTargetRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageTargetRelationshipDataWithDefaults() *OnCallPageTargetRelationshipData {
	this := OnCallPageTargetRelationshipData{}
	return &this
}

// GetId returns the Id field value.
func (o *OnCallPageTargetRelationshipData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnCallPageTargetRelationshipData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OnCallPageTargetRelationshipData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OnCallPageTargetRelationshipData) GetType() OnCallPageTargetRelationshipDataType {
	if o == nil {
		var ret OnCallPageTargetRelationshipDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnCallPageTargetRelationshipData) GetTypeOk() (*OnCallPageTargetRelationshipDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnCallPageTargetRelationshipData) SetType(v OnCallPageTargetRelationshipDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageTargetRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageTargetRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                               `json:"id"`
		Type *OnCallPageTargetRelationshipDataType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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

// NullableOnCallPageTargetRelationshipData handles when a null is used for OnCallPageTargetRelationshipData.
type NullableOnCallPageTargetRelationshipData struct {
	value *OnCallPageTargetRelationshipData
	isSet bool
}

// Get returns the associated value.
func (v NullableOnCallPageTargetRelationshipData) Get() *OnCallPageTargetRelationshipData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableOnCallPageTargetRelationshipData) Set(val *OnCallPageTargetRelationshipData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableOnCallPageTargetRelationshipData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableOnCallPageTargetRelationshipData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOnCallPageTargetRelationshipData initializes the struct as if Set has been called.
func NewNullableOnCallPageTargetRelationshipData(val *OnCallPageTargetRelationshipData) *NullableOnCallPageTargetRelationshipData {
	return &NullableOnCallPageTargetRelationshipData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableOnCallPageTargetRelationshipData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableOnCallPageTargetRelationshipData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
