/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// MonitorGroupSearchResponseCounts The counts of monitor groups per different criteria.
type MonitorGroupSearchResponseCounts struct {
	// Search facets.
	Status *[]interface{} `json:"status,omitempty"`
	// Search facets.
	Type *[]interface{} `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewMonitorGroupSearchResponseCounts instantiates a new MonitorGroupSearchResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorGroupSearchResponseCounts() *MonitorGroupSearchResponseCounts {
	this := MonitorGroupSearchResponseCounts{}
	return &this
}

// NewMonitorGroupSearchResponseCountsWithDefaults instantiates a new MonitorGroupSearchResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorGroupSearchResponseCountsWithDefaults() *MonitorGroupSearchResponseCounts {
	this := MonitorGroupSearchResponseCounts{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorGroupSearchResponseCounts) GetStatus() []interface{} {
	if o == nil || o.Status == nil {
		var ret []interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroupSearchResponseCounts) GetStatusOk() (*[]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCounts) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []interface{} and assigns it to the Status field.
func (o *MonitorGroupSearchResponseCounts) SetStatus(v []interface{}) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorGroupSearchResponseCounts) GetType() []interface{} {
	if o == nil || o.Type == nil {
		var ret []interface{}
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroupSearchResponseCounts) GetTypeOk() (*[]interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCounts) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []interface{} and assigns it to the Type field.
func (o *MonitorGroupSearchResponseCounts) SetType(v []interface{}) {
	o.Type = &v
}

func (o MonitorGroupSearchResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *MonitorGroupSearchResponseCounts) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Status *[]interface{} `json:"status,omitempty"`
		Type   *[]interface{} `json:"type,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Status = all.Status
	o.Type = all.Type
	return nil
}
