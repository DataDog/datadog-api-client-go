// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleDataAttributesRequest Attributes for creating an incident rule.
type IncidentRuleDataAttributesRequest struct {
	// A query-based condition for an incident rule.
	Condition IncidentRuleQueryCondition `json:"condition"`
	// The condition table type. 1 = raw query.
	ConditionTableType int32 `json:"condition_table_type"`
	// List of field-based conditions.
	Conditions []IncidentRuleCondition `json:"conditions,omitempty"`
	// Whether the rule is enabled.
	Enabled bool `json:"enabled"`
	// The execution type of an incident rule.
	ExecutionType IncidentRuleExecutionType `json:"execution_type"`
	// The UUID of the incident type this rule applies to.
	IncidentTypeUuid datadog.NullableUUID `json:"incident_type_uuid,omitempty"`
	// Whether any condition (OR logic) should match instead of all (AND logic).
	MatchAnyCondition *bool `json:"match_any_condition,omitempty"`
	// The task ID for an incident rule.
	TaskId IncidentRuleTaskIDType `json:"task_id"`
	// The JSON-encoded payload for the task.
	TaskPayload string `json:"task_payload"`
	// The trigger event for an incident rule.
	Trigger *IncidentRuleTriggerType `json:"trigger,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRuleDataAttributesRequest instantiates a new IncidentRuleDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRuleDataAttributesRequest(condition IncidentRuleQueryCondition, conditionTableType int32, enabled bool, executionType IncidentRuleExecutionType, taskId IncidentRuleTaskIDType, taskPayload string) *IncidentRuleDataAttributesRequest {
	this := IncidentRuleDataAttributesRequest{}
	this.Condition = condition
	this.ConditionTableType = conditionTableType
	this.Enabled = enabled
	this.ExecutionType = executionType
	this.TaskId = taskId
	this.TaskPayload = taskPayload
	return &this
}

// NewIncidentRuleDataAttributesRequestWithDefaults instantiates a new IncidentRuleDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRuleDataAttributesRequestWithDefaults() *IncidentRuleDataAttributesRequest {
	this := IncidentRuleDataAttributesRequest{}
	return &this
}

// GetCondition returns the Condition field value.
func (o *IncidentRuleDataAttributesRequest) GetCondition() IncidentRuleQueryCondition {
	if o == nil {
		var ret IncidentRuleQueryCondition
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetConditionOk() (*IncidentRuleQueryCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value.
func (o *IncidentRuleDataAttributesRequest) SetCondition(v IncidentRuleQueryCondition) {
	o.Condition = v
}

// GetConditionTableType returns the ConditionTableType field value.
func (o *IncidentRuleDataAttributesRequest) GetConditionTableType() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ConditionTableType
}

// GetConditionTableTypeOk returns a tuple with the ConditionTableType field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetConditionTableTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionTableType, true
}

// SetConditionTableType sets field value.
func (o *IncidentRuleDataAttributesRequest) SetConditionTableType(v int32) {
	o.ConditionTableType = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesRequest) GetConditions() []IncidentRuleCondition {
	if o == nil || o.Conditions == nil {
		var ret []IncidentRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetConditionsOk() (*[]IncidentRuleCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesRequest) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []IncidentRuleCondition and assigns it to the Conditions field.
func (o *IncidentRuleDataAttributesRequest) SetConditions(v []IncidentRuleCondition) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value.
func (o *IncidentRuleDataAttributesRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *IncidentRuleDataAttributesRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExecutionType returns the ExecutionType field value.
func (o *IncidentRuleDataAttributesRequest) GetExecutionType() IncidentRuleExecutionType {
	if o == nil {
		var ret IncidentRuleExecutionType
		return ret
	}
	return o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetExecutionTypeOk() (*IncidentRuleExecutionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionType, true
}

// SetExecutionType sets field value.
func (o *IncidentRuleDataAttributesRequest) SetExecutionType(v IncidentRuleExecutionType) {
	o.ExecutionType = v
}

// GetIncidentTypeUuid returns the IncidentTypeUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleDataAttributesRequest) GetIncidentTypeUuid() uuid.UUID {
	if o == nil || o.IncidentTypeUuid.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.IncidentTypeUuid.Get()
}

// GetIncidentTypeUuidOk returns a tuple with the IncidentTypeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleDataAttributesRequest) GetIncidentTypeUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentTypeUuid.Get(), o.IncidentTypeUuid.IsSet()
}

// HasIncidentTypeUuid returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesRequest) HasIncidentTypeUuid() bool {
	return o != nil && o.IncidentTypeUuid.IsSet()
}

// SetIncidentTypeUuid gets a reference to the given datadog.NullableUUID and assigns it to the IncidentTypeUuid field.
func (o *IncidentRuleDataAttributesRequest) SetIncidentTypeUuid(v uuid.UUID) {
	o.IncidentTypeUuid.Set(&v)
}

// SetIncidentTypeUuidNil sets the value for IncidentTypeUuid to be an explicit nil.
func (o *IncidentRuleDataAttributesRequest) SetIncidentTypeUuidNil() {
	o.IncidentTypeUuid.Set(nil)
}

// UnsetIncidentTypeUuid ensures that no value is present for IncidentTypeUuid, not even an explicit nil.
func (o *IncidentRuleDataAttributesRequest) UnsetIncidentTypeUuid() {
	o.IncidentTypeUuid.Unset()
}

// GetMatchAnyCondition returns the MatchAnyCondition field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesRequest) GetMatchAnyCondition() bool {
	if o == nil || o.MatchAnyCondition == nil {
		var ret bool
		return ret
	}
	return *o.MatchAnyCondition
}

// GetMatchAnyConditionOk returns a tuple with the MatchAnyCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetMatchAnyConditionOk() (*bool, bool) {
	if o == nil || o.MatchAnyCondition == nil {
		return nil, false
	}
	return o.MatchAnyCondition, true
}

// HasMatchAnyCondition returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesRequest) HasMatchAnyCondition() bool {
	return o != nil && o.MatchAnyCondition != nil
}

// SetMatchAnyCondition gets a reference to the given bool and assigns it to the MatchAnyCondition field.
func (o *IncidentRuleDataAttributesRequest) SetMatchAnyCondition(v bool) {
	o.MatchAnyCondition = &v
}

// GetTaskId returns the TaskId field value.
func (o *IncidentRuleDataAttributesRequest) GetTaskId() IncidentRuleTaskIDType {
	if o == nil {
		var ret IncidentRuleTaskIDType
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetTaskIdOk() (*IncidentRuleTaskIDType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *IncidentRuleDataAttributesRequest) SetTaskId(v IncidentRuleTaskIDType) {
	o.TaskId = v
}

// GetTaskPayload returns the TaskPayload field value.
func (o *IncidentRuleDataAttributesRequest) GetTaskPayload() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskPayload
}

// GetTaskPayloadOk returns a tuple with the TaskPayload field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetTaskPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskPayload, true
}

// SetTaskPayload sets field value.
func (o *IncidentRuleDataAttributesRequest) SetTaskPayload(v string) {
	o.TaskPayload = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesRequest) GetTrigger() IncidentRuleTriggerType {
	if o == nil || o.Trigger == nil {
		var ret IncidentRuleTriggerType
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesRequest) GetTriggerOk() (*IncidentRuleTriggerType, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesRequest) HasTrigger() bool {
	return o != nil && o.Trigger != nil
}

// SetTrigger gets a reference to the given IncidentRuleTriggerType and assigns it to the Trigger field.
func (o *IncidentRuleDataAttributesRequest) SetTrigger(v IncidentRuleTriggerType) {
	o.Trigger = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRuleDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["condition"] = o.Condition
	toSerialize["condition_table_type"] = o.ConditionTableType
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["execution_type"] = o.ExecutionType
	if o.IncidentTypeUuid.IsSet() {
		toSerialize["incident_type_uuid"] = o.IncidentTypeUuid.Get()
	}
	if o.MatchAnyCondition != nil {
		toSerialize["match_any_condition"] = o.MatchAnyCondition
	}
	toSerialize["task_id"] = o.TaskId
	toSerialize["task_payload"] = o.TaskPayload
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRuleDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition          *IncidentRuleQueryCondition `json:"condition"`
		ConditionTableType *int32                      `json:"condition_table_type"`
		Conditions         []IncidentRuleCondition     `json:"conditions,omitempty"`
		Enabled            *bool                       `json:"enabled"`
		ExecutionType      *IncidentRuleExecutionType  `json:"execution_type"`
		IncidentTypeUuid   datadog.NullableUUID        `json:"incident_type_uuid,omitempty"`
		MatchAnyCondition  *bool                       `json:"match_any_condition,omitempty"`
		TaskId             *IncidentRuleTaskIDType     `json:"task_id"`
		TaskPayload        *string                     `json:"task_payload"`
		Trigger            *IncidentRuleTriggerType    `json:"trigger,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Condition == nil {
		return fmt.Errorf("required field condition missing")
	}
	if all.ConditionTableType == nil {
		return fmt.Errorf("required field condition_table_type missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ExecutionType == nil {
		return fmt.Errorf("required field execution_type missing")
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field task_id missing")
	}
	if all.TaskPayload == nil {
		return fmt.Errorf("required field task_payload missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "condition_table_type", "conditions", "enabled", "execution_type", "incident_type_uuid", "match_any_condition", "task_id", "task_payload", "trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Condition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Condition = *all.Condition
	o.ConditionTableType = *all.ConditionTableType
	o.Conditions = all.Conditions
	o.Enabled = *all.Enabled
	if !all.ExecutionType.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionType = *all.ExecutionType
	}
	o.IncidentTypeUuid = all.IncidentTypeUuid
	o.MatchAnyCondition = all.MatchAnyCondition
	if !all.TaskId.IsValid() {
		hasInvalidField = true
	} else {
		o.TaskId = *all.TaskId
	}
	o.TaskPayload = *all.TaskPayload
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
