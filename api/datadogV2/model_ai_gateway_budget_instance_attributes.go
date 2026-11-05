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

// AiGatewayBudgetInstanceAttributes One budget instance counter.
type AiGatewayBudgetInstanceAttributes struct {
	// Budget identifier.
	BudgetId uuid.UUID `json:"budget_id"`
	// Capacity enforced for this instance, in microcents.
	EffectiveCapacityMicrocents int64 `json:"effective_capacity_microcents"`
	// Opaque canonical counter key. Empty for legacy pooled counters.
	InstanceKey string `json:"instance_key"`
	// Time the instance counter was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Spend counted against this instance, in microcents.
	SpendMicrocents int64 `json:"spend_microcents"`
	// Dimension values decoded from instance_key. User and Team values are raw UUIDs.
	SplitDims map[string]string `json:"split_dims"`
	// Start of the accounting window.
	WindowStart time.Time `json:"window_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetInstanceAttributes instantiates a new AiGatewayBudgetInstanceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetInstanceAttributes(budgetId uuid.UUID, effectiveCapacityMicrocents int64, instanceKey string, modifiedAt time.Time, spendMicrocents int64, splitDims map[string]string, windowStart time.Time) *AiGatewayBudgetInstanceAttributes {
	this := AiGatewayBudgetInstanceAttributes{}
	this.BudgetId = budgetId
	this.EffectiveCapacityMicrocents = effectiveCapacityMicrocents
	this.InstanceKey = instanceKey
	this.ModifiedAt = modifiedAt
	this.SpendMicrocents = spendMicrocents
	this.SplitDims = splitDims
	this.WindowStart = windowStart
	return &this
}

// NewAiGatewayBudgetInstanceAttributesWithDefaults instantiates a new AiGatewayBudgetInstanceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetInstanceAttributesWithDefaults() *AiGatewayBudgetInstanceAttributes {
	this := AiGatewayBudgetInstanceAttributes{}
	return &this
}

// GetBudgetId returns the BudgetId field value.
func (o *AiGatewayBudgetInstanceAttributes) GetBudgetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetBudgetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetId, true
}

// SetBudgetId sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetBudgetId(v uuid.UUID) {
	o.BudgetId = v
}

// GetEffectiveCapacityMicrocents returns the EffectiveCapacityMicrocents field value.
func (o *AiGatewayBudgetInstanceAttributes) GetEffectiveCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EffectiveCapacityMicrocents
}

// GetEffectiveCapacityMicrocentsOk returns a tuple with the EffectiveCapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetEffectiveCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveCapacityMicrocents, true
}

// SetEffectiveCapacityMicrocents sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetEffectiveCapacityMicrocents(v int64) {
	o.EffectiveCapacityMicrocents = v
}

// GetInstanceKey returns the InstanceKey field value.
func (o *AiGatewayBudgetInstanceAttributes) GetInstanceKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceKey
}

// GetInstanceKeyOk returns a tuple with the InstanceKey field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetInstanceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceKey, true
}

// SetInstanceKey sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetInstanceKey(v string) {
	o.InstanceKey = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *AiGatewayBudgetInstanceAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetSpendMicrocents returns the SpendMicrocents field value.
func (o *AiGatewayBudgetInstanceAttributes) GetSpendMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SpendMicrocents
}

// GetSpendMicrocentsOk returns a tuple with the SpendMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetSpendMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpendMicrocents, true
}

// SetSpendMicrocents sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetSpendMicrocents(v int64) {
	o.SpendMicrocents = v
}

// GetSplitDims returns the SplitDims field value.
func (o *AiGatewayBudgetInstanceAttributes) GetSplitDims() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.SplitDims
}

// GetSplitDimsOk returns a tuple with the SplitDims field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetSplitDimsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitDims, true
}

// SetSplitDims sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetSplitDims(v map[string]string) {
	o.SplitDims = v
}

// GetWindowStart returns the WindowStart field value.
func (o *AiGatewayBudgetInstanceAttributes) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceAttributes) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value.
func (o *AiGatewayBudgetInstanceAttributes) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetInstanceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_id"] = o.BudgetId
	toSerialize["effective_capacity_microcents"] = o.EffectiveCapacityMicrocents
	toSerialize["instance_key"] = o.InstanceKey
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["spend_microcents"] = o.SpendMicrocents
	toSerialize["split_dims"] = o.SplitDims
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
func (o *AiGatewayBudgetInstanceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetId                    *uuid.UUID         `json:"budget_id"`
		EffectiveCapacityMicrocents *int64             `json:"effective_capacity_microcents"`
		InstanceKey                 *string            `json:"instance_key"`
		ModifiedAt                  *time.Time         `json:"modified_at"`
		SpendMicrocents             *int64             `json:"spend_microcents"`
		SplitDims                   *map[string]string `json:"split_dims"`
		WindowStart                 *time.Time         `json:"window_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetId == nil {
		return fmt.Errorf("required field budget_id missing")
	}
	if all.EffectiveCapacityMicrocents == nil {
		return fmt.Errorf("required field effective_capacity_microcents missing")
	}
	if all.InstanceKey == nil {
		return fmt.Errorf("required field instance_key missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.SpendMicrocents == nil {
		return fmt.Errorf("required field spend_microcents missing")
	}
	if all.SplitDims == nil {
		return fmt.Errorf("required field split_dims missing")
	}
	if all.WindowStart == nil {
		return fmt.Errorf("required field window_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_id", "effective_capacity_microcents", "instance_key", "modified_at", "spend_microcents", "split_dims", "window_start"})
	} else {
		return err
	}
	o.BudgetId = *all.BudgetId
	o.EffectiveCapacityMicrocents = *all.EffectiveCapacityMicrocents
	o.InstanceKey = *all.InstanceKey
	o.ModifiedAt = *all.ModifiedAt
	o.SpendMicrocents = *all.SpendMicrocents
	o.SplitDims = *all.SplitDims
	o.WindowStart = *all.WindowStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
