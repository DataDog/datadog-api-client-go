// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationFeedbackRequestAttributes Attributes for investigation feedback.
type SecurityMonitoringSignalInvestigationFeedbackRequestAttributes struct {
	// The feedback text.
	Feedback string `json:"feedback"`
	// Structured feedback content.
	FeedbackContent []SecurityMonitoringSignalInvestigationFeedbackSection `json:"feedback_content,omitempty"`
	// Whether the feedback is incomplete.
	Incomplete *bool `json:"incomplete,omitempty"`
	// The rating value.
	Rating *string `json:"rating,omitempty"`
	// The unique ID of the security signal.
	SignalId string `json:"signal_id"`
	// The type of feedback.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationFeedbackRequestAttributes instantiates a new SecurityMonitoringSignalInvestigationFeedbackRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationFeedbackRequestAttributes(feedback string, signalId string) *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes {
	this := SecurityMonitoringSignalInvestigationFeedbackRequestAttributes{}
	this.Feedback = feedback
	this.SignalId = signalId
	return &this
}

// NewSecurityMonitoringSignalInvestigationFeedbackRequestAttributesWithDefaults instantiates a new SecurityMonitoringSignalInvestigationFeedbackRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationFeedbackRequestAttributesWithDefaults() *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes {
	this := SecurityMonitoringSignalInvestigationFeedbackRequestAttributes{}
	return &this
}

// GetFeedback returns the Feedback field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetFeedback() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetFeedbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedback, true
}

// SetFeedback sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetFeedback(v string) {
	o.Feedback = v
}

// GetFeedbackContent returns the FeedbackContent field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetFeedbackContent() []SecurityMonitoringSignalInvestigationFeedbackSection {
	if o == nil || o.FeedbackContent == nil {
		var ret []SecurityMonitoringSignalInvestigationFeedbackSection
		return ret
	}
	return o.FeedbackContent
}

// GetFeedbackContentOk returns a tuple with the FeedbackContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetFeedbackContentOk() (*[]SecurityMonitoringSignalInvestigationFeedbackSection, bool) {
	if o == nil || o.FeedbackContent == nil {
		return nil, false
	}
	return &o.FeedbackContent, true
}

// HasFeedbackContent returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) HasFeedbackContent() bool {
	return o != nil && o.FeedbackContent != nil
}

// SetFeedbackContent gets a reference to the given []SecurityMonitoringSignalInvestigationFeedbackSection and assigns it to the FeedbackContent field.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetFeedbackContent(v []SecurityMonitoringSignalInvestigationFeedbackSection) {
	o.FeedbackContent = v
}

// GetIncomplete returns the Incomplete field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetIncomplete() bool {
	if o == nil || o.Incomplete == nil {
		var ret bool
		return ret
	}
	return *o.Incomplete
}

// GetIncompleteOk returns a tuple with the Incomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetIncompleteOk() (*bool, bool) {
	if o == nil || o.Incomplete == nil {
		return nil, false
	}
	return o.Incomplete, true
}

// HasIncomplete returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) HasIncomplete() bool {
	return o != nil && o.Incomplete != nil
}

// SetIncomplete gets a reference to the given bool and assigns it to the Incomplete field.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetIncomplete(v bool) {
	o.Incomplete = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetRating() string {
	if o == nil || o.Rating == nil {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetRatingOk() (*string, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) HasRating() bool {
	return o != nil && o.Rating != nil
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetRating(v string) {
	o.Rating = &v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feedback"] = o.Feedback
	if o.FeedbackContent != nil {
		toSerialize["feedback_content"] = o.FeedbackContent
	}
	if o.Incomplete != nil {
		toSerialize["incomplete"] = o.Incomplete
	}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	toSerialize["signal_id"] = o.SignalId
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationFeedbackRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Feedback        *string                                                `json:"feedback"`
		FeedbackContent []SecurityMonitoringSignalInvestigationFeedbackSection `json:"feedback_content,omitempty"`
		Incomplete      *bool                                                  `json:"incomplete,omitempty"`
		Rating          *string                                                `json:"rating,omitempty"`
		SignalId        *string                                                `json:"signal_id"`
		Type            *string                                                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Feedback == nil {
		return fmt.Errorf("required field feedback missing")
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feedback", "feedback_content", "incomplete", "rating", "signal_id", "type"})
	} else {
		return err
	}
	o.Feedback = *all.Feedback
	o.FeedbackContent = all.FeedbackContent
	o.Incomplete = all.Incomplete
	o.Rating = all.Rating
	o.SignalId = *all.SignalId
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
