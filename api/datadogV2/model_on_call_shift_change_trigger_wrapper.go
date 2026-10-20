// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallShiftChangeTriggerWrapper Schema for an On-Call shift change trigger.
type OnCallShiftChangeTriggerWrapper struct {
	// Trigger a workflow when an On-Call shift changes. The workflow must be published.
	OnCallShiftChangeTrigger OnCallShiftChangeTrigger `json:"onCallShiftChangeTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallShiftChangeTriggerWrapper instantiates a new OnCallShiftChangeTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallShiftChangeTriggerWrapper(onCallShiftChangeTrigger OnCallShiftChangeTrigger) *OnCallShiftChangeTriggerWrapper {
	this := OnCallShiftChangeTriggerWrapper{}
	this.OnCallShiftChangeTrigger = onCallShiftChangeTrigger
	return &this
}

// NewOnCallShiftChangeTriggerWrapperWithDefaults instantiates a new OnCallShiftChangeTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallShiftChangeTriggerWrapperWithDefaults() *OnCallShiftChangeTriggerWrapper {
	this := OnCallShiftChangeTriggerWrapper{}
	return &this
}

// GetOnCallShiftChangeTrigger returns the OnCallShiftChangeTrigger field value.
func (o *OnCallShiftChangeTriggerWrapper) GetOnCallShiftChangeTrigger() OnCallShiftChangeTrigger {
	if o == nil {
		var ret OnCallShiftChangeTrigger
		return ret
	}
	return o.OnCallShiftChangeTrigger
}

// GetOnCallShiftChangeTriggerOk returns a tuple with the OnCallShiftChangeTrigger field value
// and a boolean to check if the value has been set.
func (o *OnCallShiftChangeTriggerWrapper) GetOnCallShiftChangeTriggerOk() (*OnCallShiftChangeTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnCallShiftChangeTrigger, true
}

// SetOnCallShiftChangeTrigger sets field value.
func (o *OnCallShiftChangeTriggerWrapper) SetOnCallShiftChangeTrigger(v OnCallShiftChangeTrigger) {
	o.OnCallShiftChangeTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *OnCallShiftChangeTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallShiftChangeTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *OnCallShiftChangeTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *OnCallShiftChangeTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallShiftChangeTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["onCallShiftChangeTrigger"] = o.OnCallShiftChangeTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallShiftChangeTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OnCallShiftChangeTrigger *OnCallShiftChangeTrigger `json:"onCallShiftChangeTrigger"`
		StartStepNames           []string                  `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OnCallShiftChangeTrigger == nil {
		return fmt.Errorf("required field onCallShiftChangeTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"onCallShiftChangeTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OnCallShiftChangeTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OnCallShiftChangeTrigger = *all.OnCallShiftChangeTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
