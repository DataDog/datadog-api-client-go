// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelCapabilities Model features reported by the enrichment source. Omitted properties are unknown.
type AiGatewayModelCapabilities struct {
	// Whether attachment input is supported.
	Attachment *bool `json:"attachment,omitempty"`
	// Whether reasoning can be interleaved with tool use.
	InterleavedReasoning *bool `json:"interleaved_reasoning,omitempty"`
	// Whether reasoning is supported.
	Reasoning *bool `json:"reasoning,omitempty"`
	// Whether structured output is supported.
	StructuredOutput *bool `json:"structured_output,omitempty"`
	// Whether temperature can be configured.
	Temperature *bool `json:"temperature,omitempty"`
	// Whether tool calls are supported.
	ToolCall *bool `json:"tool_call,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelCapabilities instantiates a new AiGatewayModelCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelCapabilities() *AiGatewayModelCapabilities {
	this := AiGatewayModelCapabilities{}
	return &this
}

// NewAiGatewayModelCapabilitiesWithDefaults instantiates a new AiGatewayModelCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelCapabilitiesWithDefaults() *AiGatewayModelCapabilities {
	this := AiGatewayModelCapabilities{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetAttachment() bool {
	if o == nil || o.Attachment == nil {
		var ret bool
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetAttachmentOk() (*bool, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasAttachment() bool {
	return o != nil && o.Attachment != nil
}

// SetAttachment gets a reference to the given bool and assigns it to the Attachment field.
func (o *AiGatewayModelCapabilities) SetAttachment(v bool) {
	o.Attachment = &v
}

// GetInterleavedReasoning returns the InterleavedReasoning field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetInterleavedReasoning() bool {
	if o == nil || o.InterleavedReasoning == nil {
		var ret bool
		return ret
	}
	return *o.InterleavedReasoning
}

// GetInterleavedReasoningOk returns a tuple with the InterleavedReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetInterleavedReasoningOk() (*bool, bool) {
	if o == nil || o.InterleavedReasoning == nil {
		return nil, false
	}
	return o.InterleavedReasoning, true
}

// HasInterleavedReasoning returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasInterleavedReasoning() bool {
	return o != nil && o.InterleavedReasoning != nil
}

// SetInterleavedReasoning gets a reference to the given bool and assigns it to the InterleavedReasoning field.
func (o *AiGatewayModelCapabilities) SetInterleavedReasoning(v bool) {
	o.InterleavedReasoning = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetReasoning() bool {
	if o == nil || o.Reasoning == nil {
		var ret bool
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetReasoningOk() (*bool, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given bool and assigns it to the Reasoning field.
func (o *AiGatewayModelCapabilities) SetReasoning(v bool) {
	o.Reasoning = &v
}

// GetStructuredOutput returns the StructuredOutput field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetStructuredOutput() bool {
	if o == nil || o.StructuredOutput == nil {
		var ret bool
		return ret
	}
	return *o.StructuredOutput
}

// GetStructuredOutputOk returns a tuple with the StructuredOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetStructuredOutputOk() (*bool, bool) {
	if o == nil || o.StructuredOutput == nil {
		return nil, false
	}
	return o.StructuredOutput, true
}

// HasStructuredOutput returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasStructuredOutput() bool {
	return o != nil && o.StructuredOutput != nil
}

// SetStructuredOutput gets a reference to the given bool and assigns it to the StructuredOutput field.
func (o *AiGatewayModelCapabilities) SetStructuredOutput(v bool) {
	o.StructuredOutput = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetTemperature() bool {
	if o == nil || o.Temperature == nil {
		var ret bool
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetTemperatureOk() (*bool, bool) {
	if o == nil || o.Temperature == nil {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasTemperature() bool {
	return o != nil && o.Temperature != nil
}

// SetTemperature gets a reference to the given bool and assigns it to the Temperature field.
func (o *AiGatewayModelCapabilities) SetTemperature(v bool) {
	o.Temperature = &v
}

// GetToolCall returns the ToolCall field value if set, zero value otherwise.
func (o *AiGatewayModelCapabilities) GetToolCall() bool {
	if o == nil || o.ToolCall == nil {
		var ret bool
		return ret
	}
	return *o.ToolCall
}

// GetToolCallOk returns a tuple with the ToolCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelCapabilities) GetToolCallOk() (*bool, bool) {
	if o == nil || o.ToolCall == nil {
		return nil, false
	}
	return o.ToolCall, true
}

// HasToolCall returns a boolean if a field has been set.
func (o *AiGatewayModelCapabilities) HasToolCall() bool {
	return o != nil && o.ToolCall != nil
}

// SetToolCall gets a reference to the given bool and assigns it to the ToolCall field.
func (o *AiGatewayModelCapabilities) SetToolCall(v bool) {
	o.ToolCall = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}
	if o.InterleavedReasoning != nil {
		toSerialize["interleaved_reasoning"] = o.InterleavedReasoning
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.StructuredOutput != nil {
		toSerialize["structured_output"] = o.StructuredOutput
	}
	if o.Temperature != nil {
		toSerialize["temperature"] = o.Temperature
	}
	if o.ToolCall != nil {
		toSerialize["tool_call"] = o.ToolCall
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attachment           *bool `json:"attachment,omitempty"`
		InterleavedReasoning *bool `json:"interleaved_reasoning,omitempty"`
		Reasoning            *bool `json:"reasoning,omitempty"`
		StructuredOutput     *bool `json:"structured_output,omitempty"`
		Temperature          *bool `json:"temperature,omitempty"`
		ToolCall             *bool `json:"tool_call,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attachment", "interleaved_reasoning", "reasoning", "structured_output", "temperature", "tool_call"})
	} else {
		return err
	}
	o.Attachment = all.Attachment
	o.InterleavedReasoning = all.InterleavedReasoning
	o.Reasoning = all.Reasoning
	o.StructuredOutput = all.StructuredOutput
	o.Temperature = all.Temperature
	o.ToolCall = all.ToolCall

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
