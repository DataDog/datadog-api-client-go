// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitUsageAttributes Live consumption of one AI Gateway rate limit.
type AiGatewayRateLimitUsageAttributes struct {
	// Input tokens counted in the rolling one-minute window.
	InputTokens int64 `json:"input_tokens"`
	// Requests this rate limit rejected in the last 15 minutes. Applies when mode is ENFORCE.
	LimitedRequests int64 `json:"limited_requests"`
	// Output tokens counted in the rolling one-minute window.
	OutputTokens int64 `json:"output_tokens"`
	// Requests counted in the rolling one-minute window.
	Requests int64 `json:"requests"`
	// Requests this rate limit would have rejected in the last 15 minutes. Applies when mode is ALERT.
	WouldLimitRequests int64 `json:"would_limit_requests"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitUsageAttributes instantiates a new AiGatewayRateLimitUsageAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitUsageAttributes(inputTokens int64, limitedRequests int64, outputTokens int64, requests int64, wouldLimitRequests int64) *AiGatewayRateLimitUsageAttributes {
	this := AiGatewayRateLimitUsageAttributes{}
	this.InputTokens = inputTokens
	this.LimitedRequests = limitedRequests
	this.OutputTokens = outputTokens
	this.Requests = requests
	this.WouldLimitRequests = wouldLimitRequests
	return &this
}

// NewAiGatewayRateLimitUsageAttributesWithDefaults instantiates a new AiGatewayRateLimitUsageAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitUsageAttributesWithDefaults() *AiGatewayRateLimitUsageAttributes {
	this := AiGatewayRateLimitUsageAttributes{}
	return &this
}

// GetInputTokens returns the InputTokens field value.
func (o *AiGatewayRateLimitUsageAttributes) GetInputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUsageAttributes) GetInputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputTokens, true
}

// SetInputTokens sets field value.
func (o *AiGatewayRateLimitUsageAttributes) SetInputTokens(v int64) {
	o.InputTokens = v
}

// GetLimitedRequests returns the LimitedRequests field value.
func (o *AiGatewayRateLimitUsageAttributes) GetLimitedRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LimitedRequests
}

// GetLimitedRequestsOk returns a tuple with the LimitedRequests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUsageAttributes) GetLimitedRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitedRequests, true
}

// SetLimitedRequests sets field value.
func (o *AiGatewayRateLimitUsageAttributes) SetLimitedRequests(v int64) {
	o.LimitedRequests = v
}

// GetOutputTokens returns the OutputTokens field value.
func (o *AiGatewayRateLimitUsageAttributes) GetOutputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUsageAttributes) GetOutputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputTokens, true
}

// SetOutputTokens sets field value.
func (o *AiGatewayRateLimitUsageAttributes) SetOutputTokens(v int64) {
	o.OutputTokens = v
}

// GetRequests returns the Requests field value.
func (o *AiGatewayRateLimitUsageAttributes) GetRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUsageAttributes) GetRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *AiGatewayRateLimitUsageAttributes) SetRequests(v int64) {
	o.Requests = v
}

// GetWouldLimitRequests returns the WouldLimitRequests field value.
func (o *AiGatewayRateLimitUsageAttributes) GetWouldLimitRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WouldLimitRequests
}

// GetWouldLimitRequestsOk returns a tuple with the WouldLimitRequests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUsageAttributes) GetWouldLimitRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WouldLimitRequests, true
}

// SetWouldLimitRequests sets field value.
func (o *AiGatewayRateLimitUsageAttributes) SetWouldLimitRequests(v int64) {
	o.WouldLimitRequests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitUsageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["input_tokens"] = o.InputTokens
	toSerialize["limited_requests"] = o.LimitedRequests
	toSerialize["output_tokens"] = o.OutputTokens
	toSerialize["requests"] = o.Requests
	toSerialize["would_limit_requests"] = o.WouldLimitRequests

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitUsageAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InputTokens        *int64 `json:"input_tokens"`
		LimitedRequests    *int64 `json:"limited_requests"`
		OutputTokens       *int64 `json:"output_tokens"`
		Requests           *int64 `json:"requests"`
		WouldLimitRequests *int64 `json:"would_limit_requests"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InputTokens == nil {
		return fmt.Errorf("required field input_tokens missing")
	}
	if all.LimitedRequests == nil {
		return fmt.Errorf("required field limited_requests missing")
	}
	if all.OutputTokens == nil {
		return fmt.Errorf("required field output_tokens missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.WouldLimitRequests == nil {
		return fmt.Errorf("required field would_limit_requests missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"input_tokens", "limited_requests", "output_tokens", "requests", "would_limit_requests"})
	} else {
		return err
	}
	o.InputTokens = *all.InputTokens
	o.LimitedRequests = *all.LimitedRequests
	o.OutputTokens = *all.OutputTokens
	o.Requests = *all.Requests
	o.WouldLimitRequests = *all.WouldLimitRequests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
