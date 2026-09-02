// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagAttributesStalenessDetails Details about the feature flag's staleness status.
type FeatureFlagAttributesStalenessDetails struct {
	// Code references associated with the feature flag.
	CodeReferences []map[string]interface{} `json:"code_references,omitempty"`
	// The ID of the user who dismissed the staleness notification.
	DismissedBy datadog.NullableUUID `json:"dismissed_by,omitempty"`
	// The unique identifier of the staleness details record.
	Id *uuid.UUID `json:"id,omitempty"`
	// Recommended actions to address the feature flag's staleness.
	RecommendedActions []map[string]interface{} `json:"recommended_actions,omitempty"`
	// The timestamp until which staleness checks are skipped.
	SkipStateCheckUntil datadog.NullableTime `json:"skip_state_check_until,omitempty"`
	// The reason the feature flag is considered stale.
	StaleReason datadog.NullableString `json:"stale_reason,omitempty"`
	// The staleness status of the feature flag.
	StalenessStatus *string `json:"staleness_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagAttributesStalenessDetails instantiates a new FeatureFlagAttributesStalenessDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagAttributesStalenessDetails() *FeatureFlagAttributesStalenessDetails {
	this := FeatureFlagAttributesStalenessDetails{}
	return &this
}

// NewFeatureFlagAttributesStalenessDetailsWithDefaults instantiates a new FeatureFlagAttributesStalenessDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagAttributesStalenessDetailsWithDefaults() *FeatureFlagAttributesStalenessDetails {
	this := FeatureFlagAttributesStalenessDetails{}
	return &this
}

// GetCodeReferences returns the CodeReferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributesStalenessDetails) GetCodeReferences() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.CodeReferences
}

// GetCodeReferencesOk returns a tuple with the CodeReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributesStalenessDetails) GetCodeReferencesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.CodeReferences == nil {
		return nil, false
	}
	return &o.CodeReferences, true
}

// HasCodeReferences returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasCodeReferences() bool {
	return o != nil && o.CodeReferences != nil
}

// SetCodeReferences gets a reference to the given []map[string]interface{} and assigns it to the CodeReferences field.
func (o *FeatureFlagAttributesStalenessDetails) SetCodeReferences(v []map[string]interface{}) {
	o.CodeReferences = v
}

// GetDismissedBy returns the DismissedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributesStalenessDetails) GetDismissedBy() uuid.UUID {
	if o == nil || o.DismissedBy.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.DismissedBy.Get()
}

// GetDismissedByOk returns a tuple with the DismissedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributesStalenessDetails) GetDismissedByOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissedBy.Get(), o.DismissedBy.IsSet()
}

// HasDismissedBy returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasDismissedBy() bool {
	return o != nil && o.DismissedBy.IsSet()
}

// SetDismissedBy gets a reference to the given datadog.NullableUUID and assigns it to the DismissedBy field.
func (o *FeatureFlagAttributesStalenessDetails) SetDismissedBy(v uuid.UUID) {
	o.DismissedBy.Set(&v)
}

// SetDismissedByNil sets the value for DismissedBy to be an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) SetDismissedByNil() {
	o.DismissedBy.Set(nil)
}

// UnsetDismissedBy ensures that no value is present for DismissedBy, not even an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) UnsetDismissedBy() {
	o.DismissedBy.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeatureFlagAttributesStalenessDetails) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributesStalenessDetails) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *FeatureFlagAttributesStalenessDetails) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetRecommendedActions returns the RecommendedActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributesStalenessDetails) GetRecommendedActions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.RecommendedActions
}

// GetRecommendedActionsOk returns a tuple with the RecommendedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributesStalenessDetails) GetRecommendedActionsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.RecommendedActions == nil {
		return nil, false
	}
	return &o.RecommendedActions, true
}

// HasRecommendedActions returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasRecommendedActions() bool {
	return o != nil && o.RecommendedActions != nil
}

// SetRecommendedActions gets a reference to the given []map[string]interface{} and assigns it to the RecommendedActions field.
func (o *FeatureFlagAttributesStalenessDetails) SetRecommendedActions(v []map[string]interface{}) {
	o.RecommendedActions = v
}

// GetSkipStateCheckUntil returns the SkipStateCheckUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributesStalenessDetails) GetSkipStateCheckUntil() time.Time {
	if o == nil || o.SkipStateCheckUntil.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SkipStateCheckUntil.Get()
}

// GetSkipStateCheckUntilOk returns a tuple with the SkipStateCheckUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributesStalenessDetails) GetSkipStateCheckUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipStateCheckUntil.Get(), o.SkipStateCheckUntil.IsSet()
}

// HasSkipStateCheckUntil returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasSkipStateCheckUntil() bool {
	return o != nil && o.SkipStateCheckUntil.IsSet()
}

// SetSkipStateCheckUntil gets a reference to the given datadog.NullableTime and assigns it to the SkipStateCheckUntil field.
func (o *FeatureFlagAttributesStalenessDetails) SetSkipStateCheckUntil(v time.Time) {
	o.SkipStateCheckUntil.Set(&v)
}

// SetSkipStateCheckUntilNil sets the value for SkipStateCheckUntil to be an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) SetSkipStateCheckUntilNil() {
	o.SkipStateCheckUntil.Set(nil)
}

// UnsetSkipStateCheckUntil ensures that no value is present for SkipStateCheckUntil, not even an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) UnsetSkipStateCheckUntil() {
	o.SkipStateCheckUntil.Unset()
}

// GetStaleReason returns the StaleReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributesStalenessDetails) GetStaleReason() string {
	if o == nil || o.StaleReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.StaleReason.Get()
}

// GetStaleReasonOk returns a tuple with the StaleReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributesStalenessDetails) GetStaleReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaleReason.Get(), o.StaleReason.IsSet()
}

// HasStaleReason returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasStaleReason() bool {
	return o != nil && o.StaleReason.IsSet()
}

// SetStaleReason gets a reference to the given datadog.NullableString and assigns it to the StaleReason field.
func (o *FeatureFlagAttributesStalenessDetails) SetStaleReason(v string) {
	o.StaleReason.Set(&v)
}

// SetStaleReasonNil sets the value for StaleReason to be an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) SetStaleReasonNil() {
	o.StaleReason.Set(nil)
}

// UnsetStaleReason ensures that no value is present for StaleReason, not even an explicit nil.
func (o *FeatureFlagAttributesStalenessDetails) UnsetStaleReason() {
	o.StaleReason.Unset()
}

// GetStalenessStatus returns the StalenessStatus field value if set, zero value otherwise.
func (o *FeatureFlagAttributesStalenessDetails) GetStalenessStatus() string {
	if o == nil || o.StalenessStatus == nil {
		var ret string
		return ret
	}
	return *o.StalenessStatus
}

// GetStalenessStatusOk returns a tuple with the StalenessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributesStalenessDetails) GetStalenessStatusOk() (*string, bool) {
	if o == nil || o.StalenessStatus == nil {
		return nil, false
	}
	return o.StalenessStatus, true
}

// HasStalenessStatus returns a boolean if a field has been set.
func (o *FeatureFlagAttributesStalenessDetails) HasStalenessStatus() bool {
	return o != nil && o.StalenessStatus != nil
}

// SetStalenessStatus gets a reference to the given string and assigns it to the StalenessStatus field.
func (o *FeatureFlagAttributesStalenessDetails) SetStalenessStatus(v string) {
	o.StalenessStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagAttributesStalenessDetails) MarshalJSON() ([]byte, error) {
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
func (o *FeatureFlagAttributesStalenessDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeReferences      []map[string]interface{} `json:"code_references,omitempty"`
		DismissedBy         datadog.NullableUUID     `json:"dismissed_by,omitempty"`
		Id                  *uuid.UUID               `json:"id,omitempty"`
		RecommendedActions  []map[string]interface{} `json:"recommended_actions,omitempty"`
		SkipStateCheckUntil datadog.NullableTime     `json:"skip_state_check_until,omitempty"`
		StaleReason         datadog.NullableString   `json:"stale_reason,omitempty"`
		StalenessStatus     *string                  `json:"staleness_status,omitempty"`
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

// NullableFeatureFlagAttributesStalenessDetails handles when a null is used for FeatureFlagAttributesStalenessDetails.
type NullableFeatureFlagAttributesStalenessDetails struct {
	value *FeatureFlagAttributesStalenessDetails
	isSet bool
}

// Get returns the associated value.
func (v NullableFeatureFlagAttributesStalenessDetails) Get() *FeatureFlagAttributesStalenessDetails {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFeatureFlagAttributesStalenessDetails) Set(val *FeatureFlagAttributesStalenessDetails) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFeatureFlagAttributesStalenessDetails) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableFeatureFlagAttributesStalenessDetails) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFeatureFlagAttributesStalenessDetails initializes the struct as if Set has been called.
func NewNullableFeatureFlagAttributesStalenessDetails(val *FeatureFlagAttributesStalenessDetails) *NullableFeatureFlagAttributesStalenessDetails {
	return &NullableFeatureFlagAttributesStalenessDetails{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFeatureFlagAttributesStalenessDetails) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFeatureFlagAttributesStalenessDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
