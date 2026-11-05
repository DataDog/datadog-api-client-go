// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelLimits Token limits reported by the enrichment source.
type AiGatewayModelLimits struct {
	// Total model context window in tokens.
	ContextWindowTokens *int64 `json:"context_window_tokens,omitempty"`
	// Maximum input tokens when reported separately.
	MaxInputTokens *int64 `json:"max_input_tokens,omitempty"`
	// Maximum output tokens.
	MaxOutputTokens *int64 `json:"max_output_tokens,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelLimits instantiates a new AiGatewayModelLimits object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelLimits() *AiGatewayModelLimits {
	this := AiGatewayModelLimits{}
	return &this
}

// NewAiGatewayModelLimitsWithDefaults instantiates a new AiGatewayModelLimits object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelLimitsWithDefaults() *AiGatewayModelLimits {
	this := AiGatewayModelLimits{}
	return &this
}

// GetContextWindowTokens returns the ContextWindowTokens field value if set, zero value otherwise.
func (o *AiGatewayModelLimits) GetContextWindowTokens() int64 {
	if o == nil || o.ContextWindowTokens == nil {
		var ret int64
		return ret
	}
	return *o.ContextWindowTokens
}

// GetContextWindowTokensOk returns a tuple with the ContextWindowTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelLimits) GetContextWindowTokensOk() (*int64, bool) {
	if o == nil || o.ContextWindowTokens == nil {
		return nil, false
	}
	return o.ContextWindowTokens, true
}

// HasContextWindowTokens returns a boolean if a field has been set.
func (o *AiGatewayModelLimits) HasContextWindowTokens() bool {
	return o != nil && o.ContextWindowTokens != nil
}

// SetContextWindowTokens gets a reference to the given int64 and assigns it to the ContextWindowTokens field.
func (o *AiGatewayModelLimits) SetContextWindowTokens(v int64) {
	o.ContextWindowTokens = &v
}

// GetMaxInputTokens returns the MaxInputTokens field value if set, zero value otherwise.
func (o *AiGatewayModelLimits) GetMaxInputTokens() int64 {
	if o == nil || o.MaxInputTokens == nil {
		var ret int64
		return ret
	}
	return *o.MaxInputTokens
}

// GetMaxInputTokensOk returns a tuple with the MaxInputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelLimits) GetMaxInputTokensOk() (*int64, bool) {
	if o == nil || o.MaxInputTokens == nil {
		return nil, false
	}
	return o.MaxInputTokens, true
}

// HasMaxInputTokens returns a boolean if a field has been set.
func (o *AiGatewayModelLimits) HasMaxInputTokens() bool {
	return o != nil && o.MaxInputTokens != nil
}

// SetMaxInputTokens gets a reference to the given int64 and assigns it to the MaxInputTokens field.
func (o *AiGatewayModelLimits) SetMaxInputTokens(v int64) {
	o.MaxInputTokens = &v
}

// GetMaxOutputTokens returns the MaxOutputTokens field value if set, zero value otherwise.
func (o *AiGatewayModelLimits) GetMaxOutputTokens() int64 {
	if o == nil || o.MaxOutputTokens == nil {
		var ret int64
		return ret
	}
	return *o.MaxOutputTokens
}

// GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelLimits) GetMaxOutputTokensOk() (*int64, bool) {
	if o == nil || o.MaxOutputTokens == nil {
		return nil, false
	}
	return o.MaxOutputTokens, true
}

// HasMaxOutputTokens returns a boolean if a field has been set.
func (o *AiGatewayModelLimits) HasMaxOutputTokens() bool {
	return o != nil && o.MaxOutputTokens != nil
}

// SetMaxOutputTokens gets a reference to the given int64 and assigns it to the MaxOutputTokens field.
func (o *AiGatewayModelLimits) SetMaxOutputTokens(v int64) {
	o.MaxOutputTokens = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContextWindowTokens != nil {
		toSerialize["context_window_tokens"] = o.ContextWindowTokens
	}
	if o.MaxInputTokens != nil {
		toSerialize["max_input_tokens"] = o.MaxInputTokens
	}
	if o.MaxOutputTokens != nil {
		toSerialize["max_output_tokens"] = o.MaxOutputTokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelLimits) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContextWindowTokens *int64 `json:"context_window_tokens,omitempty"`
		MaxInputTokens      *int64 `json:"max_input_tokens,omitempty"`
		MaxOutputTokens     *int64 `json:"max_output_tokens,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"context_window_tokens", "max_input_tokens", "max_output_tokens"})
	} else {
		return err
	}
	o.ContextWindowTokens = all.ContextWindowTokens
	o.MaxInputTokens = all.MaxInputTokens
	o.MaxOutputTokens = all.MaxOutputTokens

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
