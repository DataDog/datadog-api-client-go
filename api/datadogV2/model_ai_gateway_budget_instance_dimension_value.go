// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstanceDimensionValue One value of a budget split dimension.
type AiGatewayBudgetInstanceDimensionValue struct {
	// Number of instance counters carrying this value.
	InstanceCount int64 `json:"instance_count"`
	// Spend summed across those instance counters, in microcents.
	SpendMicrocents int64 `json:"spend_microcents"`
	// Raw dimension value. User and Team values are UUIDs. A counter with no value for the dimension carries the null character (U+0000).
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetInstanceDimensionValue instantiates a new AiGatewayBudgetInstanceDimensionValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetInstanceDimensionValue(instanceCount int64, spendMicrocents int64, value string) *AiGatewayBudgetInstanceDimensionValue {
	this := AiGatewayBudgetInstanceDimensionValue{}
	this.InstanceCount = instanceCount
	this.SpendMicrocents = spendMicrocents
	this.Value = value
	return &this
}

// NewAiGatewayBudgetInstanceDimensionValueWithDefaults instantiates a new AiGatewayBudgetInstanceDimensionValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetInstanceDimensionValueWithDefaults() *AiGatewayBudgetInstanceDimensionValue {
	this := AiGatewayBudgetInstanceDimensionValue{}
	return &this
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AiGatewayBudgetInstanceDimensionValue) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionValue) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AiGatewayBudgetInstanceDimensionValue) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetSpendMicrocents returns the SpendMicrocents field value.
func (o *AiGatewayBudgetInstanceDimensionValue) GetSpendMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SpendMicrocents
}

// GetSpendMicrocentsOk returns a tuple with the SpendMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionValue) GetSpendMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpendMicrocents, true
}

// SetSpendMicrocents sets field value.
func (o *AiGatewayBudgetInstanceDimensionValue) SetSpendMicrocents(v int64) {
	o.SpendMicrocents = v
}

// GetValue returns the Value field value.
func (o *AiGatewayBudgetInstanceDimensionValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *AiGatewayBudgetInstanceDimensionValue) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetInstanceDimensionValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["spend_microcents"] = o.SpendMicrocents
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetInstanceDimensionValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceCount   *int64  `json:"instance_count"`
		SpendMicrocents *int64  `json:"spend_microcents"`
		Value           *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if all.SpendMicrocents == nil {
		return fmt.Errorf("required field spend_microcents missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_count", "spend_microcents", "value"})
	} else {
		return err
	}
	o.InstanceCount = *all.InstanceCount
	o.SpendMicrocents = *all.SpendMicrocents
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
