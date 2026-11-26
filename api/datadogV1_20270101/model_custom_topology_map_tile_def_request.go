// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapTileDefRequest Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapTileDefRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection CustomTopologyMapProjection `json:"projection"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat CustomTopologyMapTileDefRequestResponseFormat `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapTileDefRequest instantiates a new CustomTopologyMapTileDefRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapTileDefRequest(projection CustomTopologyMapProjection, queries []FormulaAndFunctionQueryDefinition, responseFormat CustomTopologyMapTileDefRequestResponseFormat) *CustomTopologyMapTileDefRequest {
	this := CustomTopologyMapTileDefRequest{}
	this.Projection = projection
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewCustomTopologyMapTileDefRequestWithDefaults instantiates a new CustomTopologyMapTileDefRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapTileDefRequestWithDefaults() *CustomTopologyMapTileDefRequest {
	this := CustomTopologyMapTileDefRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *CustomTopologyMapTileDefRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapTileDefRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *CustomTopologyMapTileDefRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *CustomTopologyMapTileDefRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetProjection returns the Projection field value.
func (o *CustomTopologyMapTileDefRequest) GetProjection() CustomTopologyMapProjection {
	if o == nil {
		var ret CustomTopologyMapProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapTileDefRequest) GetProjectionOk() (*CustomTopologyMapProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *CustomTopologyMapTileDefRequest) SetProjection(v CustomTopologyMapProjection) {
	o.Projection = v
}

// GetQueries returns the Queries field value.
func (o *CustomTopologyMapTileDefRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapTileDefRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *CustomTopologyMapTileDefRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *CustomTopologyMapTileDefRequest) GetResponseFormat() CustomTopologyMapTileDefRequestResponseFormat {
	if o == nil {
		var ret CustomTopologyMapTileDefRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapTileDefRequest) GetResponseFormatOk() (*CustomTopologyMapTileDefRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *CustomTopologyMapTileDefRequest) SetResponseFormat(v CustomTopologyMapTileDefRequestResponseFormat) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CustomTopologyMapTileDefRequest) GetSort() ScalarFormulaSort {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapTileDefRequest) GetSortOk() (*ScalarFormulaSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CustomTopologyMapTileDefRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort and assigns it to the Sort field.
func (o *CustomTopologyMapTileDefRequest) SetSort(v ScalarFormulaSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapTileDefRequest) MarshalJSON() ([]byte, error) {
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
func (o *CustomTopologyMapTileDefRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                                `json:"formulas,omitempty"`
		Projection     *CustomTopologyMapProjection                   `json:"projection"`
		Queries        *[]FormulaAndFunctionQueryDefinition           `json:"queries"`
		ResponseFormat *CustomTopologyMapTileDefRequestResponseFormat `json:"response_format"`
		Sort           *ScalarFormulaSort                             `json:"sort,omitempty"`
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
