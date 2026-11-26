// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatch20270101OneOf4462867328 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatch20270101OneOf4462867328 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TextMatchAnyOfAnyOfItemType4452735776_20270101 `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTextMatch20270101OneOf4462867328 instantiates a new TextMatch20270101OneOf4462867328 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTextMatch20270101OneOf4462867328(typeVar TextMatchAnyOfAnyOfItemType4452735776_20270101, value string) *TextMatch20270101OneOf4462867328 {
	this := TextMatch20270101OneOf4462867328{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewTextMatch20270101OneOf4462867328WithDefaults instantiates a new TextMatch20270101OneOf4462867328 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTextMatch20270101OneOf4462867328WithDefaults() *TextMatch20270101OneOf4462867328 {
	this := TextMatch20270101OneOf4462867328{}
	return &this
}

// GetType returns the Type field value.
func (o *TextMatch20270101OneOf4462867328) GetType() TextMatchAnyOfAnyOfItemType4452735776_20270101 {
	if o == nil {
		var ret TextMatchAnyOfAnyOfItemType4452735776_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextMatch20270101OneOf4462867328) GetTypeOk() (*TextMatchAnyOfAnyOfItemType4452735776_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TextMatch20270101OneOf4462867328) SetType(v TextMatchAnyOfAnyOfItemType4452735776_20270101) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *TextMatch20270101OneOf4462867328) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TextMatch20270101OneOf4462867328) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *TextMatch20270101OneOf4462867328) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TextMatch20270101OneOf4462867328) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TextMatch20270101OneOf4462867328) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *TextMatchAnyOfAnyOfItemType4452735776_20270101 `json:"type"`
		Value *string                                         `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
