// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTileDefRequest Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridTileDefRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat HeatgridTileDefRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridTileDefRequest instantiates a new HeatgridTileDefRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridTileDefRequest(queries []FormulaAndFunctionQueryDefinition, responseFormat HeatgridTileDefRequestResponseFormat) *HeatgridTileDefRequest {
	this := HeatgridTileDefRequest{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewHeatgridTileDefRequestWithDefaults instantiates a new HeatgridTileDefRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridTileDefRequestWithDefaults() *HeatgridTileDefRequest {
	this := HeatgridTileDefRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *HeatgridTileDefRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridTileDefRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *HeatgridTileDefRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *HeatgridTileDefRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *HeatgridTileDefRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *HeatgridTileDefRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *HeatgridTileDefRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *HeatgridTileDefRequest) GetResponseFormat() HeatgridTileDefRequestResponseFormat {
	if o == nil {
		var ret HeatgridTileDefRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *HeatgridTileDefRequest) GetResponseFormatOk() (*HeatgridTileDefRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *HeatgridTileDefRequest) SetResponseFormat(v HeatgridTileDefRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridTileDefRequest) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridTileDefRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                       `json:"formulas,omitempty"`
		Queries        *[]FormulaAndFunctionQueryDefinition  `json:"queries"`
		ResponseFormat *HeatgridTileDefRequestResponseFormat `json:"response_format"`
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
