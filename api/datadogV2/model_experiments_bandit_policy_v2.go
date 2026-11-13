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

// ExperimentsBanditPolicyV2 The policy the bandit serves. Null until the bandit starts.
type ExperimentsBanditPolicyV2 struct {
	// When the policy started being served.
	AppliedAt time.Time `json:"applied_at"`
	// The configuration version the policy was made for.
	ConfigVersion int64 `json:"config_version"`
	// The policy ID. Each exposure logs it as `bandit.policy_id`.
	Id uuid.UUID `json:"id"`
	// WARMUP serves the arms equally. ADAPTIVE serves the initial split or a split from a refresh. FALLBACK serves one variant after the bandit ends.
	Kind ExperimentsBanditPolicyKind `json:"kind"`
	// Why the policy was applied.
	Reason ExperimentsBanditPolicyReason `json:"reason"`
	// Each variant's share of traffic. The probabilities add up to 1.
	Variants []ExperimentsBanditPolicyVariantV2 `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditPolicyV2 instantiates a new ExperimentsBanditPolicyV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditPolicyV2(appliedAt time.Time, configVersion int64, id uuid.UUID, kind ExperimentsBanditPolicyKind, reason ExperimentsBanditPolicyReason, variants []ExperimentsBanditPolicyVariantV2) *ExperimentsBanditPolicyV2 {
	this := ExperimentsBanditPolicyV2{}
	this.AppliedAt = appliedAt
	this.ConfigVersion = configVersion
	this.Id = id
	this.Kind = kind
	this.Reason = reason
	this.Variants = variants
	return &this
}

// NewExperimentsBanditPolicyV2WithDefaults instantiates a new ExperimentsBanditPolicyV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditPolicyV2WithDefaults() *ExperimentsBanditPolicyV2 {
	this := ExperimentsBanditPolicyV2{}
	return &this
}

// GetAppliedAt returns the AppliedAt field value.
func (o *ExperimentsBanditPolicyV2) GetAppliedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.AppliedAt
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetAppliedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppliedAt, true
}

// SetAppliedAt sets field value.
func (o *ExperimentsBanditPolicyV2) SetAppliedAt(v time.Time) {
	o.AppliedAt = v
}

// GetConfigVersion returns the ConfigVersion field value.
func (o *ExperimentsBanditPolicyV2) GetConfigVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetConfigVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigVersion, true
}

// SetConfigVersion sets field value.
func (o *ExperimentsBanditPolicyV2) SetConfigVersion(v int64) {
	o.ConfigVersion = v
}

// GetId returns the Id field value.
func (o *ExperimentsBanditPolicyV2) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsBanditPolicyV2) SetId(v uuid.UUID) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *ExperimentsBanditPolicyV2) GetKind() ExperimentsBanditPolicyKind {
	if o == nil {
		var ret ExperimentsBanditPolicyKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetKindOk() (*ExperimentsBanditPolicyKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *ExperimentsBanditPolicyV2) SetKind(v ExperimentsBanditPolicyKind) {
	o.Kind = v
}

// GetReason returns the Reason field value.
func (o *ExperimentsBanditPolicyV2) GetReason() ExperimentsBanditPolicyReason {
	if o == nil {
		var ret ExperimentsBanditPolicyReason
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetReasonOk() (*ExperimentsBanditPolicyReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ExperimentsBanditPolicyV2) SetReason(v ExperimentsBanditPolicyReason) {
	o.Reason = v
}

// GetVariants returns the Variants field value.
func (o *ExperimentsBanditPolicyV2) GetVariants() []ExperimentsBanditPolicyVariantV2 {
	if o == nil {
		var ret []ExperimentsBanditPolicyVariantV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyV2) GetVariantsOk() (*[]ExperimentsBanditPolicyVariantV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *ExperimentsBanditPolicyV2) SetVariants(v []ExperimentsBanditPolicyVariantV2) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditPolicyV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppliedAt.Nanosecond() == 0 {
		toSerialize["applied_at"] = o.AppliedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["applied_at"] = o.AppliedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["config_version"] = o.ConfigVersion
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["reason"] = o.Reason
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditPolicyV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppliedAt     *time.Time                          `json:"applied_at"`
		ConfigVersion *int64                              `json:"config_version"`
		Id            *uuid.UUID                          `json:"id"`
		Kind          *ExperimentsBanditPolicyKind        `json:"kind"`
		Reason        *ExperimentsBanditPolicyReason      `json:"reason"`
		Variants      *[]ExperimentsBanditPolicyVariantV2 `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppliedAt == nil {
		return fmt.Errorf("required field applied_at missing")
	}
	if all.ConfigVersion == nil {
		return fmt.Errorf("required field config_version missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applied_at", "config_version", "id", "kind", "reason", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AppliedAt = *all.AppliedAt
	o.ConfigVersion = *all.ConfigVersion
	o.Id = *all.Id
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	if !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = *all.Reason
	}
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableExperimentsBanditPolicyV2 handles when a null is used for ExperimentsBanditPolicyV2.
type NullableExperimentsBanditPolicyV2 struct {
	value *ExperimentsBanditPolicyV2
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditPolicyV2) Get() *ExperimentsBanditPolicyV2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditPolicyV2) Set(val *ExperimentsBanditPolicyV2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditPolicyV2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsBanditPolicyV2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditPolicyV2 initializes the struct as if Set has been called.
func NewNullableExperimentsBanditPolicyV2(val *ExperimentsBanditPolicyV2) *NullableExperimentsBanditPolicyV2 {
	return &NullableExperimentsBanditPolicyV2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditPolicyV2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditPolicyV2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
