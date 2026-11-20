// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetFormulasItems Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetFormulasItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDatasetFormulasItems instantiates a new FormulaDatasetFormulasItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDatasetFormulasItems(formula string) *FormulaDatasetFormulasItems {
	this := FormulaDatasetFormulasItems{}
	this.Formula = formula
	return &this
}

// NewFormulaDatasetFormulasItemsWithDefaults instantiates a new FormulaDatasetFormulasItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDatasetFormulasItemsWithDefaults() *FormulaDatasetFormulasItems {
	this := FormulaDatasetFormulasItems{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *FormulaDatasetFormulasItems) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetFormulasItems) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *FormulaDatasetFormulasItems) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *FormulaDatasetFormulasItems) SetAlias(v string) {
	o.Alias = &v
}

// GetFormula returns the Formula field value.
func (o *FormulaDatasetFormulasItems) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetFormulasItems) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *FormulaDatasetFormulasItems) SetFormula(v string) {
	o.Formula = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDatasetFormulasItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["formula"] = o.Formula
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDatasetFormulasItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias   *string `json:"alias,omitempty"`
		Formula *string `json:"formula"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}
	o.Alias = all.Alias
	o.Formula = *all.Formula

	return nil
}
