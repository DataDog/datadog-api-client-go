// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamColumn Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Field IssueStreamColumnName `json:"field"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width IssueStreamColumn20270101Width `json:"width"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewIssueStreamColumn instantiates a new IssueStreamColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueStreamColumn(field IssueStreamColumnName, width IssueStreamColumn20270101Width) *IssueStreamColumn {
	this := IssueStreamColumn{}
	this.Field = field
	this.Width = width
	return &this
}

// NewIssueStreamColumnWithDefaults instantiates a new IssueStreamColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueStreamColumnWithDefaults() *IssueStreamColumn {
	this := IssueStreamColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *IssueStreamColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStreamColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *IssueStreamColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *IssueStreamColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetField returns the Field field value.
func (o *IssueStreamColumn) GetField() IssueStreamColumnName {
	if o == nil {
		var ret IssueStreamColumnName
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *IssueStreamColumn) GetFieldOk() (*IssueStreamColumnName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *IssueStreamColumn) SetField(v IssueStreamColumnName) {
	o.Field = v
}

// GetWidth returns the Width field value.
func (o *IssueStreamColumn) GetWidth() IssueStreamColumn20270101Width {
	if o == nil {
		var ret IssueStreamColumn20270101Width
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *IssueStreamColumn) GetWidthOk() (*IssueStreamColumn20270101Width, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *IssueStreamColumn) SetWidth(v IssueStreamColumn20270101Width) {
	o.Width = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueStreamColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["field"] = o.Field
	toSerialize["width"] = o.Width
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueStreamColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias *string                         `json:"alias,omitempty"`
		Field *IssueStreamColumnName          `json:"field"`
		Width *IssueStreamColumn20270101Width `json:"width"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.Field.IsValid() {
		hasInvalidField = true
	} else {
		o.Field = *all.Field
	}
	o.Width = *all.Width

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
