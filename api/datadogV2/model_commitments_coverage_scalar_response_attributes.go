// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageScalarResponseAttributes Attributes of a scalar coverage response.
type CommitmentsCoverageScalarResponseAttributes struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageScalarResponseAttributes instantiates a new CommitmentsCoverageScalarResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageScalarResponseAttributes(columns []CommitmentsScalarColumn) *CommitmentsCoverageScalarResponseAttributes {
	this := CommitmentsCoverageScalarResponseAttributes{}
	this.Columns = columns
	return &this
}

// NewCommitmentsCoverageScalarResponseAttributesWithDefaults instantiates a new CommitmentsCoverageScalarResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageScalarResponseAttributesWithDefaults() *CommitmentsCoverageScalarResponseAttributes {
	this := CommitmentsCoverageScalarResponseAttributes{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsCoverageScalarResponseAttributes) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponseAttributes) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsCoverageScalarResponseAttributes) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsCoverageScalarResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsCoverageScalarResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsCoverageScalarResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageScalarResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageScalarResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns         *[]CommitmentsScalarColumn  `json:"columns"`
		QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "query_definition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
