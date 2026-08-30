// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderCreatedTriggerWrapper Schema for an incident responder created trigger.
type IncidentResponderCreatedTriggerWrapper struct {
	// Trigger a workflow when a responder is created for an incident.
	IncidentResponderCreatedTrigger IncidentResponderCreatedTrigger `json:"incidentResponderCreatedTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponderCreatedTriggerWrapper instantiates a new IncidentResponderCreatedTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponderCreatedTriggerWrapper(incidentResponderCreatedTrigger IncidentResponderCreatedTrigger) *IncidentResponderCreatedTriggerWrapper {
	this := IncidentResponderCreatedTriggerWrapper{}
	this.IncidentResponderCreatedTrigger = incidentResponderCreatedTrigger
	return &this
}

// NewIncidentResponderCreatedTriggerWrapperWithDefaults instantiates a new IncidentResponderCreatedTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponderCreatedTriggerWrapperWithDefaults() *IncidentResponderCreatedTriggerWrapper {
	this := IncidentResponderCreatedTriggerWrapper{}
	return &this
}

// GetIncidentResponderCreatedTrigger returns the IncidentResponderCreatedTrigger field value.
func (o *IncidentResponderCreatedTriggerWrapper) GetIncidentResponderCreatedTrigger() IncidentResponderCreatedTrigger {
	if o == nil {
		var ret IncidentResponderCreatedTrigger
		return ret
	}
	return o.IncidentResponderCreatedTrigger
}

// GetIncidentResponderCreatedTriggerOk returns a tuple with the IncidentResponderCreatedTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderCreatedTriggerWrapper) GetIncidentResponderCreatedTriggerOk() (*IncidentResponderCreatedTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentResponderCreatedTrigger, true
}

// SetIncidentResponderCreatedTrigger sets field value.
func (o *IncidentResponderCreatedTriggerWrapper) SetIncidentResponderCreatedTrigger(v IncidentResponderCreatedTrigger) {
	o.IncidentResponderCreatedTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentResponderCreatedTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponderCreatedTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentResponderCreatedTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentResponderCreatedTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponderCreatedTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentResponderCreatedTrigger"] = o.IncidentResponderCreatedTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponderCreatedTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentResponderCreatedTrigger *IncidentResponderCreatedTrigger `json:"incidentResponderCreatedTrigger"`
		StartStepNames                  []string                         `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentResponderCreatedTrigger == nil {
		return fmt.Errorf("required field incidentResponderCreatedTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentResponderCreatedTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentResponderCreatedTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentResponderCreatedTrigger = *all.IncidentResponderCreatedTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
