// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorConfigPolicyDowntimePolicyCreateRequest Downtime duration attributes of a monitor configuration policy.
type MonitorConfigPolicyDowntimePolicyCreateRequest struct {
	// The maximum allowed downtime duration, in milliseconds.
	MaxDurationMs int64 `json:"max_duration_ms"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorConfigPolicyDowntimePolicyCreateRequest instantiates a new MonitorConfigPolicyDowntimePolicyCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorConfigPolicyDowntimePolicyCreateRequest(maxDurationMs int64) *MonitorConfigPolicyDowntimePolicyCreateRequest {
	this := MonitorConfigPolicyDowntimePolicyCreateRequest{}
	this.MaxDurationMs = maxDurationMs
	return &this
}

// NewMonitorConfigPolicyDowntimePolicyCreateRequestWithDefaults instantiates a new MonitorConfigPolicyDowntimePolicyCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorConfigPolicyDowntimePolicyCreateRequestWithDefaults() *MonitorConfigPolicyDowntimePolicyCreateRequest {
	this := MonitorConfigPolicyDowntimePolicyCreateRequest{}
	return &this
}

// GetMaxDurationMs returns the MaxDurationMs field value.
func (o *MonitorConfigPolicyDowntimePolicyCreateRequest) GetMaxDurationMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxDurationMs
}

// GetMaxDurationMsOk returns a tuple with the MaxDurationMs field value
// and a boolean to check if the value has been set.
func (o *MonitorConfigPolicyDowntimePolicyCreateRequest) GetMaxDurationMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxDurationMs, true
}

// SetMaxDurationMs sets field value.
func (o *MonitorConfigPolicyDowntimePolicyCreateRequest) SetMaxDurationMs(v int64) {
	o.MaxDurationMs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorConfigPolicyDowntimePolicyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_duration_ms"] = o.MaxDurationMs
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorConfigPolicyDowntimePolicyCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxDurationMs *int64 `json:"max_duration_ms"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxDurationMs == nil {
		return fmt.Errorf("required field max_duration_ms missing")
	}
	o.MaxDurationMs = *all.MaxDurationMs

	return nil
}
