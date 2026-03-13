// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableFormulaColumn A guided table column that displays the result of evaluating a formula over other columns.
type GuidedTableFormulaColumn struct {
	// Display name shown as the column header.
	Alias *string `json:"alias,omitempty"`
	// Comparison to display in a guided table column.
	Comparison *GuidedTableColumnComparison `json:"comparison,omitempty"`
	// Visual formatting applied to values in a guided table column. Supports number/bar mode and trend mode.
	Format GuidedTableColumnFormat `json:"format"`
	// The formula expression to evaluate.
	Formula string `json:"formula"`
	// Whether this column is hidden in the rendered table.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Auto-generated name used to refer to this column.
	Name string `json:"name"`
	// Set the sort type to formula.
	Type FormulaType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableFormulaColumn instantiates a new GuidedTableFormulaColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableFormulaColumn(format GuidedTableColumnFormat, formula string, name string, typeVar FormulaType) *GuidedTableFormulaColumn {
	this := GuidedTableFormulaColumn{}
	this.Format = format
	this.Formula = formula
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewGuidedTableFormulaColumnWithDefaults instantiates a new GuidedTableFormulaColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableFormulaColumnWithDefaults() *GuidedTableFormulaColumn {
	this := GuidedTableFormulaColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GuidedTableFormulaColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GuidedTableFormulaColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GuidedTableFormulaColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *GuidedTableFormulaColumn) GetComparison() GuidedTableColumnComparison {
	if o == nil || o.Comparison == nil {
		var ret GuidedTableColumnComparison
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetComparisonOk() (*GuidedTableColumnComparison, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *GuidedTableFormulaColumn) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given GuidedTableColumnComparison and assigns it to the Comparison field.
func (o *GuidedTableFormulaColumn) SetComparison(v GuidedTableColumnComparison) {
	o.Comparison = &v
}

// GetFormat returns the Format field value.
func (o *GuidedTableFormulaColumn) GetFormat() GuidedTableColumnFormat {
	if o == nil {
		var ret GuidedTableColumnFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetFormatOk() (*GuidedTableColumnFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *GuidedTableFormulaColumn) SetFormat(v GuidedTableColumnFormat) {
	o.Format = v
}

// GetFormula returns the Formula field value.
func (o *GuidedTableFormulaColumn) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *GuidedTableFormulaColumn) SetFormula(v string) {
	o.Formula = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *GuidedTableFormulaColumn) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *GuidedTableFormulaColumn) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *GuidedTableFormulaColumn) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *GuidedTableFormulaColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableFormulaColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *GuidedTableFormulaColumn) GetType() FormulaType {
	if o == nil {
		var ret FormulaType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuidedTableFormulaColumn) GetTypeOk() (*FormulaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GuidedTableFormulaColumn) SetType(v FormulaType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableFormulaColumn) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableFormulaColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                      `json:"alias,omitempty"`
		Comparison *GuidedTableColumnComparison `json:"comparison,omitempty"`
		Format     *GuidedTableColumnFormat     `json:"format"`
		Formula    *string                      `json:"formula"`
		IsHidden   *bool                        `json:"is_hidden,omitempty"`
		Name       *string                      `json:"name"`
		Type       *FormulaType                 `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "comparison", "format", "formula", "is_hidden", "name", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
