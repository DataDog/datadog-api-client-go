// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StringTableColumnFormat Test-only generated schema for the dashboard versioning proof of concept.
type StringTableColumnFormat struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode StringTableColumnFormatMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStringTableColumnFormat instantiates a new StringTableColumnFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStringTableColumnFormat(mode StringTableColumnFormatMode) *StringTableColumnFormat {
	this := StringTableColumnFormat{}
	this.Mode = mode
	return &this
}

// NewStringTableColumnFormatWithDefaults instantiates a new StringTableColumnFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStringTableColumnFormatWithDefaults() *StringTableColumnFormat {
	this := StringTableColumnFormat{}
	return &this
}

// GetMode returns the Mode field value.
func (o *StringTableColumnFormat) GetMode() StringTableColumnFormatMode {
	if o == nil {
		var ret StringTableColumnFormatMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *StringTableColumnFormat) GetModeOk() (*StringTableColumnFormatMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *StringTableColumnFormat) SetMode(v StringTableColumnFormatMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StringTableColumnFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StringTableColumnFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode *StringTableColumnFormatMode `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
