// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardMeta Metadata about the evaluated conversation, used to select the policy and evaluator behavior.
type AIGuardMeta struct {
	// The coding agent that produced the conversation. By default, a non-empty
	// value selects the coding-agent evaluator profile, while an omitted or empty
	// value selects the default profile. Known identifiers include `claude-code`,
	// `codex`, `cursor`, and `github-copilot`; other identifiers are accepted.
	// The coding-agent profile assumes the caller can ask the user to confirm
	// Supply Chain Firewall warnings; the default profile assumes it cannot.
	CodingAgent *string `json:"coding_agent,omitempty"`
	// Request-only confidence threshold override. A number sets the threshold
	// for every threat category; an object supplies a global threshold and
	// optional per-category overrides. Omit this field to use the stored
	// sensitivity policy. This field is not returned in evaluation responses.
	ConfidenceThreshold *AIGuardConfidenceThresholdOverride `json:"confidence_threshold,omitempty"`
	// The environment the conversation belongs to.
	Env *string `json:"env,omitempty"`
	// Whether to run the Sensitive Data Scanner for this request, overriding the stored policy.
	IsSdsEnabledOverride *bool `json:"is_sds_enabled_override,omitempty"`
	// The name of the service the conversation belongs to.
	Service *string `json:"service,omitempty"`
	// A non-empty map replaces stored tool policies for this request only,
	// keyed by tool name. Any called tool not present resolves to EVALUATE.
	// Only takes effect when tool policies are enabled for the organization.
	ToolPolicies map[string]AIGuardToolPolicyAction `json:"tool_policies,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardMeta instantiates a new AIGuardMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardMeta() *AIGuardMeta {
	this := AIGuardMeta{}
	return &this
}

// NewAIGuardMetaWithDefaults instantiates a new AIGuardMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardMetaWithDefaults() *AIGuardMeta {
	this := AIGuardMeta{}
	return &this
}

// GetCodingAgent returns the CodingAgent field value if set, zero value otherwise.
func (o *AIGuardMeta) GetCodingAgent() string {
	if o == nil || o.CodingAgent == nil {
		var ret string
		return ret
	}
	return *o.CodingAgent
}

// GetCodingAgentOk returns a tuple with the CodingAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetCodingAgentOk() (*string, bool) {
	if o == nil || o.CodingAgent == nil {
		return nil, false
	}
	return o.CodingAgent, true
}

// HasCodingAgent returns a boolean if a field has been set.
func (o *AIGuardMeta) HasCodingAgent() bool {
	return o != nil && o.CodingAgent != nil
}

// SetCodingAgent gets a reference to the given string and assigns it to the CodingAgent field.
func (o *AIGuardMeta) SetCodingAgent(v string) {
	o.CodingAgent = &v
}

// GetConfidenceThreshold returns the ConfidenceThreshold field value if set, zero value otherwise.
func (o *AIGuardMeta) GetConfidenceThreshold() AIGuardConfidenceThresholdOverride {
	if o == nil || o.ConfidenceThreshold == nil {
		var ret AIGuardConfidenceThresholdOverride
		return ret
	}
	return *o.ConfidenceThreshold
}

// GetConfidenceThresholdOk returns a tuple with the ConfidenceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetConfidenceThresholdOk() (*AIGuardConfidenceThresholdOverride, bool) {
	if o == nil || o.ConfidenceThreshold == nil {
		return nil, false
	}
	return o.ConfidenceThreshold, true
}

// HasConfidenceThreshold returns a boolean if a field has been set.
func (o *AIGuardMeta) HasConfidenceThreshold() bool {
	return o != nil && o.ConfidenceThreshold != nil
}

// SetConfidenceThreshold gets a reference to the given AIGuardConfidenceThresholdOverride and assigns it to the ConfidenceThreshold field.
func (o *AIGuardMeta) SetConfidenceThreshold(v AIGuardConfidenceThresholdOverride) {
	o.ConfidenceThreshold = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *AIGuardMeta) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *AIGuardMeta) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *AIGuardMeta) SetEnv(v string) {
	o.Env = &v
}

// GetIsSdsEnabledOverride returns the IsSdsEnabledOverride field value if set, zero value otherwise.
func (o *AIGuardMeta) GetIsSdsEnabledOverride() bool {
	if o == nil || o.IsSdsEnabledOverride == nil {
		var ret bool
		return ret
	}
	return *o.IsSdsEnabledOverride
}

// GetIsSdsEnabledOverrideOk returns a tuple with the IsSdsEnabledOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetIsSdsEnabledOverrideOk() (*bool, bool) {
	if o == nil || o.IsSdsEnabledOverride == nil {
		return nil, false
	}
	return o.IsSdsEnabledOverride, true
}

// HasIsSdsEnabledOverride returns a boolean if a field has been set.
func (o *AIGuardMeta) HasIsSdsEnabledOverride() bool {
	return o != nil && o.IsSdsEnabledOverride != nil
}

// SetIsSdsEnabledOverride gets a reference to the given bool and assigns it to the IsSdsEnabledOverride field.
func (o *AIGuardMeta) SetIsSdsEnabledOverride(v bool) {
	o.IsSdsEnabledOverride = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AIGuardMeta) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AIGuardMeta) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *AIGuardMeta) SetService(v string) {
	o.Service = &v
}

// GetToolPolicies returns the ToolPolicies field value if set, zero value otherwise.
func (o *AIGuardMeta) GetToolPolicies() map[string]AIGuardToolPolicyAction {
	if o == nil || o.ToolPolicies == nil {
		var ret map[string]AIGuardToolPolicyAction
		return ret
	}
	return o.ToolPolicies
}

// GetToolPoliciesOk returns a tuple with the ToolPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMeta) GetToolPoliciesOk() (*map[string]AIGuardToolPolicyAction, bool) {
	if o == nil || o.ToolPolicies == nil {
		return nil, false
	}
	return &o.ToolPolicies, true
}

// HasToolPolicies returns a boolean if a field has been set.
func (o *AIGuardMeta) HasToolPolicies() bool {
	return o != nil && o.ToolPolicies != nil
}

// SetToolPolicies gets a reference to the given map[string]AIGuardToolPolicyAction and assigns it to the ToolPolicies field.
func (o *AIGuardMeta) SetToolPolicies(v map[string]AIGuardToolPolicyAction) {
	o.ToolPolicies = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodingAgent != nil {
		toSerialize["coding_agent"] = o.CodingAgent
	}
	if o.ConfidenceThreshold != nil {
		toSerialize["confidence_threshold"] = o.ConfidenceThreshold
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.IsSdsEnabledOverride != nil {
		toSerialize["is_sds_enabled_override"] = o.IsSdsEnabledOverride
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ToolPolicies != nil {
		toSerialize["tool_policies"] = o.ToolPolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodingAgent          *string                             `json:"coding_agent,omitempty"`
		ConfidenceThreshold  *AIGuardConfidenceThresholdOverride `json:"confidence_threshold,omitempty"`
		Env                  *string                             `json:"env,omitempty"`
		IsSdsEnabledOverride *bool                               `json:"is_sds_enabled_override,omitempty"`
		Service              *string                             `json:"service,omitempty"`
		ToolPolicies         map[string]AIGuardToolPolicyAction  `json:"tool_policies,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"coding_agent", "confidence_threshold", "env", "is_sds_enabled_override", "service", "tool_policies"})
	} else {
		return err
	}
	o.CodingAgent = all.CodingAgent
	o.ConfidenceThreshold = all.ConfidenceThreshold
	o.Env = all.Env
	o.IsSdsEnabledOverride = all.IsSdsEnabledOverride
	o.Service = all.Service
	o.ToolPolicies = all.ToolPolicies

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
