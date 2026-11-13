// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitInstanceAttributes One counter of a template rate limit and its live usage.
type AiGatewayRateLimitInstanceAttributes struct {
	// When the gateway first recorded a request for this counter.
	FirstSeen time.Time `json:"first_seen"`
	// Input tokens counted in the rolling one-minute window.
	InputTokens int64 `json:"input_tokens"`
	// The rate limit's per-counter input token limit. Null when it doesn't limit input tokens.
	InputTokensPerMinute datadog.NullableInt64 `json:"input_tokens_per_minute,omitempty"`
	// Opaque canonical key identifying the counter.
	InstanceKey string `json:"instance_key"`
	// When the gateway last recorded a request for this counter; up to a few minutes behind live traffic.
	LastSeen time.Time `json:"last_seen"`
	// Requests this counter rejected in the last 15 minutes (ENFORCE).
	LimitedRequests int64 `json:"limited_requests"`
	// Output tokens counted in the rolling one-minute window.
	OutputTokens int64 `json:"output_tokens"`
	// The rate limit's per-counter output token limit. Null when it doesn't limit output tokens.
	OutputTokensPerMinute datadog.NullableInt64 `json:"output_tokens_per_minute,omitempty"`
	// Rate limit UUID.
	RateLimitId uuid.UUID `json:"rate_limit_id"`
	// Requests counted in the rolling one-minute window.
	Requests int64 `json:"requests"`
	// The rate limit's per-counter request limit. Null when it doesn't limit requests.
	RequestsPerMinute datadog.NullableInt64 `json:"requests_per_minute,omitempty"`
	// The counter's value for each per-value dimension. user_id and team_id values are raw UUIDs.
	SplitDims map[string]string `json:"split_dims"`
	// False when this counter's usage wasn't read because the request already read the most counters it may. Its usage then reads as zero.
	UsageKnown bool `json:"usage_known"`
	// The fullest of the rate limit's limits in the rolling window, as a fraction. 1 or more means the counter is at its limit.
	Utilization float64 `json:"utilization"`
	// Requests this counter would have rejected in the last 15 minutes (ALERT).
	WouldLimitRequests int64 `json:"would_limit_requests"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitInstanceAttributes instantiates a new AiGatewayRateLimitInstanceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitInstanceAttributes(firstSeen time.Time, inputTokens int64, instanceKey string, lastSeen time.Time, limitedRequests int64, outputTokens int64, rateLimitId uuid.UUID, requests int64, splitDims map[string]string, usageKnown bool, utilization float64, wouldLimitRequests int64) *AiGatewayRateLimitInstanceAttributes {
	this := AiGatewayRateLimitInstanceAttributes{}
	this.FirstSeen = firstSeen
	this.InputTokens = inputTokens
	this.InstanceKey = instanceKey
	this.LastSeen = lastSeen
	this.LimitedRequests = limitedRequests
	this.OutputTokens = outputTokens
	this.RateLimitId = rateLimitId
	this.Requests = requests
	this.SplitDims = splitDims
	this.UsageKnown = usageKnown
	this.Utilization = utilization
	this.WouldLimitRequests = wouldLimitRequests
	return &this
}

// NewAiGatewayRateLimitInstanceAttributesWithDefaults instantiates a new AiGatewayRateLimitInstanceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitInstanceAttributesWithDefaults() *AiGatewayRateLimitInstanceAttributes {
	this := AiGatewayRateLimitInstanceAttributes{}
	return &this
}

// GetFirstSeen returns the FirstSeen field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetFirstSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetFirstSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSeen, true
}

// SetFirstSeen sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetFirstSeen(v time.Time) {
	o.FirstSeen = v
}

// GetInputTokens returns the InputTokens field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetInputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetInputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputTokens, true
}

// SetInputTokens sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetInputTokens(v int64) {
	o.InputTokens = v
}

// GetInputTokensPerMinute returns the InputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitInstanceAttributes) GetInputTokensPerMinute() int64 {
	if o == nil || o.InputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InputTokensPerMinute.Get()
}

// GetInputTokensPerMinuteOk returns a tuple with the InputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitInstanceAttributes) GetInputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputTokensPerMinute.Get(), o.InputTokensPerMinute.IsSet()
}

// HasInputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstanceAttributes) HasInputTokensPerMinute() bool {
	return o != nil && o.InputTokensPerMinute.IsSet()
}

// SetInputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the InputTokensPerMinute field.
func (o *AiGatewayRateLimitInstanceAttributes) SetInputTokensPerMinute(v int64) {
	o.InputTokensPerMinute.Set(&v)
}

// SetInputTokensPerMinuteNil sets the value for InputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) SetInputTokensPerMinuteNil() {
	o.InputTokensPerMinute.Set(nil)
}

// UnsetInputTokensPerMinute ensures that no value is present for InputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) UnsetInputTokensPerMinute() {
	o.InputTokensPerMinute.Unset()
}

// GetInstanceKey returns the InstanceKey field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetInstanceKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceKey
}

// GetInstanceKeyOk returns a tuple with the InstanceKey field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetInstanceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceKey, true
}

// SetInstanceKey sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetInstanceKey(v string) {
	o.InstanceKey = v
}

// GetLastSeen returns the LastSeen field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetLastSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetLastSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetLastSeen(v time.Time) {
	o.LastSeen = v
}

// GetLimitedRequests returns the LimitedRequests field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetLimitedRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LimitedRequests
}

// GetLimitedRequestsOk returns a tuple with the LimitedRequests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetLimitedRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitedRequests, true
}

// SetLimitedRequests sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetLimitedRequests(v int64) {
	o.LimitedRequests = v
}

// GetOutputTokens returns the OutputTokens field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetOutputTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetOutputTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputTokens, true
}

// SetOutputTokens sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetOutputTokens(v int64) {
	o.OutputTokens = v
}

// GetOutputTokensPerMinute returns the OutputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitInstanceAttributes) GetOutputTokensPerMinute() int64 {
	if o == nil || o.OutputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OutputTokensPerMinute.Get()
}

// GetOutputTokensPerMinuteOk returns a tuple with the OutputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitInstanceAttributes) GetOutputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputTokensPerMinute.Get(), o.OutputTokensPerMinute.IsSet()
}

// HasOutputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstanceAttributes) HasOutputTokensPerMinute() bool {
	return o != nil && o.OutputTokensPerMinute.IsSet()
}

// SetOutputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the OutputTokensPerMinute field.
func (o *AiGatewayRateLimitInstanceAttributes) SetOutputTokensPerMinute(v int64) {
	o.OutputTokensPerMinute.Set(&v)
}

// SetOutputTokensPerMinuteNil sets the value for OutputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) SetOutputTokensPerMinuteNil() {
	o.OutputTokensPerMinute.Set(nil)
}

// UnsetOutputTokensPerMinute ensures that no value is present for OutputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) UnsetOutputTokensPerMinute() {
	o.OutputTokensPerMinute.Unset()
}

// GetRateLimitId returns the RateLimitId field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetRateLimitId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RateLimitId
}

// GetRateLimitIdOk returns a tuple with the RateLimitId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetRateLimitIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimitId, true
}

// SetRateLimitId sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetRateLimitId(v uuid.UUID) {
	o.RateLimitId = v
}

// GetRequests returns the Requests field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetRequests(v int64) {
	o.Requests = v
}

// GetRequestsPerMinute returns the RequestsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitInstanceAttributes) GetRequestsPerMinute() int64 {
	if o == nil || o.RequestsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RequestsPerMinute.Get()
}

// GetRequestsPerMinuteOk returns a tuple with the RequestsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitInstanceAttributes) GetRequestsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestsPerMinute.Get(), o.RequestsPerMinute.IsSet()
}

// HasRequestsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstanceAttributes) HasRequestsPerMinute() bool {
	return o != nil && o.RequestsPerMinute.IsSet()
}

// SetRequestsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the RequestsPerMinute field.
func (o *AiGatewayRateLimitInstanceAttributes) SetRequestsPerMinute(v int64) {
	o.RequestsPerMinute.Set(&v)
}

// SetRequestsPerMinuteNil sets the value for RequestsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) SetRequestsPerMinuteNil() {
	o.RequestsPerMinute.Set(nil)
}

// UnsetRequestsPerMinute ensures that no value is present for RequestsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitInstanceAttributes) UnsetRequestsPerMinute() {
	o.RequestsPerMinute.Unset()
}

// GetSplitDims returns the SplitDims field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetSplitDims() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.SplitDims
}

// GetSplitDimsOk returns a tuple with the SplitDims field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetSplitDimsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitDims, true
}

// SetSplitDims sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetSplitDims(v map[string]string) {
	o.SplitDims = v
}

// GetUsageKnown returns the UsageKnown field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetUsageKnown() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.UsageKnown
}

// GetUsageKnownOk returns a tuple with the UsageKnown field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetUsageKnownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageKnown, true
}

// SetUsageKnown sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetUsageKnown(v bool) {
	o.UsageKnown = v
}

// GetUtilization returns the Utilization field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetUtilization() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetUtilizationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Utilization, true
}

// SetUtilization sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetUtilization(v float64) {
	o.Utilization = v
}

// GetWouldLimitRequests returns the WouldLimitRequests field value.
func (o *AiGatewayRateLimitInstanceAttributes) GetWouldLimitRequests() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WouldLimitRequests
}

// GetWouldLimitRequestsOk returns a tuple with the WouldLimitRequests field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstanceAttributes) GetWouldLimitRequestsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WouldLimitRequests, true
}

// SetWouldLimitRequests sets field value.
func (o *AiGatewayRateLimitInstanceAttributes) SetWouldLimitRequests(v int64) {
	o.WouldLimitRequests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitInstanceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstSeen.Nanosecond() == 0 {
		toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["input_tokens"] = o.InputTokens
	if o.InputTokensPerMinute.IsSet() {
		toSerialize["input_tokens_per_minute"] = o.InputTokensPerMinute.Get()
	}
	toSerialize["instance_key"] = o.InstanceKey
	if o.LastSeen.Nanosecond() == 0 {
		toSerialize["last_seen"] = o.LastSeen.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_seen"] = o.LastSeen.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["limited_requests"] = o.LimitedRequests
	toSerialize["output_tokens"] = o.OutputTokens
	if o.OutputTokensPerMinute.IsSet() {
		toSerialize["output_tokens_per_minute"] = o.OutputTokensPerMinute.Get()
	}
	toSerialize["rate_limit_id"] = o.RateLimitId
	toSerialize["requests"] = o.Requests
	if o.RequestsPerMinute.IsSet() {
		toSerialize["requests_per_minute"] = o.RequestsPerMinute.Get()
	}
	toSerialize["split_dims"] = o.SplitDims
	toSerialize["usage_known"] = o.UsageKnown
	toSerialize["utilization"] = o.Utilization
	toSerialize["would_limit_requests"] = o.WouldLimitRequests

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitInstanceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstSeen             *time.Time            `json:"first_seen"`
		InputTokens           *int64                `json:"input_tokens"`
		InputTokensPerMinute  datadog.NullableInt64 `json:"input_tokens_per_minute,omitempty"`
		InstanceKey           *string               `json:"instance_key"`
		LastSeen              *time.Time            `json:"last_seen"`
		LimitedRequests       *int64                `json:"limited_requests"`
		OutputTokens          *int64                `json:"output_tokens"`
		OutputTokensPerMinute datadog.NullableInt64 `json:"output_tokens_per_minute,omitempty"`
		RateLimitId           *uuid.UUID            `json:"rate_limit_id"`
		Requests              *int64                `json:"requests"`
		RequestsPerMinute     datadog.NullableInt64 `json:"requests_per_minute,omitempty"`
		SplitDims             *map[string]string    `json:"split_dims"`
		UsageKnown            *bool                 `json:"usage_known"`
		Utilization           *float64              `json:"utilization"`
		WouldLimitRequests    *int64                `json:"would_limit_requests"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FirstSeen == nil {
		return fmt.Errorf("required field first_seen missing")
	}
	if all.InputTokens == nil {
		return fmt.Errorf("required field input_tokens missing")
	}
	if all.InstanceKey == nil {
		return fmt.Errorf("required field instance_key missing")
	}
	if all.LastSeen == nil {
		return fmt.Errorf("required field last_seen missing")
	}
	if all.LimitedRequests == nil {
		return fmt.Errorf("required field limited_requests missing")
	}
	if all.OutputTokens == nil {
		return fmt.Errorf("required field output_tokens missing")
	}
	if all.RateLimitId == nil {
		return fmt.Errorf("required field rate_limit_id missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.SplitDims == nil {
		return fmt.Errorf("required field split_dims missing")
	}
	if all.UsageKnown == nil {
		return fmt.Errorf("required field usage_known missing")
	}
	if all.Utilization == nil {
		return fmt.Errorf("required field utilization missing")
	}
	if all.WouldLimitRequests == nil {
		return fmt.Errorf("required field would_limit_requests missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_seen", "input_tokens", "input_tokens_per_minute", "instance_key", "last_seen", "limited_requests", "output_tokens", "output_tokens_per_minute", "rate_limit_id", "requests", "requests_per_minute", "split_dims", "usage_known", "utilization", "would_limit_requests"})
	} else {
		return err
	}
	o.FirstSeen = *all.FirstSeen
	o.InputTokens = *all.InputTokens
	o.InputTokensPerMinute = all.InputTokensPerMinute
	o.InstanceKey = *all.InstanceKey
	o.LastSeen = *all.LastSeen
	o.LimitedRequests = *all.LimitedRequests
	o.OutputTokens = *all.OutputTokens
	o.OutputTokensPerMinute = all.OutputTokensPerMinute
	o.RateLimitId = *all.RateLimitId
	o.Requests = *all.Requests
	o.RequestsPerMinute = all.RequestsPerMinute
	o.SplitDims = *all.SplitDims
	o.UsageKnown = *all.UsageKnown
	o.Utilization = *all.Utilization
	o.WouldLimitRequests = *all.WouldLimitRequests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
