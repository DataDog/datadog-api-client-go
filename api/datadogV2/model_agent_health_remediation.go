// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthRemediation Remediation for the issue. Omitted when no remediation is available.
type AgentHealthRemediation struct {
	// A script that can be used to remediate the issue.
	Script *AgentHealthRemediationScript `json:"script,omitempty"`
	// Steps to remediate the issue.
	Steps []AgentHealthRemediationStep `json:"steps,omitempty"`
	// A summary of how to remediate the issue.
	Summary *string `json:"summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthRemediation instantiates a new AgentHealthRemediation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthRemediation() *AgentHealthRemediation {
	this := AgentHealthRemediation{}
	return &this
}

// NewAgentHealthRemediationWithDefaults instantiates a new AgentHealthRemediation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthRemediationWithDefaults() *AgentHealthRemediation {
	this := AgentHealthRemediation{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *AgentHealthRemediation) GetScript() AgentHealthRemediationScript {
	if o == nil || o.Script == nil {
		var ret AgentHealthRemediationScript
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediation) GetScriptOk() (*AgentHealthRemediationScript, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *AgentHealthRemediation) HasScript() bool {
	return o != nil && o.Script != nil
}

// SetScript gets a reference to the given AgentHealthRemediationScript and assigns it to the Script field.
func (o *AgentHealthRemediation) SetScript(v AgentHealthRemediationScript) {
	o.Script = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *AgentHealthRemediation) GetSteps() []AgentHealthRemediationStep {
	if o == nil || o.Steps == nil {
		var ret []AgentHealthRemediationStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediation) GetStepsOk() (*[]AgentHealthRemediationStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AgentHealthRemediation) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []AgentHealthRemediationStep and assigns it to the Steps field.
func (o *AgentHealthRemediation) SetSteps(v []AgentHealthRemediationStep) {
	o.Steps = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AgentHealthRemediation) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediation) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AgentHealthRemediation) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AgentHealthRemediation) SetSummary(v string) {
	o.Summary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthRemediation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthRemediation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Script  *AgentHealthRemediationScript `json:"script,omitempty"`
		Steps   []AgentHealthRemediationStep  `json:"steps,omitempty"`
		Summary *string                       `json:"summary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"script", "steps", "summary"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Script != nil && all.Script.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Script = all.Script
	o.Steps = all.Steps
	o.Summary = all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
