// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetReportingSummary Health across all counters in the active accounting window. Includes hidden counters; contains no scope values or identities.
type AiGatewayBudgetReportingSummary struct {
	// Number of counters at or above capacity.
	ExceededCapacityCount int64 `json:"exceeded_capacity_count"`
	// Number of counters in the active window. Counters are created on first use.
	InstanceCount int64 `json:"instance_count"`
	// How the budget limit applies. `shared` is one limit for all matching traffic. `per_instance` is one
	// limit for each counter of a template budget that splits by value.
	LimitMode AiGatewayBudgetLimitMode `json:"limit_mode"`
	// Number of counters at or above the warning threshold.
	OverWarningThresholdCount int64 `json:"over_warning_threshold_count"`
	// Combined capacity of the counters in the active window, in microcents. A shared budget reports its configured limit before its first counter exists.
	TotalCapacityMicrocents int64 `json:"total_capacity_microcents"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetReportingSummary instantiates a new AiGatewayBudgetReportingSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetReportingSummary(exceededCapacityCount int64, instanceCount int64, limitMode AiGatewayBudgetLimitMode, overWarningThresholdCount int64, totalCapacityMicrocents int64) *AiGatewayBudgetReportingSummary {
	this := AiGatewayBudgetReportingSummary{}
	this.ExceededCapacityCount = exceededCapacityCount
	this.InstanceCount = instanceCount
	this.LimitMode = limitMode
	this.OverWarningThresholdCount = overWarningThresholdCount
	this.TotalCapacityMicrocents = totalCapacityMicrocents
	return &this
}

// NewAiGatewayBudgetReportingSummaryWithDefaults instantiates a new AiGatewayBudgetReportingSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetReportingSummaryWithDefaults() *AiGatewayBudgetReportingSummary {
	this := AiGatewayBudgetReportingSummary{}
	return &this
}

// GetExceededCapacityCount returns the ExceededCapacityCount field value.
func (o *AiGatewayBudgetReportingSummary) GetExceededCapacityCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExceededCapacityCount
}

// GetExceededCapacityCountOk returns a tuple with the ExceededCapacityCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetReportingSummary) GetExceededCapacityCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededCapacityCount, true
}

// SetExceededCapacityCount sets field value.
func (o *AiGatewayBudgetReportingSummary) SetExceededCapacityCount(v int64) {
	o.ExceededCapacityCount = v
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AiGatewayBudgetReportingSummary) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetReportingSummary) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AiGatewayBudgetReportingSummary) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetLimitMode returns the LimitMode field value.
func (o *AiGatewayBudgetReportingSummary) GetLimitMode() AiGatewayBudgetLimitMode {
	if o == nil {
		var ret AiGatewayBudgetLimitMode
		return ret
	}
	return o.LimitMode
}

// GetLimitModeOk returns a tuple with the LimitMode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetReportingSummary) GetLimitModeOk() (*AiGatewayBudgetLimitMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitMode, true
}

// SetLimitMode sets field value.
func (o *AiGatewayBudgetReportingSummary) SetLimitMode(v AiGatewayBudgetLimitMode) {
	o.LimitMode = v
}

// GetOverWarningThresholdCount returns the OverWarningThresholdCount field value.
func (o *AiGatewayBudgetReportingSummary) GetOverWarningThresholdCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OverWarningThresholdCount
}

// GetOverWarningThresholdCountOk returns a tuple with the OverWarningThresholdCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetReportingSummary) GetOverWarningThresholdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverWarningThresholdCount, true
}

// SetOverWarningThresholdCount sets field value.
func (o *AiGatewayBudgetReportingSummary) SetOverWarningThresholdCount(v int64) {
	o.OverWarningThresholdCount = v
}

// GetTotalCapacityMicrocents returns the TotalCapacityMicrocents field value.
func (o *AiGatewayBudgetReportingSummary) GetTotalCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCapacityMicrocents
}

// GetTotalCapacityMicrocentsOk returns a tuple with the TotalCapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetReportingSummary) GetTotalCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCapacityMicrocents, true
}

// SetTotalCapacityMicrocents sets field value.
func (o *AiGatewayBudgetReportingSummary) SetTotalCapacityMicrocents(v int64) {
	o.TotalCapacityMicrocents = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetReportingSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["exceeded_capacity_count"] = o.ExceededCapacityCount
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["limit_mode"] = o.LimitMode
	toSerialize["over_warning_threshold_count"] = o.OverWarningThresholdCount
	toSerialize["total_capacity_microcents"] = o.TotalCapacityMicrocents

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetReportingSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExceededCapacityCount     *int64                    `json:"exceeded_capacity_count"`
		InstanceCount             *int64                    `json:"instance_count"`
		LimitMode                 *AiGatewayBudgetLimitMode `json:"limit_mode"`
		OverWarningThresholdCount *int64                    `json:"over_warning_threshold_count"`
		TotalCapacityMicrocents   *int64                    `json:"total_capacity_microcents"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExceededCapacityCount == nil {
		return fmt.Errorf("required field exceeded_capacity_count missing")
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if all.LimitMode == nil {
		return fmt.Errorf("required field limit_mode missing")
	}
	if all.OverWarningThresholdCount == nil {
		return fmt.Errorf("required field over_warning_threshold_count missing")
	}
	if all.TotalCapacityMicrocents == nil {
		return fmt.Errorf("required field total_capacity_microcents missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exceeded_capacity_count", "instance_count", "limit_mode", "over_warning_threshold_count", "total_capacity_microcents"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExceededCapacityCount = *all.ExceededCapacityCount
	o.InstanceCount = *all.InstanceCount
	if !all.LimitMode.IsValid() {
		hasInvalidField = true
	} else {
		o.LimitMode = *all.LimitMode
	}
	o.OverWarningThresholdCount = *all.OverWarningThresholdCount
	o.TotalCapacityMicrocents = *all.TotalCapacityMicrocents

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
