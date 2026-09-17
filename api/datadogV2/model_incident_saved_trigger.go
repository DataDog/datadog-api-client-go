// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSavedTrigger Trigger a workflow when an incident is declared or updated.
type IncidentSavedTrigger struct {
	// The maximum number of times to execute a workflow for an incident.
	ExecutionLimit *ExecutionLimit `json:"executionLimit,omitempty"`
	// The type of incident that triggers the workflow.
	IncidentType *string `json:"incidentType,omitempty"`
	// Whether to execute the workflow serially for an incident.
	SerialExecution *SerialExecution `json:"serialExecution,omitempty"`
	// Conditions that determine which incidents trigger the workflow.
	TagCondition *IncidentCondition `json:"tagCondition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSavedTrigger instantiates a new IncidentSavedTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSavedTrigger() *IncidentSavedTrigger {
	this := IncidentSavedTrigger{}
	return &this
}

// NewIncidentSavedTriggerWithDefaults instantiates a new IncidentSavedTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSavedTriggerWithDefaults() *IncidentSavedTrigger {
	this := IncidentSavedTrigger{}
	return &this
}

// GetExecutionLimit returns the ExecutionLimit field value if set, zero value otherwise.
func (o *IncidentSavedTrigger) GetExecutionLimit() ExecutionLimit {
	if o == nil || o.ExecutionLimit == nil {
		var ret ExecutionLimit
		return ret
	}
	return *o.ExecutionLimit
}

// GetExecutionLimitOk returns a tuple with the ExecutionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSavedTrigger) GetExecutionLimitOk() (*ExecutionLimit, bool) {
	if o == nil || o.ExecutionLimit == nil {
		return nil, false
	}
	return o.ExecutionLimit, true
}

// HasExecutionLimit returns a boolean if a field has been set.
func (o *IncidentSavedTrigger) HasExecutionLimit() bool {
	return o != nil && o.ExecutionLimit != nil
}

// SetExecutionLimit gets a reference to the given ExecutionLimit and assigns it to the ExecutionLimit field.
func (o *IncidentSavedTrigger) SetExecutionLimit(v ExecutionLimit) {
	o.ExecutionLimit = &v
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IncidentSavedTrigger) GetIncidentType() string {
	if o == nil || o.IncidentType == nil {
		var ret string
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSavedTrigger) GetIncidentTypeOk() (*string, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentSavedTrigger) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given string and assigns it to the IncidentType field.
func (o *IncidentSavedTrigger) SetIncidentType(v string) {
	o.IncidentType = &v
}

// GetSerialExecution returns the SerialExecution field value if set, zero value otherwise.
func (o *IncidentSavedTrigger) GetSerialExecution() SerialExecution {
	if o == nil || o.SerialExecution == nil {
		var ret SerialExecution
		return ret
	}
	return *o.SerialExecution
}

// GetSerialExecutionOk returns a tuple with the SerialExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSavedTrigger) GetSerialExecutionOk() (*SerialExecution, bool) {
	if o == nil || o.SerialExecution == nil {
		return nil, false
	}
	return o.SerialExecution, true
}

// HasSerialExecution returns a boolean if a field has been set.
func (o *IncidentSavedTrigger) HasSerialExecution() bool {
	return o != nil && o.SerialExecution != nil
}

// SetSerialExecution gets a reference to the given SerialExecution and assigns it to the SerialExecution field.
func (o *IncidentSavedTrigger) SetSerialExecution(v SerialExecution) {
	o.SerialExecution = &v
}

// GetTagCondition returns the TagCondition field value if set, zero value otherwise.
func (o *IncidentSavedTrigger) GetTagCondition() IncidentCondition {
	if o == nil || o.TagCondition == nil {
		var ret IncidentCondition
		return ret
	}
	return *o.TagCondition
}

// GetTagConditionOk returns a tuple with the TagCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSavedTrigger) GetTagConditionOk() (*IncidentCondition, bool) {
	if o == nil || o.TagCondition == nil {
		return nil, false
	}
	return o.TagCondition, true
}

// HasTagCondition returns a boolean if a field has been set.
func (o *IncidentSavedTrigger) HasTagCondition() bool {
	return o != nil && o.TagCondition != nil
}

// SetTagCondition gets a reference to the given IncidentCondition and assigns it to the TagCondition field.
func (o *IncidentSavedTrigger) SetTagCondition(v IncidentCondition) {
	o.TagCondition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSavedTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExecutionLimit != nil {
		toSerialize["executionLimit"] = o.ExecutionLimit
	}
	if o.IncidentType != nil {
		toSerialize["incidentType"] = o.IncidentType
	}
	if o.SerialExecution != nil {
		toSerialize["serialExecution"] = o.SerialExecution
	}
	if o.TagCondition != nil {
		toSerialize["tagCondition"] = o.TagCondition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSavedTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExecutionLimit  *ExecutionLimit    `json:"executionLimit,omitempty"`
		IncidentType    *string            `json:"incidentType,omitempty"`
		SerialExecution *SerialExecution   `json:"serialExecution,omitempty"`
		TagCondition    *IncidentCondition `json:"tagCondition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"executionLimit", "incidentType", "serialExecution", "tagCondition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ExecutionLimit != nil && all.ExecutionLimit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExecutionLimit = all.ExecutionLimit
	o.IncidentType = all.IncidentType
	if all.SerialExecution != nil && all.SerialExecution.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SerialExecution = all.SerialExecution
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
