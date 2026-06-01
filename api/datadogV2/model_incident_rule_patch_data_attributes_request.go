// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRulePatchDataAttributesRequest Attributes for patching an incident rule. All fields are optional.
type IncidentRulePatchDataAttributesRequest struct {
	// A query-based condition for an incident rule.
	Condition *IncidentRuleQueryCondition `json:"condition,omitempty"`
	// List of field-based conditions.
	Conditions []IncidentRuleCondition `json:"conditions,omitempty"`
	// Whether the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The JSON-encoded payload for the task.
	TaskPayload *string `json:"task_payload,omitempty"`
	// The trigger event for an incident rule.
	Trigger *IncidentRuleTriggerType `json:"trigger,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRulePatchDataAttributesRequest instantiates a new IncidentRulePatchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRulePatchDataAttributesRequest() *IncidentRulePatchDataAttributesRequest {
	this := IncidentRulePatchDataAttributesRequest{}
	return &this
}

// NewIncidentRulePatchDataAttributesRequestWithDefaults instantiates a new IncidentRulePatchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRulePatchDataAttributesRequestWithDefaults() *IncidentRulePatchDataAttributesRequest {
	this := IncidentRulePatchDataAttributesRequest{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *IncidentRulePatchDataAttributesRequest) GetCondition() IncidentRuleQueryCondition {
	if o == nil || o.Condition == nil {
		var ret IncidentRuleQueryCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataAttributesRequest) GetConditionOk() (*IncidentRuleQueryCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *IncidentRulePatchDataAttributesRequest) HasCondition() bool {
	return o != nil && o.Condition != nil
}

// SetCondition gets a reference to the given IncidentRuleQueryCondition and assigns it to the Condition field.
func (o *IncidentRulePatchDataAttributesRequest) SetCondition(v IncidentRuleQueryCondition) {
	o.Condition = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IncidentRulePatchDataAttributesRequest) GetConditions() []IncidentRuleCondition {
	if o == nil || o.Conditions == nil {
		var ret []IncidentRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataAttributesRequest) GetConditionsOk() (*[]IncidentRuleCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IncidentRulePatchDataAttributesRequest) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []IncidentRuleCondition and assigns it to the Conditions field.
func (o *IncidentRulePatchDataAttributesRequest) SetConditions(v []IncidentRuleCondition) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IncidentRulePatchDataAttributesRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataAttributesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IncidentRulePatchDataAttributesRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IncidentRulePatchDataAttributesRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTaskPayload returns the TaskPayload field value if set, zero value otherwise.
func (o *IncidentRulePatchDataAttributesRequest) GetTaskPayload() string {
	if o == nil || o.TaskPayload == nil {
		var ret string
		return ret
	}
	return *o.TaskPayload
}

// GetTaskPayloadOk returns a tuple with the TaskPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataAttributesRequest) GetTaskPayloadOk() (*string, bool) {
	if o == nil || o.TaskPayload == nil {
		return nil, false
	}
	return o.TaskPayload, true
}

// HasTaskPayload returns a boolean if a field has been set.
func (o *IncidentRulePatchDataAttributesRequest) HasTaskPayload() bool {
	return o != nil && o.TaskPayload != nil
}

// SetTaskPayload gets a reference to the given string and assigns it to the TaskPayload field.
func (o *IncidentRulePatchDataAttributesRequest) SetTaskPayload(v string) {
	o.TaskPayload = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *IncidentRulePatchDataAttributesRequest) GetTrigger() IncidentRuleTriggerType {
	if o == nil || o.Trigger == nil {
		var ret IncidentRuleTriggerType
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataAttributesRequest) GetTriggerOk() (*IncidentRuleTriggerType, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *IncidentRulePatchDataAttributesRequest) HasTrigger() bool {
	return o != nil && o.Trigger != nil
}

// SetTrigger gets a reference to the given IncidentRuleTriggerType and assigns it to the Trigger field.
func (o *IncidentRulePatchDataAttributesRequest) SetTrigger(v IncidentRuleTriggerType) {
	o.Trigger = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRulePatchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.TaskPayload != nil {
		toSerialize["task_payload"] = o.TaskPayload
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRulePatchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition   *IncidentRuleQueryCondition `json:"condition,omitempty"`
		Conditions  []IncidentRuleCondition     `json:"conditions,omitempty"`
		Enabled     *bool                       `json:"enabled,omitempty"`
		TaskPayload *string                     `json:"task_payload,omitempty"`
		Trigger     *IncidentRuleTriggerType    `json:"trigger,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "conditions", "enabled", "task_payload", "trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Condition != nil && all.Condition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Condition = all.Condition
	o.Conditions = all.Conditions
	o.Enabled = all.Enabled
	o.TaskPayload = all.TaskPayload
	if all.Trigger != nil && !all.Trigger.IsValid() {
		hasInvalidField = true
	} else {
		o.Trigger = all.Trigger
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
