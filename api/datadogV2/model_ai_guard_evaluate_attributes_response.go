// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardEvaluateAttributesResponse The evaluation result with the recommended action and detected threat information.
type AIGuardEvaluateAttributesResponse struct {
	// The recommended action for the evaluated conversation.
	Action AIGuardAction `json:"action"`
	// The global threat probability, between `0` and `1`. Included only when computed.
	GlobalProb *float64 `json:"global_prob,omitempty"`
	// Whether blocking is enabled for the organization.
	IsBlockingEnabled bool `json:"is_blocking_enabled"`
	// The explanation for the recommended action, for logging rather than user-facing display.
	Reason string `json:"reason"`
	// The redacted replacements computed for the request. Included only when redaction applies.
	RedactionReplacements []AIGuardRedactionReplacement `json:"redaction_replacements,omitempty"`
	// The Supply Chain Firewall verdict for the version-pinned install packages
	// found in the request. Evaluation requires Supply Chain Firewall to be enabled
	// for the organization and the caller to have the `appsec_vm_read` permission.
	// Without that permission, Supply Chain Firewall evaluation is skipped; the
	// AI Guard evaluation request itself is not rejected. Findings are omitted
	// when Supply Chain Firewall does not produce a result.
	// The recommended action combines the package outcomes.
	// Confirmation support is inferred from the evaluator profile, not from a
	// separate confirmation-capability flag. By default, a non-empty
	// `meta.coding_agent` selects the coding-agent profile, which assumes the caller
	// can ask for confirmation and retains WARN in `recommended_action`.
	// Omitting `meta.coding_agent` or setting it to an empty string selects the
	// default profile, which assumes the caller cannot ask for confirmation and
	// returns ALLOW instead of WARN in `recommended_action`. Package outcomes
	// still retain WARN, and BLOCK is never changed to ALLOW by this rule.
	ScfwFindings *AIGuardScfwFindings `json:"scfw_findings,omitempty"`
	// The sensitive data findings from the Sensitive Data Scanner. Included only when findings exist.
	SdsFindings []AIGuardSdsFinding `json:"sds_findings,omitempty"`
	// The detection confidence for each threat tag, between `0` and `1`. May be null when the evaluation does not produce per-category scores.
	TagProbs map[string]float64 `json:"tag_probs"`
	// The threat tags detected in the conversation.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardEvaluateAttributesResponse instantiates a new AIGuardEvaluateAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardEvaluateAttributesResponse(action AIGuardAction, isBlockingEnabled bool, reason string, tagProbs map[string]float64, tags []string) *AIGuardEvaluateAttributesResponse {
	this := AIGuardEvaluateAttributesResponse{}
	this.Action = action
	this.IsBlockingEnabled = isBlockingEnabled
	this.Reason = reason
	this.TagProbs = tagProbs
	this.Tags = tags
	return &this
}

// NewAIGuardEvaluateAttributesResponseWithDefaults instantiates a new AIGuardEvaluateAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardEvaluateAttributesResponseWithDefaults() *AIGuardEvaluateAttributesResponse {
	this := AIGuardEvaluateAttributesResponse{}
	return &this
}

// GetAction returns the Action field value.
func (o *AIGuardEvaluateAttributesResponse) GetAction() AIGuardAction {
	if o == nil {
		var ret AIGuardAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetActionOk() (*AIGuardAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *AIGuardEvaluateAttributesResponse) SetAction(v AIGuardAction) {
	o.Action = v
}

// GetGlobalProb returns the GlobalProb field value if set, zero value otherwise.
func (o *AIGuardEvaluateAttributesResponse) GetGlobalProb() float64 {
	if o == nil || o.GlobalProb == nil {
		var ret float64
		return ret
	}
	return *o.GlobalProb
}

// GetGlobalProbOk returns a tuple with the GlobalProb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetGlobalProbOk() (*float64, bool) {
	if o == nil || o.GlobalProb == nil {
		return nil, false
	}
	return o.GlobalProb, true
}

// HasGlobalProb returns a boolean if a field has been set.
func (o *AIGuardEvaluateAttributesResponse) HasGlobalProb() bool {
	return o != nil && o.GlobalProb != nil
}

// SetGlobalProb gets a reference to the given float64 and assigns it to the GlobalProb field.
func (o *AIGuardEvaluateAttributesResponse) SetGlobalProb(v float64) {
	o.GlobalProb = &v
}

// GetIsBlockingEnabled returns the IsBlockingEnabled field value.
func (o *AIGuardEvaluateAttributesResponse) GetIsBlockingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsBlockingEnabled
}

// GetIsBlockingEnabledOk returns a tuple with the IsBlockingEnabled field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetIsBlockingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlockingEnabled, true
}

// SetIsBlockingEnabled sets field value.
func (o *AIGuardEvaluateAttributesResponse) SetIsBlockingEnabled(v bool) {
	o.IsBlockingEnabled = v
}

// GetReason returns the Reason field value.
func (o *AIGuardEvaluateAttributesResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *AIGuardEvaluateAttributesResponse) SetReason(v string) {
	o.Reason = v
}

// GetRedactionReplacements returns the RedactionReplacements field value if set, zero value otherwise.
func (o *AIGuardEvaluateAttributesResponse) GetRedactionReplacements() []AIGuardRedactionReplacement {
	if o == nil || o.RedactionReplacements == nil {
		var ret []AIGuardRedactionReplacement
		return ret
	}
	return o.RedactionReplacements
}

// GetRedactionReplacementsOk returns a tuple with the RedactionReplacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetRedactionReplacementsOk() (*[]AIGuardRedactionReplacement, bool) {
	if o == nil || o.RedactionReplacements == nil {
		return nil, false
	}
	return &o.RedactionReplacements, true
}

// HasRedactionReplacements returns a boolean if a field has been set.
func (o *AIGuardEvaluateAttributesResponse) HasRedactionReplacements() bool {
	return o != nil && o.RedactionReplacements != nil
}

// SetRedactionReplacements gets a reference to the given []AIGuardRedactionReplacement and assigns it to the RedactionReplacements field.
func (o *AIGuardEvaluateAttributesResponse) SetRedactionReplacements(v []AIGuardRedactionReplacement) {
	o.RedactionReplacements = v
}

// GetScfwFindings returns the ScfwFindings field value if set, zero value otherwise.
func (o *AIGuardEvaluateAttributesResponse) GetScfwFindings() AIGuardScfwFindings {
	if o == nil || o.ScfwFindings == nil {
		var ret AIGuardScfwFindings
		return ret
	}
	return *o.ScfwFindings
}

// GetScfwFindingsOk returns a tuple with the ScfwFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetScfwFindingsOk() (*AIGuardScfwFindings, bool) {
	if o == nil || o.ScfwFindings == nil {
		return nil, false
	}
	return o.ScfwFindings, true
}

// HasScfwFindings returns a boolean if a field has been set.
func (o *AIGuardEvaluateAttributesResponse) HasScfwFindings() bool {
	return o != nil && o.ScfwFindings != nil
}

// SetScfwFindings gets a reference to the given AIGuardScfwFindings and assigns it to the ScfwFindings field.
func (o *AIGuardEvaluateAttributesResponse) SetScfwFindings(v AIGuardScfwFindings) {
	o.ScfwFindings = &v
}

// GetSdsFindings returns the SdsFindings field value if set, zero value otherwise.
func (o *AIGuardEvaluateAttributesResponse) GetSdsFindings() []AIGuardSdsFinding {
	if o == nil || o.SdsFindings == nil {
		var ret []AIGuardSdsFinding
		return ret
	}
	return o.SdsFindings
}

// GetSdsFindingsOk returns a tuple with the SdsFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetSdsFindingsOk() (*[]AIGuardSdsFinding, bool) {
	if o == nil || o.SdsFindings == nil {
		return nil, false
	}
	return &o.SdsFindings, true
}

// HasSdsFindings returns a boolean if a field has been set.
func (o *AIGuardEvaluateAttributesResponse) HasSdsFindings() bool {
	return o != nil && o.SdsFindings != nil
}

// SetSdsFindings gets a reference to the given []AIGuardSdsFinding and assigns it to the SdsFindings field.
func (o *AIGuardEvaluateAttributesResponse) SetSdsFindings(v []AIGuardSdsFinding) {
	o.SdsFindings = v
}

// GetTagProbs returns the TagProbs field value.
// If the value is explicit nil, the zero value for map[string]float64 will be returned.
func (o *AIGuardEvaluateAttributesResponse) GetTagProbs() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.TagProbs
}

// GetTagProbsOk returns a tuple with the TagProbs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AIGuardEvaluateAttributesResponse) GetTagProbsOk() (*map[string]float64, bool) {
	if o == nil || o.TagProbs == nil {
		return nil, false
	}
	return &o.TagProbs, true
}

// SetTagProbs sets field value.
func (o *AIGuardEvaluateAttributesResponse) SetTagProbs(v map[string]float64) {
	o.TagProbs = v
}

// GetTags returns the Tags field value.
func (o *AIGuardEvaluateAttributesResponse) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AIGuardEvaluateAttributesResponse) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *AIGuardEvaluateAttributesResponse) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardEvaluateAttributesResponse) MarshalJSON() ([]byte, error) {
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
	if o.RedactionReplacements != nil {
		toSerialize["redaction_replacements"] = o.RedactionReplacements
	}
	if o.ScfwFindings != nil {
		toSerialize["scfw_findings"] = o.ScfwFindings
	}
	if o.SdsFindings != nil {
		toSerialize["sds_findings"] = o.SdsFindings
	}
	if o.TagProbs != nil {
		toSerialize["tag_probs"] = o.TagProbs
	}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardEvaluateAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action                *AIGuardAction                `json:"action"`
		GlobalProb            *float64                      `json:"global_prob,omitempty"`
		IsBlockingEnabled     *bool                         `json:"is_blocking_enabled"`
		Reason                *string                       `json:"reason"`
		RedactionReplacements []AIGuardRedactionReplacement `json:"redaction_replacements,omitempty"`
		ScfwFindings          *AIGuardScfwFindings          `json:"scfw_findings,omitempty"`
		SdsFindings           []AIGuardSdsFinding           `json:"sds_findings,omitempty"`
		TagProbs              map[string]float64            `json:"tag_probs"`
		Tags                  *[]string                     `json:"tags"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "global_prob", "is_blocking_enabled", "reason", "redaction_replacements", "scfw_findings", "sds_findings", "tag_probs", "tags"})
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
	o.RedactionReplacements = all.RedactionReplacements
	if all.ScfwFindings != nil && all.ScfwFindings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScfwFindings = all.ScfwFindings
	o.SdsFindings = all.SdsFindings
	o.TagProbs = all.TagProbs
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
