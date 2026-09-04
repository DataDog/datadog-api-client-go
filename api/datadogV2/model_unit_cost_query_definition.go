// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnitCostQueryDefinition A timeseries object containing `queries` and `formulas` arrays.
type UnitCostQueryDefinition struct {
	// The list of formulas applied to the queries for this side of the ratio.
	Formulas []map[string]interface{} `json:"formulas"`
	// The list of queries evaluated for this side of the ratio.
	Queries []map[string]interface{} `json:"queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUnitCostQueryDefinition instantiates a new UnitCostQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnitCostQueryDefinition(formulas []map[string]interface{}, queries []map[string]interface{}) *UnitCostQueryDefinition {
	this := UnitCostQueryDefinition{}
	this.Formulas = formulas
	this.Queries = queries
	return &this
}

// NewUnitCostQueryDefinitionWithDefaults instantiates a new UnitCostQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnitCostQueryDefinitionWithDefaults() *UnitCostQueryDefinition {
	this := UnitCostQueryDefinition{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *UnitCostQueryDefinition) GetFormulas() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *UnitCostQueryDefinition) GetFormulasOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *UnitCostQueryDefinition) SetFormulas(v []map[string]interface{}) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *UnitCostQueryDefinition) GetQueries() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *UnitCostQueryDefinition) GetQueriesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *UnitCostQueryDefinition) SetQueries(v []map[string]interface{}) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UnitCostQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formulas"] = o.Formulas
	toSerialize["queries"] = o.Queries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UnitCostQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas *[]map[string]interface{} `json:"formulas"`
		Queries  *[]map[string]interface{} `json:"queries"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "queries"})
	} else {
		return err
	}
	o.Formulas = *all.Formulas
	o.Queries = *all.Queries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
