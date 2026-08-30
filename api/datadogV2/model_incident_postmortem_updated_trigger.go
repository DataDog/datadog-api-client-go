// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemUpdatedTrigger Trigger a workflow when a postmortem is updated for an incident.
type IncidentPostmortemUpdatedTrigger struct {
	// The maximum number of times to execute a workflow for an incident.
	ExecutionLimit *ExecutionLimit `json:"executionLimit,omitempty"`
	// The type of incident that triggers the workflow.
	IncidentType *string `json:"incidentType,omitempty"`
	// Conditions that determine which incidents trigger the workflow.
	TagCondition *IncidentCondition `json:"tagCondition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemUpdatedTrigger instantiates a new IncidentPostmortemUpdatedTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemUpdatedTrigger() *IncidentPostmortemUpdatedTrigger {
	this := IncidentPostmortemUpdatedTrigger{}
	return &this
}

// NewIncidentPostmortemUpdatedTriggerWithDefaults instantiates a new IncidentPostmortemUpdatedTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemUpdatedTriggerWithDefaults() *IncidentPostmortemUpdatedTrigger {
	this := IncidentPostmortemUpdatedTrigger{}
	return &this
}

// GetExecutionLimit returns the ExecutionLimit field value if set, zero value otherwise.
func (o *IncidentPostmortemUpdatedTrigger) GetExecutionLimit() ExecutionLimit {
	if o == nil || o.ExecutionLimit == nil {
		var ret ExecutionLimit
		return ret
	}
	return *o.ExecutionLimit
}

// GetExecutionLimitOk returns a tuple with the ExecutionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdatedTrigger) GetExecutionLimitOk() (*ExecutionLimit, bool) {
	if o == nil || o.ExecutionLimit == nil {
		return nil, false
	}
	return o.ExecutionLimit, true
}

// HasExecutionLimit returns a boolean if a field has been set.
func (o *IncidentPostmortemUpdatedTrigger) HasExecutionLimit() bool {
	return o != nil && o.ExecutionLimit != nil
}

// SetExecutionLimit gets a reference to the given ExecutionLimit and assigns it to the ExecutionLimit field.
func (o *IncidentPostmortemUpdatedTrigger) SetExecutionLimit(v ExecutionLimit) {
	o.ExecutionLimit = &v
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IncidentPostmortemUpdatedTrigger) GetIncidentType() string {
	if o == nil || o.IncidentType == nil {
		var ret string
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdatedTrigger) GetIncidentTypeOk() (*string, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentPostmortemUpdatedTrigger) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given string and assigns it to the IncidentType field.
func (o *IncidentPostmortemUpdatedTrigger) SetIncidentType(v string) {
	o.IncidentType = &v
}

// GetTagCondition returns the TagCondition field value if set, zero value otherwise.
func (o *IncidentPostmortemUpdatedTrigger) GetTagCondition() IncidentCondition {
	if o == nil || o.TagCondition == nil {
		var ret IncidentCondition
		return ret
	}
	return *o.TagCondition
}

// GetTagConditionOk returns a tuple with the TagCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdatedTrigger) GetTagConditionOk() (*IncidentCondition, bool) {
	if o == nil || o.TagCondition == nil {
		return nil, false
	}
	return o.TagCondition, true
}

// HasTagCondition returns a boolean if a field has been set.
func (o *IncidentPostmortemUpdatedTrigger) HasTagCondition() bool {
	return o != nil && o.TagCondition != nil
}

// SetTagCondition gets a reference to the given IncidentCondition and assigns it to the TagCondition field.
func (o *IncidentPostmortemUpdatedTrigger) SetTagCondition(v IncidentCondition) {
	o.TagCondition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemUpdatedTrigger) MarshalJSON() ([]byte, error) {
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
	if o.TagCondition != nil {
		toSerialize["tagCondition"] = o.TagCondition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemUpdatedTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExecutionLimit *ExecutionLimit    `json:"executionLimit,omitempty"`
		IncidentType   *string            `json:"incidentType,omitempty"`
		TagCondition   *IncidentCondition `json:"tagCondition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"executionLimit", "incidentType", "tagCondition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ExecutionLimit != nil && all.ExecutionLimit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExecutionLimit = all.ExecutionLimit
	o.IncidentType = all.IncidentType
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
