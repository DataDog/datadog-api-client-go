// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStalenessDetails The feature flag's current staleness state and suggested actions.
type FeatureFlagStalenessDetails struct {
	// Repositories and files where the flag is referenced in source code.
	CodeReferences []FeatureFlagStalenessCodeReference `json:"code_references,omitempty"`
	// The ID of the user who dismissed the staleness recommendation.
	DismissedBy datadog.NullableString `json:"dismissed_by,omitempty"`
	// The ID of the feature flag whose staleness state is returned.
	Id *string `json:"id,omitempty"`
	// Suggested actions for the flag. The first action is the primary recommendation.
	RecommendedActions []FeatureFlagStalenessRecommendedAction `json:"recommended_actions,omitempty"`
	// Time until which staleness checks are paused for the flag.
	SkipStateCheckUntil datadog.NullableTime `json:"skip_state_check_until,omitempty"`
	// Why the flag is stale or has a manually selected state. Values include `FULLY_ROLLED_OUT`, `NO_EVALUATIONS`, `NO_ACTIVITY`, and `USER_SET`.
	StaleReason datadog.NullableString `json:"stale_reason,omitempty"`
	// The current state, such as `ACTIVE`, `STALE`, or `PERMANENT`.
	StalenessStatus *string `json:"staleness_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagStalenessDetails instantiates a new FeatureFlagStalenessDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagStalenessDetails() *FeatureFlagStalenessDetails {
	this := FeatureFlagStalenessDetails{}
	return &this
}

// NewFeatureFlagStalenessDetailsWithDefaults instantiates a new FeatureFlagStalenessDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagStalenessDetailsWithDefaults() *FeatureFlagStalenessDetails {
	this := FeatureFlagStalenessDetails{}
	return &this
}

// GetCodeReferences returns the CodeReferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagStalenessDetails) GetCodeReferences() []FeatureFlagStalenessCodeReference {
	if o == nil {
		var ret []FeatureFlagStalenessCodeReference
		return ret
	}
	return o.CodeReferences
}

// GetCodeReferencesOk returns a tuple with the CodeReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagStalenessDetails) GetCodeReferencesOk() (*[]FeatureFlagStalenessCodeReference, bool) {
	if o == nil || o.CodeReferences == nil {
		return nil, false
	}
	return &o.CodeReferences, true
}

// HasCodeReferences returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasCodeReferences() bool {
	return o != nil && o.CodeReferences != nil
}

// SetCodeReferences gets a reference to the given []FeatureFlagStalenessCodeReference and assigns it to the CodeReferences field.
func (o *FeatureFlagStalenessDetails) SetCodeReferences(v []FeatureFlagStalenessCodeReference) {
	o.CodeReferences = v
}

// GetDismissedBy returns the DismissedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagStalenessDetails) GetDismissedBy() string {
	if o == nil || o.DismissedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DismissedBy.Get()
}

// GetDismissedByOk returns a tuple with the DismissedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagStalenessDetails) GetDismissedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissedBy.Get(), o.DismissedBy.IsSet()
}

// HasDismissedBy returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasDismissedBy() bool {
	return o != nil && o.DismissedBy.IsSet()
}

// SetDismissedBy gets a reference to the given datadog.NullableString and assigns it to the DismissedBy field.
func (o *FeatureFlagStalenessDetails) SetDismissedBy(v string) {
	o.DismissedBy.Set(&v)
}

// SetDismissedByNil sets the value for DismissedBy to be an explicit nil.
func (o *FeatureFlagStalenessDetails) SetDismissedByNil() {
	o.DismissedBy.Set(nil)
}

// UnsetDismissedBy ensures that no value is present for DismissedBy, not even an explicit nil.
func (o *FeatureFlagStalenessDetails) UnsetDismissedBy() {
	o.DismissedBy.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeatureFlagStalenessDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeatureFlagStalenessDetails) SetId(v string) {
	o.Id = &v
}

// GetRecommendedActions returns the RecommendedActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagStalenessDetails) GetRecommendedActions() []FeatureFlagStalenessRecommendedAction {
	if o == nil {
		var ret []FeatureFlagStalenessRecommendedAction
		return ret
	}
	return o.RecommendedActions
}

// GetRecommendedActionsOk returns a tuple with the RecommendedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagStalenessDetails) GetRecommendedActionsOk() (*[]FeatureFlagStalenessRecommendedAction, bool) {
	if o == nil || o.RecommendedActions == nil {
		return nil, false
	}
	return &o.RecommendedActions, true
}

// HasRecommendedActions returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasRecommendedActions() bool {
	return o != nil && o.RecommendedActions != nil
}

// SetRecommendedActions gets a reference to the given []FeatureFlagStalenessRecommendedAction and assigns it to the RecommendedActions field.
func (o *FeatureFlagStalenessDetails) SetRecommendedActions(v []FeatureFlagStalenessRecommendedAction) {
	o.RecommendedActions = v
}

// GetSkipStateCheckUntil returns the SkipStateCheckUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagStalenessDetails) GetSkipStateCheckUntil() time.Time {
	if o == nil || o.SkipStateCheckUntil.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SkipStateCheckUntil.Get()
}

// GetSkipStateCheckUntilOk returns a tuple with the SkipStateCheckUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagStalenessDetails) GetSkipStateCheckUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipStateCheckUntil.Get(), o.SkipStateCheckUntil.IsSet()
}

// HasSkipStateCheckUntil returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasSkipStateCheckUntil() bool {
	return o != nil && o.SkipStateCheckUntil.IsSet()
}

// SetSkipStateCheckUntil gets a reference to the given datadog.NullableTime and assigns it to the SkipStateCheckUntil field.
func (o *FeatureFlagStalenessDetails) SetSkipStateCheckUntil(v time.Time) {
	o.SkipStateCheckUntil.Set(&v)
}

// SetSkipStateCheckUntilNil sets the value for SkipStateCheckUntil to be an explicit nil.
func (o *FeatureFlagStalenessDetails) SetSkipStateCheckUntilNil() {
	o.SkipStateCheckUntil.Set(nil)
}

// UnsetSkipStateCheckUntil ensures that no value is present for SkipStateCheckUntil, not even an explicit nil.
func (o *FeatureFlagStalenessDetails) UnsetSkipStateCheckUntil() {
	o.SkipStateCheckUntil.Unset()
}

// GetStaleReason returns the StaleReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagStalenessDetails) GetStaleReason() string {
	if o == nil || o.StaleReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.StaleReason.Get()
}

// GetStaleReasonOk returns a tuple with the StaleReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagStalenessDetails) GetStaleReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaleReason.Get(), o.StaleReason.IsSet()
}

// HasStaleReason returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasStaleReason() bool {
	return o != nil && o.StaleReason.IsSet()
}

// SetStaleReason gets a reference to the given datadog.NullableString and assigns it to the StaleReason field.
func (o *FeatureFlagStalenessDetails) SetStaleReason(v string) {
	o.StaleReason.Set(&v)
}

// SetStaleReasonNil sets the value for StaleReason to be an explicit nil.
func (o *FeatureFlagStalenessDetails) SetStaleReasonNil() {
	o.StaleReason.Set(nil)
}

// UnsetStaleReason ensures that no value is present for StaleReason, not even an explicit nil.
func (o *FeatureFlagStalenessDetails) UnsetStaleReason() {
	o.StaleReason.Unset()
}

// GetStalenessStatus returns the StalenessStatus field value if set, zero value otherwise.
func (o *FeatureFlagStalenessDetails) GetStalenessStatus() string {
	if o == nil || o.StalenessStatus == nil {
		var ret string
		return ret
	}
	return *o.StalenessStatus
}

// GetStalenessStatusOk returns a tuple with the StalenessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessDetails) GetStalenessStatusOk() (*string, bool) {
	if o == nil || o.StalenessStatus == nil {
		return nil, false
	}
	return o.StalenessStatus, true
}

// HasStalenessStatus returns a boolean if a field has been set.
func (o *FeatureFlagStalenessDetails) HasStalenessStatus() bool {
	return o != nil && o.StalenessStatus != nil
}

// SetStalenessStatus gets a reference to the given string and assigns it to the StalenessStatus field.
func (o *FeatureFlagStalenessDetails) SetStalenessStatus(v string) {
	o.StalenessStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagStalenessDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodeReferences != nil {
		toSerialize["code_references"] = o.CodeReferences
	}
	if o.DismissedBy.IsSet() {
		toSerialize["dismissed_by"] = o.DismissedBy.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RecommendedActions != nil {
		toSerialize["recommended_actions"] = o.RecommendedActions
	}
	if o.SkipStateCheckUntil.IsSet() {
		toSerialize["skip_state_check_until"] = o.SkipStateCheckUntil.Get()
	}
	if o.StaleReason.IsSet() {
		toSerialize["stale_reason"] = o.StaleReason.Get()
	}
	if o.StalenessStatus != nil {
		toSerialize["staleness_status"] = o.StalenessStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagStalenessDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeReferences      []FeatureFlagStalenessCodeReference     `json:"code_references,omitempty"`
		DismissedBy         datadog.NullableString                  `json:"dismissed_by,omitempty"`
		Id                  *string                                 `json:"id,omitempty"`
		RecommendedActions  []FeatureFlagStalenessRecommendedAction `json:"recommended_actions,omitempty"`
		SkipStateCheckUntil datadog.NullableTime                    `json:"skip_state_check_until,omitempty"`
		StaleReason         datadog.NullableString                  `json:"stale_reason,omitempty"`
		StalenessStatus     *string                                 `json:"staleness_status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_references", "dismissed_by", "id", "recommended_actions", "skip_state_check_until", "stale_reason", "staleness_status"})
	} else {
		return err
	}
	o.CodeReferences = all.CodeReferences
	o.DismissedBy = all.DismissedBy
	o.Id = all.Id
	o.RecommendedActions = all.RecommendedActions
	o.SkipStateCheckUntil = all.SkipStateCheckUntil
	o.StaleReason = all.StaleReason
	o.StalenessStatus = all.StalenessStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
