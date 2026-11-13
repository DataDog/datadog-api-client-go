// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitInstancesSummaryAttributes Aggregate of a template rate limit's counters.
type AiGatewayRateLimitInstancesSummaryAttributes struct {
	// Number of counters.
	InstanceCount int64 `json:"instance_count"`
	// Requests the counters rejected in the last 15 minutes (ENFORCE), totaled over the counters whose usage was read.
	LimitedRequests *int64 `json:"limited_requests,omitempty"`
	// Counters that limited, or would have limited, a request in the last 15 minutes.
	LimitingCount int64 `json:"limiting_count"`
	// Utilization of the fullest counter.
	MaxUtilization float64 `json:"max_utilization"`
	// Counters at 80% or more of a limit in the rolling window.
	NearLimitCount int64 `json:"near_limit_count"`
	// Rate limit UUID.
	RateLimitId uuid.UUID `json:"rate_limit_id"`
	// True when the caller can't see the template's selector because it lists a provider or model the caller can't access. The counters are then hidden, and every count is zero.
	ScopeRestricted *bool `json:"scope_restricted,omitempty"`
	// True when some recently active counters' usage wasn't read, so the counts may be low.
	Truncated bool `json:"truncated"`
	// Requests the counters would have rejected in the last 15 minutes (ALERT), totaled over the counters whose usage was read.
	WouldLimitRequests *int64 `json:"would_limit_requests,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitInstancesSummaryAttributes instantiates a new AiGatewayRateLimitInstancesSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitInstancesSummaryAttributes(instanceCount int64, limitingCount int64, maxUtilization float64, nearLimitCount int64, rateLimitId uuid.UUID, truncated bool) *AiGatewayRateLimitInstancesSummaryAttributes {
	this := AiGatewayRateLimitInstancesSummaryAttributes{}
	this.InstanceCount = instanceCount
	this.LimitingCount = limitingCount
	this.MaxUtilization = maxUtilization
	this.NearLimitCount = nearLimitCount
	this.RateLimitId = rateLimitId
	this.Truncated = truncated
	return &this
}

// NewAiGatewayRateLimitInstancesSummaryAttributesWithDefaults instantiates a new AiGatewayRateLimitInstancesSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitInstancesSummaryAttributesWithDefaults() *AiGatewayRateLimitInstancesSummaryAttributes {
	this := AiGatewayRateLimitInstancesSummaryAttributes{}
	return &this
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetLimitedRequests returns the LimitedRequests field value if set, zero value otherwise.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetLimitedRequests() int64 {
	if o == nil || o.LimitedRequests == nil {
		var ret int64
		return ret
	}
	return *o.LimitedRequests
}

// GetLimitedRequestsOk returns a tuple with the LimitedRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetLimitedRequestsOk() (*int64, bool) {
	if o == nil || o.LimitedRequests == nil {
		return nil, false
	}
	return o.LimitedRequests, true
}

// HasLimitedRequests returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) HasLimitedRequests() bool {
	return o != nil && o.LimitedRequests != nil
}

// SetLimitedRequests gets a reference to the given int64 and assigns it to the LimitedRequests field.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetLimitedRequests(v int64) {
	o.LimitedRequests = &v
}

// GetLimitingCount returns the LimitingCount field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetLimitingCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LimitingCount
}

// GetLimitingCountOk returns a tuple with the LimitingCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetLimitingCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitingCount, true
}

// SetLimitingCount sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetLimitingCount(v int64) {
	o.LimitingCount = v
}

// GetMaxUtilization returns the MaxUtilization field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetMaxUtilization() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxUtilization
}

// GetMaxUtilizationOk returns a tuple with the MaxUtilization field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetMaxUtilizationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUtilization, true
}

// SetMaxUtilization sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetMaxUtilization(v float64) {
	o.MaxUtilization = v
}

// GetNearLimitCount returns the NearLimitCount field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetNearLimitCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NearLimitCount
}

// GetNearLimitCountOk returns a tuple with the NearLimitCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetNearLimitCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NearLimitCount, true
}

// SetNearLimitCount sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetNearLimitCount(v int64) {
	o.NearLimitCount = v
}

// GetRateLimitId returns the RateLimitId field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetRateLimitId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RateLimitId
}

// GetRateLimitIdOk returns a tuple with the RateLimitId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetRateLimitIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimitId, true
}

// SetRateLimitId sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetRateLimitId(v uuid.UUID) {
	o.RateLimitId = v
}

// GetScopeRestricted returns the ScopeRestricted field value if set, zero value otherwise.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetScopeRestricted() bool {
	if o == nil || o.ScopeRestricted == nil {
		var ret bool
		return ret
	}
	return *o.ScopeRestricted
}

// GetScopeRestrictedOk returns a tuple with the ScopeRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetScopeRestrictedOk() (*bool, bool) {
	if o == nil || o.ScopeRestricted == nil {
		return nil, false
	}
	return o.ScopeRestricted, true
}

// HasScopeRestricted returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) HasScopeRestricted() bool {
	return o != nil && o.ScopeRestricted != nil
}

// SetScopeRestricted gets a reference to the given bool and assigns it to the ScopeRestricted field.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetScopeRestricted(v bool) {
	o.ScopeRestricted = &v
}

// GetTruncated returns the Truncated field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetTruncated(v bool) {
	o.Truncated = v
}

// GetWouldLimitRequests returns the WouldLimitRequests field value if set, zero value otherwise.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetWouldLimitRequests() int64 {
	if o == nil || o.WouldLimitRequests == nil {
		var ret int64
		return ret
	}
	return *o.WouldLimitRequests
}

// GetWouldLimitRequestsOk returns a tuple with the WouldLimitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) GetWouldLimitRequestsOk() (*int64, bool) {
	if o == nil || o.WouldLimitRequests == nil {
		return nil, false
	}
	return o.WouldLimitRequests, true
}

// HasWouldLimitRequests returns a boolean if a field has been set.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) HasWouldLimitRequests() bool {
	return o != nil && o.WouldLimitRequests != nil
}

// SetWouldLimitRequests gets a reference to the given int64 and assigns it to the WouldLimitRequests field.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) SetWouldLimitRequests(v int64) {
	o.WouldLimitRequests = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitInstancesSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_count"] = o.InstanceCount
	if o.LimitedRequests != nil {
		toSerialize["limited_requests"] = o.LimitedRequests
	}
	toSerialize["limiting_count"] = o.LimitingCount
	toSerialize["max_utilization"] = o.MaxUtilization
	toSerialize["near_limit_count"] = o.NearLimitCount
	toSerialize["rate_limit_id"] = o.RateLimitId
	if o.ScopeRestricted != nil {
		toSerialize["scope_restricted"] = o.ScopeRestricted
	}
	toSerialize["truncated"] = o.Truncated
	if o.WouldLimitRequests != nil {
		toSerialize["would_limit_requests"] = o.WouldLimitRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceCount      *int64     `json:"instance_count"`
		LimitedRequests    *int64     `json:"limited_requests,omitempty"`
		LimitingCount      *int64     `json:"limiting_count"`
		MaxUtilization     *float64   `json:"max_utilization"`
		NearLimitCount     *int64     `json:"near_limit_count"`
		RateLimitId        *uuid.UUID `json:"rate_limit_id"`
		ScopeRestricted    *bool      `json:"scope_restricted,omitempty"`
		Truncated          *bool      `json:"truncated"`
		WouldLimitRequests *int64     `json:"would_limit_requests,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if all.LimitingCount == nil {
		return fmt.Errorf("required field limiting_count missing")
	}
	if all.MaxUtilization == nil {
		return fmt.Errorf("required field max_utilization missing")
	}
	if all.NearLimitCount == nil {
		return fmt.Errorf("required field near_limit_count missing")
	}
	if all.RateLimitId == nil {
		return fmt.Errorf("required field rate_limit_id missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_count", "limited_requests", "limiting_count", "max_utilization", "near_limit_count", "rate_limit_id", "scope_restricted", "truncated", "would_limit_requests"})
	} else {
		return err
	}
	o.InstanceCount = *all.InstanceCount
	o.LimitedRequests = all.LimitedRequests
	o.LimitingCount = *all.LimitingCount
	o.MaxUtilization = *all.MaxUtilization
	o.NearLimitCount = *all.NearLimitCount
	o.RateLimitId = *all.RateLimitId
	o.ScopeRestricted = all.ScopeRestricted
	o.Truncated = *all.Truncated
	o.WouldLimitRequests = all.WouldLimitRequests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
