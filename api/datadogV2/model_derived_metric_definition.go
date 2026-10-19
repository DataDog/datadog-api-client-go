// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DerivedMetricDefinition The definition of a derived metric as a formula over one or more queries.
type DerivedMetricDefinition struct {
	// The formula that combines the queries into the value of the derived metric.
	Formula string `json:"formula"`
	// The list of queries referenced by the formula.
	Queries []DerivedMetricQuery `json:"queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDerivedMetricDefinition instantiates a new DerivedMetricDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDerivedMetricDefinition(formula string, queries []DerivedMetricQuery) *DerivedMetricDefinition {
	this := DerivedMetricDefinition{}
	this.Formula = formula
	this.Queries = queries
	return &this
}

// NewDerivedMetricDefinitionWithDefaults instantiates a new DerivedMetricDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDerivedMetricDefinitionWithDefaults() *DerivedMetricDefinition {
	this := DerivedMetricDefinition{}
	return &this
}

// GetFormula returns the Formula field value.
func (o *DerivedMetricDefinition) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricDefinition) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *DerivedMetricDefinition) SetFormula(v string) {
	o.Formula = v
}

// GetQueries returns the Queries field value.
func (o *DerivedMetricDefinition) GetQueries() []DerivedMetricQuery {
	if o == nil {
		var ret []DerivedMetricQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricDefinition) GetQueriesOk() (*[]DerivedMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *DerivedMetricDefinition) SetQueries(v []DerivedMetricQuery) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DerivedMetricDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formula"] = o.Formula
	toSerialize["queries"] = o.Queries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DerivedMetricDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formula *string               `json:"formula"`
		Queries *[]DerivedMetricQuery `json:"queries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formula", "queries"})
	} else {
		return err
	}
	o.Formula = *all.Formula
	o.Queries = *all.Queries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
