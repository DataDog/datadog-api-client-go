// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleDataAttributesResponse Attributes of an incident rule in a response.
type IncidentRuleDataAttributesResponse struct {
	// A query-based condition for an incident rule.
	Condition *IncidentRuleQueryCondition `json:"condition,omitempty"`
	// The condition table type.
	ConditionTableType *int64 `json:"condition_table_type,omitempty"`
	// List of field-based conditions.
	Conditions []IncidentRuleCondition `json:"conditions,omitempty"`
	// Timestamp when the rule was created.
	Created *time.Time `json:"created,omitempty"`
	// UUID of the user who created the rule.
	CreatedByUuid *uuid.UUID `json:"created_by_uuid,omitempty"`
	// Timestamp when the rule was deleted.
	Deleted datadog.NullableTime `json:"deleted,omitempty"`
	// Whether the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The execution type of the rule.
	ExecutionType *int64 `json:"execution_type,omitempty"`
	// The incident settings association UUID.
	IncidentSettingsAssociationUuid datadog.NullableUUID `json:"incident_settings_association_uuid,omitempty"`
	// Whether any condition should match.
	MatchAnyCondition *bool `json:"match_any_condition,omitempty"`
	// Timestamp when the rule was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// UUID of the user who last modified the rule.
	ModifiedByUuid *uuid.UUID `json:"modified_by_uuid,omitempty"`
	// The organization ID.
	OrgId *int64 `json:"org_id,omitempty"`
	// The task ID.
	TaskId datadog.NullableString `json:"task_id,omitempty"`
	// The JSON-encoded task payload.
	TaskPayload datadog.NullableString `json:"task_payload,omitempty"`
	// The trigger event for the rule.
	Trigger *string `json:"trigger,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRuleDataAttributesResponse instantiates a new IncidentRuleDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRuleDataAttributesResponse() *IncidentRuleDataAttributesResponse {
	this := IncidentRuleDataAttributesResponse{}
	return &this
}

// NewIncidentRuleDataAttributesResponseWithDefaults instantiates a new IncidentRuleDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRuleDataAttributesResponseWithDefaults() *IncidentRuleDataAttributesResponse {
	this := IncidentRuleDataAttributesResponse{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetCondition() IncidentRuleQueryCondition {
	if o == nil || o.Condition == nil {
		var ret IncidentRuleQueryCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetConditionOk() (*IncidentRuleQueryCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasCondition() bool {
	return o != nil && o.Condition != nil
}

// SetCondition gets a reference to the given IncidentRuleQueryCondition and assigns it to the Condition field.
func (o *IncidentRuleDataAttributesResponse) SetCondition(v IncidentRuleQueryCondition) {
	o.Condition = &v
}

// GetConditionTableType returns the ConditionTableType field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetConditionTableType() int64 {
	if o == nil || o.ConditionTableType == nil {
		var ret int64
		return ret
	}
	return *o.ConditionTableType
}

// GetConditionTableTypeOk returns a tuple with the ConditionTableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetConditionTableTypeOk() (*int64, bool) {
	if o == nil || o.ConditionTableType == nil {
		return nil, false
	}
	return o.ConditionTableType, true
}

// HasConditionTableType returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasConditionTableType() bool {
	return o != nil && o.ConditionTableType != nil
}

// SetConditionTableType gets a reference to the given int64 and assigns it to the ConditionTableType field.
func (o *IncidentRuleDataAttributesResponse) SetConditionTableType(v int64) {
	o.ConditionTableType = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetConditions() []IncidentRuleCondition {
	if o == nil || o.Conditions == nil {
		var ret []IncidentRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetConditionsOk() (*[]IncidentRuleCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []IncidentRuleCondition and assigns it to the Conditions field.
func (o *IncidentRuleDataAttributesResponse) SetConditions(v []IncidentRuleCondition) {
	o.Conditions = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentRuleDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedByUuid returns the CreatedByUuid field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetCreatedByUuid() uuid.UUID {
	if o == nil || o.CreatedByUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedByUuid
}

// GetCreatedByUuidOk returns a tuple with the CreatedByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetCreatedByUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedByUuid == nil {
		return nil, false
	}
	return o.CreatedByUuid, true
}

// HasCreatedByUuid returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasCreatedByUuid() bool {
	return o != nil && o.CreatedByUuid != nil
}

// SetCreatedByUuid gets a reference to the given uuid.UUID and assigns it to the CreatedByUuid field.
func (o *IncidentRuleDataAttributesResponse) SetCreatedByUuid(v uuid.UUID) {
	o.CreatedByUuid = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleDataAttributesResponse) GetDeleted() time.Time {
	if o == nil || o.Deleted.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleDataAttributesResponse) GetDeletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasDeleted() bool {
	return o != nil && o.Deleted.IsSet()
}

// SetDeleted gets a reference to the given datadog.NullableTime and assigns it to the Deleted field.
func (o *IncidentRuleDataAttributesResponse) SetDeleted(v time.Time) {
	o.Deleted.Set(&v)
}

// SetDeletedNil sets the value for Deleted to be an explicit nil.
func (o *IncidentRuleDataAttributesResponse) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil.
func (o *IncidentRuleDataAttributesResponse) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IncidentRuleDataAttributesResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExecutionType returns the ExecutionType field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetExecutionType() int64 {
	if o == nil || o.ExecutionType == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetExecutionTypeOk() (*int64, bool) {
	if o == nil || o.ExecutionType == nil {
		return nil, false
	}
	return o.ExecutionType, true
}

// HasExecutionType returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasExecutionType() bool {
	return o != nil && o.ExecutionType != nil
}

// SetExecutionType gets a reference to the given int64 and assigns it to the ExecutionType field.
func (o *IncidentRuleDataAttributesResponse) SetExecutionType(v int64) {
	o.ExecutionType = &v
}

// GetIncidentSettingsAssociationUuid returns the IncidentSettingsAssociationUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleDataAttributesResponse) GetIncidentSettingsAssociationUuid() uuid.UUID {
	if o == nil || o.IncidentSettingsAssociationUuid.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.IncidentSettingsAssociationUuid.Get()
}

// GetIncidentSettingsAssociationUuidOk returns a tuple with the IncidentSettingsAssociationUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleDataAttributesResponse) GetIncidentSettingsAssociationUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentSettingsAssociationUuid.Get(), o.IncidentSettingsAssociationUuid.IsSet()
}

// HasIncidentSettingsAssociationUuid returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasIncidentSettingsAssociationUuid() bool {
	return o != nil && o.IncidentSettingsAssociationUuid.IsSet()
}

// SetIncidentSettingsAssociationUuid gets a reference to the given datadog.NullableUUID and assigns it to the IncidentSettingsAssociationUuid field.
func (o *IncidentRuleDataAttributesResponse) SetIncidentSettingsAssociationUuid(v uuid.UUID) {
	o.IncidentSettingsAssociationUuid.Set(&v)
}

// SetIncidentSettingsAssociationUuidNil sets the value for IncidentSettingsAssociationUuid to be an explicit nil.
func (o *IncidentRuleDataAttributesResponse) SetIncidentSettingsAssociationUuidNil() {
	o.IncidentSettingsAssociationUuid.Set(nil)
}

// UnsetIncidentSettingsAssociationUuid ensures that no value is present for IncidentSettingsAssociationUuid, not even an explicit nil.
func (o *IncidentRuleDataAttributesResponse) UnsetIncidentSettingsAssociationUuid() {
	o.IncidentSettingsAssociationUuid.Unset()
}

// GetMatchAnyCondition returns the MatchAnyCondition field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetMatchAnyCondition() bool {
	if o == nil || o.MatchAnyCondition == nil {
		var ret bool
		return ret
	}
	return *o.MatchAnyCondition
}

// GetMatchAnyConditionOk returns a tuple with the MatchAnyCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetMatchAnyConditionOk() (*bool, bool) {
	if o == nil || o.MatchAnyCondition == nil {
		return nil, false
	}
	return o.MatchAnyCondition, true
}

// HasMatchAnyCondition returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasMatchAnyCondition() bool {
	return o != nil && o.MatchAnyCondition != nil
}

// SetMatchAnyCondition gets a reference to the given bool and assigns it to the MatchAnyCondition field.
func (o *IncidentRuleDataAttributesResponse) SetMatchAnyCondition(v bool) {
	o.MatchAnyCondition = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentRuleDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = &v
}

// GetModifiedByUuid returns the ModifiedByUuid field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetModifiedByUuid() uuid.UUID {
	if o == nil || o.ModifiedByUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ModifiedByUuid
}

// GetModifiedByUuidOk returns a tuple with the ModifiedByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetModifiedByUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.ModifiedByUuid == nil {
		return nil, false
	}
	return o.ModifiedByUuid, true
}

// HasModifiedByUuid returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasModifiedByUuid() bool {
	return o != nil && o.ModifiedByUuid != nil
}

// SetModifiedByUuid gets a reference to the given uuid.UUID and assigns it to the ModifiedByUuid field.
func (o *IncidentRuleDataAttributesResponse) SetModifiedByUuid(v uuid.UUID) {
	o.ModifiedByUuid = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *IncidentRuleDataAttributesResponse) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleDataAttributesResponse) GetTaskId() string {
	if o == nil || o.TaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleDataAttributesResponse) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasTaskId() bool {
	return o != nil && o.TaskId.IsSet()
}

// SetTaskId gets a reference to the given datadog.NullableString and assigns it to the TaskId field.
func (o *IncidentRuleDataAttributesResponse) SetTaskId(v string) {
	o.TaskId.Set(&v)
}

// SetTaskIdNil sets the value for TaskId to be an explicit nil.
func (o *IncidentRuleDataAttributesResponse) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil.
func (o *IncidentRuleDataAttributesResponse) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetTaskPayload returns the TaskPayload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleDataAttributesResponse) GetTaskPayload() string {
	if o == nil || o.TaskPayload.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskPayload.Get()
}

// GetTaskPayloadOk returns a tuple with the TaskPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleDataAttributesResponse) GetTaskPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskPayload.Get(), o.TaskPayload.IsSet()
}

// HasTaskPayload returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasTaskPayload() bool {
	return o != nil && o.TaskPayload.IsSet()
}

// SetTaskPayload gets a reference to the given datadog.NullableString and assigns it to the TaskPayload field.
func (o *IncidentRuleDataAttributesResponse) SetTaskPayload(v string) {
	o.TaskPayload.Set(&v)
}

// SetTaskPayloadNil sets the value for TaskPayload to be an explicit nil.
func (o *IncidentRuleDataAttributesResponse) SetTaskPayloadNil() {
	o.TaskPayload.Set(nil)
}

// UnsetTaskPayload ensures that no value is present for TaskPayload, not even an explicit nil.
func (o *IncidentRuleDataAttributesResponse) UnsetTaskPayload() {
	o.TaskPayload.Unset()
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *IncidentRuleDataAttributesResponse) GetTrigger() string {
	if o == nil || o.Trigger == nil {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRuleDataAttributesResponse) GetTriggerOk() (*string, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *IncidentRuleDataAttributesResponse) HasTrigger() bool {
	return o != nil && o.Trigger != nil
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *IncidentRuleDataAttributesResponse) SetTrigger(v string) {
	o.Trigger = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRuleDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.ConditionTableType != nil {
		toSerialize["condition_table_type"] = o.ConditionTableType
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedByUuid != nil {
		toSerialize["created_by_uuid"] = o.CreatedByUuid
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExecutionType != nil {
		toSerialize["execution_type"] = o.ExecutionType
	}
	if o.IncidentSettingsAssociationUuid.IsSet() {
		toSerialize["incident_settings_association_uuid"] = o.IncidentSettingsAssociationUuid.Get()
	}
	if o.MatchAnyCondition != nil {
		toSerialize["match_any_condition"] = o.MatchAnyCondition
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedByUuid != nil {
		toSerialize["modified_by_uuid"] = o.ModifiedByUuid
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.TaskId.IsSet() {
		toSerialize["task_id"] = o.TaskId.Get()
	}
	if o.TaskPayload.IsSet() {
		toSerialize["task_payload"] = o.TaskPayload.Get()
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
func (o *IncidentRuleDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition                       *IncidentRuleQueryCondition `json:"condition,omitempty"`
		ConditionTableType              *int64                      `json:"condition_table_type,omitempty"`
		Conditions                      []IncidentRuleCondition     `json:"conditions,omitempty"`
		Created                         *time.Time                  `json:"created,omitempty"`
		CreatedByUuid                   *uuid.UUID                  `json:"created_by_uuid,omitempty"`
		Deleted                         datadog.NullableTime        `json:"deleted,omitempty"`
		Enabled                         *bool                       `json:"enabled,omitempty"`
		ExecutionType                   *int64                      `json:"execution_type,omitempty"`
		IncidentSettingsAssociationUuid datadog.NullableUUID        `json:"incident_settings_association_uuid,omitempty"`
		MatchAnyCondition               *bool                       `json:"match_any_condition,omitempty"`
		Modified                        *time.Time                  `json:"modified,omitempty"`
		ModifiedByUuid                  *uuid.UUID                  `json:"modified_by_uuid,omitempty"`
		OrgId                           *int64                      `json:"org_id,omitempty"`
		TaskId                          datadog.NullableString      `json:"task_id,omitempty"`
		TaskPayload                     datadog.NullableString      `json:"task_payload,omitempty"`
		Trigger                         *string                     `json:"trigger,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "condition_table_type", "conditions", "created", "created_by_uuid", "deleted", "enabled", "execution_type", "incident_settings_association_uuid", "match_any_condition", "modified", "modified_by_uuid", "org_id", "task_id", "task_payload", "trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Condition != nil && all.Condition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Condition = all.Condition
	o.ConditionTableType = all.ConditionTableType
	o.Conditions = all.Conditions
	o.Created = all.Created
	o.CreatedByUuid = all.CreatedByUuid
	o.Deleted = all.Deleted
	o.Enabled = all.Enabled
	o.ExecutionType = all.ExecutionType
	o.IncidentSettingsAssociationUuid = all.IncidentSettingsAssociationUuid
	o.MatchAnyCondition = all.MatchAnyCondition
	o.Modified = all.Modified
	o.ModifiedByUuid = all.ModifiedByUuid
	o.OrgId = all.OrgId
	o.TaskId = all.TaskId
	o.TaskPayload = all.TaskPayload
	o.Trigger = all.Trigger

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
