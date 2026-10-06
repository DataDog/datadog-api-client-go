// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAlertTrigger A trigger created from a monitor alert.
type MonitorAlertTrigger struct {
	// Attributes for a monitor alert trigger.
	MonitorAlertTrigger MonitorAlertTriggerAttributes `json:"monitor_alert_trigger"`
	// The type of monitor alert trigger.
	Type MonitorAlertTriggerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorAlertTrigger instantiates a new MonitorAlertTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorAlertTrigger(monitorAlertTrigger MonitorAlertTriggerAttributes, typeVar MonitorAlertTriggerType) *MonitorAlertTrigger {
	this := MonitorAlertTrigger{}
	this.MonitorAlertTrigger = monitorAlertTrigger
	this.Type = typeVar
	return &this
}

// NewMonitorAlertTriggerWithDefaults instantiates a new MonitorAlertTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorAlertTriggerWithDefaults() *MonitorAlertTrigger {
	this := MonitorAlertTrigger{}
	return &this
}

// GetMonitorAlertTrigger returns the MonitorAlertTrigger field value.
func (o *MonitorAlertTrigger) GetMonitorAlertTrigger() MonitorAlertTriggerAttributes {
	if o == nil {
		var ret MonitorAlertTriggerAttributes
		return ret
	}
	return o.MonitorAlertTrigger
}

// GetMonitorAlertTriggerOk returns a tuple with the MonitorAlertTrigger field value
// and a boolean to check if the value has been set.
func (o *MonitorAlertTrigger) GetMonitorAlertTriggerOk() (*MonitorAlertTriggerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorAlertTrigger, true
}

// SetMonitorAlertTrigger sets field value.
func (o *MonitorAlertTrigger) SetMonitorAlertTrigger(v MonitorAlertTriggerAttributes) {
	o.MonitorAlertTrigger = v
}

// GetType returns the Type field value.
func (o *MonitorAlertTrigger) GetType() MonitorAlertTriggerType {
	if o == nil {
		var ret MonitorAlertTriggerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorAlertTrigger) GetTypeOk() (*MonitorAlertTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MonitorAlertTrigger) SetType(v MonitorAlertTriggerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorAlertTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_alert_trigger"] = o.MonitorAlertTrigger
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorAlertTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorAlertTrigger *MonitorAlertTriggerAttributes `json:"monitor_alert_trigger"`
		Type                *MonitorAlertTriggerType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorAlertTrigger == nil {
		return fmt.Errorf("required field monitor_alert_trigger missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.MonitorAlertTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorAlertTrigger = *all.MonitorAlertTrigger
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
