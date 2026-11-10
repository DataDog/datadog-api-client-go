// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchAttributes Object containing the information of a watch.
type IssueWatchAttributes struct {
	// Action applied to the issue when the watch condition is met.
	Action *IssueWatchAction `json:"action,omitempty"`
	// Condition that triggers the watch action. The fields that are set depend on `kind`:
	// `OCCURRENCES` uses `threshold`, `CARDINALITY` uses `threshold` and `projection`,
	// `ELAPSED_TIME` uses `duration`, and `ANY` and `ALL` use `operands`.
	Condition *IssueWatchCondition `json:"condition,omitempty"`
	// Timestamp of when the watch started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// Status of the watch.
	Status *IssueWatchStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueWatchAttributes instantiates a new IssueWatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueWatchAttributes() *IssueWatchAttributes {
	this := IssueWatchAttributes{}
	return &this
}

// NewIssueWatchAttributesWithDefaults instantiates a new IssueWatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueWatchAttributesWithDefaults() *IssueWatchAttributes {
	this := IssueWatchAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IssueWatchAttributes) GetAction() IssueWatchAction {
	if o == nil || o.Action == nil {
		var ret IssueWatchAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAttributes) GetActionOk() (*IssueWatchAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IssueWatchAttributes) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given IssueWatchAction and assigns it to the Action field.
func (o *IssueWatchAttributes) SetAction(v IssueWatchAction) {
	o.Action = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *IssueWatchAttributes) GetCondition() IssueWatchCondition {
	if o == nil || o.Condition == nil {
		var ret IssueWatchCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAttributes) GetConditionOk() (*IssueWatchCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *IssueWatchAttributes) HasCondition() bool {
	return o != nil && o.Condition != nil
}

// SetCondition gets a reference to the given IssueWatchCondition and assigns it to the Condition field.
func (o *IssueWatchAttributes) SetCondition(v IssueWatchCondition) {
	o.Condition = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *IssueWatchAttributes) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *IssueWatchAttributes) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *IssueWatchAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssueWatchAttributes) GetStatus() IssueWatchStatus {
	if o == nil || o.Status == nil {
		var ret IssueWatchStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAttributes) GetStatusOk() (*IssueWatchStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssueWatchAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given IssueWatchStatus and assigns it to the Status field.
func (o *IssueWatchAttributes) SetStatus(v IssueWatchStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueWatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueWatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action    *IssueWatchAction    `json:"action,omitempty"`
		Condition *IssueWatchCondition `json:"condition,omitempty"`
		StartedAt *time.Time           `json:"started_at,omitempty"`
		Status    *IssueWatchStatus    `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "condition", "started_at", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action != nil && all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = all.Action
	if all.Condition != nil && all.Condition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Condition = all.Condition
	o.StartedAt = all.StartedAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
