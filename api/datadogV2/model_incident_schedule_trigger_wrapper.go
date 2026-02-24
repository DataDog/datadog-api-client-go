// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentScheduleTriggerWrapper Schema for an Incident Schedule-based trigger.
type IncidentScheduleTriggerWrapper struct {
	// Trigger a workflow from an Incident Schedule. The workflow must be published.
	IncidentScheduleTrigger IncidentScheduleTrigger `json:"incidentScheduleTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentScheduleTriggerWrapper instantiates a new IncidentScheduleTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentScheduleTriggerWrapper(incidentScheduleTrigger IncidentScheduleTrigger) *IncidentScheduleTriggerWrapper {
	this := IncidentScheduleTriggerWrapper{}
	this.IncidentScheduleTrigger = incidentScheduleTrigger
	return &this
}

// NewIncidentScheduleTriggerWrapperWithDefaults instantiates a new IncidentScheduleTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentScheduleTriggerWrapperWithDefaults() *IncidentScheduleTriggerWrapper {
	this := IncidentScheduleTriggerWrapper{}
	return &this
}

// GetIncidentScheduleTrigger returns the IncidentScheduleTrigger field value.
func (o *IncidentScheduleTriggerWrapper) GetIncidentScheduleTrigger() IncidentScheduleTrigger {
	if o == nil {
		var ret IncidentScheduleTrigger
		return ret
	}
	return o.IncidentScheduleTrigger
}

// GetIncidentScheduleTriggerOk returns a tuple with the IncidentScheduleTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentScheduleTriggerWrapper) GetIncidentScheduleTriggerOk() (*IncidentScheduleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentScheduleTrigger, true
}

// SetIncidentScheduleTrigger sets field value.
func (o *IncidentScheduleTriggerWrapper) SetIncidentScheduleTrigger(v IncidentScheduleTrigger) {
	o.IncidentScheduleTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentScheduleTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentScheduleTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentScheduleTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentScheduleTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentScheduleTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentScheduleTrigger"] = o.IncidentScheduleTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentScheduleTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentScheduleTrigger *IncidentScheduleTrigger `json:"incidentScheduleTrigger"`
		StartStepNames          []string                 `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentScheduleTrigger == nil {
		return fmt.Errorf("required field incidentScheduleTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentScheduleTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentScheduleTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentScheduleTrigger = *all.IncidentScheduleTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
