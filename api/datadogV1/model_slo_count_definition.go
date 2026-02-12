// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountDefinition A count-based (metric) SLI specification, composed of three parts: the good events formula, the total events formula,
// and the underlying queries.
type SLOCountDefinition struct {
	// A formula that specifies how to combine the results of multiple queries.
	GoodEventsFormula SLOFormula `json:"good_events_formula"`
	//
	Queries []SLODataSourceQueryDefinition `json:"queries"`
	// A formula that specifies how to combine the results of multiple queries.
	TotalEventsFormula SLOFormula `json:"total_events_formula"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOCountDefinition instantiates a new SLOCountDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCountDefinition(goodEventsFormula SLOFormula, queries []SLODataSourceQueryDefinition, totalEventsFormula SLOFormula) *SLOCountDefinition {
	this := SLOCountDefinition{}
	this.GoodEventsFormula = goodEventsFormula
	this.Queries = queries
	this.TotalEventsFormula = totalEventsFormula
	return &this
}

// NewSLOCountDefinitionWithDefaults instantiates a new SLOCountDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCountDefinitionWithDefaults() *SLOCountDefinition {
	this := SLOCountDefinition{}
	return &this
}

// GetGoodEventsFormula returns the GoodEventsFormula field value.
func (o *SLOCountDefinition) GetGoodEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.GoodEventsFormula
}

// GetGoodEventsFormulaOk returns a tuple with the GoodEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinition) GetGoodEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoodEventsFormula, true
}

// SetGoodEventsFormula sets field value.
func (o *SLOCountDefinition) SetGoodEventsFormula(v SLOFormula) {
	o.GoodEventsFormula = v
}

// GetQueries returns the Queries field value.
func (o *SLOCountDefinition) GetQueries() []SLODataSourceQueryDefinition {
	if o == nil {
		var ret []SLODataSourceQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinition) GetQueriesOk() (*[]SLODataSourceQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *SLOCountDefinition) SetQueries(v []SLODataSourceQueryDefinition) {
	o.Queries = v
}

// GetTotalEventsFormula returns the TotalEventsFormula field value.
func (o *SLOCountDefinition) GetTotalEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.TotalEventsFormula
}

// GetTotalEventsFormulaOk returns a tuple with the TotalEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinition) GetTotalEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalEventsFormula, true
}

// SetTotalEventsFormula sets field value.
func (o *SLOCountDefinition) SetTotalEventsFormula(v SLOFormula) {
	o.TotalEventsFormula = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCountDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["good_events_formula"] = o.GoodEventsFormula
	toSerialize["queries"] = o.Queries
	toSerialize["total_events_formula"] = o.TotalEventsFormula

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCountDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GoodEventsFormula  *SLOFormula                     `json:"good_events_formula"`
		Queries            *[]SLODataSourceQueryDefinition `json:"queries"`
		TotalEventsFormula *SLOFormula                     `json:"total_events_formula"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GoodEventsFormula == nil {
		return fmt.Errorf("required field good_events_formula missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.TotalEventsFormula == nil {
		return fmt.Errorf("required field total_events_formula missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"good_events_formula", "queries", "total_events_formula"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GoodEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GoodEventsFormula = *all.GoodEventsFormula
	o.Queries = *all.Queries
	if all.TotalEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalEventsFormula = *all.TotalEventsFormula

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
