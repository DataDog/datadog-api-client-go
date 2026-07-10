// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationshipToIncidentResponder A relationship reference for a single incident responder.
type RelationshipToIncidentResponder struct {
	// Relationship to impact object.
	Data NullableNullableRelationshipToIncidentResponderData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationshipToIncidentResponder instantiates a new RelationshipToIncidentResponder object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToIncidentResponder(data NullableNullableRelationshipToIncidentResponderData) *RelationshipToIncidentResponder {
	this := RelationshipToIncidentResponder{}
	this.Data = data
	return &this
}

// NewRelationshipToIncidentResponderWithDefaults instantiates a new RelationshipToIncidentResponder object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToIncidentResponderWithDefaults() *RelationshipToIncidentResponder {
	this := RelationshipToIncidentResponder{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for NullableRelationshipToIncidentResponderData will be returned.
func (o *RelationshipToIncidentResponder) GetData() NullableRelationshipToIncidentResponderData {
	if o == nil || o.Data.Get() == nil {
		var ret NullableRelationshipToIncidentResponderData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RelationshipToIncidentResponder) GetDataOk() (*NullableRelationshipToIncidentResponderData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *RelationshipToIncidentResponder) SetData(v NullableRelationshipToIncidentResponderData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToIncidentResponder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToIncidentResponder) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableNullableRelationshipToIncidentResponderData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Data.IsSet() {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableRelationshipToIncidentResponder handles when a null is used for RelationshipToIncidentResponder.
type NullableRelationshipToIncidentResponder struct {
	value *RelationshipToIncidentResponder
	isSet bool
}

// Get returns the associated value.
func (v NullableRelationshipToIncidentResponder) Get() *RelationshipToIncidentResponder {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRelationshipToIncidentResponder) Set(val *RelationshipToIncidentResponder) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRelationshipToIncidentResponder) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableRelationshipToIncidentResponder) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipToIncidentResponder initializes the struct as if Set has been called.
func NewNullableRelationshipToIncidentResponder(val *RelationshipToIncidentResponder) *NullableRelationshipToIncidentResponder {
	return &NullableRelationshipToIncidentResponder{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRelationshipToIncidentResponder) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRelationshipToIncidentResponder) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
