// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRows Defines how to compute the rows of a guided table.
type GuidedTableRows struct {
	// Grouping dimensions that form each row. Must be non-empty.
	GroupBy []GuidedTableRowGroupBy `json:"group_by"`
	// Sort configuration for a guided table.
	Sort GuidedTableSort `json:"sort"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableRows instantiates a new GuidedTableRows object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableRows(groupBy []GuidedTableRowGroupBy, sort GuidedTableSort) *GuidedTableRows {
	this := GuidedTableRows{}
	this.GroupBy = groupBy
	this.Sort = sort
	return &this
}

// NewGuidedTableRowsWithDefaults instantiates a new GuidedTableRows object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableRowsWithDefaults() *GuidedTableRows {
	this := GuidedTableRows{}
	return &this
}

// GetGroupBy returns the GroupBy field value.
func (o *GuidedTableRows) GetGroupBy() []GuidedTableRowGroupBy {
	if o == nil {
		var ret []GuidedTableRowGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRows) GetGroupByOk() (*[]GuidedTableRowGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *GuidedTableRows) SetGroupBy(v []GuidedTableRowGroupBy) {
	o.GroupBy = v
}

// GetSort returns the Sort field value.
func (o *GuidedTableRows) GetSort() GuidedTableSort {
	if o == nil {
		var ret GuidedTableSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRows) GetSortOk() (*GuidedTableSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *GuidedTableRows) SetSort(v GuidedTableSort) {
	o.Sort = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableRows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_by"] = o.GroupBy
	toSerialize["sort"] = o.Sort

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableRows) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy *[]GuidedTableRowGroupBy `json:"group_by"`
		Sort    *GuidedTableSort         `json:"sort"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.Sort == nil {
		return fmt.Errorf("required field sort missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_by", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GroupBy = *all.GroupBy
	if all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = *all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
