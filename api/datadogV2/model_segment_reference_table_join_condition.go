// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentReferenceTableJoinCondition How reference table rows are matched to entities.
type SegmentReferenceTableJoinCondition struct {
	// Reference table column that holds the entity ID.
	ColumnName string `json:"column_name"`
	// Facet that identifies the entity, such as `@usr.id`.
	Facet *string `json:"facet,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentReferenceTableJoinCondition instantiates a new SegmentReferenceTableJoinCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentReferenceTableJoinCondition(columnName string) *SegmentReferenceTableJoinCondition {
	this := SegmentReferenceTableJoinCondition{}
	this.ColumnName = columnName
	return &this
}

// NewSegmentReferenceTableJoinConditionWithDefaults instantiates a new SegmentReferenceTableJoinCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentReferenceTableJoinConditionWithDefaults() *SegmentReferenceTableJoinCondition {
	this := SegmentReferenceTableJoinCondition{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *SegmentReferenceTableJoinCondition) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableJoinCondition) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *SegmentReferenceTableJoinCondition) SetColumnName(v string) {
	o.ColumnName = v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *SegmentReferenceTableJoinCondition) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableJoinCondition) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *SegmentReferenceTableJoinCondition) HasFacet() bool {
	return o != nil && o.Facet != nil
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *SegmentReferenceTableJoinCondition) SetFacet(v string) {
	o.Facet = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentReferenceTableJoinCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentReferenceTableJoinCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string `json:"column_name"`
		Facet      *string `json:"facet,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnName == nil {
		return fmt.Errorf("required field column_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "facet"})
	} else {
		return err
	}
	o.ColumnName = *all.ColumnName
	o.Facet = all.Facet

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
