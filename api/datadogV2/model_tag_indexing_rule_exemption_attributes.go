// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionAttributes Attributes of a tag indexing rule exemption.
type TagIndexingRuleExemptionAttributes struct {
	// Timestamp when the exemption was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Handle of the user who created the exemption.
	CreatedByHandle *string `json:"created_by_handle,omitempty"`
	// Estimated change in indexed cardinality or points if the exemption is removed and the matching tag indexing
	// rule applies to the metric. A negative value indicates a reduction in indexed volume, a positive value
	// indicates an increase, and zero indicates no change. A `null` value indicates that the estimate is unavailable.
	EstimatedChangeInVolume datadog.NullableInt64 `json:"estimated_change_in_volume,omitempty"`
	// Discriminates between an explicit exemption (`exemption`) and a pre-existing legacy tag configuration acting as an implicit exclusion (`legacy_tag_configuration`).
	Kind *string `json:"kind,omitempty"`
	// Tag keys used to query this metric in monitors, dashboards, notebooks, and SLOs that the matching rule would
	// remove or not include if the exemption is removed.
	MissingIndexedTags []string `json:"missing_indexed_tags,omitempty"`
	// The reason the metric is exempt from tag indexing rules.
	Reason *string `json:"reason,omitempty"`
	// Indicates whether `estimated_change_in_volume` is available and, if not, why. `ok` indicates a fresh estimate.
	// `unindexed_tags` indicates that the matching rule uses tags that are not indexed for the metric.
	// `unavailable` indicates that the estimate is stale, was not computed, or could not be retrieved.
	VolumeStatus *TagIndexingRuleExemptionVolumeStatus `json:"volume_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleExemptionAttributes instantiates a new TagIndexingRuleExemptionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleExemptionAttributes() *TagIndexingRuleExemptionAttributes {
	this := TagIndexingRuleExemptionAttributes{}
	return &this
}

// NewTagIndexingRuleExemptionAttributesWithDefaults instantiates a new TagIndexingRuleExemptionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleExemptionAttributesWithDefaults() *TagIndexingRuleExemptionAttributes {
	this := TagIndexingRuleExemptionAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TagIndexingRuleExemptionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *TagIndexingRuleExemptionAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetEstimatedChangeInVolume returns the EstimatedChangeInVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagIndexingRuleExemptionAttributes) GetEstimatedChangeInVolume() int64 {
	if o == nil || o.EstimatedChangeInVolume.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedChangeInVolume.Get()
}

// GetEstimatedChangeInVolumeOk returns a tuple with the EstimatedChangeInVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagIndexingRuleExemptionAttributes) GetEstimatedChangeInVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedChangeInVolume.Get(), o.EstimatedChangeInVolume.IsSet()
}

// HasEstimatedChangeInVolume returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasEstimatedChangeInVolume() bool {
	return o != nil && o.EstimatedChangeInVolume.IsSet()
}

// SetEstimatedChangeInVolume gets a reference to the given datadog.NullableInt64 and assigns it to the EstimatedChangeInVolume field.
func (o *TagIndexingRuleExemptionAttributes) SetEstimatedChangeInVolume(v int64) {
	o.EstimatedChangeInVolume.Set(&v)
}

// SetEstimatedChangeInVolumeNil sets the value for EstimatedChangeInVolume to be an explicit nil.
func (o *TagIndexingRuleExemptionAttributes) SetEstimatedChangeInVolumeNil() {
	o.EstimatedChangeInVolume.Set(nil)
}

// UnsetEstimatedChangeInVolume ensures that no value is present for EstimatedChangeInVolume, not even an explicit nil.
func (o *TagIndexingRuleExemptionAttributes) UnsetEstimatedChangeInVolume() {
	o.EstimatedChangeInVolume.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TagIndexingRuleExemptionAttributes) SetKind(v string) {
	o.Kind = &v
}

// GetMissingIndexedTags returns the MissingIndexedTags field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetMissingIndexedTags() []string {
	if o == nil || o.MissingIndexedTags == nil {
		var ret []string
		return ret
	}
	return o.MissingIndexedTags
}

// GetMissingIndexedTagsOk returns a tuple with the MissingIndexedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetMissingIndexedTagsOk() (*[]string, bool) {
	if o == nil || o.MissingIndexedTags == nil {
		return nil, false
	}
	return &o.MissingIndexedTags, true
}

// HasMissingIndexedTags returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasMissingIndexedTags() bool {
	return o != nil && o.MissingIndexedTags != nil
}

// SetMissingIndexedTags gets a reference to the given []string and assigns it to the MissingIndexedTags field.
func (o *TagIndexingRuleExemptionAttributes) SetMissingIndexedTags(v []string) {
	o.MissingIndexedTags = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TagIndexingRuleExemptionAttributes) SetReason(v string) {
	o.Reason = &v
}

// GetVolumeStatus returns the VolumeStatus field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetVolumeStatus() TagIndexingRuleExemptionVolumeStatus {
	if o == nil || o.VolumeStatus == nil {
		var ret TagIndexingRuleExemptionVolumeStatus
		return ret
	}
	return *o.VolumeStatus
}

// GetVolumeStatusOk returns a tuple with the VolumeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetVolumeStatusOk() (*TagIndexingRuleExemptionVolumeStatus, bool) {
	if o == nil || o.VolumeStatus == nil {
		return nil, false
	}
	return o.VolumeStatus, true
}

// HasVolumeStatus returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasVolumeStatus() bool {
	return o != nil && o.VolumeStatus != nil
}

// SetVolumeStatus gets a reference to the given TagIndexingRuleExemptionVolumeStatus and assigns it to the VolumeStatus field.
func (o *TagIndexingRuleExemptionAttributes) SetVolumeStatus(v TagIndexingRuleExemptionVolumeStatus) {
	o.VolumeStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleExemptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedByHandle != nil {
		toSerialize["created_by_handle"] = o.CreatedByHandle
	}
	if o.EstimatedChangeInVolume.IsSet() {
		toSerialize["estimated_change_in_volume"] = o.EstimatedChangeInVolume.Get()
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.MissingIndexedTags != nil {
		toSerialize["missing_indexed_tags"] = o.MissingIndexedTags
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.VolumeStatus != nil {
		toSerialize["volume_status"] = o.VolumeStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleExemptionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt               *time.Time                            `json:"created_at,omitempty"`
		CreatedByHandle         *string                               `json:"created_by_handle,omitempty"`
		EstimatedChangeInVolume datadog.NullableInt64                 `json:"estimated_change_in_volume,omitempty"`
		Kind                    *string                               `json:"kind,omitempty"`
		MissingIndexedTags      []string                              `json:"missing_indexed_tags,omitempty"`
		Reason                  *string                               `json:"reason,omitempty"`
		VolumeStatus            *TagIndexingRuleExemptionVolumeStatus `json:"volume_status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by_handle", "estimated_change_in_volume", "kind", "missing_indexed_tags", "reason", "volume_status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.EstimatedChangeInVolume = all.EstimatedChangeInVolume
	o.Kind = all.Kind
	o.MissingIndexedTags = all.MissingIndexedTags
	o.Reason = all.Reason
	if all.VolumeStatus != nil && !all.VolumeStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.VolumeStatus = all.VolumeStatus
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
