/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// DashboardListItems Dashboards within a list.
type DashboardListItems struct {
	// List of dashboards in the dashboard list.
	Dashboards []DashboardListItem `json:"dashboards"`
	// Number of dashboards in the dashboard list.
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

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *DashboardListItems) GetDashboardsOk() (*[]DashboardListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboards, true
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

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListItems) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
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

func (o DashboardListItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboards"] = o.Dashboards
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

func (o *DashboardListItems) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Dashboards *[]DashboardListItem `json:"dashboards"`
	}{}
	all := struct {
		Dashboards []DashboardListItem `json:"dashboards"`
		Total      *int64              `json:"total,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Dashboards == nil {
		return fmt.Errorf("Required field dashboards missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Dashboards = all.Dashboards
	o.Total = all.Total
	return nil
}

type NullableDashboardListItems struct {
	value *DashboardListItems
	isSet bool
}

func (v NullableDashboardListItems) Get() *DashboardListItems {
	return v.value
}

func (v *NullableDashboardListItems) Set(val *DashboardListItems) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardListItems) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardListItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardListItems(val *DashboardListItems) *NullableDashboardListItems {
	return &NullableDashboardListItems{value: val, isSet: true}
}

func (v NullableDashboardListItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardListItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
