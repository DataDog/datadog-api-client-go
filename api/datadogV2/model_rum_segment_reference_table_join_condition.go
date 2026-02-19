// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentReferenceTableJoinCondition The join condition for a reference table query block.
type RumSegmentReferenceTableJoinCondition struct {
	// The reference table column to join on.
	ColumnName string `json:"column_name"`
	// The RUM facet to join on.
	Facet string `json:"facet"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentReferenceTableJoinCondition instantiates a new RumSegmentReferenceTableJoinCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentReferenceTableJoinCondition(columnName string, facet string) *RumSegmentReferenceTableJoinCondition {
	this := RumSegmentReferenceTableJoinCondition{}
	this.ColumnName = columnName
	this.Facet = facet
	return &this
}

// NewRumSegmentReferenceTableJoinConditionWithDefaults instantiates a new RumSegmentReferenceTableJoinCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentReferenceTableJoinConditionWithDefaults() *RumSegmentReferenceTableJoinCondition {
	this := RumSegmentReferenceTableJoinCondition{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *RumSegmentReferenceTableJoinCondition) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTableJoinCondition) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *RumSegmentReferenceTableJoinCondition) SetColumnName(v string) {
	o.ColumnName = v
}

// GetFacet returns the Facet field value.
func (o *RumSegmentReferenceTableJoinCondition) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTableJoinCondition) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *RumSegmentReferenceTableJoinCondition) SetFacet(v string) {
	o.Facet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentReferenceTableJoinCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	toSerialize["facet"] = o.Facet

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentReferenceTableJoinCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string `json:"column_name"`
		Facet      *string `json:"facet"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnName == nil {
		return fmt.Errorf("required field column_name missing")
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "facet"})
	} else {
		return err
	}
	o.ColumnName = *all.ColumnName
	o.Facet = *all.Facet

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
