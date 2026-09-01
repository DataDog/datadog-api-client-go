// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreatedTriggerWrapper Schema for an incident declared trigger.
type IncidentCreatedTriggerWrapper struct {
	// Trigger a workflow when an incident is declared.
	IncidentCreatedTrigger IncidentCreatedTrigger `json:"incidentCreatedTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCreatedTriggerWrapper instantiates a new IncidentCreatedTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreatedTriggerWrapper(incidentCreatedTrigger IncidentCreatedTrigger) *IncidentCreatedTriggerWrapper {
	this := IncidentCreatedTriggerWrapper{}
	this.IncidentCreatedTrigger = incidentCreatedTrigger
	return &this
}

// NewIncidentCreatedTriggerWrapperWithDefaults instantiates a new IncidentCreatedTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreatedTriggerWrapperWithDefaults() *IncidentCreatedTriggerWrapper {
	this := IncidentCreatedTriggerWrapper{}
	return &this
}

// GetIncidentCreatedTrigger returns the IncidentCreatedTrigger field value.
func (o *IncidentCreatedTriggerWrapper) GetIncidentCreatedTrigger() IncidentCreatedTrigger {
	if o == nil {
		var ret IncidentCreatedTrigger
		return ret
	}
	return o.IncidentCreatedTrigger
}

// GetIncidentCreatedTriggerOk returns a tuple with the IncidentCreatedTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentCreatedTriggerWrapper) GetIncidentCreatedTriggerOk() (*IncidentCreatedTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentCreatedTrigger, true
}

// SetIncidentCreatedTrigger sets field value.
func (o *IncidentCreatedTriggerWrapper) SetIncidentCreatedTrigger(v IncidentCreatedTrigger) {
	o.IncidentCreatedTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentCreatedTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatedTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentCreatedTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentCreatedTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreatedTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentCreatedTrigger"] = o.IncidentCreatedTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCreatedTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentCreatedTrigger *IncidentCreatedTrigger `json:"incidentCreatedTrigger"`
		StartStepNames         []string                `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentCreatedTrigger == nil {
		return fmt.Errorf("required field incidentCreatedTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentCreatedTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentCreatedTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentCreatedTrigger = *all.IncidentCreatedTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
