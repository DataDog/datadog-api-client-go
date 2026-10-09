// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridWidgetRequest A request for a heatgrid widget that uses formulas and functions.
type HeatgridWidgetRequest struct {
	// The single displayed formula can combine multiple queries.
	Formulas []HeatgridWidgetFormula `json:"formulas,omitempty"`
	// Queries returned directly or combined in a formula.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Response format for heatgrid queries.
	ResponseFormat HeatgridWidgetResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridWidgetRequest instantiates a new HeatgridWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridWidgetRequest(queries []FormulaAndFunctionQueryDefinition, responseFormat HeatgridWidgetResponseFormat) *HeatgridWidgetRequest {
	this := HeatgridWidgetRequest{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewHeatgridWidgetRequestWithDefaults instantiates a new HeatgridWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridWidgetRequestWithDefaults() *HeatgridWidgetRequest {
	this := HeatgridWidgetRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *HeatgridWidgetRequest) GetFormulas() []HeatgridWidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []HeatgridWidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetRequest) GetFormulasOk() (*[]HeatgridWidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *HeatgridWidgetRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []HeatgridWidgetFormula and assigns it to the Formulas field.
func (o *HeatgridWidgetRequest) SetFormulas(v []HeatgridWidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *HeatgridWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *HeatgridWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *HeatgridWidgetRequest) GetResponseFormat() HeatgridWidgetResponseFormat {
	if o == nil {
		var ret HeatgridWidgetResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetRequest) GetResponseFormatOk() (*HeatgridWidgetResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *HeatgridWidgetRequest) SetResponseFormat(v HeatgridWidgetResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridWidgetRequest) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []HeatgridWidgetFormula              `json:"formulas,omitempty"`
		Queries        *[]FormulaAndFunctionQueryDefinition `json:"queries"`
		ResponseFormat *HeatgridWidgetResponseFormat        `json:"response_format"`
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
