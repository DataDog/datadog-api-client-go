// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaTableColumn Test-only generated schema for the dashboard versioning proof of concept.
type FormulaTableColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *TableColumnComparison `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Format TableColumnFormat `json:"format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type FormulaTableColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaTableColumn instantiates a new FormulaTableColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaTableColumn(format TableColumnFormat, formula string, name string, typeVar FormulaTableColumnType) *FormulaTableColumn {
	this := FormulaTableColumn{}
	this.Format = format
	this.Formula = formula
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewFormulaTableColumnWithDefaults instantiates a new FormulaTableColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaTableColumnWithDefaults() *FormulaTableColumn {
	this := FormulaTableColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *FormulaTableColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *FormulaTableColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *FormulaTableColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *FormulaTableColumn) GetComparison() TableColumnComparison {
	if o == nil || o.Comparison == nil {
		var ret TableColumnComparison
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetComparisonOk() (*TableColumnComparison, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *FormulaTableColumn) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given TableColumnComparison and assigns it to the Comparison field.
func (o *FormulaTableColumn) SetComparison(v TableColumnComparison) {
	o.Comparison = &v
}

// GetFormat returns the Format field value.
func (o *FormulaTableColumn) GetFormat() TableColumnFormat {
	if o == nil {
		var ret TableColumnFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetFormatOk() (*TableColumnFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *FormulaTableColumn) SetFormat(v TableColumnFormat) {
	o.Format = v
}

// GetFormula returns the Formula field value.
func (o *FormulaTableColumn) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *FormulaTableColumn) SetFormula(v string) {
	o.Formula = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *FormulaTableColumn) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *FormulaTableColumn) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *FormulaTableColumn) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *FormulaTableColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaTableColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *FormulaTableColumn) GetType() FormulaTableColumnType {
	if o == nil {
		var ret FormulaTableColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormulaTableColumn) GetTypeOk() (*FormulaTableColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FormulaTableColumn) SetType(v FormulaTableColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaTableColumn) MarshalJSON() ([]byte, error) {
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
	toSerialize["format"] = o.Format
	toSerialize["formula"] = o.Formula
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaTableColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                 `json:"alias,omitempty"`
		Comparison *TableColumnComparison  `json:"comparison,omitempty"`
		Format     *TableColumnFormat      `json:"format"`
		Formula    *string                 `json:"formula"`
		IsHidden   *bool                   `json:"is_hidden,omitempty"`
		Name       *string                 `json:"name"`
		Type       *FormulaTableColumnType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
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
	o.Format = *all.Format
	o.Formula = *all.Formula
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
