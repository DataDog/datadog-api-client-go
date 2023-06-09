// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeAttributeMonitorIdentifierId Object of the monitor identifier.
type DowntimeAttributeMonitorIdentifierId struct {
	// ID of the monitor to prevent notifications.
	MonitorId int64 `json:"monitor_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDowntimeAttributeMonitorIdentifierId instantiates a new DowntimeAttributeMonitorIdentifierId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeMonitorIdentifierId(monitorId int64) *DowntimeAttributeMonitorIdentifierId {
	this := DowntimeAttributeMonitorIdentifierId{}
	this.MonitorId = monitorId
	return &this
}

// NewDowntimeAttributeMonitorIdentifierIdWithDefaults instantiates a new DowntimeAttributeMonitorIdentifierId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeMonitorIdentifierIdWithDefaults() *DowntimeAttributeMonitorIdentifierId {
	this := DowntimeAttributeMonitorIdentifierId{}
	return &this
}

// GetMonitorId returns the MonitorId field value.
func (o *DowntimeAttributeMonitorIdentifierId) GetMonitorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeMonitorIdentifierId) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorId, true
}

// SetMonitorId sets field value.
func (o *DowntimeAttributeMonitorIdentifierId) SetMonitorId(v int64) {
	o.MonitorId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeMonitorIdentifierId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_id"] = o.MonitorId
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeAttributeMonitorIdentifierId) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		MonitorId *int64 `json:"monitor_id"`
	}{}
	all := struct {
		MonitorId int64 `json:"monitor_id"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.MonitorId == nil {
		return fmt.Errorf("required field monitor_id missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.MonitorId = all.MonitorId

	return nil
}
