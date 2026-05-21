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

// IncidentRuleExecutionStateDataAttributesResponse Attributes of an incident rule execution state.
type IncidentRuleExecutionStateDataAttributesResponse struct {
	// Timestamp when the state was created.
	Created time.Time `json:"created"`
	// The incident identifier.
	IncidentUuid uuid.UUID `json:"incident_uuid"`
	// Timestamp of the last rule execution.
	LastExecutedAt datadog.NullableTime `json:"last_executed_at,omitempty"`
	// Timestamp when the state was last modified.
	Modified time.Time `json:"modified"`
	// The rule identifier.
	RuleUuid uuid.UUID `json:"rule_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRuleExecutionStateDataAttributesResponse instantiates a new IncidentRuleExecutionStateDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRuleExecutionStateDataAttributesResponse(created time.Time, incidentUuid uuid.UUID, modified time.Time, ruleUuid uuid.UUID) *IncidentRuleExecutionStateDataAttributesResponse {
	this := IncidentRuleExecutionStateDataAttributesResponse{}
	this.Created = created
	this.IncidentUuid = incidentUuid
	this.Modified = modified
	this.RuleUuid = ruleUuid
	return &this
}

// NewIncidentRuleExecutionStateDataAttributesResponseWithDefaults instantiates a new IncidentRuleExecutionStateDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRuleExecutionStateDataAttributesResponseWithDefaults() *IncidentRuleExecutionStateDataAttributesResponse {
	this := IncidentRuleExecutionStateDataAttributesResponse{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetIncidentUuid returns the IncidentUuid field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetIncidentUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.IncidentUuid
}

// GetIncidentUuidOk returns a tuple with the IncidentUuid field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetIncidentUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentUuid, true
}

// SetIncidentUuid sets field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetIncidentUuid(v uuid.UUID) {
	o.IncidentUuid = v
}

// GetLastExecutedAt returns the LastExecutedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetLastExecutedAt() time.Time {
	if o == nil || o.LastExecutedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExecutedAt.Get()
}

// GetLastExecutedAtOk returns a tuple with the LastExecutedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetLastExecutedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutedAt.Get(), o.LastExecutedAt.IsSet()
}

// HasLastExecutedAt returns a boolean if a field has been set.
func (o *IncidentRuleExecutionStateDataAttributesResponse) HasLastExecutedAt() bool {
	return o != nil && o.LastExecutedAt.IsSet()
}

// SetLastExecutedAt gets a reference to the given datadog.NullableTime and assigns it to the LastExecutedAt field.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetLastExecutedAt(v time.Time) {
	o.LastExecutedAt.Set(&v)
}

// SetLastExecutedAtNil sets the value for LastExecutedAt to be an explicit nil.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetLastExecutedAtNil() {
	o.LastExecutedAt.Set(nil)
}

// UnsetLastExecutedAt ensures that no value is present for LastExecutedAt, not even an explicit nil.
func (o *IncidentRuleExecutionStateDataAttributesResponse) UnsetLastExecutedAt() {
	o.LastExecutedAt.Unset()
}

// GetModified returns the Modified field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetRuleUuid returns the RuleUuid field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetRuleUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RuleUuid
}

// GetRuleUuidOk returns a tuple with the RuleUuid field value
// and a boolean to check if the value has been set.
func (o *IncidentRuleExecutionStateDataAttributesResponse) GetRuleUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleUuid, true
}

// SetRuleUuid sets field value.
func (o *IncidentRuleExecutionStateDataAttributesResponse) SetRuleUuid(v uuid.UUID) {
	o.RuleUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRuleExecutionStateDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["incident_uuid"] = o.IncidentUuid
	if o.LastExecutedAt.IsSet() {
		toSerialize["last_executed_at"] = o.LastExecutedAt.Get()
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["rule_uuid"] = o.RuleUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRuleExecutionStateDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created        *time.Time           `json:"created"`
		IncidentUuid   *uuid.UUID           `json:"incident_uuid"`
		LastExecutedAt datadog.NullableTime `json:"last_executed_at,omitempty"`
		Modified       *time.Time           `json:"modified"`
		RuleUuid       *uuid.UUID           `json:"rule_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.IncidentUuid == nil {
		return fmt.Errorf("required field incident_uuid missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.RuleUuid == nil {
		return fmt.Errorf("required field rule_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "incident_uuid", "last_executed_at", "modified", "rule_uuid"})
	} else {
		return err
	}
	o.Created = *all.Created
	o.IncidentUuid = *all.IncidentUuid
	o.LastExecutedAt = all.LastExecutedAt
	o.Modified = *all.Modified
	o.RuleUuid = *all.RuleUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
