// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactUpdatedTriggerWrapper Schema for an incident impact updated trigger.
type IncidentImpactUpdatedTriggerWrapper struct {
	// Trigger a workflow when an impact is updated for an incident.
	IncidentImpactUpdatedTrigger IncidentImpactUpdatedTrigger `json:"incidentImpactUpdatedTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImpactUpdatedTriggerWrapper instantiates a new IncidentImpactUpdatedTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImpactUpdatedTriggerWrapper(incidentImpactUpdatedTrigger IncidentImpactUpdatedTrigger) *IncidentImpactUpdatedTriggerWrapper {
	this := IncidentImpactUpdatedTriggerWrapper{}
	this.IncidentImpactUpdatedTrigger = incidentImpactUpdatedTrigger
	return &this
}

// NewIncidentImpactUpdatedTriggerWrapperWithDefaults instantiates a new IncidentImpactUpdatedTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImpactUpdatedTriggerWrapperWithDefaults() *IncidentImpactUpdatedTriggerWrapper {
	this := IncidentImpactUpdatedTriggerWrapper{}
	return &this
}

// GetIncidentImpactUpdatedTrigger returns the IncidentImpactUpdatedTrigger field value.
func (o *IncidentImpactUpdatedTriggerWrapper) GetIncidentImpactUpdatedTrigger() IncidentImpactUpdatedTrigger {
	if o == nil {
		var ret IncidentImpactUpdatedTrigger
		return ret
	}
	return o.IncidentImpactUpdatedTrigger
}

// GetIncidentImpactUpdatedTriggerOk returns a tuple with the IncidentImpactUpdatedTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactUpdatedTriggerWrapper) GetIncidentImpactUpdatedTriggerOk() (*IncidentImpactUpdatedTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentImpactUpdatedTrigger, true
}

// SetIncidentImpactUpdatedTrigger sets field value.
func (o *IncidentImpactUpdatedTriggerWrapper) SetIncidentImpactUpdatedTrigger(v IncidentImpactUpdatedTrigger) {
	o.IncidentImpactUpdatedTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentImpactUpdatedTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImpactUpdatedTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentImpactUpdatedTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentImpactUpdatedTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImpactUpdatedTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentImpactUpdatedTrigger"] = o.IncidentImpactUpdatedTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImpactUpdatedTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentImpactUpdatedTrigger *IncidentImpactUpdatedTrigger `json:"incidentImpactUpdatedTrigger"`
		StartStepNames               []string                      `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentImpactUpdatedTrigger == nil {
		return fmt.Errorf("required field incidentImpactUpdatedTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentImpactUpdatedTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentImpactUpdatedTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentImpactUpdatedTrigger = *all.IncidentImpactUpdatedTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
