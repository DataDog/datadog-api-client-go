// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentTrigger Trigger a workflow from an agent via the MCP execute tool. Workflow can be executed from Bits Chat, Bits Agent Builder, Claude Code, Codex, Cursor, and any other coding agent using the Datadog MCP.
type AgentTrigger struct {
	// Defines a rate limit for a trigger.
	RateLimit *TriggerRateLimit `json:"rateLimit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentTrigger instantiates a new AgentTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentTrigger() *AgentTrigger {
	this := AgentTrigger{}
	return &this
}

// NewAgentTriggerWithDefaults instantiates a new AgentTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentTriggerWithDefaults() *AgentTrigger {
	this := AgentTrigger{}
	return &this
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *AgentTrigger) GetRateLimit() TriggerRateLimit {
	if o == nil || o.RateLimit == nil {
		var ret TriggerRateLimit
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTrigger) GetRateLimitOk() (*TriggerRateLimit, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *AgentTrigger) HasRateLimit() bool {
	return o != nil && o.RateLimit != nil
}

// SetRateLimit gets a reference to the given TriggerRateLimit and assigns it to the RateLimit field.
func (o *AgentTrigger) SetRateLimit(v TriggerRateLimit) {
	o.RateLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RateLimit != nil {
		toSerialize["rateLimit"] = o.RateLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RateLimit *TriggerRateLimit `json:"rateLimit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rateLimit"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RateLimit != nil && all.RateLimit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RateLimit = all.RateLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
