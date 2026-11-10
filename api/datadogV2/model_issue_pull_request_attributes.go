// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuePullRequestAttributes Object containing the information of a pull request linked to an issue.
type IssuePullRequestAttributes struct {
	// Timestamp of when the pull request was closed. Omitted if the pull request is not closed.
	ClosedAt *time.Time `json:"closed_at,omitempty"`
	// Timestamp of when the pull request was linked to the issue.
	LinkedAt *time.Time `json:"linked_at,omitempty"`
	// Timestamp of when the pull request was merged. Omitted if the pull request is not merged.
	MergedAt *time.Time `json:"merged_at,omitempty"`
	// Why the pull request is linked to the issue. `fix` means the pull request fixes the issue, and `trigger` means the pull request introduced the issue.
	Reason *IssuePullRequestReason `json:"reason,omitempty"`
	// URL of the pull request.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssuePullRequestAttributes instantiates a new IssuePullRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssuePullRequestAttributes() *IssuePullRequestAttributes {
	this := IssuePullRequestAttributes{}
	return &this
}

// NewIssuePullRequestAttributesWithDefaults instantiates a new IssuePullRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssuePullRequestAttributesWithDefaults() *IssuePullRequestAttributes {
	this := IssuePullRequestAttributes{}
	return &this
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *IssuePullRequestAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuePullRequestAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || o.ClosedAt == nil {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *IssuePullRequestAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt != nil
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *IssuePullRequestAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetLinkedAt returns the LinkedAt field value if set, zero value otherwise.
func (o *IssuePullRequestAttributes) GetLinkedAt() time.Time {
	if o == nil || o.LinkedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LinkedAt
}

// GetLinkedAtOk returns a tuple with the LinkedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuePullRequestAttributes) GetLinkedAtOk() (*time.Time, bool) {
	if o == nil || o.LinkedAt == nil {
		return nil, false
	}
	return o.LinkedAt, true
}

// HasLinkedAt returns a boolean if a field has been set.
func (o *IssuePullRequestAttributes) HasLinkedAt() bool {
	return o != nil && o.LinkedAt != nil
}

// SetLinkedAt gets a reference to the given time.Time and assigns it to the LinkedAt field.
func (o *IssuePullRequestAttributes) SetLinkedAt(v time.Time) {
	o.LinkedAt = &v
}

// GetMergedAt returns the MergedAt field value if set, zero value otherwise.
func (o *IssuePullRequestAttributes) GetMergedAt() time.Time {
	if o == nil || o.MergedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.MergedAt
}

// GetMergedAtOk returns a tuple with the MergedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuePullRequestAttributes) GetMergedAtOk() (*time.Time, bool) {
	if o == nil || o.MergedAt == nil {
		return nil, false
	}
	return o.MergedAt, true
}

// HasMergedAt returns a boolean if a field has been set.
func (o *IssuePullRequestAttributes) HasMergedAt() bool {
	return o != nil && o.MergedAt != nil
}

// SetMergedAt gets a reference to the given time.Time and assigns it to the MergedAt field.
func (o *IssuePullRequestAttributes) SetMergedAt(v time.Time) {
	o.MergedAt = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *IssuePullRequestAttributes) GetReason() IssuePullRequestReason {
	if o == nil || o.Reason == nil {
		var ret IssuePullRequestReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuePullRequestAttributes) GetReasonOk() (*IssuePullRequestReason, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *IssuePullRequestAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given IssuePullRequestReason and assigns it to the Reason field.
func (o *IssuePullRequestAttributes) SetReason(v IssuePullRequestReason) {
	o.Reason = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IssuePullRequestAttributes) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuePullRequestAttributes) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IssuePullRequestAttributes) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IssuePullRequestAttributes) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssuePullRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ClosedAt != nil {
		if o.ClosedAt.Nanosecond() == 0 {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LinkedAt != nil {
		if o.LinkedAt.Nanosecond() == 0 {
			toSerialize["linked_at"] = o.LinkedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["linked_at"] = o.LinkedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MergedAt != nil {
		if o.MergedAt.Nanosecond() == 0 {
			toSerialize["merged_at"] = o.MergedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["merged_at"] = o.MergedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssuePullRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClosedAt *time.Time              `json:"closed_at,omitempty"`
		LinkedAt *time.Time              `json:"linked_at,omitempty"`
		MergedAt *time.Time              `json:"merged_at,omitempty"`
		Reason   *IssuePullRequestReason `json:"reason,omitempty"`
		Url      *string                 `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"closed_at", "linked_at", "merged_at", "reason", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClosedAt = all.ClosedAt
	o.LinkedAt = all.LinkedAt
	o.MergedAt = all.MergedAt
	if all.Reason != nil && !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = all.Reason
	}
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
