/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// DashboardListItems struct for DashboardListItems
type DashboardListItems struct {
	// List of dashboards in the dashboard list
	Dashboards []DashboardListItem `json:"dashboards"`
	// Number of dashboards in the dashboard list
	Total *int64 `json:"total,omitempty"`
}

// NewDashboardListItems instantiates a new DashboardListItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardListItems(dashboards []DashboardListItem) *DashboardListItems {
	this := DashboardListItems{}
	this.Dashboards = dashboards
	return &this
}

// NewDashboardListItemsWithDefaults instantiates a new DashboardListItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardListItemsWithDefaults() *DashboardListItems {
	this := DashboardListItems{}
	return &this
}

// GetDashboards returns the Dashboards field value
func (o *DashboardListItems) GetDashboards() []DashboardListItem {
	if o == nil {
		var ret []DashboardListItem
		return ret
	}

	return o.Dashboards
}

// SetDashboards sets field value
func (o *DashboardListItems) SetDashboards(v []DashboardListItem) {
	o.Dashboards = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DashboardListItems) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItems) GetTotalOk() (int64, bool) {
	if o == nil || o.Total == nil {
		var ret int64
		return ret, false
	}
	return *o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DashboardListItems) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *DashboardListItems) SetTotal(v int64) {
	o.Total = &v
}

type NullableDashboardListItems struct {
	Value        DashboardListItems
	ExplicitNull bool
}

func (v NullableDashboardListItems) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDashboardListItems) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
