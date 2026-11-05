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

// AiGatewayBudgetInstancesSummaryAttributes Aggregate state across budget instance counters.
type AiGatewayBudgetInstancesSummaryAttributes struct {
	// Budget identifier.
	BudgetId uuid.UUID `json:"budget_id"`
	// Number of instances at or above capacity.
	ExceededCapacityCount int64 `json:"exceeded_capacity_count"`
	// Number of instance counters.
	InstanceCount int64 `json:"instance_count"`
	// Number of instances at or above the warning threshold.
	OverWarningThresholdCount int64 `json:"over_warning_threshold_count"`
	// Whether instances were withheld because their provider or model is hidden. Summary counts and totals exclude those instances.
	RestrictedDetails bool `json:"restricted_details"`
	// Spend summed across all instances, in microcents.
	TotalSpendMicrocents int64 `json:"total_spend_microcents"`
	// Start of the accounting window.
	WindowStart time.Time `json:"window_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetInstancesSummaryAttributes instantiates a new AiGatewayBudgetInstancesSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetInstancesSummaryAttributes(budgetId uuid.UUID, exceededCapacityCount int64, instanceCount int64, overWarningThresholdCount int64, restrictedDetails bool, totalSpendMicrocents int64, windowStart time.Time) *AiGatewayBudgetInstancesSummaryAttributes {
	this := AiGatewayBudgetInstancesSummaryAttributes{}
	this.BudgetId = budgetId
	this.ExceededCapacityCount = exceededCapacityCount
	this.InstanceCount = instanceCount
	this.OverWarningThresholdCount = overWarningThresholdCount
	this.RestrictedDetails = restrictedDetails
	this.TotalSpendMicrocents = totalSpendMicrocents
	this.WindowStart = windowStart
	return &this
}

// NewAiGatewayBudgetInstancesSummaryAttributesWithDefaults instantiates a new AiGatewayBudgetInstancesSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetInstancesSummaryAttributesWithDefaults() *AiGatewayBudgetInstancesSummaryAttributes {
	this := AiGatewayBudgetInstancesSummaryAttributes{}
	return &this
}

// GetBudgetId returns the BudgetId field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetBudgetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetBudgetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetId, true
}

// SetBudgetId sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetBudgetId(v uuid.UUID) {
	o.BudgetId = v
}

// GetExceededCapacityCount returns the ExceededCapacityCount field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetExceededCapacityCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExceededCapacityCount
}

// GetExceededCapacityCountOk returns a tuple with the ExceededCapacityCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetExceededCapacityCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededCapacityCount, true
}

// SetExceededCapacityCount sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetExceededCapacityCount(v int64) {
	o.ExceededCapacityCount = v
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetOverWarningThresholdCount returns the OverWarningThresholdCount field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetOverWarningThresholdCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OverWarningThresholdCount
}

// GetOverWarningThresholdCountOk returns a tuple with the OverWarningThresholdCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetOverWarningThresholdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverWarningThresholdCount, true
}

// SetOverWarningThresholdCount sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetOverWarningThresholdCount(v int64) {
	o.OverWarningThresholdCount = v
}

// GetRestrictedDetails returns the RestrictedDetails field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetRestrictedDetails() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RestrictedDetails
}

// GetRestrictedDetailsOk returns a tuple with the RestrictedDetails field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetRestrictedDetailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictedDetails, true
}

// SetRestrictedDetails sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetRestrictedDetails(v bool) {
	o.RestrictedDetails = v
}

// GetTotalSpendMicrocents returns the TotalSpendMicrocents field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetTotalSpendMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalSpendMicrocents
}

// GetTotalSpendMicrocentsOk returns a tuple with the TotalSpendMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetTotalSpendMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSpendMicrocents, true
}

// SetTotalSpendMicrocents sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetTotalSpendMicrocents(v int64) {
	o.TotalSpendMicrocents = v
}

// GetWindowStart returns the WindowStart field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstancesSummaryAttributes) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value.
func (o *AiGatewayBudgetInstancesSummaryAttributes) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetInstancesSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_id"] = o.BudgetId
	toSerialize["exceeded_capacity_count"] = o.ExceededCapacityCount
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["over_warning_threshold_count"] = o.OverWarningThresholdCount
	toSerialize["restricted_details"] = o.RestrictedDetails
	toSerialize["total_spend_microcents"] = o.TotalSpendMicrocents
	if o.WindowStart.Nanosecond() == 0 {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetInstancesSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetId                  *uuid.UUID `json:"budget_id"`
		ExceededCapacityCount     *int64     `json:"exceeded_capacity_count"`
		InstanceCount             *int64     `json:"instance_count"`
		OverWarningThresholdCount *int64     `json:"over_warning_threshold_count"`
		RestrictedDetails         *bool      `json:"restricted_details"`
		TotalSpendMicrocents      *int64     `json:"total_spend_microcents"`
		WindowStart               *time.Time `json:"window_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetId == nil {
		return fmt.Errorf("required field budget_id missing")
	}
	if all.ExceededCapacityCount == nil {
		return fmt.Errorf("required field exceeded_capacity_count missing")
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if all.OverWarningThresholdCount == nil {
		return fmt.Errorf("required field over_warning_threshold_count missing")
	}
	if all.RestrictedDetails == nil {
		return fmt.Errorf("required field restricted_details missing")
	}
	if all.TotalSpendMicrocents == nil {
		return fmt.Errorf("required field total_spend_microcents missing")
	}
	if all.WindowStart == nil {
		return fmt.Errorf("required field window_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_id", "exceeded_capacity_count", "instance_count", "over_warning_threshold_count", "restricted_details", "total_spend_microcents", "window_start"})
	} else {
		return err
	}
	o.BudgetId = *all.BudgetId
	o.ExceededCapacityCount = *all.ExceededCapacityCount
	o.InstanceCount = *all.InstanceCount
	o.OverWarningThresholdCount = *all.OverWarningThresholdCount
	o.RestrictedDetails = *all.RestrictedDetails
	o.TotalSpendMicrocents = *all.TotalSpendMicrocents
	o.WindowStart = *all.WindowStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
