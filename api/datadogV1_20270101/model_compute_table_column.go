// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComputeTableColumn Test-only generated schema for the dashboard versioning proof of concept.
type ComputeTableColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *TableColumnComparison `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute ComputeTableColumnCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Format TableColumnFormat `json:"format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Functions []TableColumnFunction `json:"functions,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ComputeTableColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComputeTableColumn instantiates a new ComputeTableColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeTableColumn(compute ComputeTableColumnCompute, format TableColumnFormat, name string, typeVar ComputeTableColumnType) *ComputeTableColumn {
	this := ComputeTableColumn{}
	this.Compute = compute
	this.Format = format
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewComputeTableColumnWithDefaults instantiates a new ComputeTableColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeTableColumnWithDefaults() *ComputeTableColumn {
	this := ComputeTableColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ComputeTableColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ComputeTableColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ComputeTableColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *ComputeTableColumn) GetComparison() TableColumnComparison {
	if o == nil || o.Comparison == nil {
		var ret TableColumnComparison
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetComparisonOk() (*TableColumnComparison, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *ComputeTableColumn) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given TableColumnComparison and assigns it to the Comparison field.
func (o *ComputeTableColumn) SetComparison(v TableColumnComparison) {
	o.Comparison = &v
}

// GetCompute returns the Compute field value.
func (o *ComputeTableColumn) GetCompute() ComputeTableColumnCompute {
	if o == nil {
		var ret ComputeTableColumnCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetComputeOk() (*ComputeTableColumnCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *ComputeTableColumn) SetCompute(v ComputeTableColumnCompute) {
	o.Compute = v
}

// GetFormat returns the Format field value.
func (o *ComputeTableColumn) GetFormat() TableColumnFormat {
	if o == nil {
		var ret TableColumnFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetFormatOk() (*TableColumnFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *ComputeTableColumn) SetFormat(v TableColumnFormat) {
	o.Format = v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *ComputeTableColumn) GetFunctions() []TableColumnFunction {
	if o == nil || o.Functions == nil {
		var ret []TableColumnFunction
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetFunctionsOk() (*[]TableColumnFunction, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return &o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *ComputeTableColumn) HasFunctions() bool {
	return o != nil && o.Functions != nil
}

// SetFunctions gets a reference to the given []TableColumnFunction and assigns it to the Functions field.
func (o *ComputeTableColumn) SetFunctions(v []TableColumnFunction) {
	o.Functions = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *ComputeTableColumn) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *ComputeTableColumn) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *ComputeTableColumn) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *ComputeTableColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ComputeTableColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ComputeTableColumn) GetType() ComputeTableColumnType {
	if o == nil {
		var ret ComputeTableColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComputeTableColumn) GetTypeOk() (*ComputeTableColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ComputeTableColumn) SetType(v ComputeTableColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Comparison != nil {
		toSerialize["comparison"] = o.Comparison
	}
	toSerialize["compute"] = o.Compute
	toSerialize["format"] = o.Format
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeTableColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                    `json:"alias,omitempty"`
		Comparison *TableColumnComparison     `json:"comparison,omitempty"`
		Compute    *ComputeTableColumnCompute `json:"compute"`
		Format     *TableColumnFormat         `json:"format"`
		Functions  []TableColumnFunction      `json:"functions,omitempty"`
		IsHidden   *bool                      `json:"is_hidden,omitempty"`
		Name       *string                    `json:"name"`
		Type       *ComputeTableColumnType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	o.Comparison = all.Comparison
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.Format = *all.Format
	o.Functions = all.Functions
	o.IsHidden = all.IsHidden
	o.Name = *all.Name
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
