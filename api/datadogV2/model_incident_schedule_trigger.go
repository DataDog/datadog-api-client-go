// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentScheduleTrigger Trigger a workflow from an Incident Schedule. The workflow must be published.
type IncidentScheduleTrigger struct {
	// Incident type filter for the schedule.
	IncidentType *string `json:"incidentType,omitempty"`
	// Recurrence rule expression for scheduling.
	Rrule string `json:"rrule"`
	// A condition evaluated against incident tags.
	TagCondition *IncidentCondition `json:"tagCondition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentScheduleTrigger instantiates a new IncidentScheduleTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentScheduleTrigger(rrule string) *IncidentScheduleTrigger {
	this := IncidentScheduleTrigger{}
	this.Rrule = rrule
	return &this
}

// NewIncidentScheduleTriggerWithDefaults instantiates a new IncidentScheduleTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentScheduleTriggerWithDefaults() *IncidentScheduleTrigger {
	this := IncidentScheduleTrigger{}
	return &this
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IncidentScheduleTrigger) GetIncidentType() string {
	if o == nil || o.IncidentType == nil {
		var ret string
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentScheduleTrigger) GetIncidentTypeOk() (*string, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentScheduleTrigger) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given string and assigns it to the IncidentType field.
func (o *IncidentScheduleTrigger) SetIncidentType(v string) {
	o.IncidentType = &v
}

// GetRrule returns the Rrule field value.
func (o *IncidentScheduleTrigger) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *IncidentScheduleTrigger) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value.
func (o *IncidentScheduleTrigger) SetRrule(v string) {
	o.Rrule = v
}

// GetTagCondition returns the TagCondition field value if set, zero value otherwise.
func (o *IncidentScheduleTrigger) GetTagCondition() IncidentCondition {
	if o == nil || o.TagCondition == nil {
		var ret IncidentCondition
		return ret
	}
	return *o.TagCondition
}

// GetTagConditionOk returns a tuple with the TagCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentScheduleTrigger) GetTagConditionOk() (*IncidentCondition, bool) {
	if o == nil || o.TagCondition == nil {
		return nil, false
	}
	return o.TagCondition, true
}

// HasTagCondition returns a boolean if a field has been set.
func (o *IncidentScheduleTrigger) HasTagCondition() bool {
	return o != nil && o.TagCondition != nil
}

// SetTagCondition gets a reference to the given IncidentCondition and assigns it to the TagCondition field.
func (o *IncidentScheduleTrigger) SetTagCondition(v IncidentCondition) {
	o.TagCondition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentScheduleTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncidentType != nil {
		toSerialize["incidentType"] = o.IncidentType
	}
	toSerialize["rrule"] = o.Rrule
	if o.TagCondition != nil {
		toSerialize["tagCondition"] = o.TagCondition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentScheduleTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentType *string            `json:"incidentType,omitempty"`
		Rrule        *string            `json:"rrule"`
		TagCondition *IncidentCondition `json:"tagCondition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rrule == nil {
		return fmt.Errorf("required field rrule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentType", "rrule", "tagCondition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IncidentType = all.IncidentType
	o.Rrule = *all.Rrule
	if all.TagCondition != nil && all.TagCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TagCondition = all.TagCondition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
