// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchSavedFilterAttributes Fields to update. At least one of `name`, `description`, or `targeting_rules` must have a non-null value. Omitted or null fields are preserved. The creation type cannot be changed.
type PatchSavedFilterAttributes struct {
	// The new description. Send an empty string to clear the description. If omitted or null, the existing description is preserved.
	Description datadog.NullableString `json:"description,omitempty"`
	// The `updated_at` value returned by a previous read, as an RFC 3339 timestamp. If it differs from the stored update time, truncated to seconds, the request returns a 409 response. For metadata-only updates, this check does not prevent all concurrent writes. This field alone does not constitute an update.
	LastUpdatedAt datadog.NullableTime `json:"last_updated_at,omitempty"`
	// The new name of the saved filter. Must not be empty.
	Name datadog.NullableString `json:"name,omitempty"`
	// The complete replacement targeting rules, combined with OR. Conditions within each rule are combined with AND. Replaces all existing rules. Values are not appended. Must contain at least one rule. Saved filters cannot reference other saved filters.
	TargetingRules []SavedFilterTargetingRuleRequest `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchSavedFilterAttributes instantiates a new PatchSavedFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchSavedFilterAttributes() *PatchSavedFilterAttributes {
	this := PatchSavedFilterAttributes{}
	return &this
}

// NewPatchSavedFilterAttributesWithDefaults instantiates a new PatchSavedFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchSavedFilterAttributesWithDefaults() *PatchSavedFilterAttributes {
	this := PatchSavedFilterAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSavedFilterAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PatchSavedFilterAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchSavedFilterAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *PatchSavedFilterAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *PatchSavedFilterAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *PatchSavedFilterAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSavedFilterAttributes) GetLastUpdatedAt() time.Time {
	if o == nil || o.LastUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt.Get()
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PatchSavedFilterAttributes) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdatedAt.Get(), o.LastUpdatedAt.IsSet()
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *PatchSavedFilterAttributes) HasLastUpdatedAt() bool {
	return o != nil && o.LastUpdatedAt.IsSet()
}

// SetLastUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the LastUpdatedAt field.
func (o *PatchSavedFilterAttributes) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt.Set(&v)
}

// SetLastUpdatedAtNil sets the value for LastUpdatedAt to be an explicit nil.
func (o *PatchSavedFilterAttributes) SetLastUpdatedAtNil() {
	o.LastUpdatedAt.Set(nil)
}

// UnsetLastUpdatedAt ensures that no value is present for LastUpdatedAt, not even an explicit nil.
func (o *PatchSavedFilterAttributes) UnsetLastUpdatedAt() {
	o.LastUpdatedAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSavedFilterAttributes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PatchSavedFilterAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchSavedFilterAttributes) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *PatchSavedFilterAttributes) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *PatchSavedFilterAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *PatchSavedFilterAttributes) UnsetName() {
	o.Name.Unset()
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSavedFilterAttributes) GetTargetingRules() []SavedFilterTargetingRuleRequest {
	if o == nil {
		var ret []SavedFilterTargetingRuleRequest
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PatchSavedFilterAttributes) GetTargetingRulesOk() (*[]SavedFilterTargetingRuleRequest, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *PatchSavedFilterAttributes) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []SavedFilterTargetingRuleRequest and assigns it to the TargetingRules field.
func (o *PatchSavedFilterAttributes) SetTargetingRules(v []SavedFilterTargetingRuleRequest) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchSavedFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.LastUpdatedAt.IsSet() {
		toSerialize["last_updated_at"] = o.LastUpdatedAt.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchSavedFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    datadog.NullableString            `json:"description,omitempty"`
		LastUpdatedAt  datadog.NullableTime              `json:"last_updated_at,omitempty"`
		Name           datadog.NullableString            `json:"name,omitempty"`
		TargetingRules []SavedFilterTargetingRuleRequest `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "last_updated_at", "name", "targeting_rules"})
	} else {
		return err
	}
	o.Description = all.Description
	o.LastUpdatedAt = all.LastUpdatedAt
	o.Name = all.Name
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
