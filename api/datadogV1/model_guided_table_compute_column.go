// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableComputeColumn A guided table column that computes an aggregation directly from one of the table's base queries.
type GuidedTableComputeColumn struct {
	// Display name shown as the column header.
	Alias *string `json:"alias,omitempty"`
	// Comparison to display in a guided table column.
	Comparison *GuidedTableColumnComparison `json:"comparison,omitempty"`
	// Aggregation configuration.
	Compute GuidedTableComputeColumnCompute `json:"compute"`
	// Visual formatting applied to values in a guided table column. Supports number/bar mode and trend mode.
	Format GuidedTableColumnFormat `json:"format"`
	// Functions to apply sequentially to the computed value.
	Functions []GuidedTableColumnFunction `json:"functions,omitempty"`
	// Whether this column is hidden in the rendered table.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Auto-generated name used to refer to this column.
	Name string `json:"name"`
	//
	Type GuidedTableComputeColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableComputeColumn instantiates a new GuidedTableComputeColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableComputeColumn(compute GuidedTableComputeColumnCompute, format GuidedTableColumnFormat, name string, typeVar GuidedTableComputeColumnType) *GuidedTableComputeColumn {
	this := GuidedTableComputeColumn{}
	this.Compute = compute
	this.Format = format
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewGuidedTableComputeColumnWithDefaults instantiates a new GuidedTableComputeColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableComputeColumnWithDefaults() *GuidedTableComputeColumn {
	this := GuidedTableComputeColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GuidedTableComputeColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GuidedTableComputeColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GuidedTableComputeColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *GuidedTableComputeColumn) GetComparison() GuidedTableColumnComparison {
	if o == nil || o.Comparison == nil {
		var ret GuidedTableColumnComparison
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetComparisonOk() (*GuidedTableColumnComparison, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *GuidedTableComputeColumn) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given GuidedTableColumnComparison and assigns it to the Comparison field.
func (o *GuidedTableComputeColumn) SetComparison(v GuidedTableColumnComparison) {
	o.Comparison = &v
}

// GetCompute returns the Compute field value.
func (o *GuidedTableComputeColumn) GetCompute() GuidedTableComputeColumnCompute {
	if o == nil {
		var ret GuidedTableComputeColumnCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetComputeOk() (*GuidedTableComputeColumnCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *GuidedTableComputeColumn) SetCompute(v GuidedTableComputeColumnCompute) {
	o.Compute = v
}

// GetFormat returns the Format field value.
func (o *GuidedTableComputeColumn) GetFormat() GuidedTableColumnFormat {
	if o == nil {
		var ret GuidedTableColumnFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetFormatOk() (*GuidedTableColumnFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *GuidedTableComputeColumn) SetFormat(v GuidedTableColumnFormat) {
	o.Format = v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *GuidedTableComputeColumn) GetFunctions() []GuidedTableColumnFunction {
	if o == nil || o.Functions == nil {
		var ret []GuidedTableColumnFunction
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetFunctionsOk() (*[]GuidedTableColumnFunction, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return &o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *GuidedTableComputeColumn) HasFunctions() bool {
	return o != nil && o.Functions != nil
}

// SetFunctions gets a reference to the given []GuidedTableColumnFunction and assigns it to the Functions field.
func (o *GuidedTableComputeColumn) SetFunctions(v []GuidedTableColumnFunction) {
	o.Functions = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *GuidedTableComputeColumn) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *GuidedTableComputeColumn) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *GuidedTableComputeColumn) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *GuidedTableComputeColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableComputeColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *GuidedTableComputeColumn) GetType() GuidedTableComputeColumnType {
	if o == nil {
		var ret GuidedTableComputeColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumn) GetTypeOk() (*GuidedTableComputeColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GuidedTableComputeColumn) SetType(v GuidedTableComputeColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableComputeColumn) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableComputeColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                          `json:"alias,omitempty"`
		Comparison *GuidedTableColumnComparison     `json:"comparison,omitempty"`
		Compute    *GuidedTableComputeColumnCompute `json:"compute"`
		Format     *GuidedTableColumnFormat         `json:"format"`
		Functions  []GuidedTableColumnFunction      `json:"functions,omitempty"`
		IsHidden   *bool                            `json:"is_hidden,omitempty"`
		Name       *string                          `json:"name"`
		Type       *GuidedTableComputeColumnType    `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "comparison", "compute", "format", "functions", "is_hidden", "name", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
