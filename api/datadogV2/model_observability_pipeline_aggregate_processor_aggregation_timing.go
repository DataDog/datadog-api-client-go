// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAggregateProcessorAggregationTiming Configures how metrics are assigned to aggregation windows. When omitted, metrics are grouped using system time.
type ObservabilityPipelineAggregateProcessorAggregationTiming struct {
	// Grace period, in seconds, for late-arriving metrics when using event time. Defaults to 10 seconds when omitted.
	AllowedLatenessSecs *int64 `json:"allowed_lateness_secs,omitempty"`
	// Determines whether metrics are assigned to aggregation windows based on when they are processed or their timestamps.
	Type ObservabilityPipelineAggregateProcessorAggregationTimingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAggregateProcessorAggregationTiming instantiates a new ObservabilityPipelineAggregateProcessorAggregationTiming object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAggregateProcessorAggregationTiming(typeVar ObservabilityPipelineAggregateProcessorAggregationTimingType) *ObservabilityPipelineAggregateProcessorAggregationTiming {
	this := ObservabilityPipelineAggregateProcessorAggregationTiming{}
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAggregateProcessorAggregationTimingWithDefaults instantiates a new ObservabilityPipelineAggregateProcessorAggregationTiming object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAggregateProcessorAggregationTimingWithDefaults() *ObservabilityPipelineAggregateProcessorAggregationTiming {
	this := ObservabilityPipelineAggregateProcessorAggregationTiming{}
	return &this
}

// GetAllowedLatenessSecs returns the AllowedLatenessSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) GetAllowedLatenessSecs() int64 {
	if o == nil || o.AllowedLatenessSecs == nil {
		var ret int64
		return ret
	}
	return *o.AllowedLatenessSecs
}

// GetAllowedLatenessSecsOk returns a tuple with the AllowedLatenessSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) GetAllowedLatenessSecsOk() (*int64, bool) {
	if o == nil || o.AllowedLatenessSecs == nil {
		return nil, false
	}
	return o.AllowedLatenessSecs, true
}

// HasAllowedLatenessSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) HasAllowedLatenessSecs() bool {
	return o != nil && o.AllowedLatenessSecs != nil
}

// SetAllowedLatenessSecs gets a reference to the given int64 and assigns it to the AllowedLatenessSecs field.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) SetAllowedLatenessSecs(v int64) {
	o.AllowedLatenessSecs = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) GetType() ObservabilityPipelineAggregateProcessorAggregationTimingType {
	if o == nil {
		var ret ObservabilityPipelineAggregateProcessorAggregationTimingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) GetTypeOk() (*ObservabilityPipelineAggregateProcessorAggregationTimingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) SetType(v ObservabilityPipelineAggregateProcessorAggregationTimingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAggregateProcessorAggregationTiming) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedLatenessSecs != nil {
		toSerialize["allowed_lateness_secs"] = o.AllowedLatenessSecs
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAggregateProcessorAggregationTiming) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedLatenessSecs *int64                                                        `json:"allowed_lateness_secs,omitempty"`
		Type                *ObservabilityPipelineAggregateProcessorAggregationTimingType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_lateness_secs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowedLatenessSecs = all.AllowedLatenessSecs
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
