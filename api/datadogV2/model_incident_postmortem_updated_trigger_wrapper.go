// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemUpdatedTriggerWrapper Schema for an incident postmortem updated trigger.
type IncidentPostmortemUpdatedTriggerWrapper struct {
	// Trigger a workflow when a postmortem is updated for an incident.
	IncidentPostmortemUpdatedTrigger IncidentPostmortemUpdatedTrigger `json:"incidentPostmortemUpdatedTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemUpdatedTriggerWrapper instantiates a new IncidentPostmortemUpdatedTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemUpdatedTriggerWrapper(incidentPostmortemUpdatedTrigger IncidentPostmortemUpdatedTrigger) *IncidentPostmortemUpdatedTriggerWrapper {
	this := IncidentPostmortemUpdatedTriggerWrapper{}
	this.IncidentPostmortemUpdatedTrigger = incidentPostmortemUpdatedTrigger
	return &this
}

// NewIncidentPostmortemUpdatedTriggerWrapperWithDefaults instantiates a new IncidentPostmortemUpdatedTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemUpdatedTriggerWrapperWithDefaults() *IncidentPostmortemUpdatedTriggerWrapper {
	this := IncidentPostmortemUpdatedTriggerWrapper{}
	return &this
}

// GetIncidentPostmortemUpdatedTrigger returns the IncidentPostmortemUpdatedTrigger field value.
func (o *IncidentPostmortemUpdatedTriggerWrapper) GetIncidentPostmortemUpdatedTrigger() IncidentPostmortemUpdatedTrigger {
	if o == nil {
		var ret IncidentPostmortemUpdatedTrigger
		return ret
	}
	return o.IncidentPostmortemUpdatedTrigger
}

// GetIncidentPostmortemUpdatedTriggerOk returns a tuple with the IncidentPostmortemUpdatedTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdatedTriggerWrapper) GetIncidentPostmortemUpdatedTriggerOk() (*IncidentPostmortemUpdatedTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentPostmortemUpdatedTrigger, true
}

// SetIncidentPostmortemUpdatedTrigger sets field value.
func (o *IncidentPostmortemUpdatedTriggerWrapper) SetIncidentPostmortemUpdatedTrigger(v IncidentPostmortemUpdatedTrigger) {
	o.IncidentPostmortemUpdatedTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentPostmortemUpdatedTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdatedTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentPostmortemUpdatedTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentPostmortemUpdatedTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemUpdatedTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentPostmortemUpdatedTrigger"] = o.IncidentPostmortemUpdatedTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemUpdatedTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentPostmortemUpdatedTrigger *IncidentPostmortemUpdatedTrigger `json:"incidentPostmortemUpdatedTrigger"`
		StartStepNames                   []string                          `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentPostmortemUpdatedTrigger == nil {
		return fmt.Errorf("required field incidentPostmortemUpdatedTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentPostmortemUpdatedTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentPostmortemUpdatedTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentPostmortemUpdatedTrigger = *all.IncidentPostmortemUpdatedTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
