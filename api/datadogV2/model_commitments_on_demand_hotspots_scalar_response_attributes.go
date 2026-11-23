// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsOnDemandHotspotsScalarResponseAttributes Attributes of a scalar on-demand hot-spots response.
type CommitmentsOnDemandHotspotsScalarResponseAttributes struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// Array of scalar columns in the response.
	Total []CommitmentsScalarColumn `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsOnDemandHotspotsScalarResponseAttributes instantiates a new CommitmentsOnDemandHotspotsScalarResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsOnDemandHotspotsScalarResponseAttributes(columns []CommitmentsScalarColumn, total []CommitmentsScalarColumn) *CommitmentsOnDemandHotspotsScalarResponseAttributes {
	this := CommitmentsOnDemandHotspotsScalarResponseAttributes{}
	this.Columns = columns
	this.Total = total
	return &this
}

// NewCommitmentsOnDemandHotspotsScalarResponseAttributesWithDefaults instantiates a new CommitmentsOnDemandHotspotsScalarResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsOnDemandHotspotsScalarResponseAttributesWithDefaults() *CommitmentsOnDemandHotspotsScalarResponseAttributes {
	this := CommitmentsOnDemandHotspotsScalarResponseAttributes{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// GetTotal returns the Total field value.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetTotal() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) GetTotalOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) SetTotal(v []CommitmentsScalarColumn) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsOnDemandHotspotsScalarResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsOnDemandHotspotsScalarResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns         *[]CommitmentsScalarColumn  `json:"columns"`
		QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
		Total           *[]CommitmentsScalarColumn  `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "query_definition", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
