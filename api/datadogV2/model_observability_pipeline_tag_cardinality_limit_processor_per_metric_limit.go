// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit A cardinality override applied to a specific metric.
type ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit struct {
	// The action to take when the cardinality limit is exceeded.
	LimitExceededAction *ObservabilityPipelineTagCardinalityLimitProcessorAction `json:"limit_exceeded_action,omitempty"`
	// The name of the metric this override applies to.
	MetricName string `json:"metric_name"`
	// How the override is applied. `limit_override` enforces a custom limit; `excluded` omits the metric or tag from cardinality tracking.
	OverrideType ObservabilityPipelineTagCardinalityLimitProcessorOverrideType `json:"override_type"`
	// A list of per-tag cardinality overrides that apply within this metric. Must be omitted when `override_type` is `excluded`.
	PerTagLimits []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit `json:"per_tag_limits,omitempty"`
	// The maximum number of distinct tag value combinations allowed for this metric. Required when `override_type` is `limit_override`. Must be omitted when `override_type` is `excluded`.
	ValueLimit *int64 `json:"value_limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit instantiates a new ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit(metricName string, overrideType ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit {
	this := ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit{}
	this.MetricName = metricName
	this.OverrideType = overrideType
	return &this
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimitWithDefaults instantiates a new ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimitWithDefaults() *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit {
	this := ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit{}
	return &this
}

// GetLimitExceededAction returns the LimitExceededAction field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetLimitExceededAction() ObservabilityPipelineTagCardinalityLimitProcessorAction {
	if o == nil || o.LimitExceededAction == nil {
		var ret ObservabilityPipelineTagCardinalityLimitProcessorAction
		return ret
	}
	return *o.LimitExceededAction
}

// GetLimitExceededActionOk returns a tuple with the LimitExceededAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetLimitExceededActionOk() (*ObservabilityPipelineTagCardinalityLimitProcessorAction, bool) {
	if o == nil || o.LimitExceededAction == nil {
		return nil, false
	}
	return o.LimitExceededAction, true
}

// HasLimitExceededAction returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) HasLimitExceededAction() bool {
	return o != nil && o.LimitExceededAction != nil
}

// SetLimitExceededAction gets a reference to the given ObservabilityPipelineTagCardinalityLimitProcessorAction and assigns it to the LimitExceededAction field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) SetLimitExceededAction(v ObservabilityPipelineTagCardinalityLimitProcessorAction) {
	o.LimitExceededAction = &v
}

// GetMetricName returns the MetricName field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) SetMetricName(v string) {
	o.MetricName = v
}

// GetOverrideType returns the OverrideType field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetOverrideType() ObservabilityPipelineTagCardinalityLimitProcessorOverrideType {
	if o == nil {
		var ret ObservabilityPipelineTagCardinalityLimitProcessorOverrideType
		return ret
	}
	return o.OverrideType
}

// GetOverrideTypeOk returns a tuple with the OverrideType field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetOverrideTypeOk() (*ObservabilityPipelineTagCardinalityLimitProcessorOverrideType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideType, true
}

// SetOverrideType sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) SetOverrideType(v ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) {
	o.OverrideType = v
}

// GetPerTagLimits returns the PerTagLimits field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetPerTagLimits() []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit {
	if o == nil || o.PerTagLimits == nil {
		var ret []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit
		return ret
	}
	return o.PerTagLimits
}

// GetPerTagLimitsOk returns a tuple with the PerTagLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetPerTagLimitsOk() (*[]ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit, bool) {
	if o == nil || o.PerTagLimits == nil {
		return nil, false
	}
	return &o.PerTagLimits, true
}

// HasPerTagLimits returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) HasPerTagLimits() bool {
	return o != nil && o.PerTagLimits != nil
}

// SetPerTagLimits gets a reference to the given []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit and assigns it to the PerTagLimits field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) SetPerTagLimits(v []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) {
	o.PerTagLimits = v
}

// GetValueLimit returns the ValueLimit field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetValueLimit() int64 {
	if o == nil || o.ValueLimit == nil {
		var ret int64
		return ret
	}
	return *o.ValueLimit
}

// GetValueLimitOk returns a tuple with the ValueLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) GetValueLimitOk() (*int64, bool) {
	if o == nil || o.ValueLimit == nil {
		return nil, false
	}
	return o.ValueLimit, true
}

// HasValueLimit returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) HasValueLimit() bool {
	return o != nil && o.ValueLimit != nil
}

// SetValueLimit gets a reference to the given int64 and assigns it to the ValueLimit field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) SetValueLimit(v int64) {
	o.ValueLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LimitExceededAction != nil {
		toSerialize["limit_exceeded_action"] = o.LimitExceededAction
	}
	toSerialize["metric_name"] = o.MetricName
	toSerialize["override_type"] = o.OverrideType
	if o.PerTagLimits != nil {
		toSerialize["per_tag_limits"] = o.PerTagLimits
	}
	if o.ValueLimit != nil {
		toSerialize["value_limit"] = o.ValueLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LimitExceededAction *ObservabilityPipelineTagCardinalityLimitProcessorAction       `json:"limit_exceeded_action,omitempty"`
		MetricName          *string                                                        `json:"metric_name"`
		OverrideType        *ObservabilityPipelineTagCardinalityLimitProcessorOverrideType `json:"override_type"`
		PerTagLimits        []ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit `json:"per_tag_limits,omitempty"`
		ValueLimit          *int64                                                         `json:"value_limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MetricName == nil {
		return fmt.Errorf("required field metric_name missing")
	}
	if all.OverrideType == nil {
		return fmt.Errorf("required field override_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit_exceeded_action", "metric_name", "override_type", "per_tag_limits", "value_limit"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LimitExceededAction != nil && !all.LimitExceededAction.IsValid() {
		hasInvalidField = true
	} else {
		o.LimitExceededAction = all.LimitExceededAction
	}
	o.MetricName = *all.MetricName
	if !all.OverrideType.IsValid() {
		hasInvalidField = true
	} else {
		o.OverrideType = *all.OverrideType
	}
	o.PerTagLimits = all.PerTagLimits
	o.ValueLimit = all.ValueLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
