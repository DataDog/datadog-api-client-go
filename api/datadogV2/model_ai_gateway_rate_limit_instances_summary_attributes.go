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
	// Counters that limited, or would have limited, a request in the last 15 minutes.
	LimitingCount int64 `json:"limiting_count"`
	// Utilization of the fullest counter.
	MaxUtilization float64 `json:"max_utilization"`
	// Counters at 80% or more of a limit in the rolling window.
	NearLimitCount int64 `json:"near_limit_count"`
	// Rate limit UUID.
	RateLimitId uuid.UUID `json:"rate_limit_id"`
	// True when some recently active counters' usage wasn't read, so the counts may be low.
	Truncated bool `json:"truncated"`
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

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitInstancesSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["limiting_count"] = o.LimitingCount
	toSerialize["max_utilization"] = o.MaxUtilization
	toSerialize["near_limit_count"] = o.NearLimitCount
	toSerialize["rate_limit_id"] = o.RateLimitId
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitInstancesSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceCount  *int64     `json:"instance_count"`
		LimitingCount  *int64     `json:"limiting_count"`
		MaxUtilization *float64   `json:"max_utilization"`
		NearLimitCount *int64     `json:"near_limit_count"`
		RateLimitId    *uuid.UUID `json:"rate_limit_id"`
		Truncated      *bool      `json:"truncated"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_count", "limiting_count", "max_utilization", "near_limit_count", "rate_limit_id", "truncated"})
	} else {
		return err
	}
	o.InstanceCount = *all.InstanceCount
	o.LimitingCount = *all.LimitingCount
	o.MaxUtilization = *all.MaxUtilization
	o.NearLimitCount = *all.NearLimitCount
	o.RateLimitId = *all.RateLimitId
	o.Truncated = *all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
