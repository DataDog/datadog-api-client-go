// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthRemediationStep A step to remediate the issue.
type AgentHealthRemediationStep struct {
	// The order in which to perform the step.
	Order int32 `json:"order"`
	// The instructions for the step.
	Text string `json:"text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthRemediationStep instantiates a new AgentHealthRemediationStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthRemediationStep(order int32, text string) *AgentHealthRemediationStep {
	this := AgentHealthRemediationStep{}
	this.Order = order
	this.Text = text
	return &this
}

// NewAgentHealthRemediationStepWithDefaults instantiates a new AgentHealthRemediationStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthRemediationStepWithDefaults() *AgentHealthRemediationStep {
	this := AgentHealthRemediationStep{}
	return &this
}

// GetOrder returns the Order field value.
func (o *AgentHealthRemediationStep) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationStep) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *AgentHealthRemediationStep) SetOrder(v int32) {
	o.Order = v
}

// GetText returns the Text field value.
func (o *AgentHealthRemediationStep) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationStep) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *AgentHealthRemediationStep) SetText(v string) {
	o.Text = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthRemediationStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["order"] = o.Order
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthRemediationStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Order *int32  `json:"order"`
		Text  *string `json:"text"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"order", "text"})
	} else {
		return err
	}
	o.Order = *all.Order
	o.Text = *all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
