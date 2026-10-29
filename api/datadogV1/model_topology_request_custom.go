// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestCustom Scalar request whose result rows become nodes.
type TopologyRequestCustom struct {
	// List of formulas that operate on queries.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Mapping from result columns to node fields. Each result row feeds one node.
	Projection TopologyRequestCustomProjection `json:"projection"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Response format of a custom topology map request.
	ResponseFormat TopologyRequestCustomResponseFormat `json:"response_format"`
	// The controls for sorting the widget.
	Sort *WidgetSortBy `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyRequestCustom instantiates a new TopologyRequestCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyRequestCustom(projection TopologyRequestCustomProjection, queries []FormulaAndFunctionQueryDefinition, responseFormat TopologyRequestCustomResponseFormat) *TopologyRequestCustom {
	this := TopologyRequestCustom{}
	this.Projection = projection
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewTopologyRequestCustomWithDefaults instantiates a new TopologyRequestCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyRequestCustomWithDefaults() *TopologyRequestCustom {
	this := TopologyRequestCustom{}
	var responseFormat TopologyRequestCustomResponseFormat = TOPOLOGYREQUESTCUSTOMRESPONSEFORMAT_SCALAR
	this.ResponseFormat = responseFormat
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *TopologyRequestCustom) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustom) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *TopologyRequestCustom) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *TopologyRequestCustom) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetProjection returns the Projection field value.
func (o *TopologyRequestCustom) GetProjection() TopologyRequestCustomProjection {
	if o == nil {
		var ret TopologyRequestCustomProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustom) GetProjectionOk() (*TopologyRequestCustomProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *TopologyRequestCustom) SetProjection(v TopologyRequestCustomProjection) {
	o.Projection = v
}

// GetQueries returns the Queries field value.
func (o *TopologyRequestCustom) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustom) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *TopologyRequestCustom) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TopologyRequestCustom) GetResponseFormat() TopologyRequestCustomResponseFormat {
	if o == nil {
		var ret TopologyRequestCustomResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustom) GetResponseFormatOk() (*TopologyRequestCustomResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TopologyRequestCustom) SetResponseFormat(v TopologyRequestCustomResponseFormat) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TopologyRequestCustom) GetSort() WidgetSortBy {
	if o == nil || o.Sort == nil {
		var ret WidgetSortBy
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustom) GetSortOk() (*WidgetSortBy, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TopologyRequestCustom) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WidgetSortBy and assigns it to the Sort field.
func (o *TopologyRequestCustom) SetSort(v WidgetSortBy) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyRequestCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	toSerialize["projection"] = o.Projection
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyRequestCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                      `json:"formulas,omitempty"`
		Projection     *TopologyRequestCustomProjection     `json:"projection"`
		Queries        *[]FormulaAndFunctionQueryDefinition `json:"queries"`
		ResponseFormat *TopologyRequestCustomResponseFormat `json:"response_format"`
		Sort           *WidgetSortBy                        `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Formulas = all.Formulas
	if all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = *all.Projection
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
