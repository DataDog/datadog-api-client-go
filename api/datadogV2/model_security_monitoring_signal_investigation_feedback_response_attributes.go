// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationFeedbackResponseAttributes Attributes of investigation feedback.
type SecurityMonitoringSignalInvestigationFeedbackResponseAttributes struct {
	// The feedback text.
	Feedback string `json:"feedback"`
	// Structured feedback content.
	FeedbackContent []SecurityMonitoringSignalInvestigationFeedbackSection `json:"feedback_content,omitempty"`
	// The unique ID of the investigation.
	InvestigationId string `json:"investigation_id"`
	// The rating value.
	Rating *string `json:"rating,omitempty"`
	// The unique ID of the security signal.
	SignalId string `json:"signal_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationFeedbackResponseAttributes instantiates a new SecurityMonitoringSignalInvestigationFeedbackResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationFeedbackResponseAttributes(feedback string, investigationId string, signalId string) *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes {
	this := SecurityMonitoringSignalInvestigationFeedbackResponseAttributes{}
	this.Feedback = feedback
	this.InvestigationId = investigationId
	this.SignalId = signalId
	return &this
}

// NewSecurityMonitoringSignalInvestigationFeedbackResponseAttributesWithDefaults instantiates a new SecurityMonitoringSignalInvestigationFeedbackResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationFeedbackResponseAttributesWithDefaults() *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes {
	this := SecurityMonitoringSignalInvestigationFeedbackResponseAttributes{}
	return &this
}

// GetFeedback returns the Feedback field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetFeedback() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetFeedbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedback, true
}

// SetFeedback sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) SetFeedback(v string) {
	o.Feedback = v
}

// GetFeedbackContent returns the FeedbackContent field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetFeedbackContent() []SecurityMonitoringSignalInvestigationFeedbackSection {
	if o == nil || o.FeedbackContent == nil {
		var ret []SecurityMonitoringSignalInvestigationFeedbackSection
		return ret
	}
	return o.FeedbackContent
}

// GetFeedbackContentOk returns a tuple with the FeedbackContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetFeedbackContentOk() (*[]SecurityMonitoringSignalInvestigationFeedbackSection, bool) {
	if o == nil || o.FeedbackContent == nil {
		return nil, false
	}
	return &o.FeedbackContent, true
}

// HasFeedbackContent returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) HasFeedbackContent() bool {
	return o != nil && o.FeedbackContent != nil
}

// SetFeedbackContent gets a reference to the given []SecurityMonitoringSignalInvestigationFeedbackSection and assigns it to the FeedbackContent field.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) SetFeedbackContent(v []SecurityMonitoringSignalInvestigationFeedbackSection) {
	o.FeedbackContent = v
}

// GetInvestigationId returns the InvestigationId field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetInvestigationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InvestigationId
}

// GetInvestigationIdOk returns a tuple with the InvestigationId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetInvestigationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestigationId, true
}

// SetInvestigationId sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) SetInvestigationId(v string) {
	o.InvestigationId = v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetRating() string {
	if o == nil || o.Rating == nil {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetRatingOk() (*string, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) HasRating() bool {
	return o != nil && o.Rating != nil
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) SetRating(v string) {
	o.Rating = &v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feedback"] = o.Feedback
	if o.FeedbackContent != nil {
		toSerialize["feedback_content"] = o.FeedbackContent
	}
	toSerialize["investigation_id"] = o.InvestigationId
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	toSerialize["signal_id"] = o.SignalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationFeedbackResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Feedback        *string                                                `json:"feedback"`
		FeedbackContent []SecurityMonitoringSignalInvestigationFeedbackSection `json:"feedback_content,omitempty"`
		InvestigationId *string                                                `json:"investigation_id"`
		Rating          *string                                                `json:"rating,omitempty"`
		SignalId        *string                                                `json:"signal_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Feedback == nil {
		return fmt.Errorf("required field feedback missing")
	}
	if all.InvestigationId == nil {
		return fmt.Errorf("required field investigation_id missing")
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feedback", "feedback_content", "investigation_id", "rating", "signal_id"})
	} else {
		return err
	}
	o.Feedback = *all.Feedback
	o.FeedbackContent = all.FeedbackContent
	o.InvestigationId = *all.InvestigationId
	o.Rating = all.Rating
	o.SignalId = *all.SignalId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
