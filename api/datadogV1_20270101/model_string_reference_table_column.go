// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StringReferenceTableColumn Test-only generated schema for the dashboard versioning proof of concept.
type StringReferenceTableColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Format StringTableColumnFormat `json:"format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Functions []ExcludeNullTableColumnFunction `json:"functions,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Reference StringReferenceTableColumnReference `json:"reference"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StringReferenceTableColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStringReferenceTableColumn instantiates a new StringReferenceTableColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStringReferenceTableColumn(format StringTableColumnFormat, name string, reference StringReferenceTableColumnReference, typeVar StringReferenceTableColumnType) *StringReferenceTableColumn {
	this := StringReferenceTableColumn{}
	this.Format = format
	this.Name = name
	this.Reference = reference
	this.Type = typeVar
	return &this
}

// NewStringReferenceTableColumnWithDefaults instantiates a new StringReferenceTableColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStringReferenceTableColumnWithDefaults() *StringReferenceTableColumn {
	this := StringReferenceTableColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *StringReferenceTableColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *StringReferenceTableColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *StringReferenceTableColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetFormat returns the Format field value.
func (o *StringReferenceTableColumn) GetFormat() StringTableColumnFormat {
	if o == nil {
		var ret StringTableColumnFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetFormatOk() (*StringTableColumnFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *StringReferenceTableColumn) SetFormat(v StringTableColumnFormat) {
	o.Format = v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *StringReferenceTableColumn) GetFunctions() []ExcludeNullTableColumnFunction {
	if o == nil || o.Functions == nil {
		var ret []ExcludeNullTableColumnFunction
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetFunctionsOk() (*[]ExcludeNullTableColumnFunction, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return &o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *StringReferenceTableColumn) HasFunctions() bool {
	return o != nil && o.Functions != nil
}

// SetFunctions gets a reference to the given []ExcludeNullTableColumnFunction and assigns it to the Functions field.
func (o *StringReferenceTableColumn) SetFunctions(v []ExcludeNullTableColumnFunction) {
	o.Functions = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *StringReferenceTableColumn) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *StringReferenceTableColumn) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *StringReferenceTableColumn) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *StringReferenceTableColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StringReferenceTableColumn) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value.
func (o *StringReferenceTableColumn) GetReference() StringReferenceTableColumnReference {
	if o == nil {
		var ret StringReferenceTableColumnReference
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetReferenceOk() (*StringReferenceTableColumnReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value.
func (o *StringReferenceTableColumn) SetReference(v StringReferenceTableColumnReference) {
	o.Reference = v
}

// GetType returns the Type field value.
func (o *StringReferenceTableColumn) GetType() StringReferenceTableColumnType {
	if o == nil {
		var ret StringReferenceTableColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumn) GetTypeOk() (*StringReferenceTableColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StringReferenceTableColumn) SetType(v StringReferenceTableColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StringReferenceTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["format"] = o.Format
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	toSerialize["name"] = o.Name
	toSerialize["reference"] = o.Reference
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StringReferenceTableColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias     *string                              `json:"alias,omitempty"`
		Format    *StringTableColumnFormat             `json:"format"`
		Functions []ExcludeNullTableColumnFunction     `json:"functions,omitempty"`
		IsHidden  *bool                                `json:"is_hidden,omitempty"`
		Name      *string                              `json:"name"`
		Reference *StringReferenceTableColumnReference `json:"reference"`
		Type      *StringReferenceTableColumnType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Reference == nil {
		return fmt.Errorf("required field reference missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if all.Format.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Format = *all.Format
	o.Functions = all.Functions
	o.IsHidden = all.IsHidden
	o.Name = *all.Name
	if all.Reference.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Reference = *all.Reference
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
