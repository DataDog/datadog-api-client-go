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

// ExperimentsBanditPolicyHistoryV2DTODataAttributes A traffic split the bandit served, with the same fields as `active_policy`.
type ExperimentsBanditPolicyHistoryV2DTODataAttributes struct {
	// When the policy was saved. SDKs pick it up on their next configuration fetch.
	AppliedAt time.Time `json:"applied_at"`
	// The configuration version the policy was made for.
	ConfigVersion int64 `json:"config_version"`
	// Whether this is the policy the bandit serves now, the one Get bandit returns as `active_policy`. A paused bandit keeps serving it, and an ended bandit serves its fallback.
	IsActive bool `json:"is_active"`
	// `WARMUP` serves the arms equally. `ADAPTIVE` serves the initial split or a split from a refresh. `FALLBACK` serves one variant after the bandit ends.
	Kind ExperimentsBanditPolicyKind `json:"kind"`
	// The policy's place in the bandit's policy history. It follows `applied_at` and orders policies saved at the same time.
	PolicyNumber int64 `json:"policy_number"`
	// Why the policy was applied.
	Reason ExperimentsBanditPolicyReason `json:"reason"`
	// The refresh that proposed the policy. `null` for policies no refresh proposed, such as the first split and the policy served after the bandit ends.
	RefreshId datadog.NullableUUID `json:"refresh_id"`
	// Each variant's share of traffic. The probabilities add up to 1. A variant served as an arm keeps the key the arm had, even if the arm was removed later.
	Variants []ExperimentsBanditPolicyVariantV2 `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditPolicyHistoryV2DTODataAttributes instantiates a new ExperimentsBanditPolicyHistoryV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditPolicyHistoryV2DTODataAttributes(appliedAt time.Time, configVersion int64, isActive bool, kind ExperimentsBanditPolicyKind, policyNumber int64, reason ExperimentsBanditPolicyReason, refreshId datadog.NullableUUID, variants []ExperimentsBanditPolicyVariantV2) *ExperimentsBanditPolicyHistoryV2DTODataAttributes {
	this := ExperimentsBanditPolicyHistoryV2DTODataAttributes{}
	this.AppliedAt = appliedAt
	this.ConfigVersion = configVersion
	this.IsActive = isActive
	this.Kind = kind
	this.PolicyNumber = policyNumber
	this.Reason = reason
	this.RefreshId = refreshId
	this.Variants = variants
	return &this
}

// NewExperimentsBanditPolicyHistoryV2DTODataAttributesWithDefaults instantiates a new ExperimentsBanditPolicyHistoryV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditPolicyHistoryV2DTODataAttributesWithDefaults() *ExperimentsBanditPolicyHistoryV2DTODataAttributes {
	this := ExperimentsBanditPolicyHistoryV2DTODataAttributes{}
	return &this
}

// GetAppliedAt returns the AppliedAt field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetAppliedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.AppliedAt
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetAppliedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppliedAt, true
}

// SetAppliedAt sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetAppliedAt(v time.Time) {
	o.AppliedAt = v
}

// GetConfigVersion returns the ConfigVersion field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetConfigVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigVersion, true
}

// SetConfigVersion sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion = v
}

// GetIsActive returns the IsActive field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetIsActive(v bool) {
	o.IsActive = v
}

// GetKind returns the Kind field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetKind() ExperimentsBanditPolicyKind {
	if o == nil {
		var ret ExperimentsBanditPolicyKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetKindOk() (*ExperimentsBanditPolicyKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetKind(v ExperimentsBanditPolicyKind) {
	o.Kind = v
}

// GetPolicyNumber returns the PolicyNumber field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetPolicyNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PolicyNumber
}

// GetPolicyNumberOk returns a tuple with the PolicyNumber field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetPolicyNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyNumber, true
}

// SetPolicyNumber sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetPolicyNumber(v int64) {
	o.PolicyNumber = v
}

// GetReason returns the Reason field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetReason() ExperimentsBanditPolicyReason {
	if o == nil {
		var ret ExperimentsBanditPolicyReason
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetReasonOk() (*ExperimentsBanditPolicyReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetReason(v ExperimentsBanditPolicyReason) {
	o.Reason = v
}

// GetRefreshId returns the RefreshId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetRefreshId() uuid.UUID {
	if o == nil || o.RefreshId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.RefreshId.Get()
}

// GetRefreshIdOk returns a tuple with the RefreshId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetRefreshIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshId.Get(), o.RefreshId.IsSet()
}

// SetRefreshId sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetRefreshId(v uuid.UUID) {
	o.RefreshId.Set(&v)
}

// GetVariants returns the Variants field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetVariants() []ExperimentsBanditPolicyVariantV2 {
	if o == nil {
		var ret []ExperimentsBanditPolicyVariantV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) GetVariantsOk() (*[]ExperimentsBanditPolicyVariantV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) SetVariants(v []ExperimentsBanditPolicyVariantV2) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditPolicyHistoryV2DTODataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["is_active"] = o.IsActive
	toSerialize["kind"] = o.Kind
	toSerialize["policy_number"] = o.PolicyNumber
	toSerialize["reason"] = o.Reason
	toSerialize["refresh_id"] = o.RefreshId.Get()
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditPolicyHistoryV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppliedAt     *time.Time                          `json:"applied_at"`
		ConfigVersion *int64                              `json:"config_version"`
		IsActive      *bool                               `json:"is_active"`
		Kind          *ExperimentsBanditPolicyKind        `json:"kind"`
		PolicyNumber  *int64                              `json:"policy_number"`
		Reason        *ExperimentsBanditPolicyReason      `json:"reason"`
		RefreshId     datadog.NullableUUID                `json:"refresh_id"`
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
	if all.IsActive == nil {
		return fmt.Errorf("required field is_active missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.PolicyNumber == nil {
		return fmt.Errorf("required field policy_number missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if !all.RefreshId.IsSet() {
		return fmt.Errorf("required field refresh_id missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applied_at", "config_version", "is_active", "kind", "policy_number", "reason", "refresh_id", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AppliedAt = *all.AppliedAt
	o.ConfigVersion = *all.ConfigVersion
	o.IsActive = *all.IsActive
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	o.PolicyNumber = *all.PolicyNumber
	if !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = *all.Reason
	}
	o.RefreshId = all.RefreshId
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
