/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
)

// CanceledDowntimesIds struct for CanceledDowntimesIds
type CanceledDowntimesIds struct {
	CancelledIds *[]int32 `json:"cancelled_ids,omitempty"`
}

// GetCancelledIds returns the CancelledIds field if non-nil, zero value otherwise.
func (o *CanceledDowntimesIds) GetCancelledIds() []int32 {
	if o == nil || o.CancelledIds == nil {
		var ret []int32
		return ret
	}
	return *o.CancelledIds
}

// GetCancelledIdsOk returns a tuple with the CancelledIds field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CanceledDowntimesIds) GetCancelledIdsOk() ([]int32, bool) {
	if o == nil || o.CancelledIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.CancelledIds, true
}

// HasCancelledIds returns a boolean if a field has been set.
func (o *CanceledDowntimesIds) HasCancelledIds() bool {
	if o != nil && o.CancelledIds != nil {
		return true
	}

	return false
}

// SetCancelledIds gets a reference to the given []int32 and assigns it to the CancelledIds field.
func (o *CanceledDowntimesIds) SetCancelledIds(v []int32) {
	o.CancelledIds = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o CanceledDowntimesIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CancelledIds != nil {
		toSerialize["cancelled_ids"] = o.CancelledIds
	}
	return json.Marshal(toSerialize)
}
