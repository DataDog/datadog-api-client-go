/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
	"errors"
)

// ServiceLevelObjectiveDeleted struct for ServiceLevelObjectiveDeleted
type ServiceLevelObjectiveDeleted struct {
	// An array containing the ID of the deleted service level objective object.
	Data *[]string `json:"data,omitempty"`
}

// GetData returns the Data field if non-nil, zero value otherwise.
func (o *ServiceLevelObjectiveDeleted) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveDeleted) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		var ret []string
		return ret, false
	}
	return *o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceLevelObjectiveDeleted) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ServiceLevelObjectiveDeleted) SetData(v []string) {
	o.Data = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o ServiceLevelObjectiveDeleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data == nil {
		return nil, errors.New("Data is required and not nullable, but was not set on ServiceLevelObjectiveDeleted")
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}
