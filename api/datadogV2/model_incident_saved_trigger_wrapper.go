// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSavedTriggerWrapper Schema for an incident declared or updated trigger.
type IncidentSavedTriggerWrapper struct {
	// Trigger a workflow when an incident is declared or updated.
	IncidentSavedTrigger IncidentSavedTrigger `json:"incidentSavedTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSavedTriggerWrapper instantiates a new IncidentSavedTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSavedTriggerWrapper(incidentSavedTrigger IncidentSavedTrigger) *IncidentSavedTriggerWrapper {
	this := IncidentSavedTriggerWrapper{}
	this.IncidentSavedTrigger = incidentSavedTrigger
	return &this
}

// NewIncidentSavedTriggerWrapperWithDefaults instantiates a new IncidentSavedTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSavedTriggerWrapperWithDefaults() *IncidentSavedTriggerWrapper {
	this := IncidentSavedTriggerWrapper{}
	return &this
}

// GetIncidentSavedTrigger returns the IncidentSavedTrigger field value.
func (o *IncidentSavedTriggerWrapper) GetIncidentSavedTrigger() IncidentSavedTrigger {
	if o == nil {
		var ret IncidentSavedTrigger
		return ret
	}
	return o.IncidentSavedTrigger
}

// GetIncidentSavedTriggerOk returns a tuple with the IncidentSavedTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentSavedTriggerWrapper) GetIncidentSavedTriggerOk() (*IncidentSavedTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentSavedTrigger, true
}

// SetIncidentSavedTrigger sets field value.
func (o *IncidentSavedTriggerWrapper) SetIncidentSavedTrigger(v IncidentSavedTrigger) {
	o.IncidentSavedTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentSavedTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSavedTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentSavedTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentSavedTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSavedTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentSavedTrigger"] = o.IncidentSavedTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSavedTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentSavedTrigger *IncidentSavedTrigger `json:"incidentSavedTrigger"`
		StartStepNames       []string              `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentSavedTrigger == nil {
		return fmt.Errorf("required field incidentSavedTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentSavedTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentSavedTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentSavedTrigger = *all.IncidentSavedTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
