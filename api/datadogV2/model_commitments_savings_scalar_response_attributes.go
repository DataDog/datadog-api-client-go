// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsSavingsScalarResponseAttributes Attributes of a scalar savings response.
type CommitmentsSavingsScalarResponseAttributes struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// Metadata reporting that the cost data lacks the unused costs of some commitments, as with Azure exports
	// scoped below the billing profile. Utilization is then overstated, and realized savings and effective
	// savings rate are upper bounds. Omitted when the unused costs are available.
	Meta *CommitmentsUnusedCostMeta `json:"meta,omitempty"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsSavingsScalarResponseAttributes instantiates a new CommitmentsSavingsScalarResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsSavingsScalarResponseAttributes(columns []CommitmentsScalarColumn) *CommitmentsSavingsScalarResponseAttributes {
	this := CommitmentsSavingsScalarResponseAttributes{}
	this.Columns = columns
	return &this
}

// NewCommitmentsSavingsScalarResponseAttributesWithDefaults instantiates a new CommitmentsSavingsScalarResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsSavingsScalarResponseAttributesWithDefaults() *CommitmentsSavingsScalarResponseAttributes {
	this := CommitmentsSavingsScalarResponseAttributes{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsSavingsScalarResponseAttributes) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsScalarResponseAttributes) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsSavingsScalarResponseAttributes) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsSavingsScalarResponseAttributes) GetMeta() CommitmentsUnusedCostMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsUnusedCostMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsScalarResponseAttributes) GetMetaOk() (*CommitmentsUnusedCostMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsSavingsScalarResponseAttributes) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsUnusedCostMeta and assigns it to the Meta field.
func (o *CommitmentsSavingsScalarResponseAttributes) SetMeta(v CommitmentsUnusedCostMeta) {
	o.Meta = &v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsSavingsScalarResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsScalarResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsSavingsScalarResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsSavingsScalarResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsSavingsScalarResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsSavingsScalarResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns         *[]CommitmentsScalarColumn  `json:"columns"`
		Meta            *CommitmentsUnusedCostMeta  `json:"meta,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "meta", "query_definition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
