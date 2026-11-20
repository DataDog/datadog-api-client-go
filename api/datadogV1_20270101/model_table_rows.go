// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRows Test-only generated schema for the dashboard versioning proof of concept.
type TableRows struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []TableRowGroupBy `json:"group_by"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinType *TableRowsJoinType `json:"join_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort TableSort `json:"sort"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableRows instantiates a new TableRows object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRows(groupBy []TableRowGroupBy, sort TableSort) *TableRows {
	this := TableRows{}
	this.GroupBy = groupBy
	this.Sort = sort
	return &this
}

// NewTableRowsWithDefaults instantiates a new TableRows object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowsWithDefaults() *TableRows {
	this := TableRows{}
	return &this
}

// GetGroupBy returns the GroupBy field value.
func (o *TableRows) GetGroupBy() []TableRowGroupBy {
	if o == nil {
		var ret []TableRowGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *TableRows) GetGroupByOk() (*[]TableRowGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *TableRows) SetGroupBy(v []TableRowGroupBy) {
	o.GroupBy = v
}

// GetJoinType returns the JoinType field value if set, zero value otherwise.
func (o *TableRows) GetJoinType() TableRowsJoinType {
	if o == nil || o.JoinType == nil {
		var ret TableRowsJoinType
		return ret
	}
	return *o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRows) GetJoinTypeOk() (*TableRowsJoinType, bool) {
	if o == nil || o.JoinType == nil {
		return nil, false
	}
	return o.JoinType, true
}

// HasJoinType returns a boolean if a field has been set.
func (o *TableRows) HasJoinType() bool {
	return o != nil && o.JoinType != nil
}

// SetJoinType gets a reference to the given TableRowsJoinType and assigns it to the JoinType field.
func (o *TableRows) SetJoinType(v TableRowsJoinType) {
	o.JoinType = &v
}

// GetSort returns the Sort field value.
func (o *TableRows) GetSort() TableSort {
	if o == nil {
		var ret TableSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *TableRows) GetSortOk() (*TableSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *TableRows) SetSort(v TableSort) {
	o.Sort = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_by"] = o.GroupBy
	if o.JoinType != nil {
		toSerialize["join_type"] = o.JoinType
	}
	toSerialize["sort"] = o.Sort
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableRows) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy  *[]TableRowGroupBy `json:"group_by"`
		JoinType *TableRowsJoinType `json:"join_type,omitempty"`
		Sort     *TableSort         `json:"sort"`
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

	hasInvalidField := false
	o.GroupBy = *all.GroupBy
	if all.JoinType != nil && !all.JoinType.IsValid() {
		hasInvalidField = true
	} else {
		o.JoinType = all.JoinType
	}
	if all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = *all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
