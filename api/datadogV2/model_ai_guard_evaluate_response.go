// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardEvaluateResponse The result of the AI Guard evaluation.
type AIGuardEvaluateResponse struct {
	// The action recommendation from the AI Guard evaluation.
	Action AIGuardAction `json:"action"`
	// The overall threat probability score across all evaluated tags.
	GlobalProb *float64 `json:"global_prob,omitempty"`
	// Whether blocking mode is enabled for this organization.
	IsBlockingEnabled bool `json:"is_blocking_enabled"`
	// A human-readable explanation of the action recommendation.
	Reason string `json:"reason"`
	// Sensitive data findings detected in the evaluated conversation.
	SdsFindings []AIGuardSdsFinding `json:"sds_findings,omitempty"`
	// Probability scores for each evaluated threat tag.
	TagProbs map[string]float64 `json:"tag_probs"`
	// Security threat tags detected in the evaluated conversation.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardEvaluateResponse instantiates a new AIGuardEvaluateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardEvaluateResponse(action AIGuardAction, isBlockingEnabled bool, reason string, tagProbs map[string]float64, tags []string) *AIGuardEvaluateResponse {
	this := AIGuardEvaluateResponse{}
	this.Action = action
	this.IsBlockingEnabled = isBlockingEnabled
	this.Reason = reason
	this.TagProbs = tagProbs
	this.Tags = tags
	return &this
}

// NewAIGuardEvaluateResponseWithDefaults instantiates a new AIGuardEvaluateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardEvaluateResponseWithDefaults() *AIGuardEvaluateResponse {
	this := AIGuardEvaluateResponse{}
	return &this
}

// GetAction returns the Action field value.
func (o *AIGuardEvaluateResponse) GetAction() AIGuardAction {
	if o == nil {
		var ret AIGuardAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetActionOk() (*AIGuardAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *AIGuardEvaluateResponse) SetAction(v AIGuardAction) {
	o.Action = v
}

// GetGlobalProb returns the GlobalProb field value if set, zero value otherwise.
func (o *AIGuardEvaluateResponse) GetGlobalProb() float64 {
	if o == nil || o.GlobalProb == nil {
		var ret float64
		return ret
	}
	return *o.GlobalProb
}

// GetGlobalProbOk returns a tuple with the GlobalProb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetGlobalProbOk() (*float64, bool) {
	if o == nil || o.GlobalProb == nil {
		return nil, false
	}
	return o.GlobalProb, true
}

// HasGlobalProb returns a boolean if a field has been set.
func (o *AIGuardEvaluateResponse) HasGlobalProb() bool {
	return o != nil && o.GlobalProb != nil
}

// SetGlobalProb gets a reference to the given float64 and assigns it to the GlobalProb field.
func (o *AIGuardEvaluateResponse) SetGlobalProb(v float64) {
	o.GlobalProb = &v
}

// GetIsBlockingEnabled returns the IsBlockingEnabled field value.
func (o *AIGuardEvaluateResponse) GetIsBlockingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsBlockingEnabled
}

// GetIsBlockingEnabledOk returns a tuple with the IsBlockingEnabled field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetIsBlockingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlockingEnabled, true
}

// SetIsBlockingEnabled sets field value.
func (o *AIGuardEvaluateResponse) SetIsBlockingEnabled(v bool) {
	o.IsBlockingEnabled = v
}

// GetReason returns the Reason field value.
func (o *AIGuardEvaluateResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *AIGuardEvaluateResponse) SetReason(v string) {
	o.Reason = v
}

// GetSdsFindings returns the SdsFindings field value if set, zero value otherwise.
func (o *AIGuardEvaluateResponse) GetSdsFindings() []AIGuardSdsFinding {
	if o == nil || o.SdsFindings == nil {
		var ret []AIGuardSdsFinding
		return ret
	}
	return o.SdsFindings
}

// GetSdsFindingsOk returns a tuple with the SdsFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetSdsFindingsOk() (*[]AIGuardSdsFinding, bool) {
	if o == nil || o.SdsFindings == nil {
		return nil, false
	}
	return &o.SdsFindings, true
}

// HasSdsFindings returns a boolean if a field has been set.
func (o *AIGuardEvaluateResponse) HasSdsFindings() bool {
	return o != nil && o.SdsFindings != nil
}

// SetSdsFindings gets a reference to the given []AIGuardSdsFinding and assigns it to the SdsFindings field.
func (o *AIGuardEvaluateResponse) SetSdsFindings(v []AIGuardSdsFinding) {
	o.SdsFindings = v
}

// GetTagProbs returns the TagProbs field value.
func (o *AIGuardEvaluateResponse) GetTagProbs() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.TagProbs
}

// GetTagProbsOk returns a tuple with the TagProbs field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetTagProbsOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagProbs, true
}

// SetTagProbs sets field value.
func (o *AIGuardEvaluateResponse) SetTagProbs(v map[string]float64) {
	o.TagProbs = v
}

// GetTags returns the Tags field value.
func (o *AIGuardEvaluateResponse) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateResponse) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *AIGuardEvaluateResponse) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardEvaluateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.GlobalProb != nil {
		toSerialize["global_prob"] = o.GlobalProb
	}
	toSerialize["is_blocking_enabled"] = o.IsBlockingEnabled
	toSerialize["reason"] = o.Reason
	if o.SdsFindings != nil {
		toSerialize["sds_findings"] = o.SdsFindings
	}
	toSerialize["tag_probs"] = o.TagProbs
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardEvaluateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action            *AIGuardAction      `json:"action"`
		GlobalProb        *float64            `json:"global_prob,omitempty"`
		IsBlockingEnabled *bool               `json:"is_blocking_enabled"`
		Reason            *string             `json:"reason"`
		SdsFindings       []AIGuardSdsFinding `json:"sds_findings,omitempty"`
		TagProbs          *map[string]float64 `json:"tag_probs"`
		Tags              *[]string           `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.IsBlockingEnabled == nil {
		return fmt.Errorf("required field is_blocking_enabled missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.TagProbs == nil {
		return fmt.Errorf("required field tag_probs missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "global_prob", "is_blocking_enabled", "reason", "sds_findings", "tag_probs", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.GlobalProb = all.GlobalProb
	o.IsBlockingEnabled = *all.IsBlockingEnabled
	o.Reason = *all.Reason
	o.SdsFindings = all.SdsFindings
	o.TagProbs = *all.TagProbs
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
