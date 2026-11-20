// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapScalarFormulaRequest Test-only generated schema for the dashboard versioning proof of concept.
type HostmapScalarFormulaRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []HostmapFormula `json:"formulas"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat HostmapScalarFormulaRequestResponseFormat `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapScalarFormulaRequest instantiates a new HostmapScalarFormulaRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapScalarFormulaRequest(formulas []HostmapFormula, queries []FormulaAndFunctionQueryDefinition, responseFormat HostmapScalarFormulaRequestResponseFormat) *HostmapScalarFormulaRequest {
	this := HostmapScalarFormulaRequest{}
	this.Formulas = formulas
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewHostmapScalarFormulaRequestWithDefaults instantiates a new HostmapScalarFormulaRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapScalarFormulaRequestWithDefaults() *HostmapScalarFormulaRequest {
	this := HostmapScalarFormulaRequest{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *HostmapScalarFormulaRequest) GetFormulas() []HostmapFormula {
	if o == nil {
		var ret []HostmapFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *HostmapScalarFormulaRequest) GetFormulasOk() (*[]HostmapFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *HostmapScalarFormulaRequest) SetFormulas(v []HostmapFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *HostmapScalarFormulaRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *HostmapScalarFormulaRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *HostmapScalarFormulaRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *HostmapScalarFormulaRequest) GetResponseFormat() HostmapScalarFormulaRequestResponseFormat {
	if o == nil {
		var ret HostmapScalarFormulaRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *HostmapScalarFormulaRequest) GetResponseFormatOk() (*HostmapScalarFormulaRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *HostmapScalarFormulaRequest) SetResponseFormat(v HostmapScalarFormulaRequestResponseFormat) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *HostmapScalarFormulaRequest) GetSort() ScalarFormulaSort {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapScalarFormulaRequest) GetSortOk() (*ScalarFormulaSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *HostmapScalarFormulaRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort and assigns it to the Sort field.
func (o *HostmapScalarFormulaRequest) SetSort(v ScalarFormulaSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapScalarFormulaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formulas"] = o.Formulas
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostmapScalarFormulaRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       *[]HostmapFormula                          `json:"formulas"`
		Queries        *[]FormulaAndFunctionQueryDefinition       `json:"queries"`
		ResponseFormat *HostmapScalarFormulaRequestResponseFormat `json:"response_format"`
		Sort           *ScalarFormulaSort                         `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formulas == nil {
		return fmt.Errorf("required field formulas missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Formulas = *all.Formulas
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
