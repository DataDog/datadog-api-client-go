// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnModifier Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnModifier struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Args []DashboardInlineSchema48 `json:"args,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name TableColumnModifierName `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableColumnModifier instantiates a new TableColumnModifier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableColumnModifier(name TableColumnModifierName) *TableColumnModifier {
	this := TableColumnModifier{}
	this.Name = name
	return &this
}

// NewTableColumnModifierWithDefaults instantiates a new TableColumnModifier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableColumnModifierWithDefaults() *TableColumnModifier {
	this := TableColumnModifier{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *TableColumnModifier) GetArgs() []DashboardInlineSchema48 {
	if o == nil || o.Args == nil {
		var ret []DashboardInlineSchema48
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableColumnModifier) GetArgsOk() (*[]DashboardInlineSchema48, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *TableColumnModifier) HasArgs() bool {
	return o != nil && o.Args != nil
}

// SetArgs gets a reference to the given []DashboardInlineSchema48 and assigns it to the Args field.
func (o *TableColumnModifier) SetArgs(v []DashboardInlineSchema48) {
	o.Args = v
}

// GetName returns the Name field value.
func (o *TableColumnModifier) GetName() TableColumnModifierName {
	if o == nil {
		var ret TableColumnModifierName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TableColumnModifier) GetNameOk() (*TableColumnModifierName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TableColumnModifier) SetName(v TableColumnModifierName) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableColumnModifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableColumnModifier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Args []DashboardInlineSchema48 `json:"args,omitempty"`
		Name *TableColumnModifierName  `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Args = all.Args
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
