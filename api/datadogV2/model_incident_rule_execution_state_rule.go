// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleExecutionStateRule A rule in a batch request.
type IncidentRuleExecutionStateRule struct {
	// Timestamp of the last rule execution.
	LastExecutedAt datadog.NullableTime `json:"last_executed_at,omitempty"`
	// The rule identifier.
	RuleUuid uuid.UUID `json:"rule_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRuleExecutionStateRule instantiates a new IncidentRuleExecutionStateRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRuleExecutionStateRule(ruleUuid uuid.UUID) *IncidentRuleExecutionStateRule {
	this := IncidentRuleExecutionStateRule{}
	this.RuleUuid = ruleUuid
	return &this
}

// NewIncidentRuleExecutionStateRuleWithDefaults instantiates a new IncidentRuleExecutionStateRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRuleExecutionStateRuleWithDefaults() *IncidentRuleExecutionStateRule {
	this := IncidentRuleExecutionStateRule{}
	return &this
}

// GetLastExecutedAt returns the LastExecutedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleExecutionStateRule) GetLastExecutedAt() time.Time {
	if o == nil || o.LastExecutedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExecutedAt.Get()
}

// GetLastExecutedAtOk returns a tuple with the LastExecutedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleExecutionStateRule) GetLastExecutedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutedAt.Get(), o.LastExecutedAt.IsSet()
}

// HasLastExecutedAt returns a boolean if a field has been set.
func (o *IncidentRuleExecutionStateRule) HasLastExecutedAt() bool {
	return o != nil && o.LastExecutedAt.IsSet()
}

// SetLastExecutedAt gets a reference to the given datadog.NullableTime and assigns it to the LastExecutedAt field.
func (o *IncidentRuleExecutionStateRule) SetLastExecutedAt(v time.Time) {
	o.LastExecutedAt.Set(&v)
}

// SetLastExecutedAtNil sets the value for LastExecutedAt to be an explicit nil.
func (o *IncidentRuleExecutionStateRule) SetLastExecutedAtNil() {
	o.LastExecutedAt.Set(nil)
}

// UnsetLastExecutedAt ensures that no value is present for LastExecutedAt, not even an explicit nil.
func (o *IncidentRuleExecutionStateRule) UnsetLastExecutedAt() {
	o.LastExecutedAt.Unset()
}

// GetRuleUuid returns the RuleUuid field value.
func (o *IncidentRuleExecutionStateRule) GetRuleUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RuleUuid
}

// GetRuleUuidOk returns a tuple with the RuleUuid field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleExecutionStateRule) GetRuleUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleUuid, true
}

// SetRuleUuid sets field value.
func (o *IncidentRuleExecutionStateRule) SetRuleUuid(v uuid.UUID) {
	o.RuleUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRuleExecutionStateRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LastExecutedAt.IsSet() {
		toSerialize["last_executed_at"] = o.LastExecutedAt.Get()
	}
	toSerialize["rule_uuid"] = o.RuleUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRuleExecutionStateRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastExecutedAt datadog.NullableTime `json:"last_executed_at,omitempty"`
		RuleUuid       *uuid.UUID           `json:"rule_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RuleUuid == nil {
		return fmt.Errorf("required field rule_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_executed_at", "rule_uuid"})
	} else {
		return err
	}
	o.LastExecutedAt = all.LastExecutedAt
	o.RuleUuid = *all.RuleUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
