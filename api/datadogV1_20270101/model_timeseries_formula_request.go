// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesFormulaRequest Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesFormulaRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TimeseriesFormulaRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesFormulaRequest instantiates a new TimeseriesFormulaRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesFormulaRequest(queries []FormulaAndFunctionQueryDefinition, responseFormat TimeseriesFormulaRequestResponseFormat) *TimeseriesFormulaRequest {
	this := TimeseriesFormulaRequest{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewTimeseriesFormulaRequestWithDefaults instantiates a new TimeseriesFormulaRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesFormulaRequestWithDefaults() *TimeseriesFormulaRequest {
	this := TimeseriesFormulaRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *TimeseriesFormulaRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesFormulaRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *TimeseriesFormulaRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *TimeseriesFormulaRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *TimeseriesFormulaRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *TimeseriesFormulaRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *TimeseriesFormulaRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TimeseriesFormulaRequest) GetResponseFormat() TimeseriesFormulaRequestResponseFormat {
	if o == nil {
		var ret TimeseriesFormulaRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TimeseriesFormulaRequest) GetResponseFormatOk() (*TimeseriesFormulaRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TimeseriesFormulaRequest) SetResponseFormat(v TimeseriesFormulaRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesFormulaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesFormulaRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                         `json:"formulas,omitempty"`
		Queries        *[]FormulaAndFunctionQueryDefinition    `json:"queries"`
		ResponseFormat *TimeseriesFormulaRequestResponseFormat `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Formulas = all.Formulas
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
