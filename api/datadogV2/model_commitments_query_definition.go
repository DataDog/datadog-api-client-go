// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsQueryDefinition The query that produces the response values, in the same format as dashboard and notebook widget
// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
// that it produces.
type CommitmentsQueryDefinition struct {
	// The formulas of a query definition.
	Formulas []CommitmentsQueryDefinitionFormula `json:"formulas"`
	// The metric queries of a query definition.
	Queries []CommitmentsQueryDefinitionQuery `json:"queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsQueryDefinition instantiates a new CommitmentsQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsQueryDefinition(formulas []CommitmentsQueryDefinitionFormula, queries []CommitmentsQueryDefinitionQuery) *CommitmentsQueryDefinition {
	this := CommitmentsQueryDefinition{}
	this.Formulas = formulas
	this.Queries = queries
	return &this
}

// NewCommitmentsQueryDefinitionWithDefaults instantiates a new CommitmentsQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsQueryDefinitionWithDefaults() *CommitmentsQueryDefinition {
	this := CommitmentsQueryDefinition{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *CommitmentsQueryDefinition) GetFormulas() []CommitmentsQueryDefinitionFormula {
	if o == nil {
		var ret []CommitmentsQueryDefinitionFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinition) GetFormulasOk() (*[]CommitmentsQueryDefinitionFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *CommitmentsQueryDefinition) SetFormulas(v []CommitmentsQueryDefinitionFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *CommitmentsQueryDefinition) GetQueries() []CommitmentsQueryDefinitionQuery {
	if o == nil {
		var ret []CommitmentsQueryDefinitionQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinition) GetQueriesOk() (*[]CommitmentsQueryDefinitionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *CommitmentsQueryDefinition) SetQueries(v []CommitmentsQueryDefinitionQuery) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsQueryDefinition) MarshalJSON() ([]byte, error) {
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
func (o *CommitmentsQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas *[]CommitmentsQueryDefinitionFormula `json:"formulas"`
		Queries  *[]CommitmentsQueryDefinitionQuery   `json:"queries"`
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
