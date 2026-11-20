// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf4461524480 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf4461524480 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []ScatterplotFormula `json:"formulas"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf4461524480 instantiates a new ScatterplotBackendTableRequest20270101OneOf4461524480 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf4461524480(formulas []ScatterplotFormula, queries []FormulaAndFunctionQueryDefinition, responseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) *ScatterplotBackendTableRequest20270101OneOf4461524480 {
	this := ScatterplotBackendTableRequest20270101OneOf4461524480{}
	this.Formulas = formulas
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOf4461524480WithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf4461524480 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOf4461524480WithDefaults() *ScatterplotBackendTableRequest20270101OneOf4461524480 {
	this := ScatterplotBackendTableRequest20270101OneOf4461524480{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetFormulas() []ScatterplotFormula {
	if o == nil {
		var ret []ScatterplotFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetFormulasOk() (*[]ScatterplotFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) SetFormulas(v []ScatterplotFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetResponseFormat() ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetResponseFormatOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) SetResponseFormat(v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetSort() ScalarFormulaSort {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) GetSortOk() (*ScalarFormulaSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort and assigns it to the Sort field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) SetSort(v ScalarFormulaSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf4461524480) MarshalJSON() ([]byte, error) {
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
func (o *ScatterplotBackendTableRequest20270101OneOf4461524480) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       *[]ScatterplotFormula                                                 `json:"formulas"`
		Queries        *[]FormulaAndFunctionQueryDefinition                                  `json:"queries"`
		ResponseFormat *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 `json:"response_format"`
		Sort           *ScalarFormulaSort                                                    `json:"sort,omitempty"`
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
