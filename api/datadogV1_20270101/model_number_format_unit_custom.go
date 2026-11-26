// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitCustom Test-only generated schema for the dashboard versioning proof of concept.
type NumberFormatUnitCustom struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label string `json:"label"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type NumberFormatUnitCustomType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNumberFormatUnitCustom instantiates a new NumberFormatUnitCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberFormatUnitCustom(label string, typeVar NumberFormatUnitCustomType) *NumberFormatUnitCustom {
	this := NumberFormatUnitCustom{}
	this.Label = label
	this.Type = typeVar
	return &this
}

// NewNumberFormatUnitCustomWithDefaults instantiates a new NumberFormatUnitCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberFormatUnitCustomWithDefaults() *NumberFormatUnitCustom {
	this := NumberFormatUnitCustom{}
	return &this
}

// GetLabel returns the Label field value.
func (o *NumberFormatUnitCustom) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCustom) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *NumberFormatUnitCustom) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value.
func (o *NumberFormatUnitCustom) GetType() NumberFormatUnitCustomType {
	if o == nil {
		var ret NumberFormatUnitCustomType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCustom) GetTypeOk() (*NumberFormatUnitCustomType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NumberFormatUnitCustom) SetType(v NumberFormatUnitCustomType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberFormatUnitCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NumberFormatUnitCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label *string                     `json:"label"`
		Type  *NumberFormatUnitCustomType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Label = *all.Label
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
