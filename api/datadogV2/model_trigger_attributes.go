// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerAttributes The trigger definition for starting an investigation.
type TriggerAttributes struct {
	// Attributes for a general investigation, not tied to a specific monitor alert.
	GeneralInvestigation *GeneralInvestigationAttributes `json:"general_investigation,omitempty"`
	// Attributes for a monitor alert trigger.
	MonitorAlertTrigger *MonitorAlertTriggerAttributes `json:"monitor_alert_trigger,omitempty"`
	// The type of trigger for the investigation.
	Type *TriggerType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTriggerAttributes instantiates a new TriggerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTriggerAttributes() *TriggerAttributes {
	this := TriggerAttributes{}
	return &this
}

// NewTriggerAttributesWithDefaults instantiates a new TriggerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTriggerAttributesWithDefaults() *TriggerAttributes {
	this := TriggerAttributes{}
	return &this
}

// GetGeneralInvestigation returns the GeneralInvestigation field value if set, zero value otherwise.
func (o *TriggerAttributes) GetGeneralInvestigation() GeneralInvestigationAttributes {
	if o == nil || o.GeneralInvestigation == nil {
		var ret GeneralInvestigationAttributes
		return ret
	}
	return *o.GeneralInvestigation
}

// GetGeneralInvestigationOk returns a tuple with the GeneralInvestigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerAttributes) GetGeneralInvestigationOk() (*GeneralInvestigationAttributes, bool) {
	if o == nil || o.GeneralInvestigation == nil {
		return nil, false
	}
	return o.GeneralInvestigation, true
}

// HasGeneralInvestigation returns a boolean if a field has been set.
func (o *TriggerAttributes) HasGeneralInvestigation() bool {
	return o != nil && o.GeneralInvestigation != nil
}

// SetGeneralInvestigation gets a reference to the given GeneralInvestigationAttributes and assigns it to the GeneralInvestigation field.
func (o *TriggerAttributes) SetGeneralInvestigation(v GeneralInvestigationAttributes) {
	o.GeneralInvestigation = &v
}

// GetMonitorAlertTrigger returns the MonitorAlertTrigger field value if set, zero value otherwise.
func (o *TriggerAttributes) GetMonitorAlertTrigger() MonitorAlertTriggerAttributes {
	if o == nil || o.MonitorAlertTrigger == nil {
		var ret MonitorAlertTriggerAttributes
		return ret
	}
	return *o.MonitorAlertTrigger
}

// GetMonitorAlertTriggerOk returns a tuple with the MonitorAlertTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerAttributes) GetMonitorAlertTriggerOk() (*MonitorAlertTriggerAttributes, bool) {
	if o == nil || o.MonitorAlertTrigger == nil {
		return nil, false
	}
	return o.MonitorAlertTrigger, true
}

// HasMonitorAlertTrigger returns a boolean if a field has been set.
func (o *TriggerAttributes) HasMonitorAlertTrigger() bool {
	return o != nil && o.MonitorAlertTrigger != nil
}

// SetMonitorAlertTrigger gets a reference to the given MonitorAlertTriggerAttributes and assigns it to the MonitorAlertTrigger field.
func (o *TriggerAttributes) SetMonitorAlertTrigger(v MonitorAlertTriggerAttributes) {
	o.MonitorAlertTrigger = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TriggerAttributes) GetType() TriggerType {
	if o == nil || o.Type == nil {
		var ret TriggerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerAttributes) GetTypeOk() (*TriggerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TriggerAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given TriggerType and assigns it to the Type field.
func (o *TriggerAttributes) SetType(v TriggerType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TriggerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GeneralInvestigation != nil {
		toSerialize["general_investigation"] = o.GeneralInvestigation
	}
	if o.MonitorAlertTrigger != nil {
		toSerialize["monitor_alert_trigger"] = o.MonitorAlertTrigger
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TriggerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GeneralInvestigation *GeneralInvestigationAttributes `json:"general_investigation,omitempty"`
		MonitorAlertTrigger  *MonitorAlertTriggerAttributes  `json:"monitor_alert_trigger,omitempty"`
		Type                 *TriggerType                    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"general_investigation", "monitor_alert_trigger", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GeneralInvestigation = all.GeneralInvestigation
	if all.MonitorAlertTrigger != nil && all.MonitorAlertTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorAlertTrigger = all.MonitorAlertTrigger
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
