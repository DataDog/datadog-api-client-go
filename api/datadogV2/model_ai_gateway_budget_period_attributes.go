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

// AiGatewayBudgetPeriodAttributes Accounting state enforced for an AI Gateway budget's active period.
type AiGatewayBudgetPeriodAttributes struct {
	// Budget resource identifier.
	BudgetId uuid.UUID `json:"budget_id"`
	// Spend currently counted against the active accounting period.
	CurrentSpendMicrocents int64 `json:"current_spend_microcents"`
	// Capacity that new accounting periods inherit from the budget definition.
	DefaultCapacityMicrocents int64 `json:"default_capacity_microcents"`
	// Capacity enforced for the active accounting period.
	EffectiveCapacityMicrocents int64 `json:"effective_capacity_microcents"`
	// Datadog monitor ID provisioned for this budget. Null when monitor provisioning was skipped or failed.
	MonitorId datadog.NullableInt64 `json:"monitor_id"`
	// End of the active accounting window.
	WindowEnd time.Time `json:"window_end"`
	// Start of the active accounting window.
	WindowStart time.Time `json:"window_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetPeriodAttributes instantiates a new AiGatewayBudgetPeriodAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetPeriodAttributes(budgetId uuid.UUID, currentSpendMicrocents int64, defaultCapacityMicrocents int64, effectiveCapacityMicrocents int64, monitorId datadog.NullableInt64, windowEnd time.Time, windowStart time.Time) *AiGatewayBudgetPeriodAttributes {
	this := AiGatewayBudgetPeriodAttributes{}
	this.BudgetId = budgetId
	this.CurrentSpendMicrocents = currentSpendMicrocents
	this.DefaultCapacityMicrocents = defaultCapacityMicrocents
	this.EffectiveCapacityMicrocents = effectiveCapacityMicrocents
	this.MonitorId = monitorId
	this.WindowEnd = windowEnd
	this.WindowStart = windowStart
	return &this
}

// NewAiGatewayBudgetPeriodAttributesWithDefaults instantiates a new AiGatewayBudgetPeriodAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetPeriodAttributesWithDefaults() *AiGatewayBudgetPeriodAttributes {
	this := AiGatewayBudgetPeriodAttributes{}
	return &this
}

// GetBudgetId returns the BudgetId field value.
func (o *AiGatewayBudgetPeriodAttributes) GetBudgetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetBudgetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetId, true
}

// SetBudgetId sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetBudgetId(v uuid.UUID) {
	o.BudgetId = v
}

// GetCurrentSpendMicrocents returns the CurrentSpendMicrocents field value.
func (o *AiGatewayBudgetPeriodAttributes) GetCurrentSpendMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CurrentSpendMicrocents
}

// GetCurrentSpendMicrocentsOk returns a tuple with the CurrentSpendMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetCurrentSpendMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentSpendMicrocents, true
}

// SetCurrentSpendMicrocents sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetCurrentSpendMicrocents(v int64) {
	o.CurrentSpendMicrocents = v
}

// GetDefaultCapacityMicrocents returns the DefaultCapacityMicrocents field value.
func (o *AiGatewayBudgetPeriodAttributes) GetDefaultCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DefaultCapacityMicrocents
}

// GetDefaultCapacityMicrocentsOk returns a tuple with the DefaultCapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetDefaultCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultCapacityMicrocents, true
}

// SetDefaultCapacityMicrocents sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetDefaultCapacityMicrocents(v int64) {
	o.DefaultCapacityMicrocents = v
}

// GetEffectiveCapacityMicrocents returns the EffectiveCapacityMicrocents field value.
func (o *AiGatewayBudgetPeriodAttributes) GetEffectiveCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EffectiveCapacityMicrocents
}

// GetEffectiveCapacityMicrocentsOk returns a tuple with the EffectiveCapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetEffectiveCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveCapacityMicrocents, true
}

// SetEffectiveCapacityMicrocents sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetEffectiveCapacityMicrocents(v int64) {
	o.EffectiveCapacityMicrocents = v
}

// GetMonitorId returns the MonitorId field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *AiGatewayBudgetPeriodAttributes) GetMonitorId() int64 {
	if o == nil || o.MonitorId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId.Get()
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayBudgetPeriodAttributes) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorId.Get(), o.MonitorId.IsSet()
}

// SetMonitorId sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetMonitorId(v int64) {
	o.MonitorId.Set(&v)
}

// GetWindowEnd returns the WindowEnd field value.
func (o *AiGatewayBudgetPeriodAttributes) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetWindowStart returns the WindowStart field value.
func (o *AiGatewayBudgetPeriodAttributes) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetPeriodAttributes) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value.
func (o *AiGatewayBudgetPeriodAttributes) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetPeriodAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_id"] = o.BudgetId
	toSerialize["current_spend_microcents"] = o.CurrentSpendMicrocents
	toSerialize["default_capacity_microcents"] = o.DefaultCapacityMicrocents
	toSerialize["effective_capacity_microcents"] = o.EffectiveCapacityMicrocents
	toSerialize["monitor_id"] = o.MonitorId.Get()
	if o.WindowEnd.Nanosecond() == 0 {
		toSerialize["window_end"] = o.WindowEnd.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["window_end"] = o.WindowEnd.Format("2006-01-02T15:04:05.000Z07:00")
	}
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
func (o *AiGatewayBudgetPeriodAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetId                    *uuid.UUID            `json:"budget_id"`
		CurrentSpendMicrocents      *int64                `json:"current_spend_microcents"`
		DefaultCapacityMicrocents   *int64                `json:"default_capacity_microcents"`
		EffectiveCapacityMicrocents *int64                `json:"effective_capacity_microcents"`
		MonitorId                   datadog.NullableInt64 `json:"monitor_id"`
		WindowEnd                   *time.Time            `json:"window_end"`
		WindowStart                 *time.Time            `json:"window_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetId == nil {
		return fmt.Errorf("required field budget_id missing")
	}
	if all.CurrentSpendMicrocents == nil {
		return fmt.Errorf("required field current_spend_microcents missing")
	}
	if all.DefaultCapacityMicrocents == nil {
		return fmt.Errorf("required field default_capacity_microcents missing")
	}
	if all.EffectiveCapacityMicrocents == nil {
		return fmt.Errorf("required field effective_capacity_microcents missing")
	}
	if !all.MonitorId.IsSet() {
		return fmt.Errorf("required field monitor_id missing")
	}
	if all.WindowEnd == nil {
		return fmt.Errorf("required field window_end missing")
	}
	if all.WindowStart == nil {
		return fmt.Errorf("required field window_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_id", "current_spend_microcents", "default_capacity_microcents", "effective_capacity_microcents", "monitor_id", "window_end", "window_start"})
	} else {
		return err
	}
	o.BudgetId = *all.BudgetId
	o.CurrentSpendMicrocents = *all.CurrentSpendMicrocents
	o.DefaultCapacityMicrocents = *all.DefaultCapacityMicrocents
	o.EffectiveCapacityMicrocents = *all.EffectiveCapacityMicrocents
	o.MonitorId = all.MonitorId
	o.WindowEnd = *all.WindowEnd
	o.WindowStart = *all.WindowStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
