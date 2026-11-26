// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type TableRowGroupBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Format *TableRowGroupByFormat `json:"format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupKeys []TableGroupKey `json:"group_keys"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableRowGroupBy instantiates a new TableRowGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowGroupBy(groupKeys []TableGroupKey, name string) *TableRowGroupBy {
	this := TableRowGroupBy{}
	this.GroupKeys = groupKeys
	this.Name = name
	return &this
}

// NewTableRowGroupByWithDefaults instantiates a new TableRowGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowGroupByWithDefaults() *TableRowGroupBy {
	this := TableRowGroupBy{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *TableRowGroupBy) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowGroupBy) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *TableRowGroupBy) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *TableRowGroupBy) SetAlias(v string) {
	o.Alias = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *TableRowGroupBy) GetFormat() TableRowGroupByFormat {
	if o == nil || o.Format == nil {
		var ret TableRowGroupByFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowGroupBy) GetFormatOk() (*TableRowGroupByFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *TableRowGroupBy) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given TableRowGroupByFormat and assigns it to the Format field.
func (o *TableRowGroupBy) SetFormat(v TableRowGroupByFormat) {
	o.Format = &v
}

// GetGroupKeys returns the GroupKeys field value.
func (o *TableRowGroupBy) GetGroupKeys() []TableGroupKey {
	if o == nil {
		var ret []TableGroupKey
		return ret
	}
	return o.GroupKeys
}

// GetGroupKeysOk returns a tuple with the GroupKeys field value
// and a boolean to check if the value has been set.
func (o *TableRowGroupBy) GetGroupKeysOk() (*[]TableGroupKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupKeys, true
}

// SetGroupKeys sets field value.
func (o *TableRowGroupBy) SetGroupKeys(v []TableGroupKey) {
	o.GroupKeys = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *TableRowGroupBy) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowGroupBy) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *TableRowGroupBy) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *TableRowGroupBy) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *TableRowGroupBy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TableRowGroupBy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TableRowGroupBy) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	toSerialize["group_keys"] = o.GroupKeys
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableRowGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias     *string                `json:"alias,omitempty"`
		Format    *TableRowGroupByFormat `json:"format,omitempty"`
		GroupKeys *[]TableGroupKey       `json:"group_keys"`
		IsHidden  *bool                  `json:"is_hidden,omitempty"`
		Name      *string                `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupKeys == nil {
		return fmt.Errorf("required field group_keys missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if all.Format != nil && all.Format.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Format = all.Format
	o.GroupKeys = *all.GroupKeys
	o.IsHidden = all.IsHidden
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
