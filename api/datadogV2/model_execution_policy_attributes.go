// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyAttributes An execution policy.
type ExecutionPolicyAttributes struct {
	// The set of actions this policy applies to.
	ActionPattern ExecutionPolicyActionPattern `json:"action_pattern"`
	// The date and time the execution policy was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user who created the execution policy.
	CreatedBy string `json:"created_by"`
	// Whether the policy allows or denies matching actions.
	Effect ExecutionPolicyEffect `json:"effect"`
	// The name of the execution policy.
	Name string `json:"name"`
	// Restricts where the policy applies. At most one of `kubernetes`, `scripts`,
	// or `remote_action_rshell` can be set. An empty object means the policy has
	// no scope restriction.
	Scope *ExecutionPolicyScope `json:"scope,omitempty"`
	// The targets this policy applies to.
	Targets []ExecutionPolicyTarget `json:"targets"`
	// The date and time the execution policy was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The ID of the user who last updated the execution policy.
	UpdatedBy string `json:"updated_by"`
	// The version of the execution policy. Incremented on every update.
	Version int32 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyAttributes instantiates a new ExecutionPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyAttributes(actionPattern ExecutionPolicyActionPattern, createdAt time.Time, createdBy string, effect ExecutionPolicyEffect, name string, targets []ExecutionPolicyTarget, updatedAt time.Time, updatedBy string, version int32) *ExecutionPolicyAttributes {
	this := ExecutionPolicyAttributes{}
	this.ActionPattern = actionPattern
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Effect = effect
	this.Name = name
	this.Targets = targets
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	this.Version = version
	return &this
}

// NewExecutionPolicyAttributesWithDefaults instantiates a new ExecutionPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyAttributesWithDefaults() *ExecutionPolicyAttributes {
	this := ExecutionPolicyAttributes{}
	return &this
}

// GetActionPattern returns the ActionPattern field value.
func (o *ExecutionPolicyAttributes) GetActionPattern() ExecutionPolicyActionPattern {
	if o == nil {
		var ret ExecutionPolicyActionPattern
		return ret
	}
	return o.ActionPattern
}

// GetActionPatternOk returns a tuple with the ActionPattern field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetActionPatternOk() (*ExecutionPolicyActionPattern, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPattern, true
}

// SetActionPattern sets field value.
func (o *ExecutionPolicyAttributes) SetActionPattern(v ExecutionPolicyActionPattern) {
	o.ActionPattern = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ExecutionPolicyAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ExecutionPolicyAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *ExecutionPolicyAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *ExecutionPolicyAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEffect returns the Effect field value.
func (o *ExecutionPolicyAttributes) GetEffect() ExecutionPolicyEffect {
	if o == nil {
		var ret ExecutionPolicyEffect
		return ret
	}
	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetEffectOk() (*ExecutionPolicyEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value.
func (o *ExecutionPolicyAttributes) SetEffect(v ExecutionPolicyEffect) {
	o.Effect = v
}

// GetName returns the Name field value.
func (o *ExecutionPolicyAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExecutionPolicyAttributes) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ExecutionPolicyAttributes) GetScope() ExecutionPolicyScope {
	if o == nil || o.Scope == nil {
		var ret ExecutionPolicyScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetScopeOk() (*ExecutionPolicyScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ExecutionPolicyAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given ExecutionPolicyScope and assigns it to the Scope field.
func (o *ExecutionPolicyAttributes) SetScope(v ExecutionPolicyScope) {
	o.Scope = &v
}

// GetTargets returns the Targets field value.
func (o *ExecutionPolicyAttributes) GetTargets() []ExecutionPolicyTarget {
	if o == nil {
		var ret []ExecutionPolicyTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetTargetsOk() (*[]ExecutionPolicyTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *ExecutionPolicyAttributes) SetTargets(v []ExecutionPolicyTarget) {
	o.Targets = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ExecutionPolicyAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ExecutionPolicyAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *ExecutionPolicyAttributes) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *ExecutionPolicyAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetVersion returns the Version field value.
func (o *ExecutionPolicyAttributes) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyAttributes) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ExecutionPolicyAttributes) SetVersion(v int32) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_pattern"] = o.ActionPattern
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["effect"] = o.Effect
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["targets"] = o.Targets
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["updated_by"] = o.UpdatedBy
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionPattern *ExecutionPolicyActionPattern `json:"action_pattern"`
		CreatedAt     *time.Time                    `json:"created_at"`
		CreatedBy     *string                       `json:"created_by"`
		Effect        *ExecutionPolicyEffect        `json:"effect"`
		Name          *string                       `json:"name"`
		Scope         *ExecutionPolicyScope         `json:"scope,omitempty"`
		Targets       *[]ExecutionPolicyTarget      `json:"targets"`
		UpdatedAt     *time.Time                    `json:"updated_at"`
		UpdatedBy     *string                       `json:"updated_by"`
		Version       *int32                        `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionPattern == nil {
		return fmt.Errorf("required field action_pattern missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Effect == nil {
		return fmt.Errorf("required field effect missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_pattern", "created_at", "created_by", "effect", "name", "scope", "targets", "updated_at", "updated_by", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ActionPattern.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActionPattern = *all.ActionPattern
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.Effect.IsValid() {
		hasInvalidField = true
	} else {
		o.Effect = *all.Effect
	}
	o.Name = *all.Name
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.Targets = *all.Targets
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
