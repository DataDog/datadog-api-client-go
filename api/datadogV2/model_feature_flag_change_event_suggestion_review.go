// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventSuggestionReview Review details of a change request that needed approval.
// Present on review events, which `include_suggestion_events` adds.
// Also present on a change that an approved change request applied, even when `include_suggestion_events` is `false`.
type FeatureFlagChangeEventSuggestionReview struct {
	// Where the change request is in its lifecycle.
	LifecycleStatus FeatureFlagChangeEventSuggestionLifecycleStatus `json:"lifecycle_status"`
	// The comment left by the person who requested the change.
	RequesterComment *string `json:"requester_comment,omitempty"`
	// The comment left by the reviewer.
	ReviewerComment *string `json:"reviewer_comment,omitempty"`
	// The review status of the change request.
	Status FeatureFlagChangeEventSuggestionStatus `json:"status"`
	// The ID of the user who requested the change.
	SuggestionCreatedBy *string `json:"suggestion_created_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagChangeEventSuggestionReview instantiates a new FeatureFlagChangeEventSuggestionReview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagChangeEventSuggestionReview(lifecycleStatus FeatureFlagChangeEventSuggestionLifecycleStatus, status FeatureFlagChangeEventSuggestionStatus) *FeatureFlagChangeEventSuggestionReview {
	this := FeatureFlagChangeEventSuggestionReview{}
	this.LifecycleStatus = lifecycleStatus
	this.Status = status
	return &this
}

// NewFeatureFlagChangeEventSuggestionReviewWithDefaults instantiates a new FeatureFlagChangeEventSuggestionReview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagChangeEventSuggestionReviewWithDefaults() *FeatureFlagChangeEventSuggestionReview {
	this := FeatureFlagChangeEventSuggestionReview{}
	return &this
}

// GetLifecycleStatus returns the LifecycleStatus field value.
func (o *FeatureFlagChangeEventSuggestionReview) GetLifecycleStatus() FeatureFlagChangeEventSuggestionLifecycleStatus {
	if o == nil {
		var ret FeatureFlagChangeEventSuggestionLifecycleStatus
		return ret
	}
	return o.LifecycleStatus
}

// GetLifecycleStatusOk returns a tuple with the LifecycleStatus field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventSuggestionReview) GetLifecycleStatusOk() (*FeatureFlagChangeEventSuggestionLifecycleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifecycleStatus, true
}

// SetLifecycleStatus sets field value.
func (o *FeatureFlagChangeEventSuggestionReview) SetLifecycleStatus(v FeatureFlagChangeEventSuggestionLifecycleStatus) {
	o.LifecycleStatus = v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise.
func (o *FeatureFlagChangeEventSuggestionReview) GetRequesterComment() string {
	if o == nil || o.RequesterComment == nil {
		var ret string
		return ret
	}
	return *o.RequesterComment
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventSuggestionReview) GetRequesterCommentOk() (*string, bool) {
	if o == nil || o.RequesterComment == nil {
		return nil, false
	}
	return o.RequesterComment, true
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventSuggestionReview) HasRequesterComment() bool {
	return o != nil && o.RequesterComment != nil
}

// SetRequesterComment gets a reference to the given string and assigns it to the RequesterComment field.
func (o *FeatureFlagChangeEventSuggestionReview) SetRequesterComment(v string) {
	o.RequesterComment = &v
}

// GetReviewerComment returns the ReviewerComment field value if set, zero value otherwise.
func (o *FeatureFlagChangeEventSuggestionReview) GetReviewerComment() string {
	if o == nil || o.ReviewerComment == nil {
		var ret string
		return ret
	}
	return *o.ReviewerComment
}

// GetReviewerCommentOk returns a tuple with the ReviewerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventSuggestionReview) GetReviewerCommentOk() (*string, bool) {
	if o == nil || o.ReviewerComment == nil {
		return nil, false
	}
	return o.ReviewerComment, true
}

// HasReviewerComment returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventSuggestionReview) HasReviewerComment() bool {
	return o != nil && o.ReviewerComment != nil
}

// SetReviewerComment gets a reference to the given string and assigns it to the ReviewerComment field.
func (o *FeatureFlagChangeEventSuggestionReview) SetReviewerComment(v string) {
	o.ReviewerComment = &v
}

// GetStatus returns the Status field value.
func (o *FeatureFlagChangeEventSuggestionReview) GetStatus() FeatureFlagChangeEventSuggestionStatus {
	if o == nil {
		var ret FeatureFlagChangeEventSuggestionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventSuggestionReview) GetStatusOk() (*FeatureFlagChangeEventSuggestionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *FeatureFlagChangeEventSuggestionReview) SetStatus(v FeatureFlagChangeEventSuggestionStatus) {
	o.Status = v
}

// GetSuggestionCreatedBy returns the SuggestionCreatedBy field value if set, zero value otherwise.
func (o *FeatureFlagChangeEventSuggestionReview) GetSuggestionCreatedBy() string {
	if o == nil || o.SuggestionCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.SuggestionCreatedBy
}

// GetSuggestionCreatedByOk returns a tuple with the SuggestionCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventSuggestionReview) GetSuggestionCreatedByOk() (*string, bool) {
	if o == nil || o.SuggestionCreatedBy == nil {
		return nil, false
	}
	return o.SuggestionCreatedBy, true
}

// HasSuggestionCreatedBy returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventSuggestionReview) HasSuggestionCreatedBy() bool {
	return o != nil && o.SuggestionCreatedBy != nil
}

// SetSuggestionCreatedBy gets a reference to the given string and assigns it to the SuggestionCreatedBy field.
func (o *FeatureFlagChangeEventSuggestionReview) SetSuggestionCreatedBy(v string) {
	o.SuggestionCreatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagChangeEventSuggestionReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["lifecycle_status"] = o.LifecycleStatus
	if o.RequesterComment != nil {
		toSerialize["requester_comment"] = o.RequesterComment
	}
	if o.ReviewerComment != nil {
		toSerialize["reviewer_comment"] = o.ReviewerComment
	}
	toSerialize["status"] = o.Status
	if o.SuggestionCreatedBy != nil {
		toSerialize["suggestion_created_by"] = o.SuggestionCreatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagChangeEventSuggestionReview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LifecycleStatus     *FeatureFlagChangeEventSuggestionLifecycleStatus `json:"lifecycle_status"`
		RequesterComment    *string                                          `json:"requester_comment,omitempty"`
		ReviewerComment     *string                                          `json:"reviewer_comment,omitempty"`
		Status              *FeatureFlagChangeEventSuggestionStatus          `json:"status"`
		SuggestionCreatedBy *string                                          `json:"suggestion_created_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LifecycleStatus == nil {
		return fmt.Errorf("required field lifecycle_status missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lifecycle_status", "requester_comment", "reviewer_comment", "status", "suggestion_created_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.LifecycleStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.LifecycleStatus = *all.LifecycleStatus
	}
	o.RequesterComment = all.RequesterComment
	o.ReviewerComment = all.ReviewerComment
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.SuggestionCreatedBy = all.SuggestionCreatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
