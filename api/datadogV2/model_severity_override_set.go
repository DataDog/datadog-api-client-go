// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideSet Applies a manual severity override to the findings.
type SeverityOverrideSet struct {
	// The action that applies a manual severity override.
	Action SeverityOverrideSetActionType `json:"action"`
	// Additional information about the severity change. This field has a limit of 280 characters.
	Description *string `json:"description,omitempty"`
	// Severity to apply to the findings.
	// `info` sets the lowest severity the finding type allows.
	Value SeverityOverrideValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideSet instantiates a new SeverityOverrideSet object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideSet(action SeverityOverrideSetActionType, value SeverityOverrideValue) *SeverityOverrideSet {
	this := SeverityOverrideSet{}
	this.Action = action
	this.Value = value
	return &this
}

// NewSeverityOverrideSetWithDefaults instantiates a new SeverityOverrideSet object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideSetWithDefaults() *SeverityOverrideSet {
	this := SeverityOverrideSet{}
	var action SeverityOverrideSetActionType = SEVERITYOVERRIDESETACTIONTYPE_SET
	this.Action = action
	return &this
}

// GetAction returns the Action field value.
func (o *SeverityOverrideSet) GetAction() SeverityOverrideSetActionType {
	if o == nil {
		var ret SeverityOverrideSetActionType
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideSet) GetActionOk() (*SeverityOverrideSetActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *SeverityOverrideSet) SetAction(v SeverityOverrideSetActionType) {
	o.Action = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SeverityOverrideSet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityOverrideSet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SeverityOverrideSet) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SeverityOverrideSet) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value.
func (o *SeverityOverrideSet) GetValue() SeverityOverrideValue {
	if o == nil {
		var ret SeverityOverrideValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideSet) GetValueOk() (*SeverityOverrideValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SeverityOverrideSet) SetValue(v SeverityOverrideValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideSet) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action      *SeverityOverrideSetActionType `json:"action"`
		Description *string                        `json:"description,omitempty"`
		Value       *SeverityOverrideValue         `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "description", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.Description = all.Description
	if !all.Value.IsValid() {
		hasInvalidField = true
	} else {
		o.Value = *all.Value
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
