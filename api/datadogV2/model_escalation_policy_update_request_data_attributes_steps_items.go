// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataAttributesStepsItems Defines a single escalation step within an escalation policy update request. Contains assignment strategy, escalation timeout, an optional step ID, and a list of targets.
type EscalationPolicyUpdateRequestDataAttributesStepsItems struct {
	// Specifies how this escalation step will assign targets (example `default`).
	Assignment *EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment `json:"assignment,omitempty"`
	// Defines how many seconds to wait before escalating to the next step.
	EscalateAfterSeconds *int64 `json:"escalate_after_seconds,omitempty"`
	// Specifies the unique identifier of this step.
	Id *string `json:"id,omitempty"`
	// Specifies the collection of escalation targets for this step.
	Targets []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyUpdateRequestDataAttributesStepsItems instantiates a new EscalationPolicyUpdateRequestDataAttributesStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyUpdateRequestDataAttributesStepsItems(targets []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems) *EscalationPolicyUpdateRequestDataAttributesStepsItems {
	this := EscalationPolicyUpdateRequestDataAttributesStepsItems{}
	this.Targets = targets
	return &this
}

// NewEscalationPolicyUpdateRequestDataAttributesStepsItemsWithDefaults instantiates a new EscalationPolicyUpdateRequestDataAttributesStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyUpdateRequestDataAttributesStepsItemsWithDefaults() *EscalationPolicyUpdateRequestDataAttributesStepsItems {
	this := EscalationPolicyUpdateRequestDataAttributesStepsItems{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetAssignment() EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment {
	if o == nil || o.Assignment == nil {
		var ret EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetAssignmentOk() (*EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment, bool) {
	if o == nil || o.Assignment == nil {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) HasAssignment() bool {
	return o != nil && o.Assignment != nil
}

// SetAssignment gets a reference to the given EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment and assigns it to the Assignment field.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) SetAssignment(v EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment) {
	o.Assignment = &v
}

// GetEscalateAfterSeconds returns the EscalateAfterSeconds field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetEscalateAfterSeconds() int64 {
	if o == nil || o.EscalateAfterSeconds == nil {
		var ret int64
		return ret
	}
	return *o.EscalateAfterSeconds
}

// GetEscalateAfterSecondsOk returns a tuple with the EscalateAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetEscalateAfterSecondsOk() (*int64, bool) {
	if o == nil || o.EscalateAfterSeconds == nil {
		return nil, false
	}
	return o.EscalateAfterSeconds, true
}

// HasEscalateAfterSeconds returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) HasEscalateAfterSeconds() bool {
	return o != nil && o.EscalateAfterSeconds != nil
}

// SetEscalateAfterSeconds gets a reference to the given int64 and assigns it to the EscalateAfterSeconds field.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) SetEscalateAfterSeconds(v int64) {
	o.EscalateAfterSeconds = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) SetId(v string) {
	o.Id = &v
}

// GetTargets returns the Targets field value.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetTargets() []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems {
	if o == nil {
		var ret []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) GetTargetsOk() (*[]EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) SetTargets(v []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyUpdateRequestDataAttributesStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignment != nil {
		toSerialize["assignment"] = o.Assignment
	}
	if o.EscalateAfterSeconds != nil {
		toSerialize["escalate_after_seconds"] = o.EscalateAfterSeconds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyUpdateRequestDataAttributesStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignment           *EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment     `json:"assignment,omitempty"`
		EscalateAfterSeconds *int64                                                               `json:"escalate_after_seconds,omitempty"`
		Id                   *string                                                              `json:"id,omitempty"`
		Targets              *[]EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItems `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment", "escalate_after_seconds", "id", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assignment != nil && !all.Assignment.IsValid() {
		hasInvalidField = true
	} else {
		o.Assignment = all.Assignment
	}
	o.EscalateAfterSeconds = all.EscalateAfterSeconds
	o.Id = all.Id
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
