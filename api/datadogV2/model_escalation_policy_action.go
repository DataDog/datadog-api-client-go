// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyAction Triggers an escalation policy when the routing rule matches.
type EscalationPolicyAction struct {
	// The ID of the escalation policy to trigger.
	PolicyId string `json:"policy_id"`
	// Holds time zone information and a list of time restrictions for a routing rule.
	SupportHours *TimeRestrictions `json:"support_hours,omitempty"`
	// Indicates that the action is an escalation policy action.
	Type EscalationPolicyActionType `json:"type"`
	// Specifies the level of urgency for a routing rule (low, high, or dynamic).
	Urgency *Urgency `json:"urgency,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyAction instantiates a new EscalationPolicyAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyAction(policyId string, typeVar EscalationPolicyActionType) *EscalationPolicyAction {
	this := EscalationPolicyAction{}
	this.PolicyId = policyId
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyActionWithDefaults instantiates a new EscalationPolicyAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyActionWithDefaults() *EscalationPolicyAction {
	this := EscalationPolicyAction{}
	var typeVar EscalationPolicyActionType = ESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY
	this.Type = typeVar
	return &this
}

// GetPolicyId returns the PolicyId field value.
func (o *EscalationPolicyAction) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyAction) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value.
func (o *EscalationPolicyAction) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetSupportHours returns the SupportHours field value if set, zero value otherwise.
func (o *EscalationPolicyAction) GetSupportHours() TimeRestrictions {
	if o == nil || o.SupportHours == nil {
		var ret TimeRestrictions
		return ret
	}
	return *o.SupportHours
}

// GetSupportHoursOk returns a tuple with the SupportHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyAction) GetSupportHoursOk() (*TimeRestrictions, bool) {
	if o == nil || o.SupportHours == nil {
		return nil, false
	}
	return o.SupportHours, true
}

// HasSupportHours returns a boolean if a field has been set.
func (o *EscalationPolicyAction) HasSupportHours() bool {
	return o != nil && o.SupportHours != nil
}

// SetSupportHours gets a reference to the given TimeRestrictions and assigns it to the SupportHours field.
func (o *EscalationPolicyAction) SetSupportHours(v TimeRestrictions) {
	o.SupportHours = &v
}

// GetType returns the Type field value.
func (o *EscalationPolicyAction) GetType() EscalationPolicyActionType {
	if o == nil {
		var ret EscalationPolicyActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyAction) GetTypeOk() (*EscalationPolicyActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyAction) SetType(v EscalationPolicyActionType) {
	o.Type = v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *EscalationPolicyAction) GetUrgency() Urgency {
	if o == nil || o.Urgency == nil {
		var ret Urgency
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyAction) GetUrgencyOk() (*Urgency, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *EscalationPolicyAction) HasUrgency() bool {
	return o != nil && o.Urgency != nil
}

// SetUrgency gets a reference to the given Urgency and assigns it to the Urgency field.
func (o *EscalationPolicyAction) SetUrgency(v Urgency) {
	o.Urgency = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["policy_id"] = o.PolicyId
	if o.SupportHours != nil {
		toSerialize["support_hours"] = o.SupportHours
	}
	toSerialize["type"] = o.Type
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PolicyId     *string                     `json:"policy_id"`
		SupportHours *TimeRestrictions           `json:"support_hours,omitempty"`
		Type         *EscalationPolicyActionType `json:"type"`
		Urgency      *Urgency                    `json:"urgency,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PolicyId == nil {
		return fmt.Errorf("required field policy_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"policy_id", "support_hours", "type", "urgency"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PolicyId = *all.PolicyId
	if all.SupportHours != nil && all.SupportHours.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SupportHours = all.SupportHours
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Urgency != nil && !all.Urgency.IsValid() {
		hasInvalidField = true
	} else {
		o.Urgency = all.Urgency
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
